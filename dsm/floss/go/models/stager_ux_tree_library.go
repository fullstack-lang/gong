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
	// Systemes
	//
	systemesNode := &tree.Node{
		Name:            "Systemes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsSystemesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, systemesNode)
	systemesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSystemesNodeExpanded)
	systemesNode.OnClick = onNodeClicked(stager, library)

	// add a system to the library button
	confRootSystemes := ItemButtonConfiguration[
		System, *System,
		Library, *Library,
	]{
		parentNode:                         systemesNode,
		sliceForNewAddedItem:               &library.RootSystemes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSystemesNodeExpanded,
	}
	addCreateItemButton(stager, confRootSystemes)

	for _, system := range library.RootSystemes {
		stager.treeSystemes(system, systemesNode, &library.SystemsWhoseNodeIsExpanded)
	}

	//
	// SubLibraries
	//
	subLibrariesNode := &tree.Node{
		Name:            "Sub Libraries",
		FontStyle:       tree.ITALIC,
		BackgroundColor: "lightyellow",
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
}
