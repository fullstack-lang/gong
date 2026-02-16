package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) button() {
	buttonStage := stager.buttonStage
	buttonStage.Reset()

	layout := new(button.Layout).Stage(buttonStage)

	group1 := new(button.Group).Stage(buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	{
		buttonExportModifiedReqif := button.NewButton(
			// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
			&ResetReqifButtonProxy{
				stager: stager,
			},
			"Reset project file",
			string(buttons.BUTTON_reset_tv),
			"Reset project file",
		)

		group1.Buttons = append(group1.Buttons, buttonExportModifiedReqif)
	}

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}

type ResetReqifButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (proxy *ResetReqifButtonProxy) GetButtonsStage() *button.Stage {
	return proxy.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (proxy *ResetReqifButtonProxy) OnAfterUpdateButton() {

	proxy.stager.stage.Reset()
	proxy.stager.stage.Commit()
}
