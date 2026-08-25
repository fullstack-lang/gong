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
	// Systems
	//
	systemesNode := &tree.Node{
		Name:            "Systems",
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
		stager.treeSystemes(library, system, systemesNode, &library.SystemsWhoseNodeIsExpanded)
	}

	//
	// CompareAnalysis
	//
	compareAnalysisNode := &tree.Node{
		Name:            "Compare Analyses",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsCompareAnalysisNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, compareAnalysisNode)
	compareAnalysisNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsCompareAnalysisNodeExpanded)
	compareAnalysisNode.OnClick = onNodeClicked(stager, library)

	confCompareAnalysis := ItemButtonConfiguration[
		CompareAnalysis, *CompareAnalysis,
		Library, *Library,
	]{
		parentNode:                         compareAnalysisNode,
		sliceForNewAddedItem:               &library.RootCompareAnalysis,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsCompareAnalysisNodeExpanded,
	}
	addCreateItemButton(stager, confCompareAnalysis)

	for _, compareAnalysis := range library.RootCompareAnalysis {
		stager.treeCompareAnalysisWithinLibrary(library, compareAnalysis, compareAnalysisNode)
	}

	//
	// Complexities
	//
	complexitiesNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, complexitiesNode)
	complexitiesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsComplexitysNodeExpanded)
	complexitiesNode.OnClick = onNodeClicked(stager, library)

	confComplexities := ItemButtonConfiguration[
		Complexity, *Complexity,
		Library, *Library,
	]{
		parentNode:                         complexitiesNode,
		sliceForNewAddedItem:               &library.RootComplexitys,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsComplexitysNodeExpanded,
		parentElement:                      library,
	}
	addCreateItemButton(stager, confComplexities)

	for _, complexity := range library.RootComplexitys {
		stager.treeComplexityWithinLibrary(library, complexity, complexitiesNode)
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

	confPerformances := ItemButtonConfiguration[
		Performance, *Performance,
		Library, *Library,
	]{
		parentNode:                         performancesNode,
		sliceForNewAddedItem:               &library.RootPerformances,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsPerformancesNodeExpanded,
		parentElement:                      library,
	}
	addCreateItemButton(stager, confPerformances)

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

	confEfforts := ItemButtonConfiguration[
		Effort, *Effort,
		Library, *Library,
	]{
		parentNode:                         effortsNode,
		sliceForNewAddedItem:               &library.RootEfforts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsEffortsNodeExpanded,
		parentElement:                      library,
	}
	addCreateItemButton(stager, confEfforts)

	for _, effort := range library.RootEfforts {
		stager.treeEffortWithinLibrary(library, effort, effortsNode)
	}

	//
	// Notes
	//
	notesNode := &tree.Node{
		Name:            "Notes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsNotesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, notesNode)
	notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsNotesNodeExpanded)
	notesNode.OnClick = onNodeClicked(stager, library)

	confNotes := ItemButtonConfiguration[
		Note, *Note,
		Library, *Library,
	]{
		parentNode:                         notesNode,
		sliceForNewAddedItem:               &library.RootNotes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsNotesNodeExpanded,
		parentElement:                      library,
	}
	addCreateItemButton(stager, confNotes)

	for _, note := range library.RootNotes {
		stager.treeNote(library, note, notesNode)
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
		parentElement:                      library,
	}
	addCreateItemButton(stager, confSubLibraries)
}
