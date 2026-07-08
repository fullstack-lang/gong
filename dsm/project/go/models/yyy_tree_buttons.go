// generated code (do not edit)
package models

import (
	"encoding/base64"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
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
	IsButtonInMenu                     bool
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
	if conf.IsButtonInMenu {
		if conf.parentNode.Menu == nil {
			conf.parentNode.Menu = &tree.Menu{Name: "Add"}
		}
		conf.parentNode.Menu.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Menu.Buttons...)
	} else {
		conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)
	}

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
	if conf.IsButtonInMenu {
		if conf.parentNode.Menu == nil {
			conf.parentNode.Menu = &tree.Menu{Name: "Add"}
		}
		conf.parentNode.Menu.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Menu.Buttons...)
	} else {
		conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)
	}

	addButton.OnClick = func() {
		newAbstractElement := processAbstractItemAddition(stager, conf.ItemButtonConfiguration, callbacks)

		if conf.receivingDiagram != nil && conf.sliceForNewAddedShape != nil {
			newShapeToDiagram(newAbstractElement, conf.receivingDiagram, conf.sliceForNewAddedShape, stager, addButton.ClientOnY)
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
		LayoutConcreteType
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
	if conf.IsButtonInMenu {
		if conf.parentNode.Menu == nil {
			conf.parentNode.Menu = &tree.Menu{Name: "Add"}
		}
		conf.parentNode.Menu.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Menu.Buttons...)
	} else {
		conf.parentNode.Buttons = append([]*tree.Button{addButton}, conf.parentNode.Buttons...)
	}

	addButton.OnClick = func() {
		newAbstractElement := processAbstractItemAddition(stager, conf.ItemButtonConfiguration, callbacks)

		if conf.receivingDiagram != nil && conf.sliceForNewAddedShape != nil {
			newShape := newShapeToDiagram(newAbstractElement, conf.receivingDiagram, conf.sliceForNewAddedShape, stager, addButton.ClientOnY)

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

				if parentShape.GetConcreteLayoutDirection() == Horizontal {
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

func addRenameButton[T AbstractType](at T, node *tree.Node, stager *Stager) {
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}
	if !at.GetIsInRenameMode() {
		node.Menu.Buttons = append(node.Menu.Buttons,
			&tree.Button{
				Name: "Rename",
				Icon: string(buttons.BUTTON_edit_note),
				OnClick: func() {
					at.SetIsInRenameMode(true)
					stager.stage.Commit()
				},
				HasToolTip:      true,
				ToolTipText:     "Rename the " + GetGongstructNameFromPointer(at),
				ToolTipPosition: tree.Above,
			})
	} else {
		node.Menu.Buttons = append(node.Menu.Buttons,
			&tree.Button{
				Name: "Cancel rename",
				Icon: string(buttons.BUTTON_edit_off),
				OnClick: func() {
					at.SetIsInRenameMode(false)
					stager.stage.Commit()
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}
}

// ---------------------------------------------------------
// 1. BASE CONFIGURATION (Abstract elements only)
// ---------------------------------------------------------

type TreeNodeConfiguration[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	},
] struct {
	diagram                     DiagramType
	parentNode                  *tree.Node
	element                     AT
	parentElement               AT
	elementsWhoseNodeIsExpanded *[]AT
}

// ---------------------------------------------------------
// 2. MIDDLE CONFIGURATION (Abstract + Shape)
// ---------------------------------------------------------

type TreeNodeAndShapeConfiguration[
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
	CT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	},
] struct {
	TreeNodeConfiguration[AT, AT_, DiagramType]
	shapes    *[]CT
	shapesMap map[AT]CT
}

// ---------------------------------------------------------
// 3. COMPLEX CONFIGURATION (Abstract + Shape + Link)
// ---------------------------------------------------------

type TreeNodeShapeAndLinkConfiguration[
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
	CT_ Gongstruct,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	},
] struct {
	TreeNodeAndShapeConfiguration[AT, AT_, CT, CT_, DiagramType]
	map_Element_CompositionShape map[AT]ACT
	compositionShapes            *[]ACT
}

// ---------------------------------------------------------
// 4. CONFIGURATION WITHOUT LINK
// ---------------------------------------------------------

type TreeNodeAndShapeConfigurationWithoutLink[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	ParentAT interface {
		*ParentAT_
		AbstractType
	},
	ParentAT_ Gongstruct,
	CT interface {
		*CT_
		RectShapeInterface
		ConcreteType
	},
	CT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	},
] struct {
	diagram                     DiagramType
	parentNode                  *tree.Node
	element                     AT
	parentElement               ParentAT
	elementsWhoseNodeIsExpanded *[]AT
	shapes                      *[]CT
	shapesMap                   map[AT]CT
}

func createBaseNode[
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
	CT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	}](
	stager *Stager,
	diagram DiagramType,
	element AT,
	parentNode *tree.Node,
	elementsWhoseNodeIsExpanded *[]AT,
	shapesMap map[AT]CT,
) *tree.Node {
	stage := stager.stage
	node := &tree.Node{
		Name: element.GetName(),

		IsExpanded: slices.Contains(*elementsWhoseNodeIsExpanded, element),

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.GetIsChecked(),

		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Above,
		CheckboxToolTipText:     "Add " + GetGongstructNameFromPointer(element) + " to diagram",

		IsNodeClickable: true,

		IsInEditMode: element.GetIsInRenameMode(),
	}
	if diagram.GetIsShowPrefix() {
		node.IsWithPrefix = true
		node.Prefix = element.GetComputedPrefix()
	}
	parentNode.Children = append(parentNode.Children, node)

	// if the element is not in rename mode, then we can add a button to enter rename mode
	addRenameButton(element, node, stager)

	// if the element is in the diagram, then we can add a button to remove it from the diagram
	if shape, ok := shapesMap[element]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.Name = "Show"
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	// add a button to have the list of other diagrams where the element is present
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}

	addAllChildrenButton := &tree.Button{
		Name:            "Add item & all children to diagram",
		Icon:            string(buttons.BUTTON_add_circle),
		ToolTipText:     "Add element and all children to diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			// Suspend callbacks to avoid multiple commits to the UI
			tmp1 := stager.stage.OnInitCommitFromBackCallback
			tmp2 := stager.stage.beforeCommitHooks
			tmp3 := stager.stage.afterCommitHooks
			stager.stage.OnInitCommitFromBackCallback = nil
			stager.stage.beforeCommitHooks = nil
			stager.stage.afterCommitHooks = nil

			var toggleNodeAndChildren func(n *tree.Node, isChecked bool)
			toggleNodeAndChildren = func(n *tree.Node, isChecked bool) {
				if n.HasCheckboxButton && n.IsChecked != isChecked {
					frontNode := *n
					frontNode.IsChecked = isChecked
					if n.OnIsCheckedChanged != nil {
						n.OnIsCheckedChanged(isChecked)
					}
				}
				for _, child := range n.Children {
					toggleNodeAndChildren(child, isChecked)
				}
			}
			toggleNodeAndChildren(node, true)

			// Restore callbacks and perform final commit
			stager.stage.OnInitCommitFromBackCallback = tmp1
			stager.stage.beforeCommitHooks = tmp2
			stager.stage.afterCommitHooks = tmp3
			stager.stage.Commit()
		},
	}

	removeAllChildrenButton := &tree.Button{
		Name:            "Remove item & all children from diagram",
		Icon:            string(buttons.BUTTON_remove_circle),
		ToolTipText:     "Remove element and all children from diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			// Suspend callbacks to avoid multiple commits to the UI
			tmp1 := stager.stage.OnInitCommitFromBackCallback
			tmp2 := stager.stage.beforeCommitHooks
			tmp3 := stager.stage.afterCommitHooks
			stager.stage.OnInitCommitFromBackCallback = nil
			stager.stage.beforeCommitHooks = nil
			stager.stage.afterCommitHooks = nil

			var toggleNodeAndChildren func(n *tree.Node, isChecked bool)
			toggleNodeAndChildren = func(n *tree.Node, isChecked bool) {
				if n.HasCheckboxButton && n.IsChecked != isChecked {
					frontNode := *n
					frontNode.IsChecked = isChecked
					if n.OnIsCheckedChanged != nil {
						n.OnIsCheckedChanged(isChecked)
					}
				}
				for _, child := range n.Children {
					toggleNodeAndChildren(child, isChecked)
				}
			}
			toggleNodeAndChildren(node, false)

			// Restore callbacks and perform final commit
			stager.stage.OnInitCommitFromBackCallback = tmp1
			stager.stage.beforeCommitHooks = tmp2
			stager.stage.afterCommitHooks = tmp3
			stager.stage.Commit()
		},
	}

	node.Menu.Buttons = append(node.Menu.Buttons, addAllChildrenButton, removeAllChildrenButton)

	diagrams := stager.map_Element_Diagrams[any(element).(AbstractType)]
	if len(diagrams) > 1 {
		if diagram.GetDiagramListElement() == any(element).(AbstractType) {
			node.IsExpanded = true
			diagramsButton := &tree.Button{
				Name:            "Hide diagrams",
				Icon:            string(buttons.BUTTON_list),
				ToolTipText:     "List of other " + fmt.Sprint(len(diagrams)-1) + " diagrams where element is present",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.SetDiagramListElement(nil)
					stage.Commit()
				},
			}
			if node.Menu == nil {
				node.Menu = &tree.Menu{Name: "Menu"}
			}
			node.Menu.Buttons = append(node.Menu.Buttons, diagramsButton)

			for _, diag := range diagrams {
				if any(diag) == any(diagram) {
					continue
				}
				diagramNode := &tree.Node{
					Name:            diag.GetName(),
					ToolTipText:     "Go to diagram \"" + diag.GetName() + "\"",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					IsNodeClickable: true,
					OnClick: func(frontNode *tree.Node) {
						for diagram_ := range *GetGongstructInstancesSetFromPointerType[DiagramType](stager.stage) {
							diagram_.SetIsChecked(false)
						}
						diag.SetIsChecked(true)
						diagram.SetDiagramListElement(nil)
						stage.Commit()
					},
				}
				node.Children = append(node.Children, diagramNode)
			}
		} else {
			if node.Menu == nil {
				node.Menu = &tree.Menu{Name: "Menu"}
			}
			node.Menu.Buttons = append(node.Menu.Buttons,
				&tree.Button{
					Name:            "Show diagrams",
					Icon:            string(buttons.BUTTON_list),
					ToolTipText:     "Show list of other diagrams where \"" + element.GetName() + "\" is present",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnClick: func() {
						diagram.SetDiagramListElement(any(element).(AbstractType))
						stage.Commit()
					},
				})
		}
	}

	return node
}

