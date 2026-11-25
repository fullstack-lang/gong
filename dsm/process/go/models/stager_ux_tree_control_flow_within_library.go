package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeControlFlowWithinLibrary(
	library *Library,
	controlFlow *ControlFlow,
	parentNode *tree.Node,
) {
	controlFlowNode := &tree.Node{
		Name:            controlFlow.GetName(),
		IsExpanded:      slices.Contains(library.ControlFlowsWhoseNodeIsExpanded, controlFlow),
		IsNodeClickable: true,
		IsInEditMode:    controlFlow.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, controlFlowNode)

	addRenameButton(controlFlow, controlFlowNode, stager)

	controlFlowNode.OnNameChange = stager.onNameChange(controlFlow)
	controlFlowNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, controlFlow, &library.ControlFlowsWhoseNodeIsExpanded)
	controlFlowNode.OnClick = onNodeClicked(stager, controlFlow)

	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      controlFlow.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	controlFlowNode.Children = append(controlFlowNode.Children, datasNode)
	datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&controlFlow.IsDatasNodeExpanded)

	for _, data := range topLevelDatasInSlice(controlFlow.Datas) {
		stager.treeData(library, data, datasNode)
	}
}
