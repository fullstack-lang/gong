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
}