func addNodeToTree[
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
	CT_ Gongstruct,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	}](
	stager *Stager,
	conf TreeNodeShapeAndLinkConfiguration[AT, AT_, CT, CT_, ACT, ACT_, DiagramType],
) *tree.Node {
	stage := stager.stage
	node := createBaseNode(
		stager,
		conf.diagram,
		conf.element,
		conf.parentNode,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapesMap,
	)

	_, isParentInDiagram := conf.shapesMap[conf.parentElement]
	_, isChildInDiagram := conf.shapesMap[conf.element]
	var compositionShape ACT
	compositionShape, _ = conf.map_Element_CompositionShape[conf.element]

	// if the parent element is in the diagram and the child element is in the diagram,
	// then we can add a checkbox to add/remove the link between the parent and child element to/from the diagram
	if conf.parentElement != nil && isParentInDiagram && isChildInDiagram {
		var toggleLinkButton *tree.Button
		if _, ok := conf.map_Element_CompositionShape[conf.element]; ok {
			node.IsSecondCheckboxChecked = true
			toggleLinkButton = &tree.Button{
				Name:            "Remove link shape",
				Icon:            string(buttons.BUTTON_link_off),
				ToolTipText:     "Remove link shape \"" + conf.parentElement.GetName() + "\" to \"" + conf.element.GetName() + "\" to diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					frontNode := *node
					frontNode.IsSecondCheckboxChecked = false
					if node.OnIsSecondCheckboxCheckedChanged != nil {
						node.OnIsSecondCheckboxCheckedChanged(frontNode.IsSecondCheckboxChecked)
					}
				},
			}
		} else {
			toggleLinkButton = &tree.Button{
				Name:            "Add link shape",
				Icon:            string(buttons.BUTTON_link),
				ToolTipText:     "Add link shape \"" + conf.parentElement.GetName() + "\" to \"" + conf.element.GetName() + "\" to diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					frontNode := *node
					frontNode.IsSecondCheckboxChecked = true
					if node.OnIsSecondCheckboxCheckedChanged != nil {
						node.OnIsSecondCheckboxCheckedChanged(frontNode.IsSecondCheckboxChecked)
					}
				},
			}
		}
		if node.Menu == nil {
			node.Menu = &tree.Menu{Name: "Menu"}
		}
		node.Menu.Buttons = append(node.Menu.Buttons, toggleLinkButton)

		if compositionShape != nil {
			showHideCompositionButton := &tree.Button{
				Name:            "Hide link",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					compositionShape.SetIsHidden(!compositionShape.GetIsHidden())
					stage.Commit()
				},
			}

			if compositionShape.GetIsHidden() {
				_ = compositionShape
				showHideCompositionButton.Name = "Show link"
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
				showHideCompositionButton.ToolTipText = "Show link from \"" + conf.parentElement.GetName() +
					"\" to \"" + conf.element.GetName() + "\""
			} else {
				showHideCompositionButton.Name = "Hide link"
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
				showHideCompositionButton.ToolTipText = "Hide link from \"" + conf.parentElement.GetName() +
					"\" to \"" + conf.element.GetName() + "\""
			}
			if node.Menu == nil {
				node.Menu = &tree.Menu{Name: "Menu"}
			}
			node.Menu.Buttons = append(node.Menu.Buttons, showHideCompositionButton)
		}
	}

	//

	// what to do when the node is clicked
	setCallbacksElementInDiagram(
		stager,
		node,
		conf.diagram,
		conf.element,
		conf.parentElement,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapes,
		conf.shapesMap,
		compositionShape,
		conf.compositionShapes,
	)

	return node
}

