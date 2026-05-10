package models

import (
	"fmt"
	"log"
	"strings"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (stager *Stager) exportWebsite() {
	stager.ssgStage.Reset()

	content := ssg.Content{
		Name:           "Root to process the website",
		ContentPath:    "/tmp/process",
		MardownContent: "## Process website",
		OutputPath:     "./generated static web site",
	}

	for _, process := range GetGongstrucsSorted[*Process](stager.stage) {
		chapter := &ssg.Chapter{
			Name:           process.Name,
			MardownContent: "### " + process.Name,
		}
		content.Chapters = append(content.Chapters, chapter)

		if len(process.Participants) > 0 {
			page := &ssg.Page{
				Name:           "Participants",
				MardownContent: "### Participants\n\n| Name |\n|---|\n",
			}
			for _, p := range process.Participants {
				page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		if len(process.ExternalParticipants) > 0 {
			page := &ssg.Page{
				Name:           "External Participants",
				MardownContent: "### External Participants\n\n| Name |\n|---|\n",
			}
			for _, p := range process.ExternalParticipants {
				page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		var allTasks []*Task
		var allControlFlows []*ControlFlow
		for _, p := range process.Participants {
			allTasks = append(allTasks, p.Tasks...)
			allControlFlows = append(allControlFlows, p.ControlFlows...)
		}

		if len(allTasks) > 0 {
			page := &ssg.Page{
				Name:           "Tasks",
				MardownContent: "### Tasks\n\n| Name | Participant |\n|---|---|\n",
			}
			for _, p := range process.Participants {
				for _, t := range p.Tasks {
					page.MardownContent += fmt.Sprintf("| %s | %s |\n", t.Name, p.Name)
				}
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		if len(allControlFlows) > 0 {
			page := &ssg.Page{
				Name:           "Control Flows",
				MardownContent: "### Control Flows\n\n| Name | Start | End |\n|---|---|---|\n",
			}
			for _, c := range allControlFlows {
				start := ""
				if c.Start != nil {
					start = c.Start.Name
				}
				end := ""
				if c.End != nil {
					end = c.End.Name
				}
				page.MardownContent += fmt.Sprintf("| %s | %s | %s |\n", c.Name, start, end)
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		if len(process.DataFlows) > 0 {
			page := &ssg.Page{
				Name:           "Data Flows",
				MardownContent: "### Data Flows\n\n| Name |\n|---|\n",
			}
			for _, d := range process.DataFlows {
				page.MardownContent += fmt.Sprintf("| %s |\n", d.Name)
			}
			chapter.Pages = append(chapter.Pages, page)
		}

		if len(process.DiagramProcesss) > 0 {
			for _, diagram := range process.DiagramProcesss {
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
	}

	content.LogoSVGFile = stager.GetRootLibrary().LogoSVGFile

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
}
