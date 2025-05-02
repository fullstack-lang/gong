package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type GongstructNodeProxy struct {
	node            *tree.Node
	stager          *Stager
	classDiagram    *Classdiagram
	gongstruct      *gong.GongStruct
	gongStructShape *GongStructShape
}

func (proxy *GongstructNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		proxy.classDiagram.AddGongStructShape(proxy.stager.stage, diagramPackage, proxy.gongstruct.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongStructShape(proxy.stager.stage, proxy.gongstruct.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		proxy.gongStructShape.IsExpanded = true
		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		proxy.gongStructShape.IsExpanded = false
		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
