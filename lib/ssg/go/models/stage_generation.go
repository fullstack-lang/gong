package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (stage *Stage) Generation() {

	contents := GetGongstructInstancesSet[Content](stage)

	if len(*contents) != 1 {
		log.Println("Generation requires exactly one Content instance.") //
		return
	}

	var content *Content
	for k := range *contents { //
		content = k
		break // Assuming only one element, break after finding it
	}

	if content == nil {
		log.Println("No Content instance found.")
		return
	}

	contentPath := content.ContentPath //
	if contentPath == "" {
		log.Println("ContentPath is empty.") //
		return
	}

	// 1. Create directory
	err := os.MkdirAll(contentPath, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating directory '%s': %v\n", contentPath, err)
		return
	}
	log.Printf("Directory created or already exists: %s\n", contentPath)

	// 2. Define file path
	filePath := filepath.Join(contentPath, "_index.md")

	// 3. Define file content using content Name and Text fields
	// references the Content struct definition with Name and Text fields.

	fileContent := fmt.Sprintf("--\ntitle: \"%s\"\n--\n%s", content.Name, content.Text) //

	// 4. Write content to file
	err = os.WriteFile(filePath, []byte(fileContent), 0644) // Use 0644 for standard file permissions
	if err != nil {
		log.Printf("Error writing file '%s': %v\n", filePath, err)
		return
	}
	log.Printf("File created successfully: %s\n", filePath)

}
