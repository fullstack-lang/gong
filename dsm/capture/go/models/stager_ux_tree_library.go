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

	addRenameButton(library, libraryNode, stager)

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

	confNotes := ItemButtonConfiguration[
		Note, *Note,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Notes,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksNotes := addCreateItemButton(stager, confNotes)
	callbacksNotes.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsNotesNodeExpanded = true
		}
	}

	confConcepts := ItemButtonConfiguration[
		Concept, *Concept,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootConcepts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksConcepts := addCreateItemButton(stager, confConcepts)
	callbacksConcepts.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsConceptsNodeExpanded = true
		}
	}

	confViews := ItemButtonConfiguration[
		View, *View,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootViews,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksViews := addCreateItemButton(stager, confViews)
	callbacksViews.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsViewsNodeExpanded = true
		}
	}

	confRequirements := ItemButtonConfiguration[
		Requirement, *Requirement,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootRequirements,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksRequirements := addCreateItemButton(stager, confRequirements)
	callbacksRequirements.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsRequirementsNodeExpanded = true
		}
	}

	confDeliverables := ItemButtonConfiguration[
		Deliverable, *Deliverable,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootDeliverables,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksDeliverables := addCreateItemButton(stager, confDeliverables)
	callbacksDeliverables.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsPBSNodeExpanded = true
		}
	}

	confStakeholders := ItemButtonConfiguration[
		Stakeholder, *Stakeholder,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootStakeholders,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksStakeholders := addCreateItemButton(stager, confStakeholders)
	callbacksStakeholders.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsStakeholdersNodeExpanded = true
		}
	}

	confConcerns := ItemButtonConfiguration[
		Concern, *Concern,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootConcerns,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksConcerns := addCreateItemButton(stager, confConcerns)
	callbacksConcerns.OnBeforeCommit = func() {
		for _, diagram := range library.Diagrams {
			diagram.IsConcernsNodeExpanded = true
		}
	}

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	diagramItemAdderCallback := addCreateItemButton(stager, confDiagrams)

	diagramItemAdderCallback.OnBeforeCommit = func() {
		newDiagram := diagramItemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagram := range library.Diagrams {
		stager.treeDiagramCapture(library, diagram, libraryNode)
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}

	EnsureRenameIsFirst(libraryNode)
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(isExpanded bool) {
	return func(newIsExpanded bool) {
		if *isExpanded != newIsExpanded {
			*isExpanded = newIsExpanded
		}
		stager.stage.Commit()
	}
}
