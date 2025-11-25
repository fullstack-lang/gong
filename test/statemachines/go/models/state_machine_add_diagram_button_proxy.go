package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SateMachineAddDiagramButtonProxy struct {
	stager       *Stager
	stateMachine *StateMachine
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *SateMachineAddDiagramButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	s := p.stager.stage
	newDiagram := (&Diagram{
		Name:        "New Diagram",
		IsEditable_: true,
	}).Stage(s)

	p.stateMachine.Diagrams = append(p.stateMachine.Diagrams, newDiagram)
	p.stager.stage.Commit()
}
