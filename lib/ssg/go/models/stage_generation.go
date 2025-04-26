package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/lib/ssg/go/gen"
	// Import strconv for float to string conversion if needed directly
)

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

	contentPath := content.ContentPath
	if contentPath == "" {
		log.Println("ContentPath is empty.")
		return
	}

	log.Printf("Starting generation process for content '%s' in path '%s'", content.Name, contentPath)

	// --- Start: Remove existing contentPath directory ---
	log.Printf("Attempting to remove existing directory: %s", contentPath)
	err := os.RemoveAll(contentPath)
	if err != nil {
		// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
		// or if the error was related to something already gone.
		// If MkdirAll fails later, that error will be caught.
		log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", contentPath, err)
	} else {
		log.Printf("Successfully removed existing directory: %s", contentPath)
	}
	// --- End: Remove existing contentPath directory ---

	// --- Start: Generate root _index.md for the Content ---
	// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
	err = os.MkdirAll(contentPath, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating root content directory '%s': %v\n", contentPath, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root directory cannot be created.
		return
	}
	log.Printf("Root content directory created or already exists: %s\n", contentPath)

	// 2. Define the root _index.md file path
	rootIndexFilePath := filepath.Join(contentPath, "_index.md")

	// 3. Define file content using the Content struct's Name and Text fields
	rootFileContent := fmt.Sprintf("---\ntitle: \"%s\"\n---\n%s", content.Name, content.MardownContent)

	// 4. Write content to the root _index.md file
	err = os.WriteFile(rootIndexFilePath, []byte(rootFileContent), 0644) // Use 0644 for standard file permissions
	if err != nil {
		log.Printf("Error writing root _index.md file '%s': %v\n", rootIndexFilePath, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root index file cannot be written.
		return
	}
	log.Printf("Root _index.md file created successfully: %s\n", rootIndexFilePath)
	// --- End: Generate root _index.md for the Content ---

	// --- Start: Generate subdirectories and _index.md for each Chapter ---
	// Iterate through each chapter associated with the Content instance
	for idx, chapter := range content.Chapters {
		log.Printf("Processing chapter: %s", chapter.Name)

		// 1. Create subdirectory for the chapter
		// Use chapter.Name for the subdirectory name. Consider sanitizing the name
		// if it might contain characters invalid for directory names.
		// For simplicity, assuming chapter.Name is a valid directory name here.
		chapterDirName := chapter.Name // Might need sanitization in a real application
		chapterDirPath := filepath.Join(contentPath, chapterDirName)

		err := os.MkdirAll(chapterDirPath, 0755) // Use 0755 for standard directory permissions
		if err != nil {
			log.Printf("Error creating directory '%s' for chapter '%s': %v", chapterDirPath, chapter.Name, err)
			continue // Skip this chapter if directory creation fails
		}
		log.Printf("Directory created or already exists: %s", chapterDirPath)

		// 2. Define the _index.md file path within the chapter directory
		chapterIndexFilePath := filepath.Join(chapterDirPath, "_index.md")

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

		// 4. Write content to the _index.md file
		err = os.WriteFile(chapterIndexFilePath, []byte(chapterFileContent), 0644) // Use 0644 for standard file permissions
		if err != nil {
			log.Printf("Error writing file '%s' for chapter '%s': %v", chapterIndexFilePath, chapter.Name, err)
			continue // Skip this chapter if file writing fails
		}
		log.Printf("File created successfully: %s", chapterIndexFilePath)

		for idx, page := range chapter.Pages {
			pageIndexFilePath := filepath.Join(chapterDirPath, page.GetName()+".md")

			pageFileContent := fmt.Sprintf(`---
title: "%s"
weight: %d
---
%s`,
				page.Name,
				idx,                 // Convert float64 weight to int
				page.MardownContent) // Use Description as body content based on example

			err = os.WriteFile(pageIndexFilePath, []byte(pageFileContent), 0644) // Use 0644 for standard file permissions
			if err != nil {
				log.Printf("Error writing file '%s' for chapter '%s': %v", pageIndexFilePath, page.Name, err)
				continue // Skip this chapter if file writing fails
			}
			log.Printf("File created successfully: %s", pageIndexFilePath)
		}

	}
	// --- End: Generate subdirectories and _index.md for each Chapter ---

	log.Println("Generation process finished.")

	// --- Build Steps ---
	stage.markdown2ssg(content)
}

func (*Stage) markdown2ssg(content *Content) {
	if err := gen.CleanOutputDir(content.OutputPath); err != nil {
		log.Fatalf("Error cleaning output directory '%s': %v", content.OutputPath, err)
	}
	// log.Printf("Cleaned output directory '%s'.\n", content.OutputPath)

	templates, err := gen.LoadTemplates(content.LayoutPath)
	if err != nil {
		log.Fatalf("Error loading templates from '%s': %v", content.LayoutPath, err)
	}
	// log.Println("Loaded HTML templates.")

	site := &gen.SiteInfo{
		Pages:     make(map[string]*gen.Page),
		Templates: templates,
	}
	// Pass build target and output dir to parseContent
	err = gen.ParseContent(content.ContentPath, site, content.Target.ToString(), content.OutputPath)
	if err != nil {
		log.Fatalf("Error parsing content from '%s': %v", content.ContentPath, err)
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
