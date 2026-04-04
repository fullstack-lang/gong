package models

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing/fstest"

	"github.com/fullstack-lang/gong/lib/ssg/go/gen"
	// Import strconv for float to string conversion if needed directly
)

// Generation orchestrates the markdown and static site generation process.
// It extracts the single Content instance from the stage, sets up the
// necessary directory structure under ContentPath, and generates the
// markdown files (_index.md for the root and chapters, and individual .md
// files for pages). It also embeds sections (text, images, downloadable files)
// into the page markdown. Finally, it triggers the HTML rendering process.
func (stage *Stage) Generation(inMemory bool) (base64Zip string, err error) {

	contents := GetGongstructInstancesSet[Content](stage)

	if len(*contents) != 1 {
		err = fmt.Errorf("generation requires exactly one Content instance")
		log.Println(err)
		return "", err
	}

	var content *Content
	for k := range *contents {
		content = k
		break // Assuming only one element, break after finding it
	}

	if content == nil {
		err = fmt.Errorf("no Content instance found")
		log.Println(err)
		return "", err
	}

	memoryFS := make(fstest.MapFS)

	// --- Start: Generate root _index.md for the Content ---
	// 1. Define the root _index.md file path
	rootIndexFilePath := "_index.md"

	// 2. Define file content using the Content struct's Name and Text fields
	rootFileContent := fmt.Sprintf("---\ntitle: \"%s\"\nversioninfo: \"%s\"\n---\n%s", content.Name, content.VersionInfo, content.MardownContent)

	// 3. Write content to the root _index.md file in MapFS
	memoryFS[rootIndexFilePath] = &fstest.MapFile{Data: []byte(rootFileContent)}
	// --- End: Generate root _index.md for the Content ---

	// --- Start: Generate subdirectories and _index.md for each Chapter ---
	var processChapter func(chapter *Chapter, parentDir string, weight int)
	processChapter = func(chapter *Chapter, parentDir string, weight int) {
		// 1. Create subdirectory for the chapter
		// Use chapter.Name for the subdirectory name. Consider sanitizing the name
		// if it might contain characters invalid for directory names.
		chapterDirName := SanitizeFileName(chapter.Name, " ") // <--- MODIFIED: Sanitize the chapter name
		chapterDirPath := path.Join(parentDir, chapterDirName)

		// 2. Define the _index.md file path within the chapter directory
		chapterIndexFilePath := path.Join(chapterDirPath, "_index.md")

		// 3. Define file content using chapter fields
		//    Using chapter.Description for the body content as per the user's example.
		//    Converting weight (float64) to int for the front matter example.
		chapterFileContent := fmt.Sprintf(`---
title: "%s"
weight: %d
---
%s`,
			chapter.Name,
			weight,                 // Convert float64 weight to int
			chapter.MardownContent) // Use Description as body content based on example

		// 4. Write content to the _index.md file in MapFS
		memoryFS[chapterIndexFilePath] = &fstest.MapFile{Data: []byte(chapterFileContent)}

		for idx, page := range chapter.Pages {
			sanitizedPageName := SanitizeFileName(page.GetName(), " ") // <--- ADDED: Sanitize the page name
			pageIndexFilePath := path.Join(chapterDirPath, sanitizedPageName+".md")

			var pageBody strings.Builder
			pageBody.WriteString(page.MardownContent)

			for _, section := range page.Sections {
				pageBody.WriteString("\n\n")
				if section.Name != "" {
					pageBody.WriteString(fmt.Sprintf("## %s\n\n", section.Name))
				}
				if section.MardownContent != "" {
					pageBody.WriteString(section.MardownContent)
					pageBody.WriteString("\n\n")
				}
				if section.IsImage {
					if section.SvgImage != nil {
						// Clean the SVG content to prevent the markdown parser from treating
						// indented lines as code blocks or empty lines as paragraph breaks.
						pageBody.WriteString("<div class=\"svg-container\">\n")
						for _, line := range strings.Split(section.SvgImage.Content, "\n") {
							if trimmed := strings.TrimSpace(line); trimmed != "" {
								pageBody.WriteString(trimmed + "\n")
							}
						}
						pageBody.WriteString("</div>\n\n")
					}
					if section.PngImage != nil {
						pageBody.WriteString(fmt.Sprintf("![%s](data:image/png;base64,%s)\n\n", section.PngImage.Name, section.PngImage.Base64Content))
					}
					if section.JpgImage != nil {
						pageBody.WriteString(fmt.Sprintf("![%s](data:image/jpeg;base64,%s)\n\n", section.JpgImage.Name, section.JpgImage.Base64Content))
					}
				}
				if section.IsDownloadableFile && section.DownloadableFile != nil {
					fileName := filepath.Base(section.DownloadableFile.Name)
					pageBody.WriteString(fmt.Sprintf("<a href=\"data:application/octet-stream;base64,%s\" download=\"%s\">\n", section.DownloadableFile.Base64Content, fileName))
					pageBody.WriteString(fmt.Sprintf("<button>Download %s</button>\n", fileName))
					pageBody.WriteString("</a>\n\n")
				}
			}

			pageFileContent := fmt.Sprintf(`---
title: "%s"
weight: %d
---
%s`,
				sanitizedPageName,
				(len(chapter.SubChapters)+idx+1)*10, // Give pages a higher weight to place them after sub-chapters
				pageBody.String())                   // Use Description as body content based on example

			memoryFS[pageIndexFilePath] = &fstest.MapFile{Data: []byte(pageFileContent)}
		}

		for idx, subChapter := range chapter.SubChapters {
			processChapter(subChapter, chapterDirPath, (idx+1)*10)
		}
	}

	// Iterate through each chapter associated with the Content instance
	for idx, chapter := range content.Chapters {
		processChapter(chapter, "", (idx+1)*10)
	}
	// --- End: Generate subdirectories and _index.md for each Chapter ---

	// log.Println("Generation process finished.")

	originalOutputPath := content.OutputPath
	if inMemory {
		var tempDir string
		tempDir, err = os.MkdirTemp("", "ssg_output_*")
		if err != nil {
			return "", err
		}
		defer os.RemoveAll(tempDir)
		content.OutputPath = tempDir
	}

	// --- Build Steps ---
	err = stage.markdown2ssg(content, memoryFS)
	if err != nil {
		content.OutputPath = originalOutputPath
		return "", err
	}

	if inMemory {
		var buf bytes.Buffer
		zipWriter := zip.NewWriter(&buf)

		err = filepath.WalkDir(content.OutputPath, func(p string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if d.IsDir() {
				return nil
			}
			relPath, relErr := filepath.Rel(content.OutputPath, p)
			if relErr != nil {
				return relErr
			}
			file, openErr := os.Open(p)
			if openErr != nil {
				return openErr
			}
			defer file.Close()

			w, createErr := zipWriter.Create(filepath.ToSlash(relPath))
			if createErr != nil {
				return createErr
			}
			_, copyErr := io.Copy(w, file)
			return copyErr
		})
		if err != nil {
			content.OutputPath = originalOutputPath
			return "", err
		}

		err = zipWriter.Close()
		if err != nil {
			content.OutputPath = originalOutputPath
			return "", err
		}

		base64Zip = base64.StdEncoding.EncodeToString(buf.Bytes())
		content.OutputPath = originalOutputPath
	}

	return base64Zip, nil
}

