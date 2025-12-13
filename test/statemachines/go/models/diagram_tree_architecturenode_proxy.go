package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type diagramArchitectureNodeProxy struct {
	stager         *Stager
	architecture   *Architecture
	isNodeExpanded *bool
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *diagramArchitectureNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	p.stager.probeForm.FillUpFormFromGongstruct(p.architecture, "Architecture")

	p.stager.stage.Commit()
}