func addNodeToTreeWithoutLink[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	ParentAT interface {
		*ParentAT_
		AbstractType
	},
	ParentAT_ Gongstruct,
	CT interface {
		*CT_
		RectShapeInterface
		ConcreteType
	},
	CT_ Gongstruct,
	DiagramType interface {
		DiagramIF
		AbstractType
		comparable
	}](
	stager *Stager,
	conf TreeNodeAndShapeConfigurationWithoutLink[AT, AT_, ParentAT, ParentAT_, CT, CT_, DiagramType],
) *tree.Node {
	node := createBaseNode(
		stager,
		conf.diagram,
		conf.element,
		conf.parentNode,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapesMap,
	)

	// what to do when the node is clicked
	setCallbacksElementInDiagramWithoutLink(
		stager,
		node,
		conf.diagram,
		conf.element,
		conf.parentElement,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapes,
		conf.shapesMap,
	)

	return node
}

func (stager *Stager) GetPkgName() string {
	pkgPath := stager.stage.MetaPackageImportPath
	pkgName := ""
	parts := strings.Split(pkgPath, "/")
	if len(parts) >= 3 {
		pkgName = parts[len(parts)-3]
	}
	return pkgName
}

func (stager *Stager) button() {
	buttonStage := stager.buttonStage
	buttonStage.Reset()

	layout := new(button.Layout)

	group1 := new(button.Group)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Stop",
		Icon:  string(buttons.BUTTON_stop_circle),
		Label: "Stop",
		OnClick: func() {
			log.Println("Stop")
			os.Exit(0)
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:    "Web",
		Icon:    string(buttons.BUTTON_web),
		Label:   "Web",
		OnClick: stager.exportWebsite,
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Go",
		Icon:  string(buttons.BUTTON_file_download),
		Label: "Go",
		OnClick: func() {
			log.Println("Exporting stage as Go file")

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload)

			if stager.fileName == "" {
				pkgPath := stager.stage.MetaPackageImportPath
				pkgName := ""
				parts := strings.Split(pkgPath, "/")
				if len(parts) >= 3 {
					pkgName = parts[len(parts)-3]
				}
				stager.fileName = pkgName + "-" + stager.stage.GetName() + ".go"
			}

			prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

			stageString, err := stager.stage.MarshallToString(stager.stage.MetaPackageImportPath, "main")
			if err != nil {
				log.Println("Error serializing stage: " + err.Error())
				return
			}

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))

			load.StageBranch(stager.loadStage, fileToDownload)
			stager.loadStage.Commit()

			time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
			stager.load()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export launcher HTML",
		Icon:  string(buttons.BUTTON_launch),
		Label: "Export launcher HTML",
		OnClick: func() {
			log.Println("Exporting launcher HTML")

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload)

			pkgName := stager.GetPkgName()

			if stager.fileName == "" {
				stager.fileName = pkgName + "-" + stager.stage.GetName() + ".go"
			}

			prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName + ".html"

			stageString, err := stager.stage.MarshallToString(stager.stage.MetaPackageImportPath, "main")
			if err != nil {
				log.Println("Error serializing stage: " + err.Error())
				return
			}

			// Escape the Go string for safe inclusion in a JavaScript template literal
			jsGoCode := strings.ReplaceAll(stageString, "\\", "\\\\")
			jsGoCode = strings.ReplaceAll(jsGoCode, "`", "\\`")
			jsGoCode = strings.ReplaceAll(jsGoCode, "${", "\\${")
			jsGoCode = strings.ReplaceAll(jsGoCode, "</script>", "<\\/script>")

			htmlString := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Launch ` + pkgName + ` App</title>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; margin-top: 50px; }
        button { padding: 10px 20px; font-size: 16px; cursor: pointer; }
    </style>
