package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitRenderingTabButtonStage() {
	stager.renderingTabButtonStage.Reset()
	stage := stager.renderingTabButtonStage

	layout := new(button.Layout).Stage(stage)

	group1 := new(button.Group).Stage(stage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	buttonExportRenderingCong := button.NewButton(
		&ExportRenderingConfButtonProxy{
			stager: stager,
		},
		"Export Rendering Configuration",
		string(buttons.BUTTON_fact_check),
		"Export Rendering Configuration",
	)

	group1.Buttons = append(group1.Buttons, buttonExportRenderingCong)

	stage.Commit()
}

type ExportRenderingConfButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportRenderingConfButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.renderingTabButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportRenderingConfButtonProxy) OnAfterUpdateButton() {
	e.stager.reqifExporter.ExportRenderingConf(e.stager)
}
