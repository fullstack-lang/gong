package models

import (
	"log"
	"path/filepath"
	"time"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonGeneratesStaticWebSiteProxy struct {
	stager *Stager
}

func (proxy *ButtonGeneratesStaticWebSiteProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	stager := proxy.stager
	stage := stager.stage

	imageFilesDirectoryPath := "../../diagrams/static/images"

	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	var diagramPackage *DiagramPackage
	for k := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		log.Fatalln("There should be at least one diagram package on the stage")
	}

	stager.prepareStaticDic(imageFilesDirectoryPath)

	// grab all images of diagrams to be included in the static web site
	selectedClassdiagram := diagramPackage.SelectedClassdiagram
	for _, classDiagram := range diagramPackage.Classdiagrams {
		if !classDiagram.IsIncludedInStaticWebSite {
			continue
		}

		diagramPackage.SelectedClassdiagram = classDiagram
		stager.UpdateSVGStage()

		imageFilePath := filepath.Join(imageFilesDirectoryPath, classDiagram.Name+".svg")
		svg.GrabGeneratedSVGFile(stager.svgStage, imageFilePath, 5000*time.Millisecond)
	}
	diagramPackage.SelectedClassdiagram = selectedClassdiagram

	stager.ssgStage.Reset()

	ssg.StageBranch(
		stager.ssgStage,
		&ssg.Content{
			Name:           "Generated",
			MardownContent: "This is the web site for ...",

			LayoutPath: "../../diagrams/layout",
			StaticPath: "../../diagrams/static",

			ContentPath: "../../diagrams/content",
			OutputPath:  "../../diagrams/docs",

			Target: ssg.Target(ssg.FILE),

			// Chapters: ssgChapters,
		},
	)

	stager.ssgStage.Commit()
}
