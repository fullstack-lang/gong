package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	var libraryNode = &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.GetIsInRenameMode(),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.GetRootLibrary() {
		if !library.GetIsInRenameMode() {
			libraryNode.Buttons = append(libraryNode.Buttons,
				&tree.Button{
					Name: library.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnClick: func() {
						library.SetIsInRenameMode(true)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Rename the " + GetGongstructNameFromPointer(library),
					ToolTipPosition: tree.Above,
				})
		} else {
			libraryNode.Buttons = append(libraryNode.Buttons,
				&tree.Button{
					Name: library.GetName() + " " + string(buttons.BUTTON_edit_off),
					Icon: string(buttons.BUTTON_edit_off),
					OnClick: func() {
						library.SetIsInRenameMode(false)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Cancel renaming",
					ToolTipPosition: tree.Above,
				})
		}
	}

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
	}
	addCreateItemButton(stager, confSubLibraries)

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
	}
	addCreateItemButton(stager, confDiagrams)

	for _, diagram := range library.Diagrams {
		stager.treeDiagramCapture(library, diagram, libraryNode)
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(isExpanded bool) {
	return func(newIsExpanded bool) {
		if *isExpanded != newIsExpanded {
			*isExpanded = newIsExpanded
		}
		stager.stage.Commit()
	}
}
