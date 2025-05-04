package models

import (
	"fmt"
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonNewClassdiagramProxy struct {
	stager *Stager
}

func (proxy *ButtonNewClassdiagramProxy) ButtonUpdated(
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

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := "Default"
	newClassdiagramName := initialName
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *GetGongstructInstancesSet[Classdiagram](stage) {
			if classdiagram.Name == newClassdiagramName {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			newClassdiagramName = initialName + fmt.Sprintf("_%d", index)
		}
	}

	newClassdiagram := (&Classdiagram{
		Name: newClassdiagramName,
	}).Stage(stage)

	diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, newClassdiagram)

	stager.UpdateAndCommitTreeStage()
	stage.Commit()
}
