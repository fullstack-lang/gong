package models

import (
	"encoding/base64"
	"fmt"
	"go/parser"
	"go/token"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type FileToUploadProxy struct {
	stager *Stager
}

func (stager *Stager) load() {
	stager.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &loadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your <library>.go file here or ",
	}

	message.Stage(stager.loadStage)

	stager.loadStage.Commit()
}

type loadProxy struct {
	stager *Stager
}

func (proxy *loadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	proxy.stager.fileName = uploadedFile.GetName()

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	fmt.Println(string(decodedBytes))

	proxy.stager.stage.Reset()
	fmt.Println("OnFileUpload: after reset")
	ParseAstFromBytes(proxy.stager.stage, decodedBytes)
	fmt.Println("OnFileUpload: after parse")
	proxy.stager.stage.Commit()
	fmt.Println("OnFileUpload: after commit")

	return nil
}

func ParseAstFromBytes(stage *Stage, input []byte) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", input, parser.ParseComments)
	if errParser != nil {
		return fmt.Errorf("Unable to parse: %w", errParser)
	}
	return ParseAstFileFromAst(stage, inFile, fset, false)
}
