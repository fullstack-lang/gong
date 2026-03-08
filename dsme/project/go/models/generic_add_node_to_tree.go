package models

import (
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
	CT_ Gongstruct](
	stager *Stager,
	diagram *Diagram,
	parentNode *tree.Node,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap *map[AT]CT,
) *tree.Node {
	node := &tree.Node{
		Name: element.GetName(),

		IsExpanded: slices.Index(*elementsWhoseNodeIsExpanded, element) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add " + GetGongstructNameFromPointer(element) + " to diagram",

		IsNodeClickable: true,

		IsInEditMode: element.GetIsInRenameMode(),
	}
	if diagram.ShowPrefix {
		node.IsWithPrefix = true
		node.Prefix = element.GetComputedPrefix()
	}
	parentNode.Children = append(parentNode.Children, node)

	if !element.GetIsInRenameMode() {
		node.Buttons = append(node.Buttons,
			&tree.Button{
				Name: element.GetName() + " " + string(buttons.BUTTON_edit_note),
				Icon: string(buttons.BUTTON_edit_note),
				Impl: &tree.FunctionalButtonProxy{
					OnUpdated: func(stage *tree.Stage, button, updatedButton *tree.Button) {
						element.SetIsInRenameMode(true)
						stager.stage.Commit()
					},
				},
				HasToolTip:      true,
				ToolTipText:     "Rename the " + GetGongstructNameFromPointer(element),
				ToolTipPosition: tree.Above,
			})
	} else {
		node.Buttons = append(node.Buttons,
			&tree.Button{
				Name: element.GetName() + " " + string(buttons.BUTTON_edit_off),
				Icon: string(buttons.BUTTON_edit_off),
				Impl: &tree.FunctionalButtonProxy{
					OnUpdated: func(stage *tree.Stage, button, updatedButton *tree.Button) {
						element.SetIsInRenameMode(false)
						stager.stage.Commit()
					},
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}

	if _, ok := (*shapesMap)[element]; ok {
		node.IsChecked = true
	}

	// what to do when the node is clicked
	node.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: onUpdateElementInDiagram(
			stager,
			diagram,
			element,
			elementsWhoseNodeIsExpanded,
			shapes,
			shapesMap),
	}

	return node
}
