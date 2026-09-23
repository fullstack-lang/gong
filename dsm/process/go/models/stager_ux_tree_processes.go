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
		IsButtonInMenu:                     true,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagramProcesss)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *stager.stage.GetInstancesSet[*DiagramProcess]() {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}
	if len(processNode.Menu.Buttons) > 0 {
		processNode.Menu.Buttons[0].Name = "Add Diagram Process"
		processNode.Menu.Buttons[0].ToolTipText = "Add a Diagram Process to \"" + process.GetName() + "\""
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
		IsButtonInMenu:                     true,
	}
	callbacksSubProcesses := addCreateItemButton(stager, confSubProcesses)
	callbacksSubProcesses.OnBeforeCommit = func() {
		process.IsSubProcessNodeExpanded = true
	}
	if len(processNode.Menu.Buttons) > 0 {
		processNode.Menu.Buttons[0].Name = "Add SubProcess"
		processNode.Menu.Buttons[0].ToolTipText = "Add a SubProcess to \"" + process.GetName() + "\""
	}

	// SubProcesses category node (only if items are present)
	if len(process.SubProcesses) > 0 {
		subProcessesNode := &tree.Node{
			Name:            "SubProcesses",
			FontStyle:       tree.ITALIC,
			IsExpanded:      process.IsSubProcessNodeExpanded,
			IsNodeClickable: true,
		}
		processNode.Children = append(processNode.Children, subProcessesNode)
		subProcessesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&process.IsSubProcessNodeExpanded)
		subProcessesNode.OnClick = onNodeClicked(stager, process)

		confSubProcessesNode := ItemButtonConfiguration[
			Process, *Process,
			Process, *Process,
		]{
			parentNode:                         subProcessesNode,
			sliceForNewAddedItem:               &process.SubProcesses,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &process.IsSubProcessNodeExpanded,
			parentElement:                      process,
			IsButtonInMenu:                     true,
		}
		callbacksSubProcessesNode := addCreateItemButton(stager, confSubProcessesNode)
		callbacksSubProcessesNode.OnBeforeCommit = func() {
			process.IsSubProcessNodeExpanded = true
		}
		if len(subProcessesNode.Menu.Buttons) > 0 {
			subProcessesNode.Menu.Buttons[0].Name = "Add SubProcess"
			subProcessesNode.Menu.Buttons[0].ToolTipText = "Add a SubProcess to \"" + process.GetName() + "\""
		}

		for _, process_ := range process.SubProcesses {
			stager.treeProcesses(process_, subProcessesNode, processsWhoseNodeIsExpanded)
		}
	}
}
