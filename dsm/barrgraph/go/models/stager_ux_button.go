package models

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

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

			tempFile, err := os.CreateTemp("", fileName)
			if err != nil {
				log.Fatalf("Failed to create temporary file: %v", err)
			}

			// 2. Defer the removal of the file to ensure it's cleaned up.
			defer os.Remove(tempFile.Name())

			// 3. Marshall the data into the temporary file.
			stage.Marshall(tempFile, "github.com/fullstack-lang/gong/dsm/barrgraph/go/models", "main")

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
		Name:  "Export website",
		Icon:  string(buttons.BUTTON_web),
		Label: "Export website",
		OnClick: func() {
			stager.ssgStage.Reset()

			content := ssg.Content{
				Name:           "Root to project the website",
				MardownContent: "## Barrgraph website",
				OutputPath:     "./generated static web site",
			}

			if len(GetGongstrucsSorted[*Diagram](stager.stage)) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Diagrams",
					MardownContent: "### Diagrams\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Diagrams",
					MardownContent: "### Diagrams\n",
				}
				for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
					page.MardownContent += "\n- " + diagram.Name

					svgObject := stager.generateSvgObject(diagram)
					_ = svgObject
					svgString, maxX, maxY := svgObject.GenerateString()

					// Replace 100% width/height with exact pixel values to prevent the
					// SVG from stretching and looking "too big" in the browser.
					svgString = strings.Replace(svgString, `width="100%"`, fmt.Sprintf(`width="%f"`, maxX), 1)
					svgString = strings.Replace(svgString, `height="100%"`, fmt.Sprintf(`height="%f"`, maxY), 1)

					section := &ssg.Section{
						Name:    diagram.Name,
						IsImage: true,
						SvgImage: &ssg.SvgImage{
							Content: svgString,
						},
					}
					page.Sections = append(page.Sections, section)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			movements := GetGongstrucsSorted[*Movement](stager.stage)
			if len(movements) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Movements",
					MardownContent: "### Movements\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Movements Table",
					MardownContent: "### Movements Table\n\n| Name | Date | Taxonomic Filter | Featured | Major | Minor | Additional Name |\n|---|---|---|---|---|---|---|\n",
				}
				for _, m := range movements {
					date := ""
					if !m.HideDate && !m.Date.IsZero() {
						date = m.Date.Format("2006")
					}
					page.MardownContent += fmt.Sprintf("| %s | %s | %s | %t | %t | %t | %s |\n", m.Name, date, m.TaxonomicFilter, m.IsFeatured, m.IsMajor, m.IsMinor, m.AdditionnalName)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			artists := GetGongstrucsSorted[*Artist](stager.stage)
			if len(artists) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Artists",
					MardownContent: "### Artists\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Artists Table",
					MardownContent: "### Artists Table\n\n| Name | Dead | Date of Death | Place |\n|---|---|---|---|\n",
				}
				for _, a := range artists {
					dod := ""
					if a.IsDead && !a.DateOfDeath.IsZero() {
						dod = a.DateOfDeath.Format("2006")
					}
					place := ""
					if a.Place != nil {
						place = a.Place.Name
					}
					page.MardownContent += fmt.Sprintf("| %s | %t | %s | %s |\n", a.Name, a.IsDead, dod, place)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			artefacts := GetGongstrucsSorted[*ArtefactType](stager.stage)
			if len(artefacts) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Artefact Types",
					MardownContent: "### Artefact Types\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Artefact Types Table",
					MardownContent: "### Artefact Types Table\n\n| Name |\n|---|\n",
				}
				for _, a := range artefacts {
					page.MardownContent += fmt.Sprintf("| %s |\n", a.Name)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			influences := GetGongstrucsSorted[*Influence](stager.stage)
			if len(influences) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Influences",
					MardownContent: "### Influences\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Influences Table",
					MardownContent: "### Influences Table\n\n| Name | Source | Target | Hypothetical |\n|---|---|---|---|\n",
				}
				for _, i := range influences {
					source := ""
					if i.SourceMovement != nil {
						source = i.SourceMovement.Name
					}
					if i.SourceArtefactType != nil {
						source = i.SourceArtefactType.Name
					}
					if i.SourceArtist != nil {
						source = i.SourceArtist.Name
					}

					target := ""
					if i.TargetMovement != nil {
						target = i.TargetMovement.Name
					}
					if i.TargetArtefactType != nil {
						target = i.TargetArtefactType.Name
					}
					if i.TargetArtist != nil {
						target = i.TargetArtist.Name
					}

					page.MardownContent += fmt.Sprintf("| %s | %s | %s | %t |\n", i.Name, source, target, i.IsHypothtical)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			places := GetGongstrucsSorted[*Place](stager.stage)
			if len(places) > 0 {
				chapter := &ssg.Chapter{
					Name:           "Places",
					MardownContent: "### Places\n",
				}
				content.Chapters = append(content.Chapters, chapter)

				page := &ssg.Page{
					Name:           "Places Table",
					MardownContent: "### Places Table\n\n| Name |\n|---|\n",
				}
				for _, p := range places {
					page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
				}
				chapter.Pages = append(chapter.Pages, page)
			}

			ssg.StageBranch(stager.ssgStage, &content)

			stager.ssgStage.Commit()

			zipData, err := stager.ssgStage.Generation(true)
			if err != nil {
				log.Println(err)
			}

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)
			fileToDownload.Base64EncodedContent = zipData

			fileToDownload.Name = "site.zip"

			stager.loadStage.Commit()
		},
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}
