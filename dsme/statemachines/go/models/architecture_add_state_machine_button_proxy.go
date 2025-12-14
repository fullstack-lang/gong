package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ArchitectureAddDiagramButtonProxy struct {
	stager      *Stager
	architecure *Architecture
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *ArchitectureAddDiagramButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	s := p.stager.stage
	newDiagram := (&StateMachine{
		Name: "New StateMachine",
	}).Stage(s)

	p.architecure.StateMachines = append(p.architecure.StateMachines, newDiagram)
	p.stager.stage.Commit()
}
