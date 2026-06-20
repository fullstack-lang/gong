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
		Name:           "Root to process the website",
		ContentPath:    "/tmp/process",
		MardownContent: "## Process website",
	}

	content.LogoSVGFile = stager.GetRootLibrary().LogoSVGFile

	refChapter := &ssg.Chapter{
		Name:           "References",
		MardownContent: "## References\n",
	}
	content.Chapters = append(content.Chapters, refChapter)

	appendWebExportableChapter(stager, refChapter, "Processes", GetGongstrucsSorted[*Process](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Participants", GetGongstrucsSorted[*Participant](stager.stage))
	appendWebExportableChapter(stager, refChapter, "Tasks", GetGongstrucsSorted[*Task](stager.stage))
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

func (process *Process) GetDescription() string { return process.Description }
func (process *Process) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(process.Name, " "), " ", "%20")
}

func (process *Process) GeneratePage(stager *Stager) *ssg.Page {
	processPage := &ssg.Page{Name: process.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", process.Name, process.Description)}

	if len(process.DiagramProcesss) > 0 {
		for _, diagram := range process.DiagramProcesss {
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
			processPage.Sections = append(processPage.Sections, section)
		}
	}

	if len(process.Participants) > 0 {
		page := &ssg.Section{
			Name:           "Participants",
			MardownContent: "### Participants\n\n| Name |\n|---|\n",
		}
		for _, p := range process.Participants {
			page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
		}
		processPage.Sections = append(processPage.Sections, page)
	}

	if len(process.ExternalParticipants) > 0 {
		page := &ssg.Section{
			Name:           "External Participants",
			MardownContent: "### External Participants\n\n| Name |\n|---|\n",
		}
		for _, p := range process.ExternalParticipants {
			page.MardownContent += fmt.Sprintf("| %s |\n", p.Name)
		}
		processPage.Sections = append(processPage.Sections, page)
	}

	var allTasks []*Task
	var allControlFlows []*ControlFlow
	for _, p := range process.Participants {
		allTasks = append(allTasks, p.Tasks...)
		allControlFlows = append(allControlFlows, p.ControlFlows...)
	}

	if len(allTasks) > 0 {
		page := &ssg.Section{
			Name:           "Tasks",
			MardownContent: "### Tasks\n\n| Name | Participant |\n|---|---|\n",
		}
		for _, p := range process.Participants {
			for _, t := range p.Tasks {
				page.MardownContent += fmt.Sprintf("| %s | %s |\n", t.Name, p.Name)
			}
		}
		processPage.Sections = append(processPage.Sections, page)
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
		processPage.Sections = append(processPage.Sections, page)
	}

	if len(process.DataFlows) > 0 {
		page := &ssg.Section{
			Name:           "Data Flows",
			MardownContent: "### Data Flows\n\n| Name |\n|---|\n",
		}
		for _, d := range process.DataFlows {
			page.MardownContent += fmt.Sprintf("| %s |\n", d.Name)
		}
		processPage.Sections = append(processPage.Sections, page)
	}

	return processPage
}

func (inst *Participant) GetDescription() string { return inst.Description }
func (inst *Participant) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Participant) GeneratePage(stager *Stager) *ssg.Page {
	return &ssg.Page{Name: inst.Name, MardownContent: fmt.Sprintf("#### %s\n\n%s", inst.Name, inst.Description)}
}

func (inst *Task) GetDescription() string { return inst.Description }
func (inst *Task) GetReferencePath() string {
	return strings.ReplaceAll(ssg.SanitizeFileName(inst.Name, " "), " ", "%20")
}

func (inst *Task) GeneratePage(stager *Stager) *ssg.Page {
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
