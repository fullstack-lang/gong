package models

import (
	"fmt"
	"log"
	"strings"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

type WebExportable interface {
	GetName() string
	GetDescription() string
	GetReferencePath() string
	GeneratePage(stager *Stager) *ssg.Page
}

func (stager *Stager) exportWebsite() {
	stager.ssgStage.Reset()

	content := ssg.Content{
		Name:           "Root to floss website",
		ContentPath:    "/tmp/floss",
		MardownContent: "## Floss website",
	}

	content.LogoSVGFile = stager.GetRootLibrary().LogoSVGFile

	refChapter := &ssg.Chapter{
		Name:           "References",
		MardownContent: "## References\n",
	}
	content.Chapters = append(content.Chapters, refChapter)

	appendWebExportableChapter(stager, refChapter, "Systemes", GetGongstrucsSorted[*System](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Libraries", GetGongstrucsSorted[*Library](stager.stage))

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

func appendWebExportableChapter[T WebExportable](stager *Stager, refChapter *ssg.Chapter, title string, instances []T) {
	if len(instances) > 0 {
		sub := &ssg.Chapter{Name: title, MardownContent: "### " + title + "\n\n| Name | Description |\n|---|---|\n"}
		refChapter.SubChapters = append(refChapter.SubChapters, sub)
		for _, inst := range instances {
			sub.MardownContent += fmt.Sprintf("| [%s](%s/index.html) | %s |\n", inst.GetName(), inst.GetReferencePath(), strings.ReplaceAll(inst.GetDescription(), "\n", "<br>"))
			sub.Pages = append(sub.Pages, inst.GeneratePage(stager))
		}
	}
}

func (system *System) GetDescription() string { return system.Description }
func (system *System) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(system.Name, " "), " ", "%20")
}

func (system *System) GeneratePage(stager *Stager) *ssg.Page {
	systemPage := &ssg.Page{Name: system.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", system.Name, system.Description)}

	if len(system.DiagramFlosses) > 0 {
		for _, diagram := range system.DiagramFlosses {
			svgObject := stager.generateSvgObject(diagram)
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
			systemPage.Sections = append(systemPage.Sections, section)
		}
	}

	return systemPage
}

func (library *Library) GetDescription() string { return library.Description }
func (library *Library) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(library.Name, " "), " ", "%20")
}

func (library *Library) GeneratePage(stager *Stager) *ssg.Page {
	libraryPage := &ssg.Page{Name: library.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", library.Name, library.Description)}
	return libraryPage
}
