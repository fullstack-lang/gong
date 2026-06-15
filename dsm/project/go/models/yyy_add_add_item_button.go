// generated code (do not edit)
package models

import (
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// This file contains the implementation of the "add item" button for tree nodes.
// The main function is addAddItemButton, which is a highly generic function that can be used
// to add any type of item (e.g. Product, Task, Note) to any type of parent item (e.g. Product, Task, Participant).
// It also supports the automatic creation of the corresponding shape on the diagram and an association link to the parent shape.
//
// The implementation relies heavily on Go's generics to achieve maximum reusability and type safety.
// The configuration structs (ItemButtonConfiguration, ItemAndShapeButtonConfiguration, ItemShapeAndLinkButtonConfiguration)
// allow for flexible customization of the button's behavior and the associated diagram updates.
//
// implement is currently quite verbose but go langauages changes will significantly reduce this verbosity:
// https://github.com/golang/go/issues/61731, proposal, spec: support type inference on generic structs
// https://github.com/golang/go/issues/9859 , spec: direct reference to embedded fields in struct literals (in go 1.27)
// https://github.com/golang/go/issues/12854, spec: type inferred composite literals
type itemAdderCallback[PT AbstractType] struct {
	createdItem    PT
	OnBeforeCommit func()
}

type parentNodeExpansionType string

const (
	parentNodeExpansionTypeNone           parentNodeExpansionType = "parentNodeExpansionTypeNone"
	parentNodeExpansionTypeByBooleanValue parentNodeExpansionType = "parentNodeExpansionTypeByBooleanValue"
	parentNodeExpansionTypeBySlice        parentNodeExpansionType = "parentNodeExpansionTypeBySlice"
)

// ---------------------------------------------------------
// 1. BASE CONFIGURATION (Abstract elements only)
// ---------------------------------------------------------

type ItemButtonConfiguration[
	AT Gongstruct, PAT interface {
		*AT
		AbstractType
	},
	ParentAT Gongstruct, PParentAT interface {
		*ParentAT
		AbstractType
	},
] struct {
	parentNode                         *tree.Node
	sliceForNewAddedItem               *[]PAT
	isParentNodeExpandedByAddOperation bool
	parentNodeExpansionType            parentNodeExpansionType
	parentNodeExpansionBooleanValue    *bool
	parentNodeExpansionSliceEncoding   *[]PParentAT
	parentElement                      PParentAT
}

// ---------------------------------------------------------
// 2. MIDDLE CONFIGURATION (Abstract + Shape)
// ---------------------------------------------------------

type ItemAndShapeButtonConfiguration[
	AT Gongstruct, PAT interface {
		*AT
		AbstractType
	},
	ParentAT Gongstruct, PParentAT interface {
		*ParentAT
		AbstractType
	},
	CT Gongstruct, PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},
] struct {
	ItemButtonConfiguration[AT, PAT, ParentAT, PParentAT]

	receivingDiagram      DiagramIF
	sliceForNewAddedShape *[]PCT
}

// ---------------------------------------------------------
// 3. COMPLEX CONFIGURATION (Abstract + Shape + Link)
// ---------------------------------------------------------

type ItemShapeAndLinkButtonConfiguration[
	AT Gongstruct, PAT interface {
		*AT
		AbstractType
	},
	ParentAT Gongstruct, PParentAT interface {
		*ParentAT
		AbstractType
	},
	CT Gongstruct, PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},
	ACT Gongstruct, PACT interface {
		*ACT
		LinkShapeInterface
		AssociationConcreteType
	},
] struct {
	ItemAndShapeButtonConfiguration[AT, PAT, ParentAT, PParentAT, CT, PCT]

	sliceForNewCompositionShapes *[]PACT
}

// ---------------------------------------------------------
// 4. SHARED CORE LOGIC
// ---------------------------------------------------------

// processAbstractItemAddition handles the item creation, form updates,
// and tree expansion logic shared by the button types.
func processAbstractItemAddition[
	AT Gongstruct, PAT interface {
		*AT
		AbstractType
	},
	ParentAT Gongstruct, PParentAT interface {
		*ParentAT
		AbstractType
	},
](
	stager *Stager,
	conf ItemButtonConfiguration[AT, PAT, ParentAT, PParentAT],
	callbacks *itemAdderCallback[PAT],
) PAT {
	newAbstractElement := PAT(new(AT))
	callbacks.createdItem = newAbstractElement
	newAbstractElement.SetName("New" + GetGongstructNameFromPointer(newAbstractElement))
	newAbstractElement.SetName("") // easier to rename an item when its name is empty
	newAbstractElement.SetIsInRenameMode(true)
	newAbstractElement.StageVoid(stager.stage)
	*conf.sliceForNewAddedItem = append(*conf.sliceForNewAddedItem, newAbstractElement)
	stager.stage.ComputeReverseMaps() // this is important, otherwise, the form is not correctly initialized

	stager.probeForm.FillUpFormFromGongstruct(newAbstractElement, GetPointerToGongstructName[PAT]())

	// add the parent item to the list of items whose node is expanded
	if conf.isParentNodeExpandedByAddOperation {
		conf.parentNode.IsExpanded = true
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
	return newAbstractElement
}

// ---------------------------------------------------------
// 5. FUNCTIONS
// ---------------------------------------------------------

// addCreateItemButton appends an "add" button to the given tree node.
// This button instantiates a new abstract element of type PT.
func addCreateItemButton[
	PAT interface {
		*AT
		AbstractType
	},
	AT Gongstruct,
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},
](
	stager *Stager,
	conf ItemButtonConfiguration[AT, PAT, ParentAT, PParentAT],
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name: "Add " + GetGongstructNameFromPointer(dummyItem),
		Icon: string(buttons.BUTTON_add),
		ToolTipText: "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" +
			conf.parentNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)

	addButton.OnClick = func() {
		processAbstractItemAddition(stager, conf, callbacks)

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}
	return
}

