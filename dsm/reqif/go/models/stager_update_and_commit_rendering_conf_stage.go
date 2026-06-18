package models

import (
	"encoding/base64"
	"fmt"
	"go/parser"
	"go/token"
	"log"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

// StageAllOfTypeToAnotherStage stages all instances of type T from source to dest stage
func StageAllOfTypeToAnotherStage[T PointerToGongstruct](source, dest *Stage) {
	for o := range *GetGongstructInstancesSetFromPointerType[T](source) {
		o.StageVoid(dest)
	}
}

// UnstageAllOfType unstages all instances of type T from a stage
func UnstageAllOfType[T PointerToGongstruct](stage *Stage) {
	for o := range *GetGongstructInstancesSetFromPointerType[T](stage) {
		o.UnstageVoid(stage)
	}
}

type RenderingConfFileToUploadProxy struct {
	stager *Stager
}

func (proxy *RenderingConfFileToUploadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fileName := uploadedFile.GetName()
	_ = fileName

	log.Printf("OnFileUpload: Decoding base64 content of size %d", len(uploadedFile.Base64EncodedContent))
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		log.Printf("OnFileUpload ERROR: base64 decode failed: %v", err)
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	log.Printf("OnFileUpload: Parsed %d decoded bytes", len(decodedBytes))

	stageForRenderingConf := NewStage("renderingConf")
	log.Printf("OnFileUpload: Parsing AST from bytes...")
	err = ParseAstFromBytes(stageForRenderingConf, decodedBytes)
	if err != nil {
		log.Printf("OnFileUpload ERROR: ParseAstFromBytes failed: %v", err)
		return fmt.Errorf("ParseAstFromBytes failed: %w", err)
	}

	log.Printf("OnFileUpload: Successfully parsed AST. Processing rendering conf...")
	proxy.stager.processRenderingConf(stageForRenderingConf)

	log.Printf("OnFileUpload: Finished processing rendering conf.")
	return nil
}

func (stager *Stager) updateAndCommitLoadRenderingConfStage() {

	stager.loadRenderingConfStage.Reset()

	var fileToUpload = &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &RenderingConfFileToUploadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadRenderingConfStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your rendering conf .go file here or ",
	}

	message.Stage(stager.loadRenderingConfStage)

	stager.loadRenderingConfStage.Commit()

}

func ParseAstFromBytes(stage *Stage, input []byte) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", input, parser.ParseComments)
	if errParser != nil {
		return fmt.Errorf("unable to parse bytes: %w", errParser)
	}

	return ParseAstFileFromAst(stage, inFile, fset, false)
}

func (stager *Stager) processRenderingConf(stageForRenderingConf *Stage) {

	stage := stager.GetStage()

	// remove existing
	UnstageAllOfType[*SPECIFICATION_Rendering](stage)
	UnstageAllOfType[*SPEC_OBJECT_TYPE_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_DATE_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_INTEGER_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_REAL_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_STRING_Rendering](stage)
	UnstageAllOfType[*ATTRIBUTE_DEFINITION_XHTML_Rendering](stage)

	// stage those in storage
	StageAllOfTypeToAnotherStage[*SPECIFICATION_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*SPEC_OBJECT_TYPE_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_DATE_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_INTEGER_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_REAL_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_STRING_Rendering](stageForRenderingConf, stage)
	StageAllOfTypeToAnotherStage[*ATTRIBUTE_DEFINITION_XHTML_Rendering](stageForRenderingConf, stage)

	stager.enforceRenderingConfigurationSemantic()

	stager.GetSpecificationsTreeUpdater().UpdateAttributeDefinitionNb(stager)
	stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(stager)
	stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(stager)
	stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsTreeStage(stager)
	stager.updateAndCommitLoadRenderingConfStage()
	stager.UpdateAndCommitLoadReqifStage()
}
