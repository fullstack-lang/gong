package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.getRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}
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

	confPlants := ItemButtonConfiguration[
		Plant, *Plant, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Plants,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confPlants)

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}

	for _, plant := range library.Plants {
		stager.treePlant(plant, &libraryNode.Children)
	}
}
