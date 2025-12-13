package models

import (
	"log"
	"os"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (*Stager) prepareStaticDic(pathToGeneratedSVG string) (error, bool) {
	ssg.CopyDirectory("../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/static",
		"../../diagrams/static",
	)

	// --- Start: Remove existing pathToGeneratedSVG directory ---
	// log.Printf("Attempting to remove existing directory: %s", pathToGeneratedSVG)
	err := os.RemoveAll(pathToGeneratedSVG)
	if err != nil {
		// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
		// or if the error was related to something already gone.
		// If MkdirAll fails later, that error will be caught.
		log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", pathToGeneratedSVG, err)
	} else {
		// log.Printf("Successfully removed existing directory: %s", pathToGeneratedSVG)
	}
	// --- End: Remove existing pathToGeneratedSVG directory ---

	// --- Start: Generate root _index.md for the Content ---
	// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
	err = os.MkdirAll(pathToGeneratedSVG, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating root content directory '%s': %v\n", pathToGeneratedSVG, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root directory cannot be created.
		return nil, true
	}
	// log.Printf("Root content directory created or already exists: %s\n", pathToGeneratedSVG)

	/* copy necessary images for the geenration */

	return err, false
}
