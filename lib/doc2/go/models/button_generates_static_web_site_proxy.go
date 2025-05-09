package models

import (
	"log"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
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

	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	var diagramPackage *DiagramPackage
	for k := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		log.Fatalln("There should be at least one diagram package on the stage")
	}

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