// markdown2ssg handles the transformation of the generated markdown files
// into a static HTML site. It cleans the output directory, loads the layout
// templates, parses the markdown content, builds the site navigation structure,
// renders the final HTML pages, and copies any static assets to the output directory.
func (*Stage) markdown2ssg(content *Content, memoryFS fs.FS) error {
	if err := gen.CleanOutputDir(content.OutputPath); err != nil {
		return fmt.Errorf("error cleaning output directory '%s': %w", content.OutputPath, err)
	}
	// log.Printf("Cleaned output directory '%s'.\n", content.OutputPath)

	templates, err := gen.LoadTemplates()
	if err != nil {
		return fmt.Errorf("error loading embedded templates: %w", err)
	}
	// log.Println("Loaded HTML templates.")

	site := &gen.SiteInfo{
		Pages:                         make(map[string]*gen.Page),
		Templates:                     templates,
		IsBespokeLogoFileName:         content.IsBespokeLogoFileName,
		BespokeLogoFileName:           content.BespokeLogoFileName,
		IsBespokePageTileLogoFileName: content.IsBespokePageTileLogoFileName,
		BespokePageTileLogoFileName:   content.BespokePageTileLogoFileName,
	}
	// Pass build target and output dir to parseContent
	err = gen.ParseContent(memoryFS, site, content.Target.ToString(), content.OutputPath)
	if err != nil {
		return fmt.Errorf("error parsing content from memory: %w", err)
	}
	// log.Printf("Parsed %d content files.\n", len(site.Pages))

	gen.BuildSiteStructure(site) // <-- Call the updated function
	// log.Println("Built site structure (sections and pages).")

	// Pass output dir and build target to renderPages
	err = gen.RenderPages(site, content.OutputPath, content.Target.ToString())
	if err != nil {
		return fmt.Errorf("error rendering HTML pages: %w", err)
	}
	// log.Println("Rendered HTML pages.")

	// Pass output dir to copyStaticFiles
	if err := gen.CopyStaticFiles(content.StaticPath, content.OutputPath); err != nil {
		return fmt.Errorf("error copying static files from '%s' to '%s': %w", content.StaticPath, content.OutputPath, err)
	}
	// log.Println("Copied static files.")

	if content.LogoSVGFile != "" {
		logoFileName := "logo.svg"
		if content.IsBespokeLogoFileName && content.BespokeLogoFileName != "" {
			logoFileName = content.BespokeLogoFileName
		}
		logoPath := filepath.Join(content.OutputPath, "images", logoFileName)
		if err := os.MkdirAll(filepath.Dir(logoPath), 0755); err != nil {
			return fmt.Errorf("error creating directory for logo: %w", err)
		}
		if err := os.WriteFile(logoPath, []byte(content.LogoSVGFile), 0644); err != nil {
			return fmt.Errorf("error writing logo file: %w", err)
		}
	}

	// POST-PROCESSING: Fix nested submenu JS toggling and styling
	filepath.WalkDir(content.OutputPath, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !d.IsDir() && strings.HasSuffix(p, ".html") {
			htmlData, err := os.ReadFile(p)
			if err == nil {
				htmlStr := string(htmlData)
				// 1. Fix JS selector to target all list items, not just top-level ones
				htmlStr = strings.ReplaceAll(htmlStr, "document.querySelectorAll('.sidebar > ul > li');", "document.querySelectorAll('.sidebar li');")
				// 2. Add a class to links that have a submenu to style them as chapters
				htmlStr = strings.ReplaceAll(htmlStr, "if (chapterLink && submenu) {", "if (chapterLink && submenu) {\n                    chapterLink.classList.add('has-submenu');")

				// 3. Allow navigation by removing preventDefault so users can visit chapter pages
				htmlStr = strings.ReplaceAll(htmlStr, "event.preventDefault();", "")

				// 4. Rebuild the sidebar HTML to contain the full recursive tree
				navStartIndex := strings.Index(htmlStr, `<nav class="sidebar">`)
				if navStartIndex != -1 {
					navCloseIndex := strings.Index(htmlStr[navStartIndex:], "</nav>")
					if navCloseIndex != -1 {
						navCloseIndex += navStartIndex
						h2Index := strings.Index(htmlStr[navStartIndex:navCloseIndex], "</h2>")
						if h2Index != -1 {
							h2Index += navStartIndex

							relPath, _ := filepath.Rel(content.OutputPath, p)
							relPath = filepath.ToSlash(relPath)
							depth := strings.Count(relPath, "/")
							prefix := strings.Repeat("../", depth)

							var sb strings.Builder
							sb.WriteString("\n            <ul>\n")

							var walkChapter func(c *Chapter, currentPath string)
							walkChapter = func(c *Chapter, currentPath string) {
								chapterDir := SanitizeFileName(c.Name, " ")
								chapterPath := chapterDir
								if currentPath != "" {
									chapterPath = currentPath + "/" + chapterDir
								}

								targetRelPath := chapterPath + "/index.html"
								href := prefix + strings.ReplaceAll(targetRelPath, " ", "%20")

								activeClass := ""
								if targetRelPath == relPath {
									activeClass = ` class="active-page"`
								}

								sb.WriteString("                <li>\n")
								sb.WriteString(fmt.Sprintf(`                    <a href="%s"%s>%s</a>`+"\n", href, activeClass, c.Name))

								hasChildren := len(c.SubChapters) > 0 || len(c.Pages) > 0
								if hasChildren {
									sb.WriteString("                    <ul>\n")
									for _, sub := range c.SubChapters {
										walkChapter(sub, chapterPath)
									}
									for _, page := range c.Pages {
										pageDir := SanitizeFileName(page.GetName(), " ")
										pagePath := chapterPath + "/" + pageDir
										pageTargetRelPath := pagePath + "/index.html"
										pageHref := prefix + strings.ReplaceAll(pageTargetRelPath, " ", "%20")

										pageActiveClass := ""
										if pageTargetRelPath == relPath {
											pageActiveClass = ` class="active-page"`
										}

										sb.WriteString("                        <li>\n")
										sb.WriteString(fmt.Sprintf(`                            <a href="%s"%s>%s</a>`+"\n", pageHref, pageActiveClass, page.GetName()))
										sb.WriteString("                        </li>\n")
									}
									sb.WriteString("                    </ul>\n")
								}
								sb.WriteString("                </li>\n")
							}

							for _, c := range content.Chapters {
								walkChapter(c, "")
							}

							sb.WriteString("            </ul>\n        ")

							htmlStr = htmlStr[:h2Index+5] + sb.String() + htmlStr[navCloseIndex:]
						}
					}
				}

				os.WriteFile(p, []byte(htmlStr), 0644)
			}
		}
		return nil
	})

	// Append CSS for the .has-submenu class so sub-chapters look like main chapters
	cssPath := filepath.Join(content.OutputPath, "css", "style.css")
	cssFile, err := os.OpenFile(cssPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		cssFile.WriteString("\n/* Sub-chapter styling */\n.has-submenu { font-weight: bold !important; color: #007bff !important; }\n")
		cssFile.Close()
	}

	// log.Println("Build finished successfully!")
	return nil
}
