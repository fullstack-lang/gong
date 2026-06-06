// generated code (do not edit)
package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

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
			Name:            "Remove",
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
		if node.Menu == nil {
			node.Menu = &tree.Menu{Name: "Menu"}
		}
		node.Menu.Buttons = append(node.Menu.Buttons, visibilityButton)
	}

	// add a button to have the list of other diagrams where the element is present
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
					OnUpdate: func(_ *tree.Stage, _, _ *tree.Node) {
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
		node.HasSecondCheckboxButton = true
		node.SecondCheckboxHasToolTip = true
		node.SecondCheckboxToolTipPosition = tree.Right

		if _, ok := conf.map_Element_CompositionShape[conf.element]; ok {
			node.CheckboxToolTipText = "Remove link shape \"" + conf.parentElement.GetName() +
				"\" to \"" + conf.element.GetName() + "\" to diagram"
			node.IsSecondCheckboxChecked = true
		} else {
			node.CheckboxToolTipText = "Add link shape \"" + conf.parentElement.GetName() +
				"\" to \"" + conf.element.GetName() + "\" to diagram"
		}

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
	node.OnUpdate = onUpdateElementInDiagram(
		stager,
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
	node.OnUpdate = onUpdateElementInDiagramWithoutLink(
		stager,
		conf.diagram,
		conf.element,
		conf.parentElement,
		conf.elementsWhoseNodeIsExpanded,
		conf.shapes,
		conf.shapesMap,
	)

	return node
}
