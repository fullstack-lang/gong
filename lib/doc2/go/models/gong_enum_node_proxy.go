package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type GongEnumNodeProxy struct {
	node          *tree.Node
	stager        *Stager
	classDiagram  *Classdiagram
	gongEnum      *gong.GongEnum
	gongEnumShape *GongEnumShape
	rank          int
}

func (proxy *GongEnumNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		diagramPackage := getTheDiagramPackage(proxy.stager.stage)
		proxy.classDiagram.AddGongEnumShape(proxy.stager.stage, diagramPackage, proxy.gongEnum.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongEnumShape(proxy.stager.stage, proxy.gongEnum.Name)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	if front.IsExpanded && !staged.IsExpanded {
		ToggleNodeExpanded(&proxy.classDiagram.NodeGongEnumNodeExpansionBinaryEncoding, proxy.rank)

		front.IsExpanded = false

		proxy.stager.stage.Commit()
	}
	if !front.IsExpanded && staged.IsExpanded {
		ToggleNodeExpanded(&proxy.classDiagram.NodeGongEnumNodeExpansionBinaryEncoding, proxy.rank)

		front.IsExpanded = true

		proxy.stager.stage.Commit()
	}
}
