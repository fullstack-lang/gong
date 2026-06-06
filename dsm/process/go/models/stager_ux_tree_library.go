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

	if library != stager.getRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpandedTmp)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)

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
	subLibrariesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSubLibrariesNodeExpanded)
	subLibrariesNode.OnClick = onNodeClicked(stager, library)

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
	// Processes
	//
	processesNode := &tree.Node{
		Name:            "Processes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsProcessesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, processesNode)
	processesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsProcessesNodeExpanded)
	processesNode.OnClick = onNodeClicked(stager, library)

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
	// Data Flows
	//
	dataFlowNodes := &tree.Node{
		Name:            "Data Flows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsDataFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, dataFlowNodes)
	dataFlowNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsDataFlowsNodeExpanded)
	dataFlowNodes.OnClick = onNodeClicked(stager, library)

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
	datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsDatasNodeExpanded)
	datasNode.OnClick = onNodeClicked(stager, library)

	for _, data := range library.RootDatas {
		stager.treeDataWithinLibrary(library, data, datasNode)
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
		parentNodeExpansionBooleanValue:    &library.IsDatasNodeExpanded,
	}
	addCreateItemButton(stager, confData)

	//
	// Resources
	//
	resourcesNode := &tree.Node{
		Name:            "Resources",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsResourcesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, resourcesNode)
	resourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsResourcesNodeExpanded)
	resourcesNode.OnClick = onNodeClicked(stager, library)

	for _, resource := range library.RootResources {
		stager.treeResourceWithinLibrary(library, resource, resourcesNode)
	}

	// add resource button
	confResource := ItemButtonConfiguration[
		Resource, *Resource,
		Resource, *Resource,
	]{
		parentNode:                         resourcesNode,
		sliceForNewAddedItem:               &library.RootResources,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsResourcesNodeExpanded,
	}
	addCreateItemButton(stager, confResource)

	// Notes
	notesNode := &tree.Node{
		Name:            "Notes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsNotesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, notesNode)
	notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsNotesNodeExpanded)
	notesNode.OnClick = onNodeClicked(stager, library)

	for _, note := range library.RootNotes {
		stager.treeNote(library, note, notesNode)
	}

	// add note button
	confNote := ItemButtonConfiguration[
		Note, *Note,
		Note, *Note,
	]{
		parentNode:                         notesNode,
		sliceForNewAddedItem:               &library.RootNotes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsNotesNodeExpanded,
	}
	addCreateItemButton(stager, confNote)
}

// Helper callbacks

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
