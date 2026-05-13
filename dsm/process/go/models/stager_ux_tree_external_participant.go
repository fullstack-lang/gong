package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeExternalParticipants(
	diagramProcess *DiagramProcess,
	externalParticipant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagramProcess.map_Participant_ExternalParticipantShape[externalParticipant]
	node := &tree.Node{
		Name:                    externalParticipant.Name,
		IsExpanded:              slices.Contains(diagramProcess.ExternalParticipantWhoseNodeIsExpanded, externalParticipant),
		IsNodeClickable:         true,
		IsInEditMode:            externalParticipant.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the participant shape"
			}
			return "Click to create a participant shape for this participant within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(externalParticipant, node, stager)

	if ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnNameChange = stager.onNameChange(externalParticipant)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalParticipant, &diagramProcess.ExternalParticipantWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(externalParticipant, diagramProcess, &diagramProcess.ExternalParticipant_Shapes, stage)

			stage.Commit()
			return
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramProcess.ExternalParticipant_Shapes, shape)
			diagramProcess.ExternalParticipant_Shapes = slices.Delete(diagramProcess.ExternalParticipant_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, externalParticipant)

	// allocated resources
	allocatedResourcesNode := &tree.Node{
		Name:            "Allocated Resources",
		FontStyle:       tree.ITALIC,
		IsExpanded:      externalParticipant.IsResourcesNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, allocatedResourcesNode)
	allocatedResourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&externalParticipant.IsResourcesNodeExpanded)

	for _, resource := range externalParticipant.Resources {
		stager.treeAllocatedResourceWithinDiagramWithinParticipant(diagramProcess, resource, externalParticipant, allocatedResourcesNode)
	}

	// processes
	processesNode := &tree.Node{
		Name:            "Allocated Processes",
		FontStyle:       tree.ITALIC,
		IsExpanded:      externalParticipant.IsProcessesNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, processesNode)
	processesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&externalParticipant.IsProcessesNodeExpanded)

	for _, process := range externalParticipant.Processes {
		stager.treeProcessesWithinDiagramProcessWithinParticipant(diagramProcess, process, externalParticipant, processesNode)
	}

	// out data flows and in data flows
	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramProcess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, externalParticipant),
	}
	node.Children = append(node.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalParticipant, &diagramProcess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range externalParticipant.outDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramProcess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, externalParticipant),
	}
	node.Children = append(node.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalParticipant, &diagramProcess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range externalParticipant.inDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeInDataFlows)
	}
}
