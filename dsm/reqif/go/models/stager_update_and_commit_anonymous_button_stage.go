package models

import (
	"fmt"
	"log"
	"os"

	samples "github.com/fullstack-lang/gong/dsm/reqif/go/cmd/reqif/samples"

	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitAnonymousButtonStage() {
	stager.anonymousButtonStage.Reset()
	stage := stager.anonymousButtonStage

	layout := new(button.Layout).Stage(stage)

	group1 := new(button.Group).Stage(stage)
	group1.Percentage = 100
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
			string(buttons.BUTTON_shuffle),
			"Load sample (collecting drone)",
		))

	buttonKill := button.NewButton(
		&StopButtonProxy{
			stager: stager,
		},
		"Stop for maintenance",
		string(buttons.BUTTON_stop_circle),
		"Stop for maintenance",
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
	// 1. Load the reqifz file embedded in data samples
	content, svgImages, jpgImages, pngImages, err := extractReqifFromZip(samples.SampleReqIFz)
	if err != nil {
		fmt.Println("Error extracting reqif from zip:", err)
		return
	}

	e.stager.processReqifData(content, svgImages, jpgImages, pngImages, "ReqIF for Wheeled Tennis Ball Drone.reqifz")

	// 2. Load the rendering conf embedded in data samples
	stageForRenderingConf := NewStage("renderingConf")
	ParseAstFromBytes(stageForRenderingConf, samples.SampleRenderingConf)

	e.stager.processRenderingConf(stageForRenderingConf)
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
