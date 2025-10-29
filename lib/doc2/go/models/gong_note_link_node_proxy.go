package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type GongNoteLinkNodeProxy struct {
	node              *tree.Node
	stager            *Stager
	classDiagram      *Classdiagram
	gongNote          *gong.GongNote
	gongNoteShape     *GongNoteShape
	gongNoteLink      *gong.GongLink
	gongNoteLinkShape *GongNoteLinkShape
}

func (proxy *GongNoteLinkNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		proxy.classDiagram.AddGongNoteLinkShape(
			proxy.stager.stage,
			proxy.stager.gongStage,
			proxy.gongNote,
			proxy.gongNoteLink,
			proxy.gongNoteShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongNoteLinkShape(
			proxy.stager.stage,
			proxy.gongNoteLinkShape,
			proxy.gongNoteShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}
}
