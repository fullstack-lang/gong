package models

import (
	gong "github.com/fullstack-lang/gong/go/models"

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
		nodeClassdiagram := &tree.Node{
			Name:              classDiagram.Name,
			HasCheckboxButton: true,
			IsChecked:         selected,
		}
		nodeClassdiagram.Impl = &ClassDiagramNodeCheckProxy{
			node:         nodeClassdiagram,
			stager:       stager,
			classDiagram: classDiagram,
		}

		root.Children = append(root.Children, nodeClassdiagram)

		if diagramPackage.SelectedClassdiagram == classDiagram {
			nodeClassdiagram.IsExpanded = true
			for _, gongStruct := range gong.GetGongstrucsSorted[*gong.GongStruct](stager.gongStage) {

				nodeNamedStruct := &tree.Node{
					Name:              gongStruct.Name,
					HasCheckboxButton: true,
				}
				nodeClassdiagram.Children = append(nodeClassdiagram.Children, nodeNamedStruct)
			}
		}
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

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = proxy.classDiagram

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = nil

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()
	}
}
