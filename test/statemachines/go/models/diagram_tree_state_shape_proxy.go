package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Diagram_Tree_StateShapeProxy struct {
	stager     *Stager
	stateShape *StateShape
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *Diagram_Tree_StateShapeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {
	if frontNode.IsExpanded && !stagedNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		p.stateShape.IsExpanded = true
		p.stager.stage.Commit()
	}
	if !frontNode.IsExpanded && stagedNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		p.stateShape.IsExpanded = false
		p.stager.stage.Commit()
	}
}
