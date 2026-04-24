package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func addRenameButton[T AbstractType](at T, libraryNode *tree.Node, stager *Stager) {
	if !at.GetIsInRenameMode() {
		libraryNode.Buttons = append(libraryNode.Buttons,
			&tree.Button{
				Name: at.GetName() + " " + string(buttons.BUTTON_edit_note),
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
		libraryNode.Buttons = append(libraryNode.Buttons,
			&tree.Button{
				Name: at.GetName() + " " + string(buttons.BUTTON_edit_off),
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
