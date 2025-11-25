package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type DiagramAddStateButtonProxy struct {
	stager       *Stager
	stateMachine *StateMachine
	diagram      *Diagram
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *DiagramAddStateButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	s := p.stager.stage
	newState := new(State).Stage(s)

	newState.Name = "New State"
	p.stateMachine.States = append(p.stateMachine.States, newState)

	newStateShapeToDiagram(newState, p.diagram).Stage(p.stager.stage)

	p.stager.stage.Commit()
}
