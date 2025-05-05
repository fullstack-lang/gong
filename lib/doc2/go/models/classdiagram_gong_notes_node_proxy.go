package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ClassDiagramGongNotesNodeProxy struct {
	stager       *Stager
	classDiagram *Classdiagram
}

func (proxy *ClassDiagramGongNotesNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.NodeGongNotesIsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.NodeGongNotesIsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
