package models

import (
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
	diagramPackage := getTheDiagramPackage(stager.stage)

	for _, classDiagram := range diagramPackage.Classdiagrams {
		var selected bool
		if diagramPackage.SelectedClassdiagram == classDiagram {
			selected = true
		}
		node := &tree.Node{
			Name:              classDiagram.Name,
			HasCheckboxButton: true,
			IsChecked:         selected,
		}
		node.Impl = &ClassDiagramNodeCheckProxy{
			node:         node,
			stager:       stager,
			classDiagram: classDiagram,
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
	node         *tree.Node
	stager       *Stager
	classDiagram *Classdiagram
}

// OnAfterUpdate is called when a node is checked/unchecked by the user
func (proxy *ClassDiagramNodeCheckProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// uncheck all other diagram
	diagramPackage := getTheDiagramPackage(proxy.stager.stage)
	diagramPackage.SelectedClassdiagram = proxy.classDiagram

	proxy.stager.UpdateAndCommitTreeStage()
	proxy.stager.stage.Commit()
}
