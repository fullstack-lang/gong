// generated code - do not edit
package probe

import (
	"log"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplRefresh struct {
	probe *Probe
}

func NewButtonImplRefresh(
	probe *Probe,
) (buttonImplRefresh *ButtonImplRefresh) {

	buttonImplRefresh = new(ButtonImplRefresh)
	buttonImplRefresh.probe = probe

	return
}

func (buttonImpl *ButtonImplRefresh) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplRefresh: ButtonUpdated")

	buttonImpl.probe.Refresh()
}
