package models

import (
	"encoding/base64"
	"fmt"
	"log"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (stager *Stager) exportWebsite() {
	fmt.Println("exportWebsite: Starting website export process...")
	stager.ssgStage.Reset()

	content := ssg.Content{
		Name:           "Root to project the website",
		MardownContent: "## Project website",
	}

	for _, lib := range GetGongstrucsSorted[*Library](stager.stage) {
		fmt.Println("exportWebsite: Processing library:", lib.Name)
		chapter := &ssg.Chapter{
			Name:           lib.Name,
			MardownContent: "### " + lib.Name,
		}
		content.Chapters = append(content.Chapters, chapter)

		if len(lib.SubLibraries) > 0 {
			page := &ssg.Page{
				Name:           "SubLibraries",
				MardownContent: "### SubLibraries\n",
			}
			for _, l := range lib.SubLibraries {
				page.MardownContent += "\n- " + l.Name
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		page := &ssg.Page{
			Name:           "XL files",
			MardownContent: "### XL files\n",
		}

		fmt.Println("exportWebsite: Generating Excel files...")
		excelBytes, err := SerializeStageAsBytes(stager.stage, false)
		if err != nil {
			log.Println("Error creating Excel export in memory:", err)
			fmt.Println("exportWebsite: Error creating Excel in memory:", err)
		} else {
			fmt.Printf("exportWebsite: Excel generated successfully. Bytes: %d\n", len(excelBytes))
			encodedXL := base64.StdEncoding.EncodeToString(excelBytes)
			dataURI := "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64," + encodedXL
			page.MardownContent += fmt.Sprintf("\n[Download %s.xlsx](%s)\n", lib.Name, dataURI)

			section := &ssg.Section{
				Name: "Embedded XL file",
			}
			page.Sections = append(page.Sections, section)
		}

		chapter.Pages = append(chapter.Pages, page)
	}

	content.LogoSVGFile = stager.GetRootLibrary().LogoSVGFile

	fmt.Println("exportWebsite: Staging SSG branch...")
	ssg.StageBranch(stager.ssgStage, &content)

	fmt.Println("exportWebsite: Committing SSG stage...")
	stager.ssgStage.Commit()

	fmt.Println("exportWebsite: Calling ssgStage.Generation(true)...")
	zipData, err := stager.ssgStage.Generation(true)
	if err != nil {
		log.Println(err)
		fmt.Println("exportWebsite: ERROR during SSG Generation:", err)
	}
	fmt.Printf("exportWebsite: Generation completed. zipData length: %d\n", len(zipData))

	fmt.Println("exportWebsite: Preparing download payload...")
	stager.loadStage.Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)
	fileToDownload.Base64EncodedContent = zipData

	fileToDownload.Name = "site.zip"

	fmt.Println("exportWebsite: Committing loadStage...")
	stager.loadStage.Commit()
	fmt.Println("exportWebsite: Website export process finished.")
}
