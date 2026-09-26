package models

import (
	"log"
	"os"

	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitAnonymousButtonStage() {
	stager.anonymousButtonStage.Reset()
	stage := stager.anonymousButtonStage

	layout := new(button.Layout).Stage(stage)

	group1 := new(button.Group).Stage(stage)
	group1.Percentage = 100
	group1.NbColumns = 1
	layout.Groups = append(layout.Groups, group1)

	group1.Buttons = append(group1.Buttons,
		button.NewButton(
			&AnonymousButtonProxy{
				stager: stager,
			},
			"Export blanked version",
			string(buttons.BUTTON_shuffle),
			"Export blanked version",
		))

	group1.Buttons = append(group1.Buttons,
		button.NewButton(
			&LoadSampleButtonProxy{
				stager: stager,
			},
			"Load sample (collecting drone)",
			string(buttons.BUTTON_file_open),
			"Load sample (collecting drone)",
		))

	buttonKill := button.NewButton(
		&StopButtonProxy{
			stager: stager,
		},
		"Stop",
		string(buttons.BUTTON_stop_circle),
		"Stop",
	)

	group1.Buttons = append(group1.Buttons, buttonKill)

	stage.Commit()
}

type AnonymousButtonProxy struct {
	stager *Stager
}

func (e *AnonymousButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

func (e *AnonymousButtonProxy) OnAfterUpdateButton() {
	e.stager.reqifExporter.ExportAnonymousReqif(e.stager)
}

type LoadSampleButtonProxy struct {
	stager *Stager
}

func (e *LoadSampleButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

func (e *LoadSampleButtonProxy) OnAfterUpdateButton() {
	e.stager.LoadSampleProject()
}

type StopButtonProxy struct {
	stager *Stager
}

func (e *StopButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

func (proxy *StopButtonProxy) OnAfterUpdateButton() {
	log.Println("Arret demandé de l'application")
	os.Exit(0)
}
