package models

import (
	"fmt"
	"log"
	"os"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) updateExportXLButtonStage() {

	stage := stager.buttonExportXLStage
	stage.Reset()

	layout := &button.Layout{}

	group := &button.Group{
		NbColumns: 2,
	}
	layout.Groups = append(layout.Groups, group)

	button_ := new(button.Button)
	button_.Name = "Export XL"
	button_.Label = "Export XL"
	button_.Icon = string(buttons.BUTTON_download)
	group.Buttons = append(group.Buttons, button_)

	proxy := &ExportXLButtonProxy{
		stager: stager,
	}

	button_.Proxy = proxy

	buttonKill := button.NewButton(
		&StopButtonProxy{
			stager: stager,
		},
		"Stop for maintenance",
		string(buttons.BUTTON_stop_circle),
		"Stop for maintenance",
	)

	group.Buttons = append(group.Buttons, buttonKill)

	button.StageBranch(stage, layout)

	stage.Commit()
}

type ExportXLButtonProxy struct {
	stager *Stager
}

// Updated implements models.ButtonProxyInterface.
func (proxy *ExportXLButtonProxy) Updated() {

	filename := fmt.Sprintf("./SC-DOC-7300-0016-MOSS-%s.xlsx", time.Now().Local().Format("2006-01-02-15-04-05"))

	SerializeStageMessageTypeTransitionEtat(proxy.stager.stage, filename)
}

type StopButtonProxy struct {
	stager *Stager
}

func (e *StopButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonExportXLStage
}

func (proxy *StopButtonProxy) OnAfterUpdateButton() {

	log.Println("Arret demandé de l'application")
	os.Exit(0)
}
