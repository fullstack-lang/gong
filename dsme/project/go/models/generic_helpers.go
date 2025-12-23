package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func OnUpdateAbstractElement[AT AbstractType](stager *Stager, element AT) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			element.SetIsExpanded(frontNode.IsExpanded)
		} else {
			stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		}
		stager.stage.Commit()
	}
}

func OnAddCompositionShape[
	AT AbstractType,
	CCT interface {
		*CCT_
		LinkShapeInterface
		CompositionConcreteType
	},
	CCT_ Gongstruct](
	stager *Stager, diagram *Diagram, parent, child AT, shapes *[]CCT) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape := CCT(new(CCT_))
		compositionShape.StageVoid(stager.stage)
		compositionShape.SetName(parent.GetName() + " to " + child.GetName())
		compositionShape.SetAbstractChildElement(child)
		compositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
		compositionShape.SetEndOrientation(ORIENTATION_VERTICAL)
		compositionShape.SetCornerOffsetRatio(1.68)
		compositionShape.SetStartRatio(0.5)
		compositionShape.SetEndRatio(0.5)

		*shapes = append(*shapes, compositionShape)
		stager.stage.Commit()
	}
}

func OnRemoveCompositionShape[
	CCT interface {
		*CCT_
		LinkShapeInterface
		CompositionConcreteType
	},
	CCT_ Gongstruct](stager *Stager, diagram *Diagram, compositionShape CCT, shapes *[]CCT) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape.UnstageVoid(stager.stage)
		idx := slices.Index(*shapes, compositionShape)
		*shapes = slices.Delete(*shapes, idx, idx+1)
		stager.stage.Commit()
	}
}
