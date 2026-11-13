package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramStructure(
	system *System,
	diagramStructure *DiagramStructure,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagramStructure.Name,
		IsExpanded:        slices.Contains(system.DiagramStructureWhoseNodeIsExpanded, diagramStructure) == true,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagramStructure.IsChecked,

		IsInEditMode: diagramStructure.isInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	element := diagramStructure
	node := diagramNode

	addRenameButton(element, node, stager)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			// uncheck all diagrams
			for diagramStructure_ := range *stager.stage.GetInstancesSet[*DiagramStructure]() {
				diagramStructure_.IsChecked = false
			}

			diagramStructure.IsChecked = true
			stager.stage.Commit()
			return
		} else {
			diagramStructure.IsChecked = false
			stager.stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, diagramStructure)
	node.OnNameChange = stager.onNameChange(diagramStructure)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagramStructure, &system.DiagramStructureWhoseNodeIsExpanded)

	// editability button (toggles edit / non-edit mode, hiding control-point dots when off)
	{
		editButton := &tree.Button{
			Name:            "Diagram Editability",
			Icon:            string(buttons.BUTTON_edit),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Edit diagram",
			OnClick: func() {
				diagramStructure.IsEditable_ = !diagramStructure.IsEditable_
				stager.stage.Commit()
			},
		}
		if diagramStructure.IsEditable_ {
			editButton.Icon = string(buttons.BUTTON_edit_off)
			editButton.ToolTipText = "Stop editing diagram"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, editButton)
	}

	// prefix button
	{
		showPrefixButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagramStructure.IsShowPrefix = !diagramStructure.IsShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagramStructure.IsShowPrefix {
			showPrefixButton.Icon = string(buttons.BUTTON_label)
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
	}

	// Parts
	{
		partsNode := &tree.Node{
			Name:            "Parts",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramStructure.IsPartsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, partsNode)
		partsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramStructure.IsPartsNodeExpanded)

		visitedSystems := make(map[*System]bool)
		if system != nil {
			visitedSystems[system] = true
		}
		for _, part := range system.Parts {
			stager.treePartsWithEnclosing(
				diagramStructure,
				part,
				nil,
				false,
				&diagramStructure.PartWhoseNodeIsExpanded,
				visitedSystems,
				partsNode,
			)
		}
		confParts := ItemButtonConfiguration[
			Part, *Part,
			System, *System,
		]{
			parentNode:                         partsNode,
			sliceForNewAddedItem:               &system.Parts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramStructure.IsPartsNodeExpanded,
		}
		addCreateItemButton(stager, confParts)
	}

	// external parts
	{
		externalPartsNode := &tree.Node{
			Name:            "External Parts",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramStructure.IsExternalPartsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, externalPartsNode)
		externalPartsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramStructure.IsExternalPartsNodeExpanded)

		for _, part := range system.ExternalParts {
			stager.treeExternalParts(diagramStructure, part, externalPartsNode)
		}

		addCreateItemButton(stager, ItemButtonConfiguration[
			Part, *Part,
			System, *System,
		]{
			parentNode:                         externalPartsNode,
			sliceForNewAddedItem:               &system.ExternalParts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramStructure.IsExternalPartsNodeExpanded,
		})
	}

	{
		//
		// DataFlows
		//
		dataFlowsNode := &tree.Node{
			Name:            "Data Flows",
			FontStyle:       tree.ITALIC,
			IsExpanded:      system.IsDataFlowsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, dataFlowsNode)
		dataFlowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsDataFlowsNodeExpanded)

		for _, dataFlow := range system.DataFlows {
			stager.treeDataFlowsWithinSystemDiagram(diagramStructure, dataFlow, dataFlowsNode)
		}
	}

	{
		// Notes
		//
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramStructure.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramStructure.IsNotesNodeExpanded)

		for _, note := range system.GetOwningLibrary().RootNotes {
			stager.treeNoteWithinDiagramStructure(diagramStructure, note, notesNode)
		}
	}
}
