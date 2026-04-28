package models

import (
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type itemAdderCallback[PT AbstractType] struct {
	createdItem    PT
	OnBeforeCommit func()
}

// addAddItemButton appends an "add" button to the given tree node.
// When clicked, this button instantiates a new abstract element of type PT,
// adds it to the provided items slice, and prepares the UI for immediate renaming.
//
// Special behaviors are implemented for specific element types (e.g., Diagram, Library).
// If a receivingDiagram and shapes are provided, the function also creates
// the corresponding visual shape (of type CT) on the diagram. If a parentItem
// is provided and its shape is found, an association link (of type ACT) is automatically
// generated between the parent and the newly created child shape.
//
// The heavy use of generics allows this function to be completely agnostic to the actual
// underlying types (e.g. Product, Task, Note, and their respective shapes).
func addAddItemButton[
	AT Gongstruct,
	PAT interface {
		*AT
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
	parentItemsWhoseNodeIsExpanded *[]PAT, parentItem PAT, isNodeExpanded *bool,
	node *tree.Node, items *[]PAT,
	receivingDiagram DiagramIF,
	shapes *[]CT,
	associationShapes *[]ACT,
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(dummyItem) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.OnClick = func() {
		newAbstractElement := PAT(new(AT))
		callbacks.createdItem = newAbstractElement
		newAbstractElement.SetName("New" + GetGongstructNameFromPointer(newAbstractElement))
		newAbstractElement.SetName("") // easier to rename an item when its name is empty
		newAbstractElement.SetIsInRenameMode(true)
		newAbstractElement.StageVoid(stager.stage)
		*items = append(*items, newAbstractElement)
		stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

		// stager.probeForm.SetCommitMode(false), no need yet
		stager.probeForm.FillUpFormFromGongstruct(newAbstractElement, GetPointerToGongstructName[PAT]())
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
			newShape := newShapeToDiagram(newAbstractElement, receivingDiagram, shapes, stager.stage)

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
				addAssociationShapeToDiagram(stager, parentItem, newAbstractElement, associationShapes)

				newShape.SetX(parentShape.GetX() + float64(len(*items)-1)*parentShape.GetWidth()*1.2)
				newShape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)
			}
		}

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}

	return
}

// addAddItemButtonSimple appends an "add" button to the given tree node.
// When clicked, this button instantiates a new abstract element of type PT,
// adds it to the provided items slice, and prepares the UI for immediate renaming.
// It is a simplified version of addAddItemButton without diagram/shape logic.
func addAddItemButtonSimple[
	AT Gongstruct,
	PAT interface {
		*AT
		AbstractType
	},
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},
](
	stager *Stager,
	parentItemsWhoseNodeIsExpanded *[]PParentAT,
	parentItem PParentAT,
	isNodeExpanded *bool,
	node *tree.Node,
	items *[]PAT,
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(dummyItem) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.OnClick = func() {
		newAbstractElement := PAT(new(AT))
		callbacks.createdItem = newAbstractElement
		newAbstractElement.SetName("New" + GetGongstructNameFromPointer(newAbstractElement))
		newAbstractElement.SetName("") // easier to rename an item when its name is empty
		newAbstractElement.SetIsInRenameMode(true)
		newAbstractElement.StageVoid(stager.stage)
		*items = append(*items, newAbstractElement)
		stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

		stager.probeForm.FillUpFormFromGongstruct(newAbstractElement, GetPointerToGongstructName[PAT]())

		// add the parent item to the list of items whose node is expanded
		if parentItemsWhoseNodeIsExpanded != nil && parentItem != nil &&
			!slices.Contains(*parentItemsWhoseNodeIsExpanded, parentItem) {
			*parentItemsWhoseNodeIsExpanded = append(*parentItemsWhoseNodeIsExpanded, parentItem)
		}
		if isNodeExpanded != nil {
			*isNodeExpanded = true
		}

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}

	return
}

// addAddItemButtonVerySimple appends an "add" button to the given tree node.
// When clicked, this button instantiates a new abstract element of type PT,
// adds it to the provided items slice, and prepares the UI for immediate renaming.
// It is a simplified version of addAddItemButton without diagram/shape logic.
func addAddItemButtonVerySimple[
	AT Gongstruct,
	PAT interface {
		*AT
		AbstractType
	},
](
	stager *Stager,
	isNodeExpanded *bool,
	node *tree.Node,
	items *[]PAT,
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(dummyItem) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.OnClick = func() {
		newAbstractElement := PAT(new(AT))
		callbacks.createdItem = newAbstractElement
		newAbstractElement.SetName("New" + GetGongstructNameFromPointer(newAbstractElement))
		newAbstractElement.SetName("") // easier to rename an item when its name is empty
		newAbstractElement.SetIsInRenameMode(true)
		newAbstractElement.StageVoid(stager.stage)
		*items = append(*items, newAbstractElement)
		stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

		stager.probeForm.FillUpFormFromGongstruct(newAbstractElement, GetPointerToGongstructName[PAT]())

		if isNodeExpanded != nil {
			*isNodeExpanded = true
		}

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}

	return
}
