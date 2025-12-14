package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type DiagramCopyButtonProxy struct {
	stager  *Stager
	diagram *Diagram
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *DiagramCopyButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	s := p.stager.stage
	newDiagram := new(Diagram).Stage(s)

	if stateMachine, ok := p.stager.map_diagram_stateMachine[p.diagram]; ok {
		stateMachine.Diagrams = append(stateMachine.Diagrams, newDiagram)
	}

	newDiagram.Name = p.diagram.Name + " copy"

	for _, stateShape_ := range p.diagram.State_Shapes {
		stateShape := new(StateShape).Stage(s)
		stateShape.State = stateShape_.State
		stateShape.RectShape.X = stateShape_.RectShape.X
		stateShape.RectShape.Y = stateShape_.RectShape.Y
		stateShape.RectShape.Width = stateShape_.RectShape.Width
		stateShape.RectShape.Height = stateShape_.RectShape.Height

		newDiagram.State_Shapes = append(newDiagram.State_Shapes, stateShape)
	}

	// get all transitions
	for _, transitionShape_ := range p.diagram.Transition_Shapes {
		transitionShape := new(Transition_Shape).Stage(s)
		transitionShape.Transition = transitionShape_.Transition

		transitionShape.LinkShape.StartRatio = transitionShape_.LinkShape.StartRatio
		transitionShape.LinkShape.StartOrientation = transitionShape_.LinkShape.StartOrientation
		transitionShape.LinkShape.EndRatio = transitionShape_.LinkShape.EndRatio
		transitionShape.LinkShape.EndOrientation = transitionShape_.LinkShape.EndOrientation
		transitionShape.LinkShape.CornerOffsetRatio = transitionShape_.LinkShape.CornerOffsetRatio

		newDiagram.Transition_Shapes = append(newDiagram.Transition_Shapes, transitionShape)
	}

	p.stager.stage.Commit()
}
