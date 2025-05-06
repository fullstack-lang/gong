package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type GongNoteNodeProxy struct {
	node          *tree.Node
	stager        *Stager
	classDiagram  *Classdiagram
	gongNote      *gong.GongNote
	gongNoteShape *GongNoteShape
	rank          int
}

func (proxy *GongNoteNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)

		proxy.classDiagram.AddGongNoteShape(proxy.stager.stage, proxy.gongNote, diagramPackage, proxy.gongNote.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongNoteShape(proxy.stager.stage, proxy.gongNote.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		ToggleNodeExpanded(&proxy.classDiagram.NodeGongNoteNodeExpansionBinaryEncoding, proxy.rank)

		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		ToggleNodeExpanded(&proxy.classDiagram.NodeGongNoteNodeExpansionBinaryEncoding, proxy.rank)

		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
