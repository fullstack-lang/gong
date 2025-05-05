package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ClassDiagramGongEnumsNodeProxy struct {
	stager       *Stager
	classDiagram *Classdiagram
}

func (proxy *ClassDiagramGongEnumsNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.NodeGongEnumsIsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.NodeGongEnumsIsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
