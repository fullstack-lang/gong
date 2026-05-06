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

	diagramNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			diagramProcess.Name = frontNode.Name
			diagramProcess.isInRenameMode = false
			stager.stage.Commit()
			return
		}
		if frontNode.IsChecked && !stagedNode.IsChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			diagramProcess.IsChecked = true
			stagedNode.IsChecked = frontNode.IsChecked
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			diagramProcess.IsChecked = false
			stagedNode.IsChecked = frontNode.IsChecked
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded {
			if slices.Index(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess) == -1 {
				process.DiagramProcessWhoseNodeIsExpanded = append(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess)
			}
		} else {
			if idx := slices.Index(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess); idx != -1 {
				process.DiagramProcessWhoseNodeIsExpanded = slices.Delete(process.DiagramProcessWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		stager.probeForm.FillUpFormFromGongstruct(diagramProcess, "Diagram")
	}

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
		participantsNode.OnClick = func(frontNode *tree.Node) {
			if frontNode.IsExpanded != diagramProcess.IsParticipantsNodeExpanded {
				diagramProcess.IsParticipantsNodeExpanded = frontNode.IsExpanded
				stager.stage.Commit()
				return
			}
		}

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
		externalParticipantsNode.OnClick = func(frontNode *tree.Node) {
			if frontNode.IsExpanded != diagramProcess.IsExternalParticipantsNodeExpanded {
				diagramProcess.IsExternalParticipantsNodeExpanded = frontNode.IsExpanded
				stager.stage.Commit()
				return
			}
		}

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
		dataFlowsNode.OnClick = func(frontNode *tree.Node) {
			if frontNode.IsExpanded != process.IsDataFlowsNodeExpanded {
				process.IsDataFlowsNodeExpanded = frontNode.IsExpanded
				stager.stage.Commit()
				return
			}
			stager.stage.Commit()
		}

		for _, dataFlow := range process.DataFlows {
			stager.treeDataFlowsWithinProcessDiagram(diagramProcess, dataFlow, dataFlowsNode)
		}
	}
}
