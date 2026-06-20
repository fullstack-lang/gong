package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeSystem(system *System, parentNode *tree.Node, expandedSlice *[]*System) {
	systemNode := &tree.Node{
		Name:            system.Name,
		IsExpanded:      system.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    system.isInRenameMode,
	}
	if systemNode.Name == "" {
		systemNode.Name = "System"
	}
	parentNode.Children = append(parentNode.Children, systemNode)

	addRenameButton(system, systemNode, stager)
	systemNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsExpanded)
	systemNode.OnNameChange = stager.onNameChange(system)
	systemNode.OnClick = onNodeClicked(stager, system)

	//
	// Sub Systems
	//
	subSystemsNode := &tree.Node{
		Name:            "Sub Systems",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsSubSystemsNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, subSystemsNode)
	subSystemsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsSubSystemsNodeExpanded)
	subSystemsNode.OnClick = onNodeClicked(stager, system)

	for _, subSystem := range system.SubSystems {
		stager.treeSystem(subSystem, subSystemsNode, &system.SubSystemsWhoseNodeIsExpanded)
	}

	confSubSystems := ItemButtonConfiguration[
		System, *System,
		System, *System,
	]{
		parentNode:                         subSystemsNode,
		sliceForNewAddedItem:               &system.SubSystems,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsSubSystemsNodeExpanded,
	}
	addCreateItemButton(stager, confSubSystems)

	//
	// Diagrams
	//
	diagramsNode := &tree.Node{
		Name:            "Diagrams",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsDiagramStructuresNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, diagramsNode)
	diagramsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsDiagramStructuresNodeExpanded)
	diagramsNode.OnClick = onNodeClicked(stager, system)

	for _, diag := range system.DiagramStructures {
		stager.treeDiagramStructure(diag, diagramsNode, &system.DiagramStructuresWhoseNodeIsExpanded)
	}

	confDiagrams := ItemButtonConfiguration[
		DiagramStructure, *DiagramStructure,
		System, *System,
	]{
		parentNode:                         diagramsNode,
		sliceForNewAddedItem:               &system.DiagramStructures,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsDiagramStructuresNodeExpanded,
	}
	addCreateItemButton(stager, confDiagrams)

	//
	// Parts
	//
	partsNode := &tree.Node{
		Name:            "Parts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsPartsNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, partsNode)
	partsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsPartsNodeExpanded)
	partsNode.OnClick = onNodeClicked(stager, system)

	for _, part := range system.Parts {
		partNode := &tree.Node{
			Name:            part.Name,
			IsExpanded:      part.IsExpanded,
			IsNodeClickable: true,
			IsInEditMode:    part.isInRenameMode,
		}
		if partNode.Name == "" {
			partNode.Name = "Part"
		}
		partsNode.Children = append(partsNode.Children, partNode)
		addRenameButton(part, partNode, stager)
		partNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsExpanded)
		partNode.OnNameChange = stager.onNameChange(part)
		partNode.OnClick = onNodeClicked(stager, part)
	}

	confParts := ItemAndShapeButtonConfiguration[
		Part, *Part,
		System, *System,
		PartShape, *PartShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Part, *Part,
			System, *System,
		]{
			parentNode:                         partsNode,
			sliceForNewAddedItem:               &system.Parts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &system.IsPartsNodeExpanded,
		},
		receivingDiagram:      stager.diagram,
	}
	if stager.diagram != nil {
		confParts.sliceForNewAddedShape = &stager.diagram.Part_Shapes
	}
	addCreateItemAndShapeButton(stager, confParts)

	//
	// Links
	//
	linksNode := &tree.Node{
		Name:            "Links",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsLinksNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, linksNode)
	linksNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsLinksNodeExpanded)
	linksNode.OnClick = onNodeClicked(stager, system)

	for _, link := range system.Links {
		linkNode := &tree.Node{
			Name:            link.Name,
			IsExpanded:      link.IsExpanded,
			IsNodeClickable: true,
			IsInEditMode:    link.isInRenameMode,
		}
		if linkNode.Name == "" {
			linkNode.Name = "Link"
		}
		linksNode.Children = append(linksNode.Children, linkNode)
		addRenameButton(link, linkNode, stager)
		linkNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&link.IsExpanded)
		linkNode.OnNameChange = stager.onNameChange(link)
		linkNode.OnClick = onNodeClicked(stager, link)
	}

	confLinks := ItemShapeAndLinkButtonConfiguration[
		Link, *Link,
		System, *System,
		PartShape, *PartShape,
		LinkAssociationShape, *LinkAssociationShape,
	]{
		ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
			Link, *Link,
			System, *System,
			PartShape, *PartShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Link, *Link,
				System, *System,
			]{
				parentNode:                         linksNode,
				sliceForNewAddedItem:               &system.Links,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &system.IsLinksNodeExpanded,
			},
			receivingDiagram: stager.diagram,
		},
	}
	if stager.diagram != nil {
		confLinks.sliceForNewCompositionShapes = &stager.diagram.Link_Shapes
	}
	addCreateItemShapeAndLinkButton(stager, confLinks)
}
