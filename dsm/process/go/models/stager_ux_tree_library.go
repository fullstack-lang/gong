package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpandedTmp,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.rootLibrary {
		addRenameButton(library, libraryNode, stager)
	}

	libraryNode.OnUpdate = stager.OnUpdateLibrary(library)

	//
	// Processes
	//
	processesNode := &tree.Node{
		Name:            "Processes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsProcessesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, processesNode)
	processesNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != library.IsProcessesNodeExpanded {
			library.IsProcessesNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	// add a process to the library button
	confRootProcesses := ItemButtonConfiguration[
		Process, *Process,
		Library, *Library,
	]{
		parentNode:                         processesNode,
		sliceForNewAddedItem:               &library.RootProcesses,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsProcessesNodeExpanded,
	}
	addCreateItemButton(stager, confRootProcesses)

	for _, process := range library.RootProcesses {
		stager.treeProcesses(process, processesNode, &library.ProcesssWhoseNodeIsExpanded)
	}

	//
	// SubLibraries
	//
	subLibrariesNode := &tree.Node{
		Name:            "Sub Libraries",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsSubLibrariesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, subLibrariesNode)
	subLibrariesNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != library.IsSubLibrariesNodeExpanded {
			library.IsSubLibrariesNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(subLibrary, &subLibrariesNode.Children)
	}

	// add sub library button
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         subLibrariesNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSubLibrariesNodeExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	//
	// Data Flows
	//
	dataFlowNodes := &tree.Node{
		Name:            "Data Flows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsDataFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, dataFlowNodes)
	dataFlowNodes.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != library.IsDataFlowsNodeExpanded {
			library.IsDataFlowsNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	for _, dataFlow := range library.RootDataFlows {
		stager.treeDataFlowWithinLibrary(library, dataFlow, dataFlowNodes)
	}

	//
	// Data
	//
	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, datasNode)
	datasNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != library.IsDatasNodeExpanded {
			library.IsDatasNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	for _, data := range library.RootDatas {
		stager.treeData(library, data, datasNode)
	}

	// add data button
	confData := ItemButtonConfiguration[
		Data, *Data,
		Data, *Data,
	]{
		parentNode:                         datasNode,
		sliceForNewAddedItem:               &library.RootDatas,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
	}
	addCreateItemButton(stager, confData)
}

// Helper callbacks

func (stager *Stager) OnUpdateLibrary(library *Library) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			library.IsExpandedTmp = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			library.Name = frontNode.Name
			library.isInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(library, GetPointerToGongstructName[*Library]())
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			*isExpanded = !*isExpanded
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateDiagram(diagram *DiagramProcess) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsChecked && !stagedNode.IsChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
			stagedNode.IsChecked = frontNode.IsChecked
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			diagram.IsChecked = false
			stagedNode.IsChecked = frontNode.IsChecked
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			diagram.isExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			diagram.Name = frontNode.Name
			diagram.isInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
	}
}
