package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowWithinLibrary(
	library *Library,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	dataFlowNode := &tree.Node{
		Name:            dataFlow.GetName(),
		IsExpanded:      slices.Contains(library.DataFlowsWhoseNodeIsExpanded, dataFlow),
		IsNodeClickable: true,
		IsInEditMode:    dataFlow.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	dataFlowNode.OnNameChange = stager.onNameChange(dataFlow)
	dataFlowNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataFlow, &library.DataFlowsWhoseNodeIsExpanded)
	dataFlowNode.OnClick = onNodeClicked(stager, dataFlow)

	//
	// Data
	//
	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      dataFlow.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	dataFlowNode.Children = append(dataFlowNode.Children, datasNode)
	datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&dataFlow.IsDatasNodeExpanded)

	for _, data := range dataFlow.Datas {
		stager.treeData(library, data, datasNode)
	}
}
