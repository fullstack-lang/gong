// generated code - do not edit
package probe

import (

	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
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
	gongtreeStage *gongtree_models.Stage,
	stageButton, front *gongtree_models.Button) {

	// log.Println("ButtonImplRefresh: ButtonUpdated")

	buttonImpl.probe.Refresh()
}
