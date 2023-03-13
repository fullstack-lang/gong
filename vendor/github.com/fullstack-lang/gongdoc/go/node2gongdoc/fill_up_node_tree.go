package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	nodeCb := new(NodeCB)
	nodeCb.diagramPackage = diagramPackage

	nodeCb.FillUpDiagramNodeTree(diagramPackage)
	nodeCb.FillUpTreeOfGongObjects()
	nodeCb.computeNodesConfiguration(diagramPackage.Stage_)

	// set callbacks on node updates
	diagramPackage.Stage_.OnAfterNodeUpdateCallback = nodeCb
	diagramPackage.Stage_.OnAfterNodeCreateCallback = nodeCb
	diagramPackage.Stage_.OnAfterNodeDeleteCallback = nodeCb
}