</head>
<body>

    <button id="launchBtn">Open App and Process ` + cleanFileName + `</button>

    <script>
        const rawGoCode = ` + "`" + jsGoCode + "`" + `;
        
        // Base64 encode the UTF-8 string to pass it as fileData
        const utf8Bytes = new TextEncoder().encode(rawGoCode);
        let binary = '';
        for (let i = 0; i < utf8Bytes.length; i++) {
            binary += String.fromCharCode(utf8Bytes[i]);
        }
        const fileXContent = btoa(binary);
        
        const fileXName = "` + cleanFileName + `";
        const targetUrl = "` + pkgName + `-app-portable.html";
        const targetOrigin = "*"; 

        document.getElementById('launchBtn').addEventListener('click', () => {
            const appWindow = window.open(targetUrl, '_blank');
            const payload = {
                action: 'PROCESS_INJECTED_FILE',
                fileName: fileXName,
                fileData: fileXContent
            };

            let ackReceived = false;
            window.addEventListener('message', (event) => {
                if (event.data && event.data.action === 'FILE_PROCESSED') {
                    ackReceived = true;
                    console.log("Application successfully received and processed the file!");
                }
            });

            const sendInterval = setInterval(() => {
                if (ackReceived || appWindow.closed) {
                    clearInterval(sendInterval);
                    return;
                }
                appWindow.postMessage(payload, targetOrigin);
                console.log("File payload sent to application (waiting for ACK)...");
            }, 1000);
        });
    </script>
