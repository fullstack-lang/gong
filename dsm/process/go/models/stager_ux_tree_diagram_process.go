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
			// uncheck all diagrams
			for diagramProcess_ := range *stager.stage.GetInstancesSet[*DiagramProcess]() {
				diagramProcess_.IsChecked = false
			}

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

	//
	// Buttons on diagramNode menu
	//
	confParticipantsDiagram := ItemButtonConfiguration[
		Participant, *Participant,
		Process, *Process,
	]{
		parentNode:                         diagramNode,
		sliceForNewAddedItem:               &process.Participants,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &diagramProcess.IsParticipantsNodeExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksParticipantsDiagram := addCreateItemButton(stager, confParticipantsDiagram)
	callbacksParticipantsDiagram.OnBeforeCommit = func() {
		diagramProcess.IsParticipantsNodeExpanded = true
	}
	if len(diagramNode.Menu.Buttons) > 0 {
		diagramNode.Menu.Buttons[0].Name = "Add Participant"
		diagramNode.Menu.Buttons[0].ToolTipText = "Add a Participant to \"" + diagramProcess.Name + "\""
	}

	confExternalParticipantsDiagram := ItemButtonConfiguration[
		Participant, *Participant,
		Process, *Process,
	]{
		parentNode:                         diagramNode,
		sliceForNewAddedItem:               &process.ExternalParticipants,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &diagramProcess.IsExternalParticipantsNodeExpanded,
		IsButtonInMenu:                     true,
	}
	callbacksExternalParticipantsDiagram := addCreateItemButton(stager, confExternalParticipantsDiagram)
	callbacksExternalParticipantsDiagram.OnBeforeCommit = func() {
		diagramProcess.IsExternalParticipantsNodeExpanded = true
	}
	if len(diagramNode.Menu.Buttons) > 0 {
		diagramNode.Menu.Buttons[0].Name = "Add External Participant"
		diagramNode.Menu.Buttons[0].ToolTipText = "Add an External Participant to \"" + diagramProcess.Name + "\""
	}

	if process.GetOwningLibrary() != nil {
		confNotesDiagram := ItemButtonConfiguration[
			Note, *Note,
			Library, *Library,
		]{
			parentNode:                         diagramNode,
			sliceForNewAddedItem:               &process.GetOwningLibrary().RootNotes,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsNotesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksNotesDiagram := addCreateItemButton(stager, confNotesDiagram)
		callbacksNotesDiagram.OnBeforeCommit = func() {
			diagramProcess.IsNotesNodeExpanded = true
		}
		if len(diagramNode.Menu.Buttons) > 0 {
			diagramNode.Menu.Buttons[0].Name = "Add Note"
			diagramNode.Menu.Buttons[0].ToolTipText = "Add a Note to \"" + diagramProcess.Name + "\""
		}
	}

	// Participants
	if len(process.Participants) > 0 {
		participantsNode := &tree.Node{
			Name:            "Participants",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramProcess.IsParticipantsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, participantsNode)
		participantsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramProcess.IsParticipantsNodeExpanded)

		confParticipantsNode := ItemButtonConfiguration[
			Participant, *Participant,
			Process, *Process,
		]{
			parentNode:                         participantsNode,
			sliceForNewAddedItem:               &process.Participants,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsParticipantsNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksParticipantsNode := addCreateItemButton(stager, confParticipantsNode)
		callbacksParticipantsNode.OnBeforeCommit = func() {
			diagramProcess.IsParticipantsNodeExpanded = true
		}
		if len(participantsNode.Menu.Buttons) > 0 {
			participantsNode.Menu.Buttons[0].Name = "Add Participant"
			participantsNode.Menu.Buttons[0].ToolTipText = "Add a Participant to \"" + diagramProcess.Name + "\""
		}

		for _, participant := range process.Participants {
			stager.treeParticipants(diagramProcess, participant, participantsNode)
		}
	}

	// external participants
	if len(process.ExternalParticipants) > 0 {
		externalParticipantsNode := &tree.Node{
			Name:            "External Participants",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramProcess.IsExternalParticipantsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, externalParticipantsNode)
		externalParticipantsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramProcess.IsExternalParticipantsNodeExpanded)

		confExternalParticipantsNode := ItemButtonConfiguration[
			Participant, *Participant,
			Process, *Process,
		]{
			parentNode:                         externalParticipantsNode,
			sliceForNewAddedItem:               &process.ExternalParticipants,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsExternalParticipantsNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksExternalParticipantsNode := addCreateItemButton(stager, confExternalParticipantsNode)
		callbacksExternalParticipantsNode.OnBeforeCommit = func() {
			diagramProcess.IsExternalParticipantsNodeExpanded = true
		}
		if len(externalParticipantsNode.Menu.Buttons) > 0 {
			externalParticipantsNode.Menu.Buttons[0].Name = "Add External Participant"
			externalParticipantsNode.Menu.Buttons[0].ToolTipText = "Add an External Participant to \"" + diagramProcess.Name + "\""
		}

		for _, participant := range process.ExternalParticipants {
			stager.treeExternalParticipants(diagramProcess, participant, externalParticipantsNode)
		}
	}

	//
	// DataFlows
	//
	if len(process.DataFlows) > 0 {
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

	// Notes
	//
	if process.GetOwningLibrary() != nil && len(process.GetOwningLibrary().RootNotes) > 0 {
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramProcess.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramProcess.IsNotesNodeExpanded)

		confNotesNode := ItemButtonConfiguration[
			Note, *Note,
			Library, *Library,
		]{
			parentNode:                         notesNode,
			sliceForNewAddedItem:               &process.GetOwningLibrary().RootNotes,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramProcess.IsNotesNodeExpanded,
			IsButtonInMenu:                     true,
		}
		callbacksNotesNode := addCreateItemButton(stager, confNotesNode)
		callbacksNotesNode.OnBeforeCommit = func() {
			diagramProcess.IsNotesNodeExpanded = true
		}
		if len(notesNode.Menu.Buttons) > 0 {
			notesNode.Menu.Buttons[0].Name = "Add Note"
			notesNode.Menu.Buttons[0].ToolTipText = "Add a Note to \"" + diagramProcess.Name + "\""
		}

		for _, note := range process.GetOwningLibrary().RootNotes {
			stager.treeNoteWithinDiagramProcess(diagramProcess, note, notesNode)
		}
	}
}
