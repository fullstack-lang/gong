package models

import (
	"log"
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

func OnUpdateElementInDiagram[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	CT interface {
		*CT_
		RectShapeInterface
		ConcreteType
	},
	CT_ Gongstruct](
	stager *Stager,
	diagram *Diagram,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap *map[AT]CT,
) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		// find the shape (if any)
		shape := (*shapesMap)[element]

		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager.stage)
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stager.stage)
			idx := slices.Index(*shapes, shape)
			*shapes = slices.Delete(*shapes, idx, idx+1)
			stager.stage.Commit()
			return
		}

		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
					*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
				}
			} else {
				if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
					*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[*Product]())
		stager.stage.Commit()
	}
}
