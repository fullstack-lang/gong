package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type GongEnumNodeValueProxy struct {
	stager        *Stager
	classDiagram  *Classdiagram
	gongEnumShape *GongEnumShape
	gongEnum      *gong.GongEnum
	gongEnumValue *gong.GongEnumValue
}

func (proxy *GongEnumNodeValueProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		proxy.classDiagram.AddGongEnumValueShapeToDiagram(
			proxy.stager.stage,
			proxy.gongEnumShape,
			proxy.gongEnum,
			proxy.gongEnumValue)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveGongEnumValueShapeFromDiagram(
			proxy.stager.stage,
			proxy.gongEnumShape,
			proxy.gongEnumValue,
		)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}
}
