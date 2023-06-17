package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// common struct to node impl
// required for computing node states
type NodeImplGongObjectAbstract struct {
	diagramPackage    *gongdoc_models.DiagramPackage
	treeOfGongObjects *gongtree_models.Tree
}

func NewNodeImplGongObjectAbstract(
	diagramPackage *gongdoc_models.DiagramPackage,
	treeOfGongObjects *gongtree_models.Tree,
) (nodeImplGongObjectAbstract NodeImplGongObjectAbstract) {

	nodeImplGongObjectAbstract.diagramPackage = diagramPackage
	nodeImplGongObjectAbstract.treeOfGongObjects = treeOfGongObjects
	return
}
