package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type diagramStateMachineNodeProxy struct {
	stager         *Stager
	stateMachine   *StateMachine
	isNodeExpanded *bool
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *diagramStateMachineNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		*p.isNodeExpanded = !*p.isNodeExpanded
		p.stager.stage.Commit()
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.stateMachine, "State Machine")

	p.stager.stage.Commit()
}
