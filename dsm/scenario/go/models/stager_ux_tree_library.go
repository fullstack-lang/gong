package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(library *Library, parentNodes *[]*tree.Node) {
	name := library.Name
	if name == "" && library.IsRootLibrary {
		name = "Root Library"
	}

	libraryNode := &tree.Node{
		Name:            name,
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
	// Analyses
	//
	analysesNode := &tree.Node{
		Name:            "Analyses",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsAnalysesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, analysesNode)
	analysesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsAnalysesNodeExpanded)
	analysesNode.OnClick = onNodeClicked(stager, library)

	// add an analysis to the library button
	confAnalyses := ItemButtonConfiguration[
		Analysis, *Analysis,
		Library, *Library,
	]{
		parentNode:                         analysesNode,
		sliceForNewAddedItem:               &library.Analyses,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsAnalysesNodeExpanded,
	}
	addCreateItemButton(stager, confAnalyses)

	for _, analysis := range library.Analyses {
		stager.treeAnalysis(analysis, analysesNode)
	}

	//
	// Sub Libraries
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
