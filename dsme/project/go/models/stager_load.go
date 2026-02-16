package models

import (
	"encoding/base64"
	"fmt"
	"os"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type FileToUploadProxy struct {
	stager *Stager
}

func (stager *Stager) load() {

	stager.loadStage.Reset()

	var fileToUpload = &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &loadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your .reqif or .reqifz file here or ",
	}

	message.Stage(stager.loadStage)

	stager.loadStage.Commit()
}

type loadProxy struct {
	stager *Stager
}

func (proxy *loadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fileName := uploadedFile.GetName()
	_ = fileName

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	proxy.stager.stage.Reset()
	ParseAstFromBytes(proxy.stager.stage, decodedBytes)
	proxy.stager.stage.Commit()

	return nil
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
