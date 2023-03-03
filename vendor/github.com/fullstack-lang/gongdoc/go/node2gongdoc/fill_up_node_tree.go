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
	nodeCb.updateNodesStates(&gongdoc_models.Stage)

	// set callbacks on node updates
	gongdoc_models.Stage.OnAfterNodeUpdateCallback = nodeCb
	gongdoc_models.Stage.OnAfterNodeCreateCallback = nodeCb
	gongdoc_models.Stage.OnAfterNodeDeleteCallback = nodeCb
}
