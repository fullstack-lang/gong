package models

import (
	"fmt"
	"io/fs"
	"log"
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
func (stage *Stage) Generation() {

	contents := GetGongstructInstancesSet[Content](stage)

	if len(*contents) != 1 {
		log.Println("Generation requires exactly one Content instance.")
		return
	}

	var content *Content
	for k := range *contents {
		content = k
		break // Assuming only one element, break after finding it
	}

	if content == nil {
		log.Println("No Content instance found.")
		return
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
	// Iterate through each chapter associated with the Content instance
	for idx, chapter := range content.Chapters {
		// log.Printf("Processing chapter: %s", chapter.Name)

		// 1. Create subdirectory for the chapter
		// Use chapter.Name for the subdirectory name. Consider sanitizing the name
		// if it might contain characters invalid for directory names.
		// For simplicity, assuming chapter.Name is a valid directory name here.
		chapterDirName := SanitizeFileName(chapter.Name, " ") // <--- MODIFIED: Sanitize the chapter name

		// 2. Define the _index.md file path within the chapter directory
		chapterIndexFilePath := path.Join(chapterDirName, "_index.md")

		// 3. Define file content using chapter fields
		//    Using chapter.Description for the body content as per the user's example.
		//    Converting weight (float64) to int for the front matter example.
		chapterFileContent := fmt.Sprintf(`---
title: "%s"
weight: %d
---
%s`,
			chapter.Name,
			idx,                    // Convert float64 weight to int
			chapter.MardownContent) // Use Description as body content based on example

		// 4. Write content to the _index.md file in MapFS
		memoryFS[chapterIndexFilePath] = &fstest.MapFile{Data: []byte(chapterFileContent)}

		for idx, page := range chapter.Pages {
			sanitizedPageName := SanitizeFileName(page.GetName(), " ") // <--- ADDED: Sanitize the page name
			pageIndexFilePath := path.Join(chapterDirName, sanitizedPageName+".md")

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
				idx,               // Convert float64 weight to int
				pageBody.String()) // Use Description as body content based on example

			memoryFS[pageIndexFilePath] = &fstest.MapFile{Data: []byte(pageFileContent)}
		}

	}
	// --- End: Generate subdirectories and _index.md for each Chapter ---

	// log.Println("Generation process finished.")

	// --- Build Steps ---
	stage.markdown2ssg(content, memoryFS)
}

// markdown2ssg handles the transformation of the generated markdown files
// into a static HTML site. It cleans the output directory, loads the layout
// templates, parses the markdown content, builds the site navigation structure,
// renders the final HTML pages, and copies any static assets to the output directory.
func (*Stage) markdown2ssg(content *Content, memoryFS fs.FS) {
	if err := gen.CleanOutputDir(content.OutputPath); err != nil {
		log.Fatalf("Error cleaning output directory '%s': %v", content.OutputPath, err)
	}
	// log.Printf("Cleaned output directory '%s'.\n", content.OutputPath)

	templates, err := gen.LoadTemplates()
	if err != nil {
		log.Fatalf("Error loading embedded templates: %v", err)
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
		log.Fatalf("Error parsing content from memory: %v", err)
	}
	// log.Printf("Parsed %d content files.\n", len(site.Pages))

	gen.BuildSiteStructure(site) // <-- Call the updated function
	// log.Println("Built site structure (sections and pages).")

	// Pass output dir and build target to renderPages
	err = gen.RenderPages(site, content.OutputPath, content.Target.ToString())
	if err != nil {
		log.Fatalf("Error rendering HTML pages: %v", err)
	}
	// log.Println("Rendered HTML pages.")

	// Pass output dir to copyStaticFiles
	if err := gen.CopyStaticFiles(content.StaticPath, content.OutputPath); err != nil {
		log.Fatalf("Error copying static files from '%s' to '%s': %v", content.StaticPath, content.OutputPath, err)
	}
	// log.Println("Copied static files.")

	// log.Println("Build finished successfully!")
}
