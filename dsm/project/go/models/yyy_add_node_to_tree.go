// generated code (do not edit)
package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

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
	diagram DiagramType,
	parentNode *tree.Node,
	element AT,
	parentElement AT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
	map_Element_CompositionShape map[AT]ACT,
	compositionShapes *[]ACT,
) *tree.Node {
	stage := stager.stage
	node := &tree.Node{
		Name: element.GetName(),

		IsExpanded: slices.Index(*elementsWhoseNodeIsExpanded, element) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.GetIsChecked(),

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add " + GetGongstructNameFromPointer(element) + " to diagram",

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
			Name:            diagram.GetName(),
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
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	_, isParentInDiagram := shapesMap[parentElement]
	_, isChildInDiagram := shapesMap[element]
	var compositionShape ACT
	compositionShape, _ = map_Element_CompositionShape[element]

	// add a button to have the list of other diagrams where the element is present
	diagrams := stager.map_Element_Diagrams[element]
	if len(diagrams) > 1 {
		if diagram.GetElementWhoseDiagramListIsDisplayed() == element {
			node.IsExpanded = true
			diagramsButton := &tree.Button{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_list),
				ToolTipText:     "List of other " + fmt.Sprint(len(diagrams)-1) + " diagrams where element is present",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.SetElementWhoseDiagramListIsDisplayed(nil)
					stage.Commit()
				},
			}
			node.Buttons = append(node.Buttons, diagramsButton)

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
					OnUpdate: func(_ *tree.Stage, _, _ *tree.Node) {
						for diagram_ := range *GetGongstructInstancesSetFromPointerType[DiagramType](stager.stage) {
							diagram_.SetIsChecked(false)
						}
						diag.SetIsChecked(true)
						diagram.SetElementWhoseDiagramListIsDisplayed(nil)
						stage.Commit()
					},
				}
				node.Children = append(node.Children, diagramNode)
			}
		} else {
			node.Buttons = append(node.Buttons,
				&tree.Button{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_list),
					ToolTipText:     "Show list of other diagrams where \"" + element.GetName() + "\" is present",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnClick: func() {
						diagram.SetElementWhoseDiagramListIsDisplayed(element)
						stage.Commit()
					},
				})
		}
	}

	//

	// if the parent element is in the diagram and the child element is in the diagram,
	// then we can add a checkbox to add/remove the link between the parent and child element to/from the diagram
	if parentElement != nil && isParentInDiagram && isChildInDiagram {
		node.HasSecondCheckboxButton = true
		node.SecondCheckboxHasToolTip = true
		node.SecondCheckboxToolTipPosition = tree.Right

		if _, ok := map_Element_CompositionShape[element]; ok {
			node.CheckboxToolTipText = "Remove link shape \"" + parentElement.GetName() +
				"\" to \"" + element.GetName() + "\" to diagram"
			node.IsSecondCheckboxChecked = true
		} else {
			node.CheckboxToolTipText = "Add link shape \"" + parentElement.GetName() +
				"\" to \"" + element.GetName() + "\" to diagram"
		}

		if compositionShape != nil {
			showHideCompositionButton := &tree.Button{
				Name:            GetGongstructNameFromPointer(element) + " " + string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					compositionShape.SetIsHidden(!compositionShape.GetIsHidden())
					stage.Commit()
				},
			}

			if compositionShape.GetIsHidden() {
				_ = compositionShape
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
				showHideCompositionButton.ToolTipText = "Show link from \"" + parentElement.GetName() +
					"\" to \"" + element.GetName() + "\""
			} else {
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
				showHideCompositionButton.ToolTipText = "Hide link from \"" + parentElement.GetName() +
					"\" to \"" + element.GetName() + "\""
			}
			node.Buttons = append(node.Buttons, showHideCompositionButton)
		}
	}

	//

	// what to do when the node is clicked
	node.OnUpdate = onUpdateElementInDiagram(
		stager,
		diagram,
		element,
		parentElement,
		elementsWhoseNodeIsExpanded,
		shapes,
		shapesMap,
		compositionShape,
		compositionShapes,
	)

	return node
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

// addNodeToTreeWithConf is a wrapper around addNodeToTree that uses the TreeNodeShapeAndLinkConfiguration
func addNodeToTreeWithConf[
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
	return addNodeToTree(
		stager,
		conf.diagram,
		conf.parentNode,
		conf.element,
		conf.parentElement,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapes,
		conf.shapesMap,
		conf.map_Element_CompositionShape,
		conf.compositionShapes,
	)
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
	diagram DiagramType,
	parentNode *tree.Node,
	element AT,
	parentElement ParentAT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
) *tree.Node {
	stage := stager.stage
	node := &tree.Node{
		Name: element.GetName(),

		IsExpanded: slices.Index(*elementsWhoseNodeIsExpanded, element) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.GetIsChecked(),

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add " + GetGongstructNameFromPointer(element) + " to diagram",

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
			Name:            diagram.GetName(),
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
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	// add a button to have the list of other diagrams where the element is present
	diagrams := stager.map_Element_Diagrams[any(element).(AbstractType)]
	if len(diagrams) > 1 {
		if diagram.GetElementWhoseDiagramListIsDisplayed() == any(element).(AbstractType) {
			node.IsExpanded = true
			diagramsButton := &tree.Button{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_list),
				ToolTipText:     "List of other " + fmt.Sprint(len(diagrams)-1) + " diagrams where element is present",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.SetElementWhoseDiagramListIsDisplayed(nil)
					stage.Commit()
				},
			}
			node.Buttons = append(node.Buttons, diagramsButton)

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
					OnUpdate: func(_ *tree.Stage, _, _ *tree.Node) {
						for diagram_ := range *GetGongstructInstancesSetFromPointerType[DiagramType](stager.stage) {
							diagram_.SetIsChecked(false)
						}
						diag.SetIsChecked(true)
						diagram.SetElementWhoseDiagramListIsDisplayed(nil)
						stage.Commit()
					},
				}
				node.Children = append(node.Children, diagramNode)
			}
		} else {
			node.Buttons = append(node.Buttons,
				&tree.Button{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_list),
					ToolTipText:     "Show list of other diagrams where \"" + element.GetName() + "\" is present",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnClick: func() {
						diagram.SetElementWhoseDiagramListIsDisplayed(any(element).(AbstractType))
						stage.Commit()
					},
				})
		}
	}

	// what to do when the node is clicked
	node.OnUpdate = onUpdateElementInDiagramWithoutLink(
		stager,
		diagram,
		element,
		parentElement,
		elementsWhoseNodeIsExpanded,
		shapes,
		shapesMap,
	)

	return node
}

func addNodeToTreeWithoutLinkWithConf[
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
	return addNodeToTreeWithoutLink(
		stager,
		conf.diagram,
		conf.parentNode,
		conf.element,
		conf.parentElement,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapes,
		conf.shapesMap,
	)
}
