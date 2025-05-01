package models

import (
	"log"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()

	// put a "class diagram button at the root"
	root := &tree.Node{
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
	}

	// append a node below for each diagram
	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stager.stage)
	var diagramPackage *DiagramPackage
	for k, _ := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		log.Fatalln("There should be at least one diagram package on the stage")
	}

	for _, classDiagram := range diagramPackage.Classdiagrams {
		node := &tree.Node{
			Name:              classDiagram.Name,
			HasCheckboxButton: true,
		}
		node.Impl = &ClassDiagramNodeCheckProxy{
			node:   node,
			stager: stager,
		}

		root.Children = append(root.Children, node)
	}

	tree.StageBranch(stager.treeStage,
		&tree.Tree{
			Name:      stager.stage.GetProbeTreeSidebarStageName(),
			RootNodes: []*tree.Node{root},
		},
	)

	stager.treeStage.Commit()
}

type ClassDiagramNodeCheckProxy struct {
	node   *tree.Node
	stager *Stager
}

// OnAfterUpdate is called when a node is checked/unchecked by the user
func (impl *ClassDiagramNodeCheckProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

}
