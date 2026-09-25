package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:                 library.Name,
		IsExpanded:           library.IsExpanded,
		IsNodeClickable:      true,
		IsInEditMode:         library.isInRenameMode,
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(buttons.BUTTON_local_library),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	addRenameButton(library, libraryNode, stager)
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpanded)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confSubLibraries)

	confNotes := ItemButtonConfiguration[
		Note, *Note, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Notes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksNotes := addCreateItemButton(stager, confNotes)
	callbacksNotes.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsNotesNodeExpanded = true
		}
	}

	confResources := ItemButtonConfiguration[
		Resource, *Resource, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootResources,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksResources := addCreateItemButton(stager, confResources)
	callbacksResources.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsResourcesNodeExpanded = true
		}
	}

	confTaskGroups := ItemButtonConfiguration[
		TaskGroup, *TaskGroup, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootTaskGroups,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksTaskGroups := addCreateItemButton(stager, confTaskGroups)
	callbacksTaskGroups.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsTaskGroupsNodeExpanded = true
		}
	}
	if callbacksTaskGroups.button != nil {
		callbacksTaskGroups.button.Name = "Add Task Group"
		callbacksTaskGroups.button.ToolTipText = "Add a Task Group to \"" + library.Name + "\""
	} else if len(libraryNode.Menu.Buttons) > 0 {
		libraryNode.Menu.Buttons[0].Name = "Add Task Group"
		libraryNode.Menu.Buttons[0].ToolTipText = "Add a Task Group to \"" + library.Name + "\""
	}

	confTasks := ItemButtonConfiguration[
		Task, *Task, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootTasks,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksTasks := addCreateItemButton(stager, confTasks)
	callbacksTasks.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsWBSNodeExpanded = true
		}
		if callbacksTasks.createdItem != nil {
			callbacksTasks.createdItem.IsAllDay = true
		}
	}

	confProducts := ItemButtonConfiguration[
		Product, *Product, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootProducts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksProducts := addCreateItemButton(stager, confProducts)
	callbacksProducts.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsPBSNodeExpanded = true
		}
	}

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagrams)

	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		newDiagram.IsInAutoLayoutMode = true
		for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagram := range library.Diagrams {
		stager.treeDiagram(library, diagram, libraryNode)
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}

	EnsureRenameIsFirst(libraryNode)
}
