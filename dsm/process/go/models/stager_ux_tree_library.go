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

	addRenameButton(library, libraryNode, stager)
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpandedTmp)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)

	//
	// Buttons on libraryNode menu
	//
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
		IsButtonInMenu:                     true,
	}
	callbacksSubLibraries := addCreateItemButton(stager, confSubLibraries)
	callbacksSubLibraries.OnBeforeCommit = func() {
		library.IsSubLibrariesNodeExpanded = true
	}
	if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Sub Library"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Sub Library to \"" + library.Name + "\""
	}

	confRootProcesses := ItemButtonConfiguration[
		Process, *Process,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootProcesses,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
		IsButtonInMenu:                     true,
	}
	callbacksProcesses := addCreateItemButton(stager, confRootProcesses)
	callbacksProcesses.OnBeforeCommit = func() {
		library.IsProcessesNodeExpanded = true
	}
	if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Process"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Process to \"" + library.Name + "\""
	}

	confData := ItemButtonConfiguration[
		Data, *Data,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootDatas,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
		IsButtonInMenu:                     true,
	}
	callbacksData := addCreateItemButton(stager, confData)
	callbacksData.OnBeforeCommit = func() {
		library.IsDatasNodeExpanded = true
	}
	if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Data"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Data to \"" + library.Name + "\""
	}

	confResource := ItemButtonConfiguration[
		Resource, *Resource,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootResources,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
		IsButtonInMenu:                     true,
	}
	callbacksResource := addCreateItemButton(stager, confResource)
	callbacksResource.OnBeforeCommit = func() {
		library.IsResourcesNodeExpanded = true
	}
	if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Resource"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Resource to \"" + library.Name + "\""
	}

	confNote := ItemButtonConfiguration[
		Note, *Note,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootNotes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
		IsButtonInMenu:                     true,
	}
	callbacksNote := addCreateItemButton(stager, confNote)
	callbacksNote.OnBeforeCommit = func() {
		library.IsNotesNodeExpanded = true
	}
	if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Note"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Note to \"" + library.Name + "\""
	}

	//
	// SubLibraries
	//
	if len(library.SubLibraries) > 0 {
		subLibrariesNode := &tree.Node{
			Name:            "Sub Libraries",
			FontStyle:       tree.ITALIC,
			IsExpanded:      library.IsSubLibrariesNodeExpanded,
			IsNodeClickable: true,
		}
		libraryNode.Children = append(libraryNode.Children, subLibrariesNode)
		subLibrariesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSubLibrariesNodeExpanded)
		subLibrariesNode.OnClick = onNodeClicked(stager, library)

		confSubLibrariesNode := ItemButtonConfiguration[
			Library, *Library,
			Library, *Library,
		]{
			parentNode:                         subLibrariesNode,
			sliceForNewAddedItem:               &library.SubLibraries,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.IsSubLibrariesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksSubLibrariesNode := addCreateItemButton(stager, confSubLibrariesNode)
		callbacksSubLibrariesNode.OnBeforeCommit = func() {
			library.IsSubLibrariesNodeExpanded = true
		}
		if len(subLibrariesNode.Menu.Buttons) > 0 {
			subLibrariesNode.Menu.Buttons[0].Name = "Add Sub Library"
			subLibrariesNode.Menu.Buttons[0].ToolTipText = "Add a Sub Library to \"" + library.Name + "\""
		}

		for _, subLibrary := range library.SubLibraries {
			stager.treeLibrary(subLibrary, &subLibrariesNode.Children)
		}
	}

	//
	// Processes
	//
	if len(library.RootProcesses) > 0 {
		processesNode := &tree.Node{
			Name:            "Processes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      library.IsProcessesNodeExpanded,
			IsNodeClickable: true,
		}
		libraryNode.Children = append(libraryNode.Children, processesNode)
		processesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsProcessesNodeExpanded)
		processesNode.OnClick = onNodeClicked(stager, library)

		confRootProcessesNode := ItemButtonConfiguration[
			Process, *Process,
			Library, *Library,
		]{
			parentNode:                         processesNode,
			sliceForNewAddedItem:               &library.RootProcesses,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.IsProcessesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksProcessesNode := addCreateItemButton(stager, confRootProcessesNode)
		callbacksProcessesNode.OnBeforeCommit = func() {
			library.IsProcessesNodeExpanded = true
		}
		if len(processesNode.Menu.Buttons) > 0 {
			processesNode.Menu.Buttons[0].Name = "Add Process"
			processesNode.Menu.Buttons[0].ToolTipText = "Add a Process to \"" + library.Name + "\""
		}

		for _, process := range library.RootProcesses {
			stager.treeProcesses(process, processesNode, &library.ProcesssWhoseNodeIsExpanded)
		}
	}

	//
	// Data Flows
	//
	if len(library.RootDataFlows) > 0 {
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
	}

	//
	// Data
	//
	if len(library.RootDatas) > 0 {
		datasNode := &tree.Node{
			Name:            "Data",
			FontStyle:       tree.ITALIC,
			IsExpanded:      library.IsDatasNodeExpanded,
			IsNodeClickable: true,
		}
		libraryNode.Children = append(libraryNode.Children, datasNode)
		datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsDatasNodeExpanded)
		datasNode.OnClick = onNodeClicked(stager, library)

		confDataNode := ItemButtonConfiguration[
			Data, *Data,
			Library, *Library,
		]{
			parentNode:                         datasNode,
			sliceForNewAddedItem:               &library.RootDatas,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.IsDatasNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksDataNode := addCreateItemButton(stager, confDataNode)
		callbacksDataNode.OnBeforeCommit = func() {
			library.IsDatasNodeExpanded = true
		}
		if len(datasNode.Menu.Buttons) > 0 {
			datasNode.Menu.Buttons[0].Name = "Add Data"
			datasNode.Menu.Buttons[0].ToolTipText = "Add a Data to \"" + library.Name + "\""
		}

		for _, data := range library.RootDatas {
			stager.treeDataWithinLibrary(library, data, datasNode)
		}
	}

	//
	// Resources
	//
	if len(library.RootResources) > 0 {
		resourcesNode := &tree.Node{
			Name:            "Resources",
			FontStyle:       tree.ITALIC,
			IsExpanded:      library.IsResourcesNodeExpanded,
			IsNodeClickable: true,
		}
		libraryNode.Children = append(libraryNode.Children, resourcesNode)
		resourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsResourcesNodeExpanded)
		resourcesNode.OnClick = onNodeClicked(stager, library)

		confResourceNode := ItemButtonConfiguration[
			Resource, *Resource,
			Library, *Library,
		]{
			parentNode:                         resourcesNode,
			sliceForNewAddedItem:               &library.RootResources,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.IsResourcesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksResourceNode := addCreateItemButton(stager, confResourceNode)
		callbacksResourceNode.OnBeforeCommit = func() {
			library.IsResourcesNodeExpanded = true
		}
		if len(resourcesNode.Menu.Buttons) > 0 {
			resourcesNode.Menu.Buttons[0].Name = "Add Resource"
			resourcesNode.Menu.Buttons[0].ToolTipText = "Add a Resource to \"" + library.Name + "\""
		}

		for _, resource := range library.RootResources {
			stager.treeResourceWithinLibrary(library, resource, resourcesNode)
		}
	}

	//
	// Notes
	//
	if len(library.RootNotes) > 0 {
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      library.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		libraryNode.Children = append(libraryNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsNotesNodeExpanded)
		notesNode.OnClick = onNodeClicked(stager, library)

		confNoteNode := ItemButtonConfiguration[
			Note, *Note,
			Library, *Library,
		]{
			parentNode:                         notesNode,
			sliceForNewAddedItem:               &library.RootNotes,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.IsNotesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksNoteNode := addCreateItemButton(stager, confNoteNode)
		callbacksNoteNode.OnBeforeCommit = func() {
			library.IsNotesNodeExpanded = true
		}
		if len(notesNode.Menu.Buttons) > 0 {
			notesNode.Menu.Buttons[0].Name = "Add Note"
			notesNode.Menu.Buttons[0].ToolTipText = "Add a Note to \"" + library.Name + "\""
		}

		for _, note := range library.RootNotes {
			stager.treeNote(library, note, notesNode)
		}
	}
}
