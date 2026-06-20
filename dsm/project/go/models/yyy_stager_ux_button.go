// generated code (do not edit)
package models

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
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
			log.Println("Exporting as stage")

			loadStage := stager.loadStage
			loadStage.Reset()

			stage := stager.stage

			fileToDownload := new(load.FileToDownload).Stage(loadStage)

			if stager.fileName == "" {
				stager.fileName = stager.stage.GetName() + ".go"
			}

			// 1. Compile the regex to match "YYYYMMDD HHMM " at the start of the string
			prefixRegex := regexp.MustCompile(`^\d{8} \d{4} `)

			// 2. Remove the matched prefix if it exists, leaving just the base file name
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			// 3. Prepend the fresh timestamp
			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

			fileName := filepath.Base(fileToDownload.Name)

			// Use MarshallToString for pure in-memory export (WASM compatible)
			res, err := stage.MarshallToString(stage.MetaPackageImportPath, "main")
			if err != nil {
				log.Println("Error marshalling to string:", err)
				return
			}

			// The content is now a string in memory.
			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(res))

			stager.loadStage.Commit()

			log.Println("Finished exporting file", fileName)

			time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we delete the file.
			stager.load()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export XL file",
		Icon:  string(buttons.BUTTON_table_view),
		Label: "Export XL file",
		OnClick: func() {
			log.Println("Exporting as XL file")

			loadStage := stager.loadStage
			loadStage.Reset()

			stage := stager.stage

			fileToDownload := new(load.FileToDownload).Stage(loadStage)

			if stager.fileName == "" {
				stager.fileName = stager.stage.GetName() + ".go"
			}

			// 1. Compile the regex to match "YYYYMMDD HHMM " at the start of the string
			prefixRegex := regexp.MustCompile(`^\d{8} \d{4} `)

			// 2. Remove the matched prefix if it exists, leaving just the base file name
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")
			cleanFileName = strings.TrimSuffix(cleanFileName, ".go") + ".xlsx"

			fileToDownload.Name = time.Now().Format("20060102 0304 ") + cleanFileName

			excelBytes, err := SerializeStageAsBytes(stage, false)
			if err != nil {
				log.Println("Error serializing stage:", err)
				return
			}

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(excelBytes)

			stager.loadStage.Commit()

			log.Println("Finished exporting XL file", fileToDownload.Name)

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
