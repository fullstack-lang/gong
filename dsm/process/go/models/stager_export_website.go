package models

import (
	"log"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (stager *Stager) exportWebsite() {
	stager.ssgStage.Reset()

	content := ssg.Content{
		Name:           "Root to project the website",
		ContentPath:    "/tmp/project",
		MardownContent: "## Project website",
		OutputPath:     "./generated static web site",
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
}
