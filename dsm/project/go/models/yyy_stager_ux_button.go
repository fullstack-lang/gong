// generated code (do not edit)
package models

import (
	"log"
	"os"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) button() {
	buttonStage := stager.buttonStage
	buttonStage.Reset()

	layout := new(button.Layout)

	group1 := new(button.Group)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Stop for maintenance",
		Icon:  string(buttons.BUTTON_stop_circle),
		Label: "Stop for maintenance",
		OnClick: func() {
			log.Println("Stop")
			os.Exit(0)
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:    "Export website",
		Icon:    string(buttons.BUTTON_web),
		Label:   "Export website",
		OnClick: stager.exportWebsite,
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}
