// generated code - do not edit
package probe

import (
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonImplToggleDisplayAssociations struct {
	probe *Probe
}

func NewButtonImplToggleDisplayAssociations(
	probe *Probe,
) (buttonImplToggleDisplayAssociations *ButtonImplToggleDisplayAssociations) {

	buttonImplToggleDisplayAssociations = new(ButtonImplToggleDisplayAssociations)
	buttonImplToggleDisplayAssociations.probe = probe

	return
}

func (buttonImpl *ButtonImplToggleDisplayAssociations) ButtonUpdated(
	gongtreeStage *gongtree_models.Stage,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.probe.displayAssociationsSlicesWithNewLines =
		!buttonImpl.probe.displayAssociationsSlicesWithNewLines
}
