package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	var libraryNode = &tree.Node{
		Name:                 library.Name,
		IsExpanded:           library.IsExpanded,
		IsNodeClickable:      true,
		IsInEditMode:         library.GetIsInRenameMode(),
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(buttons.BUTTON_local_library),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.GetRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}

	libraryNode.OnNameChange = func(newName string) {
		library.Name = newName
		library.SetIsInRenameMode(false)
		stager.stage.Commit()
	}
	libraryNode.OnIsExpandedChange = func(isExpanded bool) {
		library.IsExpanded = isExpanded
		stager.stage.Commit()
	}
	libraryNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(library, GetPointerToGongstructName[*Library]())
		stager.stage.Commit()
	}

	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confSubLibraries)

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     false,
	}
	diagramItemAdderCallback := addCreateItemButton(stager, confDiagrams)

	diagramItemAdderCallback.OnBeforeCommit = func() {
		newDiagram := diagramItemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	// Also add "Add Diagram" into menu for discoverability
	if libraryNode.Menu == nil {
		libraryNode.Menu = &tree.Menu{Name: "Menu"}
	}
	libraryNode.Menu.Buttons = append([]*tree.Button{
		{
			Name:            "Add Diagram",
			Icon:            string(buttons.BUTTON_add),
			ToolTipText:     "Add a Diagram to \"" + libraryNode.Name + "\"",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				callbacks := &itemAdderCallback[*Diagram]{}
				newAbstractElement := processAbstractItemAddition(stager, confDiagrams, callbacks)
				newAbstractElement.IsEditable_ = true
				newAbstractElement.IsExpanded = true
				for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
					diagram_.IsChecked = false
				}
				newAbstractElement.IsChecked = true
				stager.stage.Commit()
			},
		},
	}, libraryNode.Menu.Buttons...)

	for _, diagram := range library.Diagrams {
		stager.treeDiagramCapture(library, diagram, libraryNode)
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(isExpanded bool) {
	return func(newIsExpanded bool) {
		if *isExpanded != newIsExpanded {
			*isExpanded = newIsExpanded
		}
		stager.stage.Commit()
	}
}
