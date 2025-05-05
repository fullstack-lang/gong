package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ClassDiagramGongStructsNodeProxy struct {
	stager       *Stager
	classDiagram *Classdiagram
}

func (proxy *ClassDiagramGongStructsNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	if front.IsExpanded && !staged.IsExpanded {
		proxy.classDiagram.NodeGongStructsIsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.classDiagram.NodeGongStructsIsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
