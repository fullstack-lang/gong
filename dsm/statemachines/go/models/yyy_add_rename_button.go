// generated code (do not edit)
package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

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
