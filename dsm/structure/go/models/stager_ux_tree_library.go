package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	if libraryNode.Name == "" {
		libraryNode.Name = "Root Library"
	}
	*parentNodes = append(*parentNodes, libraryNode)

	addRenameButton(library, libraryNode, stager)

	//
	// Sub Libraries
	//
	subLibrariesNode := &tree.Node{
		Name:            "Sub Libraries",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsSubLibrariesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, subLibrariesNode)

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(subLibrary, &subLibrariesNode.Children)
	}

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
	// Systems
	//
	systemsNode := &tree.Node{
		Name:            "Systems",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsSystemsNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, systemsNode)

	for _, system := range library.RootSystems {
		stager.treeSystem(system, systemsNode, &library.SystemsWhoseNodeIsExpanded)
	}

	confSystems := ItemButtonConfiguration[
		System, *System,
		Library, *Library,
	]{
		parentNode:                         systemsNode,
		sliceForNewAddedItem:               &library.RootSystems,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSystemsNodeExpanded,
	}
	addCreateItemButton(stager, confSystems)
}
