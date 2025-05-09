package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ClassDiagramNodeProxy struct {
	node         *tree.Node
	stager       *Stager
	classDiagram *Classdiagram
}

// OnAfterUpdate is called when a node is checked/unchecked by the user
func (proxy *ClassDiagramNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = proxy.classDiagram

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		diagramPackage.SelectedClassdiagram = nil

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.IsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.IsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}

	if front.Name != staged.Name {
		proxy.classDiagram.Name = front.Name
		proxy.classDiagram.IsInRenameMode = false

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()
	}
}
