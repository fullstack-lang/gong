package models

import (
	"log"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type DiagramTree_State_Proxy struct {
	stager     *Stager
	stateShape *StateShape
	state      *State
	diagram    *Diagram
}

func (p *DiagramTree_State_Proxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {

		stagedNode.IsChecked = frontNode.IsChecked

		// add the State_Shape
		if p.stateShape != nil {
			log.Fatalln("adding a shape to an already state shape")
		}

		stateShape := newStateShapeToDiagram(p.state, p.diagram).Stage(p.stager.stage)
		p.stateShape = stateShape

		p.stager.stage.Commit()
		return
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		stagedNode.IsChecked = frontNode.IsChecked

		// one need to remove the State_Shape
		if p.stateShape == nil {
			log.Fatalln("remove a non existing shape to state")
		}

		stateShape := p.stateShape
		p.stateShape = nil
		stateShape.Unstage(p.stager.stage)

		idx := slices.Index(p.diagram.State_Shapes, stateShape)
		p.diagram.State_Shapes = slices.Delete(p.diagram.State_Shapes, idx, idx+1)

		p.stager.stage.Commit()
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.state, "State")
}
