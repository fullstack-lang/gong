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
	processNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			process.SetName(frontNode.Name)
			process.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(*processsWhoseNodeIsExpanded, process) == -1 {
					*processsWhoseNodeIsExpanded = append(*processsWhoseNodeIsExpanded, process)
				}
			} else {
				if idx := slices.Index(*processsWhoseNodeIsExpanded, process); idx != -1 {
					*processsWhoseNodeIsExpanded = slices.Delete(*processsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(process, GetPointerToGongstructName[*Process]())
		stager.stage.Commit()
	}

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
		parentNodeExpansionBooleanValue:    &process.isExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagramProcesss)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.isExpanded = true
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
