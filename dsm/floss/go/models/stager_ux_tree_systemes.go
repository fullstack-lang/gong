package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeSystemes(
	system *System,
	parentNode *tree.Node,
	systemsWhoseNodeIsExpanded *[]*System,
) {
	systemNode := &tree.Node{
		Name:            system.GetName(),
		IsExpanded:      slices.Contains(*systemsWhoseNodeIsExpanded, system),
		IsNodeClickable: true,
		IsInEditMode:    system.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, systemNode)

	addRenameButton(system, systemNode, stager)
	systemNode.OnNameChange = stager.onNameChange(system)
	systemNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, system, systemsWhoseNodeIsExpanded)
	systemNode.OnClick = onNodeClicked(stager, system)

	// Diagrams
	for _, diagramFloss := range system.DiagramFlosses {
		stager.treeDiagramFloss(system, diagramFloss, systemNode)
	}

	confDiagramFlosses := ItemButtonConfiguration[
		DiagramFloss, *DiagramFloss,
		System, *System,
	]{
		parentNode:                         systemNode,
		sliceForNewAddedItem:               &system.DiagramFlosses,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagramFlosses)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	//
	// SubSystemes
	//
	confSubSystemes := ItemButtonConfiguration[
		System, *System,
		System, *System,
	]{
		parentNode:                         systemNode,
		sliceForNewAddedItem:               &system.SubSystemes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   systemsWhoseNodeIsExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confSubSystemes)

	// SubSystemes
	subSystemesNode := &tree.Node{
		Name:            "SubSystemes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsSubSystemNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, subSystemesNode)

	for _, system_ := range system.SubSystemes {
		stager.treeSystemes(system_, subSystemesNode, systemsWhoseNodeIsExpanded)
	}

	//
	// Complexitys
	//
	complexitysNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, complexitysNode)
	complexitysNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsComplexitysNodeExpanded)
	complexitysNode.OnClick = onNodeClicked(stager, system)

	confComplexitys := ItemButtonConfiguration[
		Complexity, *Complexity,
		System, *System,
	]{
		parentNode:                         complexitysNode,
		sliceForNewAddedItem:               &system.Complexities,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsComplexitysNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confComplexitys)

	for _, complexity := range system.Complexities {
		stager.treeComplexityWithinSystem(system, complexity, complexitysNode)
	}

	//
	// Performances
	//
	performancesNode := &tree.Node{
		Name:            "Performances",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsPerformancesNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, performancesNode)
	performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsPerformancesNodeExpanded)
	performancesNode.OnClick = onNodeClicked(stager, system)

	confPerformances := ItemButtonConfiguration[
		Performance, *Performance,
		System, *System,
	]{
		parentNode:                         performancesNode,
		sliceForNewAddedItem:               &system.Performances,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsPerformancesNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confPerformances)

	for _, performance := range system.Performances {
		stager.treePerformanceWithinSystem(system, performance, performancesNode)
	}

	//
	// Efforts
	//
	effortsNode := &tree.Node{
		Name:            "Efforts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsEffortsNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, effortsNode)
	effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsEffortsNodeExpanded)
	effortsNode.OnClick = onNodeClicked(stager, system)

	confEfforts := ItemButtonConfiguration[
		Effort, *Effort,
		System, *System,
	]{
		parentNode:                         effortsNode,
		sliceForNewAddedItem:               &system.Efforts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsEffortsNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confEfforts)

	for _, effort := range system.Efforts {
		stager.treeEffortWithinSystem(system, effort, effortsNode)
	}
}
