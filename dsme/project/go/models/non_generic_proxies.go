package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type _ProjectNodeProxy struct {
	stager  *Stager
	node    *tree.Node
	project *Project
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *_ProjectNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		p.project.IsExpanded = !p.project.IsExpanded
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.project, GetPointerToGongstructName[*Product]())

	p.stager.stage.Commit()
}

type _ProductNodeProxy struct {
	stager  *Stager
	node    *tree.Node
	product *Product
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *_ProductNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		p.product.IsExpanded = !p.product.IsExpanded
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.product, GetPointerToGongstructName[*Product]())

	p.stager.stage.Commit()
}
