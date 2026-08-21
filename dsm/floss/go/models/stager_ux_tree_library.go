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
		IsExpanded:      library.IsSystemsNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, systemesNode)
	systemesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSystemsNodeExpanded)
	systemesNode.OnClick = onNodeClicked(stager, library)

	// add a system to the library button
	confRootSystemes := ItemButtonConfiguration[
		System, *System,
		Library, *Library,
	]{
		parentNode:                         systemesNode,
		sliceForNewAddedItem:               &library.RootSystems,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSystemsNodeExpanded,
	}
	addCreateItemButton(stager, confRootSystemes)

	for _, system := range library.RootSystems {
		stager.treeSystemes(system, systemesNode, &library.SystemsWhoseNodeIsExpanded)
	}

	//
	// Complexitys
	//
	complexitysNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, complexitysNode)
	complexitysNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsComplexitysNodeExpanded)
	complexitysNode.OnClick = onNodeClicked(stager, library)

	confComplexity := ItemButtonConfiguration[
		Complexity, *Complexity,
		Library, *Library,
	]{
		parentNode:                         complexitysNode,
		sliceForNewAddedItem:               &library.RootComplexitys,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsComplexitysNodeExpanded,
	}
	addCreateItemButton(stager, confComplexity)

	for _, complexity := range library.RootComplexitys {
		stager.treeComplexityWithinLibrary(library, complexity, complexitysNode)
	}

	//
	// Performances
	//
	performancesNode := &tree.Node{
		Name:            "Performances",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsPerformancesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, performancesNode)
	performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsPerformancesNodeExpanded)
	performancesNode.OnClick = onNodeClicked(stager, library)

	confPerformance := ItemButtonConfiguration[
		Performance, *Performance,
		Library, *Library,
	]{
		parentNode:                         performancesNode,
		sliceForNewAddedItem:               &library.RootPerformances,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsPerformancesNodeExpanded,
	}
	addCreateItemButton(stager, confPerformance)

	for _, performance := range library.RootPerformances {
		stager.treePerformanceWithinLibrary(library, performance, performancesNode)
	}

	//
	// Efforts
	//
	effortsNode := &tree.Node{
		Name:            "Efforts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsEffortsNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, effortsNode)
	effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsEffortsNodeExpanded)
	effortsNode.OnClick = onNodeClicked(stager, library)

	confEffort := ItemButtonConfiguration[
		Effort, *Effort,
		Library, *Library,
	]{
		parentNode:                         effortsNode,
		sliceForNewAddedItem:               &library.RootEfforts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsEffortsNodeExpanded,
	}
	addCreateItemButton(stager, confEffort)

	for _, effort := range library.RootEfforts {
		stager.treeEffortWithinLibrary(library, effort, effortsNode)
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
