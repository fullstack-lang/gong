package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramProcess(
	process *Process,
	diagramProcess *DiagramProcess,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagramProcess.Name,
		IsExpanded:        slices.Contains(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess) == true,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagramProcess.IsChecked,

		IsInEditMode: diagramProcess.isInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	element := diagramProcess
	node := diagramNode

	addRenameButton(element, node, stager)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			diagramProcess.IsChecked = true
			stager.stage.Commit()
			return
		} else {
			diagramProcess.IsChecked = false
			stager.stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, diagramProcess)
	node.OnNameChange = stager.onNameChange(diagramProcess)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagramProcess, &process.DiagramProcessWhoseNodeIsExpanded)

	// prefix button
	{
		showPrefixButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagramProcess.IsShowPrefix = !diagramProcess.IsShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagramProcess.IsShowPrefix {
			showPrefixButton.Icon = string(buttons.BUTTON_label)
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
	}

	// Participants
	{
		participantsNode := &tree.Node{
			Name:            "Participants",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramProcess.IsParticipantsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, participantsNode)
		participantsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramProcess.IsParticipantsNodeExpanded)

		for _, participant := range process.Participants {
			stager.treeParticipants(diagramProcess, participant, participantsNode)
		}
		confParticipants := ItemButtonConfiguration[
			Participant, *Participant,
			Process, *Process,
		]{
			parentNode:                         participantsNode,
			sliceForNewAddedItem:               &process.Participants,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsParticipantsNodeExpanded,
		}
		addCreateItemButton(stager, confParticipants)
	}

	// external participants
	{
		externalParticipantsNode := &tree.Node{
			Name:            "External Participants",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramProcess.IsExternalParticipantsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, externalParticipantsNode)
		externalParticipantsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramProcess.IsExternalParticipantsNodeExpanded)

		for _, participant := range process.ExternalParticipants {
			stager.treeExternalParticipants(diagramProcess, participant, externalParticipantsNode)
		}

		addCreateItemButton(stager, ItemButtonConfiguration[
			Participant, *Participant,
			Process, *Process,
		]{
			parentNode:                         externalParticipantsNode,
			sliceForNewAddedItem:               &process.ExternalParticipants,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsExternalParticipantsNodeExpanded,
		})
	}

	{
		//
		// DataFlows
		//
		dataFlowsNode := &tree.Node{
			Name:            "Data Flows",
			FontStyle:       tree.ITALIC,
			IsExpanded:      process.IsDataFlowsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, dataFlowsNode)
		dataFlowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&process.IsDataFlowsNodeExpanded)

		for _, dataFlow := range process.DataFlows {
			stager.treeDataFlowsWithinProcessDiagram(diagramProcess, dataFlow, dataFlowsNode)
		}
	}
}
