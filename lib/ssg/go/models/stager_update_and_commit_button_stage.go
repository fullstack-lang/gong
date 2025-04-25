package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitButtonStage() {
	stager.buttonStage.Reset()

	layout := new(button.Layout).Stage(stager.buttonStage)

	group1 := new(button.Group).Stage(stager.buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	button := button.NewButton(
		// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
		stager,
		"Generate static site",
		string(buttons.BUTTON_web_asset),
		"Generate static site",
	)

	group1.Buttons = append(group1.Buttons, button)

	stager.buttonStage.Commit()
}
