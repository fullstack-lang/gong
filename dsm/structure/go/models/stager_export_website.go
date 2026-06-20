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
		Name:           "Root to system the website",
		ContentPath:    "/tmp/system",
		MardownContent: "## System website",
	}

	content.LogoSVGFile = stager.GetRootLibrary().LogoSVGFile

	refChapter := &ssg.Chapter{
		Name:           "References",
		MardownContent: "## References\n",
	}
	content.Chapters = append(content.Chapters, refChapter)

	appendWebExportableChapter(stager, refChapter, "Systemes", GetGongstrucsSorted[*System](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Parts", GetGongstrucsSorted[*Part](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Ports", GetGongstrucsSorted[*Port](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Control Flows", GetGongstrucsSorted[*ControlFlow](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Data Flows", GetGongstrucsSorted[*DataFlow](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Datas", GetGongstrucsSorted[*Data](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Resources", GetGongstrucsSorted[*Resource](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Notes", GetGongstrucsSorted[*Note](stager.stage))
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

	if len(system.DiagramStructures) > 0 {
		for _, diagram := range system.DiagramStructures {
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

	if len(system.Parts) > 0 {
		page := &ssg.Section{
			Name:           "Parts",
			MardownContent: "### Parts\n\n| Name |\n|---|\n",
		}
		for _, p := range system.Parts {
			page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
		}
		systemPage.Sections = append(systemPage.Sections, page)
	}

	if len(system.ExternalParts) > 0 {
		page := &ssg.Section{
			Name:           "External Parts",
			MardownContent: "### External Parts\n\n| Name |\n|---|\n",
		}
		for _, p := range system.ExternalParts {
			page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
		}
		systemPage.Sections = append(systemPage.Sections, page)
	}

	var allPorts []*Port
	var allControlFlows []*ControlFlow
	for _, p := range system.Parts {
		allPorts = append(allPorts, p.Ports...)
		allControlFlows = append(allControlFlows, p.ControlFlows...)
	}

	if len(allPorts) > 0 {
		page := &ssg.Section{
			Name:           "Ports",
			MardownContent: "### Ports\n\n| Name | Part |\n|---|---|\n",
		}
		for _, p := range system.Parts {
			for _, t := range p.Ports {
				page.MardownContent += fmt.Sprintf("| %s | %s |\n", t.Name, p.Name)
			}
		}
		systemPage.Sections = append(systemPage.Sections, page)
	}

	if len(allControlFlows) > 0 {
		page := &ssg.Section{
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
		systemPage.Sections = append(systemPage.Sections, page)
	}

	if len(system.DataFlows) > 0 {
		page := &ssg.Section{
			Name:           "Data Flows",
			MardownContent: "### Data Flows\n\n| Name |\n|---|\n",
		}
		for _, d := range system.DataFlows {
			page.MardownContent += fmt.Sprintf("| %s |\n", d.Name)
		}
		systemPage.Sections = append(systemPage.Sections, page)
	}

	return systemPage
}

func (inst *Part) GetDescription() string { return inst.Description }
func (inst *Part) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Part) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Port) GetDescription() string { return inst.Description }
func (inst *Port) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Port) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *ControlFlow) GetDescription() string { return inst.Description }
func (inst *ControlFlow) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *ControlFlow) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *DataFlow) GetDescription() string { return inst.Description }
func (inst *DataFlow) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *DataFlow) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Data) GetDescription() string { return inst.Description }
func (inst *Data) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Data) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Resource) GetDescription() string { return inst.Description }
func (inst *Resource) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Resource) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Note) GetDescription() string { return inst.Description }
func (inst *Note) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Note) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Library) GetDescription() string { return inst.Description }
func (inst *Library) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Library) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}
