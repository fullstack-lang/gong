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
		OnUpdate: func() {
			stager.stage.Reset()
			stager.stage.Commit()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export file",
		Icon:  string(buttons.BUTTON_fact_check),
		Label: "Export file",
		OnUpdate: func() {
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
			stage.Marshall(tempFile, "github.com/fullstack-lang/gong/dsm/project/go/models", "main")

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
		OnUpdate: func() {
			log.Println("Stop")
			os.Exit(0)
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export website",
		Icon:  string(buttons.BUTTON_web),
		Label: "Export website",
		OnUpdate: func() {
			stager.ssgStage.Reset()

			content := ssg.Content{
				Name:           "Root to project the website",
				ContentPath:    "/tmp/project",
				MardownContent: "## Project website",
				OutputPath:     "./generated static web site",
			}

			for _, lib := range GetGongstrucsSorted[*Library](stager.stage) {
				chapter := &ssg.Chapter{
					Name:           lib.Name,
					MardownContent: "### " + lib.Name,
				}
				content.Chapters = append(content.Chapters, chapter)

				if len(lib.RootProducts) > 0 {
					page := &ssg.Page{
						Name:           "Products",
						MardownContent: "### Products\n",
					}
					for _, p := range lib.RootProducts {
						page.MardownContent += productToMD(p, 0)
					}

					allProducts := flattenProducts(lib.RootProducts)
					page.MardownContent += "\n\n#### Products Table\n\n"
					page.MardownContent += "| Prefix | Name | Description |\n"
					page.MardownContent += "|---|---|---|\n"
					for _, p := range allProducts {
						desc := strings.ReplaceAll(p.Description, "\n", " ")
						desc = strings.ReplaceAll(desc, "|", "\\|")
						page.MardownContent += fmt.Sprintf("| %s | %s | %s |\n", p.GetComputedPrefix(), p.Name, desc)
					}
					chapter.Pages = append(chapter.Pages, page)
				}

				if len(lib.RootTasks) > 0 {
					page := &ssg.Page{
						Name:           "Tasks",
						MardownContent: "### Tasks\n",
					}
					for _, t := range lib.RootTasks {
						page.MardownContent += taskToMD(t, 0)
					}

					allTasks := flattenTasks(lib.RootTasks)
					page.MardownContent += "\n\n#### Tasks Table\n\n"
					page.MardownContent += "| Prefix | Name | Description | Start | End | Completion | Inputs | Outputs |\n"
					page.MardownContent += "|---|---|---|---|---|---|---|---|\n"
					for _, t := range allTasks {
						desc := strings.ReplaceAll(t.Description, "\n", " ")
						desc = strings.ReplaceAll(desc, "|", "\\|")
						start := ""
						if !t.Start.IsZero() {
							start = t.Start.Format("2006-01-02")
						}
						end := ""
						if !t.End.IsZero() {
							end = t.End.Format("2006-01-02")
						}
						comp := ""
						if t.IsWithCompletion {
							comp = string(t.Completion)
						}
						var inputs, outputs []string
						for _, in := range t.Inputs {
							inputs = append(inputs, in.Name)
						}
						for _, out := range t.Outputs {
							outputs = append(outputs, out.Name)
						}

						page.MardownContent += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n", t.GetComputedPrefix(), t.Name, desc, start, end, comp, strings.Join(inputs, ", "), strings.Join(outputs, ", "))
					}
					chapter.Pages = append(chapter.Pages, page)
				}

				if len(lib.RootResources) > 0 {
					page := &ssg.Page{
						Name:           "Resources",
						MardownContent: "### Resources\n",
					}
					for _, r := range lib.RootResources {
						page.MardownContent += resourceToMD(r, 0)
					}
					chapter.Pages = append(chapter.Pages, page)
				}

				if len(lib.Notes) > 0 {
					page := &ssg.Page{
						Name:           "Notes",
						MardownContent: "### Notes\n",
					}
					for _, n := range lib.Notes {
						page.MardownContent += noteToMD(n, 0)
					}
					chapter.Pages = append(chapter.Pages, page)
				}

				if len(lib.Diagrams) > 0 {
					page := &ssg.Page{
						Name:           "Diagrams",
						MardownContent: "### Diagrams\n",
					}
					for _, diagram := range lib.Diagrams {
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

				tempFile, err := os.CreateTemp("", "export-*.xlsx")
				if err != nil {
					log.Println("Error creating temp file for Excel export:", err)
				} else {
					defer os.Remove(tempFile.Name())
					SerializeStage2(stager.stage, tempFile.Name(), false)

					content, err := os.ReadFile(tempFile.Name())
					if err != nil {
						log.Println("Error reading temp Excel file:", err)
					} else {
						encodedXL := base64.StdEncoding.EncodeToString(content)
						dataURI := "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64," + encodedXL
						page.MardownContent += fmt.Sprintf("\n[Download %s.xlsx](%s)\n", lib.Name, dataURI)

						section := &ssg.Section{
							Name: "Embedded XL file",
						}
						page.Sections = append(page.Sections, section)
					}
				}

				chapter.Pages = append(chapter.Pages, page)
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
		},
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}

func productToMD(p *Product, depth int) string {
	indent := strings.Repeat("  ", depth)
	md := fmt.Sprintf("%s- **%s**\n", indent, p.Name)
	if p.Description != "" {
		md += fmt.Sprintf("%s  - Description: %s\n", indent, p.Description)
	}
	if len(p.SubProducts) > 0 {
		md += fmt.Sprintf("%s  - SubProducts:\n", indent)
		for _, sub := range p.SubProducts {
			md += productToMD(sub, depth+2)
		}
	}
	return md
}

func flattenProducts(products []*Product) []*Product {
	var flat []*Product
	for _, p := range products {
		flat = append(flat, p)
		flat = append(flat, flattenProducts(p.SubProducts)...)
	}
	return flat
}

func flattenTasks(tasks []*Task) []*Task {
	var flat []*Task
	for _, t := range tasks {
		flat = append(flat, t)
		flat = append(flat, flattenTasks(t.SubTasks)...)
	}
	return flat
}

func taskToMD(t *Task, depth int) string {
	indent := strings.Repeat("  ", depth)
	md := fmt.Sprintf("%s- **%s**\n", indent, t.Name)
	if t.Description != "" {
		md += fmt.Sprintf("%s  - Description: %s\n", indent, t.Description)
	}
	if !t.Start.IsZero() || !t.End.IsZero() {
		md += fmt.Sprintf("%s  - Start: %s, End: %s\n", indent, t.Start.Format("2006-01-02"), t.End.Format("2006-01-02"))
	}
	if t.IsWithCompletion {
		md += fmt.Sprintf("%s  - Completion: %s\n", indent, string(t.Completion))
	}
	if len(t.Inputs) > 0 {
		var inputs []string
		for _, in := range t.Inputs {
			inputs = append(inputs, in.Name)
		}
		md += fmt.Sprintf("%s  - Inputs: %s\n", indent, strings.Join(inputs, ", "))
	}
	if len(t.Outputs) > 0 {
		var outputs []string
		for _, out := range t.Outputs {
			outputs = append(outputs, out.Name)
		}
		md += fmt.Sprintf("%s  - Outputs: %s\n", indent, strings.Join(outputs, ", "))
	}
	if len(t.SubTasks) > 0 {
		md += fmt.Sprintf("%s  - SubTasks:\n", indent)
		for _, sub := range t.SubTasks {
			md += taskToMD(sub, depth+2)
		}
	}
	return md
}

func resourceToMD(r *Resource, depth int) string {
	indent := strings.Repeat("  ", depth)
	md := fmt.Sprintf("%s- **%s**\n", indent, r.Name)
	if r.Description != "" {
		md += fmt.Sprintf("%s  - Description: %s\n", indent, r.Description)
	}
	if len(r.Tasks) > 0 {
		var tasks []string
		for _, t := range r.Tasks {
			tasks = append(tasks, t.Name)
		}
		md += fmt.Sprintf("%s  - Tasks: %s\n", indent, strings.Join(tasks, ", "))
	}
	if len(r.SubResources) > 0 {
		md += fmt.Sprintf("%s  - SubResources:\n", indent)
		for _, sub := range r.SubResources {
			md += resourceToMD(sub, depth+2)
		}
	}
	return md
}

func noteToMD(n *Note, depth int) string {
	indent := strings.Repeat("  ", depth)
	md := fmt.Sprintf("%s- **%s**\n", indent, n.Name)
	if len(n.Products) > 0 {
		var items []string
		for _, p := range n.Products {
			items = append(items, p.Name)
		}
		md += fmt.Sprintf("%s  - Products: %s\n", indent, strings.Join(items, ", "))
	}
	if len(n.Tasks) > 0 {
		var items []string
		for _, t := range n.Tasks {
			items = append(items, t.Name)
		}
		md += fmt.Sprintf("%s  - Tasks: %s\n", indent, strings.Join(items, ", "))
	}
	if len(n.Resources) > 0 {
		var items []string
		for _, r := range n.Resources {
			items = append(items, r.Name)
		}
		md += fmt.Sprintf("%s  - Resources: %s\n", indent, strings.Join(items, ", "))
	}
	return md
}
