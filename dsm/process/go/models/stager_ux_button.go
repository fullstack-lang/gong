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

	layout := new(button.Layout)

	group1 := new(button.Group)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Reset file",
		Icon:  string(buttons.BUTTON_reset_tv),
		Label: "Reset file",
		OnClick: func() {
			stager.stage.Reset()
			stager.stage.Commit()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export file",
		Icon:  string(buttons.BUTTON_fact_check),
		Label: "Export file",
		OnClick: func() {
			log.Println("Exporting the rendering configuration")

			loadStage := stager.loadStage
			loadStage.Reset()

			stage := stager.stage

			fileToDownload := new(load.FileToDownload).Stage(loadStage)

			if stager.fileName == "" {
				stager.fileName = "library.go"
			}

			fileToDownload.Name = stager.fileName

			fileName := filepath.Base(fileToDownload.Name)

			// 2. Use MarshallToString for pure in-memory export (WASM compatible)
			res, err := stage.MarshallToString("github.com/fullstack-lang/gong/dsm/project/go/models", "main")
			if err != nil {
				log.Println("Error marshalling to string:", err)
				return
			}

			// 3. The content is now a string in memory.
			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(res))

			stager.loadStage.Commit()

			log.Println("Finished exporting file", fileName)

			time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we delete the file.
			stager.load()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Stop for maintenance",
		Icon:  string(buttons.BUTTON_stop_circle),
		Label: "Stop for maintenance",
		OnClick: func() {
			log.Println("Stop")
			os.Exit(0)
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:    "Export website",
		Icon:    string(buttons.BUTTON_web),
		Label:   "Export website",
		OnClick: stager.exportWebsite,
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}
