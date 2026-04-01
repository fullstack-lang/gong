package models

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func extractGongSsgEmbeddedFiles(
	source embed.FS,
	pathToExtractedFiles string) {

	// 1. Extract the embedded 'StaticDir' content to a temporary directory.

	// In your actual project, you would use your imported ssg_defaults.StaticDir
	// For example: tempSourcePath, err := extractEmbedFSToTemp(ssg_defaults.StaticDir, ".")
	tempSourcePath, err := extractEmbedFSToTemp(source, ".")
	if err != nil {
		log.Fatalf("Error extracting embedded directory: %v", err)
	}
	// Defer cleanup of the temporary directory
	defer os.RemoveAll(tempSourcePath)

	log.Printf("Embedded content extracted to temporary path: %s", tempSourcePath)

	// 2. Define your destination path
	destinationPath := pathToExtractedFiles

	err = os.RemoveAll(pathToExtractedFiles)
	if err != nil {
		log.Println("Extracting ssg directory, removing target dir before copy, error:", err.Error())
	} else {
		log.Println("Extracting ssg directory, removing target dir before copy:", pathToExtractedFiles)
	}

	// 3. Use the CopyDirectory function with the temporary path as source
	log.Printf("Copying from '%s' to '%s'", tempSourcePath, destinationPath)
	err = models.CopyDirectory(tempSourcePath, destinationPath) // Use your actual package
	if err != nil {
		log.Fatalf("Error using CopyDirectory: %v", err)
	}

	log.Printf("Successfully copied embedded directory from temp location to %s", destinationPath)

}

// extractEmbedFSToTemp is a helper function to extract an embed.FS to a temporary directory.
// rootInEmbed: The root directory within the embed.FS to start extracting from (e.g., "." for the whole FS).
func extractEmbedFSToTemp(embedFS embed.FS, rootInEmbed string) (string, error) {
	tempDir, err := os.MkdirTemp("", "embed_extract_*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	// Walk the embedded directory
	err = fs.WalkDir(embedFS, rootInEmbed, func(path string, d fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			return fmt.Errorf("error during walk at %s: %w", path, errWalk)
		}

		// Determine the corresponding path in the temporary directory
		relPath, err := filepath.Rel(rootInEmbed, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}
		destPath := filepath.Join(tempDir, relPath)

		if d.IsDir() {
			// Create the directory in the temporary location
			if err := os.MkdirAll(destPath, 0755); err != nil { // Standard directory permissions
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
		} else {
			// Read the embedded file content
			data, err := embedFS.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read embedded file %s: %w", path, err)
			}
			// Write the file to the temporary location
			if err := os.WriteFile(destPath, data, 0644); err != nil { // Standard file permissions
				return fmt.Errorf("failed to write file %s: %w", destPath, err)
			}
		}
		return nil
	})

	if err != nil {
		os.RemoveAll(tempDir) // Clean up the temporary directory on error
		return "", err
	}

	return tempDir, nil
}
