package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (stager *Stager) exportWebsite() {
	stager.ssgStage.Reset()

	content := ssg.Content{
		Name:           "Root to project the website",
		MardownContent: "## Barrgraph website",
	}

	if len(GetGongstrucsSorted[*Diagram](stager.stage)) > 0 {
		chapter := &ssg.Chapter{
			Name:           "Diagrams",
			MardownContent: "### Diagrams\n",
		}
		content.Chapters = append(content.Chapters, chapter)

		for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
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
			chapter.Sections = append(chapter.Sections, section)
		}
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

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + "site.zip"

	stager.loadStage.Commit()
}
