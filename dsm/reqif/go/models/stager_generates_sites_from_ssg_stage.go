package models

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/lib/ssg/go/defaults"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

const logo string = "logo.svg"

func (stager *Stager) generatesSiteFromSSGStage() {

	var siteWeb *StaticWebSite
	{
		siteWebs := GetGongstructInstancesSet[StaticWebSite](stager.stage)
		if len(*siteWebs) != 1 {
			log.Fatalln("There should be one siteWeb")
		}
		for k := range *siteWebs {
			siteWeb = k
		}
		_ = siteWeb
	}

	// check that the relative path is sanitized
	if !filepath.IsLocal(siteWeb.OutputStaticWebDir) {
		log.Fatalln("Path siteWeb.OutputStaticWebDir", siteWeb.OutputStaticWebDir, "is not local")
	}
	if !filepath.IsLocal(siteWeb.InputImagesDir) {
		log.Fatalln("Path siteWeb.InputImagesDir", siteWeb.InputImagesDir, "is not local")
	}

	extractGongSsgEmbeddedFiles(defaults.StaticDir, stager.pathToExtractedGongSsgDefaultStaticFiles)
	extractGongSsgEmbeddedFiles(defaults.LayoutsDir, stager.pathToExtractedGongSsgDefaultLayoutFiles)

	// start by copying the static directory
	_, shouldReturn := stager.copyImagesToDirectoryForSsgGeneration()
	if shouldReturn {
		return
	}

	// serialize the stage into a XL file
	filename := filepath.Join(stager.rootPathToImageInputs, siteWeb.InputImagesDir, "reqif.xlsx")
	SerializeStage(stager.stage, filename)

	stager.ssgStage.Generation()
}

func (stager *Stager) copyImagesToDirectoryForSsgGeneration() (error, bool) {
	ssg.CopyDirectory(
		stager.pathToExtractedGongSsgDefaultStaticFiles,
		stager.pathToFilesUsedForSsgGeneration,
	)

	var siteWeb *StaticWebSite
	{
		siteWebs := GetGongstructInstancesSet[StaticWebSite](stager.stage)
		if len(*siteWebs) != 1 {
			log.Fatalln("There should be one siteWeb")
		}
		for k := range *siteWebs {
			siteWeb = k
		}
		_ = siteWeb
	}

	pathToGeneratedImages := filepath.Join(stager.pathToFilesUsedForSsgGeneration, "static/images")
	// --- Start: Remove existing pathToGeneratedSVG directory ---
	// log.Printf("Attempting to remove existing directory: %s", pathToGeneratedSVG)
	err := os.RemoveAll(pathToGeneratedImages)
	if err != nil {
		// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
		// or if the error was related to something already gone.
		// If MkdirAll fails later, that error will be caught.
		log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", pathToGeneratedImages, err)
	} else {
		// log.Printf("Successfully removed existing directory: %s", pathToGeneratedSVG)
	}
	// --- End: Remove existing pathToGeneratedSVG directory ---

	// --- Start: Generate root _index.md for the Content ---
	// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
	err = os.MkdirAll(pathToGeneratedImages, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating root content directory '%s': %v\n", pathToGeneratedImages, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root directory cannot be created.
		return nil, true
	}
	// log.Printf("Root content directory created or already exists: %s\n", pathToGeneratedSVG)

	ssg.CopyFile(
		filepath.Join(stager.rootPathToImageInputs, siteWeb.InputImagesDir, logo),
		filepath.Join(pathToGeneratedImages, logo))

	images := *GetGongstructInstancesSet[StaticWebSiteImage](stager.stage)
	for image := range images {
		ssg.CopyFile(filepath.Join(image.SourceDirectoryPath, siteWeb.InputImagesDir, image.Name), filepath.Join(pathToGeneratedImages, image.Name))
	}

	return err, false
}
