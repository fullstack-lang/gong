package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func addAddItemButton[
	T Gongstruct,
	PT interface {
		*T
		AbstractType
	},
	CT interface {
		*S
		RectShapeInterface
		ConcreteType
	},
	S Gongstruct,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
](
	stager *Stager,
	parentItemsWhoseNodeIsExpanded *[]PT, parentItem PT, isNodeExpanded *bool,
	node *tree.Node, items *[]PT,
	receivingDiagram *Diagram, shapes *[]CT,
	associationShapes *[]ACT,
) {
	var dummyItem PT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(dummyItem) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
			newItem := PT(new(T))
			newItem.SetName("New" + GetGongstructNameFromPointer(newItem))
			newItem.StageVoid(stager.stage)
			*items = append(*items, newItem)
			stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

			// if the created item is a newDiagram, set the IsEditable_ field to true
			if newDiagram, ok := any(newItem).(*Diagram); ok {
				newDiagram.IsEditable_ = true
				newDiagram.IsExpanded = true
				for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
					diagram_.IsChecked = false
				}
				newDiagram.IsChecked = true
			}

			// if the created item is a project, add a diagram to it
			if newProject, ok := any(newItem).(*Project); ok {
				newProject.IsExpanded = true
				for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
					diagram_.IsChecked = false
				}
				newDiagram := &Diagram{
					Name:        "Default Diagram",
					IsChecked:   true,
					IsEditable_: true,
					ExpandableNodeObject: ExpandableNodeObject{
						IsExpanded: true,
					},
				}
				newDiagram.StageVoid(stager.stage)
				newProject.Diagrams = append(newProject.Diagrams, newDiagram)
			}

			if receivingDiagram != nil && shapes != nil {
				newShape := newShapeToDiagram(newItem, receivingDiagram, shapes, stager.stage)

				// get the parent shape to eventually create an association shape
				var parentShape CT
				if parentItem != nil {
					for _, parentShape_ := range *shapes {
						if parentShape_.GetAbstractElement() == parentItem {
							parentShape = parentShape_
							break
						}
					}
				}
				if parentShape != nil && parentItem != nil && associationShapes != nil {
					addAssociationShapeToDiagram(stager, parentItem, newItem, associationShapes)

					newShape.SetX(parentShape.GetX() + float64(len(*items)-1)*parentShape.GetWidth()*1.2)
					newShape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)
				}
			}

			stager.probeForm.FillUpFormFromGongstruct(newItem, GetPointerToGongstructName[PT]())

			// add the parent item to the list of items whose node is expanded
			if parentItemsWhoseNodeIsExpanded != nil && parentItem != nil &&
				!slices.Contains(*parentItemsWhoseNodeIsExpanded, parentItem) {
				*parentItemsWhoseNodeIsExpanded = append(*parentItemsWhoseNodeIsExpanded, parentItem)
			}
			if isNodeExpanded != nil {
				*isNodeExpanded = true
			}

			stager.stage.Commit()
		},
	}
}

func onAddAssociationShape[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
](
	stager *Stager, start ATstart, end ATend, shapes *[]ACT) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		addAssociationShapeToDiagram(stager, start, end, shapes)
	}
}

func addAssociationShapeToDiagram[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType // the association concrete type shape
	},
	ACT_ Gongstruct](stager *Stager, start ATstart, end ATend, shapes *[]ACT) {
	compositionShape := newConcreteAssociation(start, end, shapes)
	compositionShape.StageVoid(stager.stage)
	stager.stage.Commit()
}

func newConcreteAssociation[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](start ATstart, end ATend, shapes *[]ACT) ACT {
	compositionShape := ACT(new(ACT_))
	compositionShape.SetName(start.GetName() + " to " + end.GetName())
	compositionShape.SetAbstractStartElement(start)
	compositionShape.SetAbstractEndElement(end)
	compositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
	compositionShape.SetEndOrientation(ORIENTATION_VERTICAL)

	if taskInputShape, ok := any(compositionShape).(*TaskInputShape); ok {
		taskInputShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		taskInputShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
		taskInputShape.CornerOffsetRatio = 0.2
	}
	if taskOutputShape, ok := any(compositionShape).(*TaskOutputShape); ok {
		taskOutputShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		taskOutputShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
		taskOutputShape.CornerOffsetRatio = 0.2
	}

	compositionShape.SetCornerOffsetRatio(1.68)
	compositionShape.SetStartRatio(0.5)
	compositionShape.SetEndRatio(0.5)
	*shapes = append(*shapes, compositionShape)

	return compositionShape
}

func onRemoveAssociationShape[
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](stager *Stager, compositionShape ACT, shapes *[]ACT) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape.UnstageVoid(stager.stage)
		idx := slices.Index(*shapes, compositionShape)
		*shapes = slices.Delete(*shapes, idx, idx+1)
		stager.stage.Commit()
	}
}

func onUpdateElementInDiagram[
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
			stager.stage.Commit()
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}

func onUpdateExpandableNode[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
](
	stager *Stager,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
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

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}
