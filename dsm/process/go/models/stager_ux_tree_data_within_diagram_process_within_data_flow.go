package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataWithinDiagramProcessWithinDataFlow(
	dataFlowNode *tree.Node,
	diagramProcess *DiagramProcess,
	dataFlow *DataFlow,
	dataFlowShape *DataFlowShape,
	shapePresent bool,
	data *Data) {

	dataNode := &tree.Node{
		Name:            data.GetName(),
		IsExpanded:      slices.Contains(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, dataFlow),
		IsNodeClickable: true,
		IsInEditMode:    data.GetIsInRenameMode(),
	}
	dataFlowNode.Children = append(dataFlowNode.Children, dataNode)
	dataNode.OnClick = func(frontNode *tree.Node) {

	}

}
