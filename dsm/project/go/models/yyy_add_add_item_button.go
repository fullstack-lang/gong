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

// addItemButtonConfiguration defines the configuration for an "add" button
// in the tree of instance in a DSM
// instead of having a large API surface, there is one
// addAddElementButton(addItemButtonConfiguration)
type addItemButtonConfiguration[

	// Abstact Type of the added element
	AT Gongstruct,
	PAT interface {
		*AT
		AbstractType
	},

	// Abstact Type of the parent node of the added element
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},

	// Concrete type of the added element
	CT Gongstruct,
	PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},

	// Concrete type of the association concrete element
	ACT Gongstruct,
	PACT interface {
		*ACT
		LinkShapeInterface
		AssociationConcreteType
	},

] struct {
	// parentNode is where the node with be stored
	parentNode *tree.Node

	// the slice where to append the added element
	sliceForNewAddedItem *[]PAT

	// IsParentNodeExpanded indicates if it is necessary to expand the parent node
	// in most case, it is true. When you add an item, you expect the parent node to expand
	isParentNodeExpandedByAddOperation bool

	//
	parentNodeExpansionType
	parentNodeExpansionBooleanValue *bool
	// if the parent element is part of the slice, the node is exapanded
	parentNodeExpansionSliceEncoding *[]PParentAT
	parentElement                    PParentAT

	// IsWithAdditionOfConcreteShape tells if a shape is added to diagram up in the tree
	// by convention, in a DSM, the edit tree is organized by diagrams with
	// the abstract instances below the diagram
	isWithAdditionOfShape bool
	receivingDiagram      DiagramIF
	sliceForNewAddedShape *[]PCT

	// IsWithAdditionOfAssociationShape tells if an association shape is added to diagram
	isWithAdditionOfAssociationShape bool
	sliceForNewCompositionShapes     *[]PACT
}

type parentNodeExpansionType string

const (
	parentNodeExpansionTypeNone           parentNodeExpansionType = ""
	parentNodeExpansionTypeByBooleanValue parentNodeExpansionType = "parentNodeExpansionTypeByBooleanValue"
	parentNodeExpansionTypeBySlice        parentNodeExpansionType = "parentNodeExpansionTypeBySlice"
)

// look forwward to https://github.com/golang/go/issues/77273
// spec: generic methods for Go
func addAddButton[

	// Abstact Type of the added element
	PAT interface {
		*AT
		AbstractType
	},
	AT Gongstruct,

	// Abstact Type of the parent node of the added element
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},

	// Concrete type of the added element
	CT Gongstruct,
	PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},

	// Concrete type of the association concrete element
	ACT Gongstruct,
	PACT interface {
		*ACT
		LinkShapeInterface
		AssociationConcreteType
	},
](
	stager *Stager,
	conf addItemButtonConfiguration[AT, PAT, ParentAT, PParentAT, CT, PCT, ACT, PACT],

) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name: GetGongstructNameFromPointer(dummyItem) + " " + string(buttons.BUTTON_add),
		Icon: string(buttons.BUTTON_add),
		ToolTipText: "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" +
			conf.parentNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	conf.parentNode.Buttons = append(conf.parentNode.Buttons, addButton)

	addButton.OnClick = func() {
		newAbstractElement := PAT(new(AT))
		callbacks.createdItem = newAbstractElement
		newAbstractElement.SetName("New" + GetGongstructNameFromPointer(newAbstractElement))
		newAbstractElement.SetName("") // easier to rename an item when its name is empty
		newAbstractElement.SetIsInRenameMode(true)
		newAbstractElement.StageVoid(stager.stage)
		*conf.sliceForNewAddedItem = append(*conf.sliceForNewAddedItem, newAbstractElement)
		stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

		// stager.probeForm.SetCommitMode(false), no need yet
		stager.probeForm.FillUpFormFromGongstruct(newAbstractElement, GetPointerToGongstructName[PAT]())
		// stager.probeForm.SetCommitMode(true)

		// add the parent item to the list of items whose node is expanded
		if conf.isParentNodeExpandedByAddOperation {
			switch conf.parentNodeExpansionType {
			case parentNodeExpansionTypeBySlice:
				if conf.parentNodeExpansionSliceEncoding != nil && conf.parentElement != nil &&
					!slices.Contains(*conf.parentNodeExpansionSliceEncoding, conf.parentElement) {
					*conf.parentNodeExpansionSliceEncoding = append(*conf.parentNodeExpansionSliceEncoding, conf.parentElement)
				}
			case parentNodeExpansionTypeByBooleanValue:
				*conf.parentNodeExpansionBooleanValue = true
			}
		}

		if conf.isWithAdditionOfShape {
			if conf.receivingDiagram != nil && conf.sliceForNewAddedShape != nil {
				newShape := newShapeToDiagram(newAbstractElement, conf.receivingDiagram, conf.sliceForNewAddedShape, stager.stage)

				// get the parent shape to eventually create an association shape
				var parentShape PCT
				if conf.parentElement != nil {
					for _, parentShape_ := range *conf.sliceForNewAddedShape {
						if parentShape_.GetAbstractElement() == conf.parentElement {
							parentShape = parentShape_
							break
						}
					}
				}
				if conf.isWithAdditionOfAssociationShape {
					if parentShape != nil && conf.parentElement != nil && conf.sliceForNewCompositionShapes != nil {
						addAssociationShapeToDiagram(stager, conf.parentElement, newAbstractElement, conf.sliceForNewCompositionShapes)

						newShape.SetX(parentShape.GetX() + float64(len(*conf.sliceForNewAddedItem)-1)*parentShape.GetWidth()*1.2)
						newShape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)
					}
				}
			}
		}

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}
	return
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
	parentItemsWhoseNodeIsExpanded *[]PAT,
	parentItem PAT,
	isNodeExpanded *bool,
	node *tree.Node,
	items *[]PAT,
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
