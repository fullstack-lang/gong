package models

import (
	"encoding/base64"
	"fmt"
	"os"

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

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	stageForRenderingConf := NewStage("renderingConf")
	ParseAstFromBytes(stageForRenderingConf, decodedBytes)

	proxy.stager.processRenderingConf(stageForRenderingConf)

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
	// 1. Create a temporary file. The "" for dir means use the OS default.
	// The pattern "ast-*.tmp" helps in identifying the file if it's left behind.
	tempFile, err := os.CreateTemp("", "ast-*.tmp")
	if err != nil {
		return fmt.Errorf("could not create temporary file: %w", err)
	}

	// 2. IMPORTANT: Schedule the file to be removed when the function returns.
	// `defer` ensures this runs even if subsequent operations fail.
	defer os.Remove(tempFile.Name())

	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	// 3. Write the input bytes to the temporary file.
	if _, err := tempFile.Write(input); err != nil {
		return fmt.Errorf("could not write to temporary file: %w", err)
	}

	// 4. Close the file to ensure all data is flushed to disk before parsing.
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("could not close temporary file: %w", err)
	}

	// 5. Now, call your original function with the path to our new temp file.
	return ParseAstFile(stage, tempFile.Name(), false)
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
