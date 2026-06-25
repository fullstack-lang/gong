// generated code (do not edit)
package models

import (
	"encoding/base64"
	"log"
	"os"
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
		Name:  "Stop",
		Icon:  string(buttons.BUTTON_stop_circle),
		Label: "Stop",
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

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export stage as Go file",
		Icon:  string(buttons.BUTTON_file_download),
		Label: "Export stage as Go file",
		OnClick: func() {
			log.Println("Exporting stage as Go file")

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload)

			if stager.fileName == "" {
				pkgPath := stager.stage.MetaPackageImportPath
				pkgName := ""
				parts := strings.Split(pkgPath, "/")
				if len(parts) >= 3 {
					pkgName = parts[len(parts)-3]
				}
				stager.fileName = pkgName + "-" + stager.stage.GetName() + ".go"
			}

			prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

			stageString, err := stager.stage.MarshallToString(stager.stage.MetaPackageImportPath, "models")
			if err != nil {
				log.Println("Error serializing stage: " + err.Error())
				return
			}

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))

			load.StageBranch(stager.loadStage, fileToDownload)
			stager.loadStage.Commit()

			time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
			stager.load()
		},
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}
