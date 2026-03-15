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
	stage := stager.stage
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
						stage.Commit()
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
						stage.Commit()
					},
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}

	if shape, ok := (*shapesMap)[element]; ok {
		node.IsChecked = true
		node.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
					shape.SetIsHidden(!shape.GetIsHidden())
					stage.Commit()
				},
			},
		}
		if shape.GetIsHidden() {
			node.Buttons[0].Icon = string(buttons.BUTTON_visibility)
			node.Buttons[0].ToolTipText = "Show on diagram"
		}
	}

	// what to do when the node is clicked
	node.OnUpdate = onUpdateElementInDiagram(
		stager,
		diagram,
		element,
		elementsWhoseNodeIsExpanded,
		shapes,
		shapesMap)

	return node
}