// addCreateItemAndShapeButton appends an "add" button to the given tree node.
// This button instantiates a new abstract element and its visual shape.
func addCreateItemAndShapeButton[
	PAT interface {
		*AT
		AbstractType
	},
	AT Gongstruct,
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},
	CT Gongstruct,
	PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},
](
	stager *Stager,
	conf ItemAndShapeButtonConfiguration[AT, PAT, ParentAT, PParentAT, CT, PCT],
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name: "Add " + GetGongstructNameFromPointer(dummyItem),
		Icon: string(buttons.BUTTON_add),
		ToolTipText: "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" +
			conf.parentNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)

	addButton.OnClick = func() {
		newAbstractElement := processAbstractItemAddition(stager, conf.ItemButtonConfiguration, callbacks)

		if conf.receivingDiagram != nil && conf.sliceForNewAddedShape != nil {
			newShapeToDiagram(newAbstractElement, conf.receivingDiagram, conf.sliceForNewAddedShape, stager.stage)
		}

		if callbacks.OnBeforeCommit != nil {
			callbacks.OnBeforeCommit()
		}
		stager.stage.Commit()
	}
	return
}

// addCreateItemShapeAndLinkButton appends an "add" button to the given tree node.
// This button instantiates a new abstract element, its visual shape, and an association link.
func addCreateItemShapeAndLinkButton[
	PAT interface {
		*AT
		AbstractType
	},
	AT Gongstruct,
	ParentAT Gongstruct,
	PParentAT interface {
		*ParentAT
		AbstractType
	},
	CT Gongstruct,
	PCT interface {
		*CT
		RectShapeInterface
		ConcreteType
	},
	ACT Gongstruct,
	PACT interface {
		*ACT
		LinkShapeInterface
		AssociationConcreteType
	},
](
	stager *Stager,
	conf ItemShapeAndLinkButtonConfiguration[AT, PAT, ParentAT, PParentAT, CT, PCT, ACT, PACT],
) (callbacks *itemAdderCallback[PAT]) {
	callbacks = &itemAdderCallback[PAT]{}

	var dummyItem PAT
	addButton := &tree.Button{
		Name: "Add " + GetGongstructNameFromPointer(dummyItem),
		Icon: string(buttons.BUTTON_add),
		ToolTipText: "Add a " + GetGongstructNameFromPointer(dummyItem) + " to \"" +
			conf.parentNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)

	addButton.OnClick = func() {
		newAbstractElement := processAbstractItemAddition(stager, conf.ItemButtonConfiguration, callbacks)

		if conf.receivingDiagram != nil && conf.sliceForNewAddedShape != nil {
			newShape := newShapeToDiagram(newAbstractElement, conf.receivingDiagram, conf.sliceForNewAddedShape, stager.stage)

			var parentShape PCT
			if conf.parentElement != nil {
				for _, parentShape_ := range *conf.sliceForNewAddedShape {
					if parentShape_.GetAbstractElement() == conf.parentElement {
						parentShape = parentShape_
						break
					}
				}
			}
			if parentShape != nil && conf.parentElement != nil && conf.sliceForNewCompositionShapes != nil {
				addAssociationShapeToDiagram(stager, conf.parentElement, newAbstractElement, conf.sliceForNewCompositionShapes)

				var isHorizontal bool
				if pShape, ok := any(parentShape).(*ProductShape); ok {
					if pShape.LayoutDirection == Horizontal {
						isHorizontal = true
					}
				}

				if isHorizontal {
					newShape.SetX(parentShape.GetX() + parentShape.GetWidth()/2.0 + 50.0)
					newShape.SetY(parentShape.GetY() + parentShape.GetHeight() + 50.0 + float64(len(*conf.sliceForNewAddedItem)-1)*parentShape.GetHeight()*1.2)

					if len(*conf.sliceForNewCompositionShapes) > 0 {
						newCompositionShape := (*conf.sliceForNewCompositionShapes)[len(*conf.sliceForNewCompositionShapes)-1]
						newCompositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
						newCompositionShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
						newCompositionShape.SetCornerOffsetRatio(1.5)
					}
				} else {
					newShape.SetX(parentShape.GetX() + float64(len(*conf.sliceForNewAddedItem)-1)*parentShape.GetWidth()*1.2)
					newShape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)

					if len(*conf.sliceForNewCompositionShapes) > 0 {
						newCompositionShape := (*conf.sliceForNewCompositionShapes)[len(*conf.sliceForNewCompositionShapes)-1]
						newCompositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
						newCompositionShape.SetEndOrientation(ORIENTATION_VERTICAL)
						ratio := (newShape.GetY() - parentShape.GetY()) / parentShape.GetHeight()
						newCompositionShape.SetCornerOffsetRatio((ratio-1.0)/2.0 + 1.0)
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
