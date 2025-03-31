package models

import (
	"log"

	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonImplRefresh struct {
	probe *Probe2
}

func NewButtonImplRefresh(
	probe *Probe2,
) (buttonImplRefresh *ButtonImplRefresh) {

	buttonImplRefresh = new(ButtonImplRefresh)
	buttonImplRefresh.probe = probe

	return
}

func (buttonImpl *ButtonImplRefresh) ButtonUpdated(
	gongtreeStage *gongtree_models.Stage,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplRefresh: ButtonUpdated")

	buttonImpl.probe.Refresh()
}
