package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type expandableNodeProxy struct {
	stager         *Stager
	isNodeExpanded *bool
	node           *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *expandableNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		*p.isNodeExpanded = !*p.isNodeExpanded
	}

	p.stager.stage.Commit()
}
