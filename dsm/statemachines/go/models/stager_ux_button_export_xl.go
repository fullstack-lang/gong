package models

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

func (stager *Stager) buttonsExportXL() {

	stage := stager.buttonExportXLStage
	stage.Reset()

	layout := &button.Layout{}

	group := &button.Group{
		NbColumns: 2,
	}
	layout.Groups = append(layout.Groups, group)

	button_ := new(button.Button)
	button_.Name = "Export XL"
	button_.Label = "Export XL"
	button_.Icon = string(buttons.BUTTON_download)
	group.Buttons = append(group.Buttons, button_)

	proxy := &ExportXLButtonProxy{
		stager: stager,
	}

	button_.Proxy = proxy

	buttonKill := button.NewButton(
		&StopButtonProxy{
			stager: stager,
		},
		"Stop for maintenance",
		string(buttons.BUTTON_stop_circle),
		"Stop for maintenance",
	)

	group.Buttons = append(group.Buttons, buttonKill)

	button.StageBranch(stage, layout)

	stage.Commit()
}

type ExportXLButtonProxy struct {
	stager *Stager
}

// Updated implements models.ButtonProxyInterface.
func (proxy *ExportXLButtonProxy) Updated() {

	stager := proxy.stager

	loadStage := stager.loadStage
	loadStage.Reset()
	fileToDownload := new(load.FileToDownload).Stage(loadStage)

	stager.fileName = fmt.Sprintf("%s.xlsx", time.Now().Local().Format("2006-01-02-15-04-05"))

	fileToDownload.Name = stager.fileName

	tempFile, err := os.CreateTemp("", stager.fileName)
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}

	// 2. Defer the removal of the file to ensure it's cleaned up.
	defer os.Remove(tempFile.Name())

	// 3. Marshall the data into the temporary file.
	SerializeStageMessageTypeTransitionEtat(proxy.stager.stage, tempFile)

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

	log.Println("Finished exporting file", tempFile.Name())

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we delete the file.
	stager.load()
}

type StopButtonProxy struct {
	stager *Stager
}

func (e *StopButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonExportXLStage
}

func (proxy *StopButtonProxy) OnAfterUpdateButton() {

	log.Println("Arret demandé de l'application")
	os.Exit(0)
}
