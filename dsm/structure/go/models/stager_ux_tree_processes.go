package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProcesses(
	process *Process,
	parentNode *tree.Node,
	processsWhoseNodeIsExpanded *[]*Process,
) {
	processNode := &tree.Node{
		Name:            process.GetName(),
		IsExpanded:      slices.Contains(*processsWhoseNodeIsExpanded, process),
		IsNodeClickable: true,
		IsInEditMode:    process.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, processNode)

	addRenameButton(process, processNode, stager)
	processNode.OnNameChange = stager.onNameChange(process)
	processNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, process, processsWhoseNodeIsExpanded)
	processNode.OnClick = onNodeClicked(stager, process)

	// Diagrams
	for _, diagramProcess := range process.DiagramProcesss {
		stager.treeDiagramProcess(process, diagramProcess, processNode)
	}

	confDiagramProcesss := ItemButtonConfiguration[
		DiagramProcess, *DiagramProcess,
		Process, *Process,
	]{
		parentNode:                         processNode,
		sliceForNewAddedItem:               &process.DiagramProcesss,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &process.IsExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagramProcesss)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	//
	// SubProcesses
	//
	confSubProcesses := ItemButtonConfiguration[
		Process, *Process,
		Process, *Process,
	]{
		parentNode:                         processNode,
		sliceForNewAddedItem:               &process.SubProcesses,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   processsWhoseNodeIsExpanded,
		parentElement:                      process,
	}
	addCreateItemButton(stager, confSubProcesses)

	// SubProcesses
	subProcessesNode := &tree.Node{
		Name:            "SubProcesses",
		FontStyle:       tree.ITALIC,
		IsExpanded:      process.IsSubProcessNodeExpanded,
		IsNodeClickable: true,
	}
	processNode.Children = append(processNode.Children, subProcessesNode)

	for _, process_ := range process.SubProcesses {
		stager.treeProcesses(process_, subProcessesNode, processsWhoseNodeIsExpanded)
	}
}