</body>
</html>`

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(htmlString))

			load.StageBranch(stager.loadStage, fileToDownload)
			stager.loadStage.Commit()

			time.Sleep(1 * time.Second)
			stager.load()
		},
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}

type FileToUploadProxy struct {
	stager *Stager
}

func (stager *Stager) load() {
	stager.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &loadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your <library>.go file here or ",
	}

	message.Stage(stager.loadStage)

	stager.loadStage.Commit()
}

type loadProxy struct {
	stager *Stager
}

func (proxy *loadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fmt.Println("OnFileUpload: start")
	proxy.stager.fileName = uploadedFile.GetName()

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	// if the user loads a second file, we don't want the previous file to be committed
	proxy.stager.stage.OnInitCommitCallback = nil
	proxy.stager.createViews()

	proxy.stager.stage.Reset()
	fmt.Println("OnFileUpload: after reset")
	ParseAstFromBytes(proxy.stager.stage, decodedBytes)
	fmt.Println("OnFileUpload: after parse")
	proxy.stager.stage.Commit()
	fmt.Println("OnFileUpload: after commit")

	return nil
}

func ParseAstFromBytes(stage *Stage, input []byte) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", input, parser.ParseComments)
	if errParser != nil {
		return fmt.Errorf("Unable to parse: %w", errParser)
	}
	return ParseAstFileFromAst(stage, inFile, fset, false)
}

// onNameChange provides a reusable callback for tree.Node.OnNameChange
func (stager *Stager) onNameChange(element AbstractType) func(newName string) {
	return func(newName string) {
		element.SetName(newName)
		element.SetIsInRenameMode(false)
		stager.stage.Commit()
	}
}

func onNodeClicked[T AbstractType](stager *Stager, element T) func(frontNode *tree.Node) {
	return func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[T]())
		stager.stage.Commit()
	}
}

// onIsExpandedChangeSlice provides a reusable callback for tree.Node.OnIsExpandedChange backed by a slice
func onIsExpandedChangeSlice[T comparable](stager *Stager, element T, expandedSlice *[]T) func(isExpanded bool) {
	return func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*expandedSlice, element) == -1 {
				*expandedSlice = append(*expandedSlice, element)
			}
		} else {
			if idx := slices.Index(*expandedSlice, element); idx != -1 {
				*expandedSlice = slices.Delete(*expandedSlice, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}
}

// onIsExpandedChangeBool provides a reusable callback for tree.Node.OnIsExpandedChange backed by a boolean
func (stager *Stager) onIsExpandedChangeBool(isExpandedPtr *bool) func(isExpanded bool) {
	return func(isExpanded bool) {
		*isExpandedPtr = isExpanded
		log.Println("onIsExpandedChangeBool, New value", *isExpandedPtr)
		stager.stage.CommitWithSuspendedCallbacks() // important otherwise, the front is overwelmed when there is a Search and Jump" node expansion
	}
}
