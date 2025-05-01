package models

import (
	"fmt"
	"log"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()

	tree.StageBranch(stager.treeStage,
		&tree.Tree{
			Name: stager.stage.GetProbeTreeSidebarStageName(),
			RootNodes: []*tree.Node{
				{
					Name:       "Class Diagrams",
					IsExpanded: true,
					Buttons: []*tree.Button{
						tree.NewButton(
							&ButtonNewClassdiagramProxy{
								stager: stager,
							},
							"Class Diagramm Add Button",
							string(buttons.BUTTON_add)),
					},
				},
			},
		},
	)

	stager.treeStage.Commit()
}

type ButtonNewClassdiagramProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (b *ButtonNewClassdiagramProxy) GetButtonsStage() *tree.Stage {
	return b.stager.treeStage
}

// OnAfterUpdateButton implements models.Target.
func (b *ButtonNewClassdiagramProxy) OnAfterUpdateButton() {

	stager := b.stager
	stage := stager.stage

	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	var diagramPackage *DiagramPackage
	for k, _ := range diagramPackages {
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

	newClassdiagram := (&Classdiagram{Name: newClassdiagramName}).Stage(stage)

	diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, newClassdiagram)
	stage.Commit()
}
