package models

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

func (stager *Stager) button() {
	buttonStage := stager.buttonStage
	buttonStage.Reset()

	layout := new(button.Layout).Stage(buttonStage)

	group1 := new(button.Group).Stage(buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	{
		buttonExportModifiedReqif := button.NewButton(
			// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
			&ResetReqifButtonProxy{
				stager: stager,
			},
			"Reset project file",
			string(buttons.BUTTON_reset_tv),
			"Reset project file",
		)

		group1.Buttons = append(group1.Buttons, buttonExportModifiedReqif)
	}
	{
		buttonExportRenderingCong := button.NewButton(
			&downloadButtonProxy{
				stager: stager,
			},
			"Export Rendering Configuration",
			string(buttons.BUTTON_fact_check),
			"Export Rendering Configuration",
		)
		group1.Buttons = append(group1.Buttons, buttonExportRenderingCong)

	}

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}

type ResetReqifButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (proxy *ResetReqifButtonProxy) GetButtonsStage() *button.Stage {
	return proxy.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (proxy *ResetReqifButtonProxy) OnAfterUpdateButton() {

	proxy.stager.stage.Reset()
	proxy.stager.stage.Commit()
}

type downloadButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *downloadButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (proxy *downloadButtonProxy) OnAfterUpdateButton() {
	log.Println("Exporting the rendering configuration")

	stager := proxy.stager
	loadStage := stager.loadStage
	loadStage.Reset()

	stage := stager.stage

	fileToDownload := new(load.FileToDownload).Stage(loadStage)

	if stager.fileName == "" {
		stager.fileName = "project.go"
	}

	fileToDownload.Name = stager.fileName

	fileName := filepath.Base(fileToDownload.Name)

	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}

	// 2. Defer the removal of the file to ensure it's cleaned up.
	defer os.Remove(tempFile.Name())

	// 3. Marshall the data into the temporary file.
	stage.Marshall(tempFile, "github.com/fullstack-lang/gong/dsme/project/go/models", "main")

	// 4. Read the content back from the file.
	// os.ReadFile needs a file path, so we use the name.
	// First, ensure the file is closed so all data is flushed to disk.
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Failed to close temp file: %v", err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		log.Fatalf("Failed to read back temp file: %v", err)
	}

	// 5. The content is now a string in memory, and the file will be deleted.
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(content)

	stager.loadStage.Commit()

	log.Println("Finished exporting project file", tempFile.Name())

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we delete the file.
	stager.load()
}
