package models

import (
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
			newItem.SetName("") // easier to rename an item when its name is empty
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
					AbstractTypeFields: AbstractTypeFields{
						IsExpanded: true,
					},
				}
				newDiagram.StageVoid(stager.stage)
				newProject.Diagrams = append(newProject.Diagrams, newDiagram)
			}

			// stager.probeForm.SetCommitMode(false), no need yet
			stager.probeForm.FillUpFormFromGongstruct(newItem, GetPointerToGongstructName[PT]())
			// stager.probeForm.SetCommitMode(true)

			// add the parent item to the list of items whose node is expanded
			if parentItemsWhoseNodeIsExpanded != nil && parentItem != nil &&
				!slices.Contains(*parentItemsWhoseNodeIsExpanded, parentItem) {
				*parentItemsWhoseNodeIsExpanded = append(*parentItemsWhoseNodeIsExpanded, parentItem)
			}
			if isNodeExpanded != nil {
				*isNodeExpanded = true
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
			stager.stage.Commit()
		},
	}
}
