package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeParticipants(
	diagram *DiagramProcess,
	participant *Participant,
	parentNode *tree.Node,
	participantsWhoseNodeIsExpanded *[]*Participant) {

	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagram.map_Participant_ParticipantShape[participant]
	node := &tree.Node{
		Name:              participant.Name,
		IsExpanded:        slices.Contains(*participantsWhoseNodeIsExpanded, participant),
		IsNodeClickable:   true,
		IsInEditMode:      participant.isInRenameMode,
		HasCheckboxButton: true,
		IsChecked:         ok,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(participant, node, stager)

	if shape, ok := diagram.map_Participant_ParticipantShape[participant]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagram.GetName(),
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

	node.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			participant.Name = frontNode.Name
			participant.isInRenameMode = false
			stage.Commit()
			return
		}
		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(participant, diagram, &diagram.Participant_Shapes, stage)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagram.Participant_Shapes, shape)
			diagram.Participant_Shapes = slices.Delete(diagram.Participant_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if !slices.Contains(*participantsWhoseNodeIsExpanded, participant) {
					*participantsWhoseNodeIsExpanded = append(*participantsWhoseNodeIsExpanded, participant)
				}
			} else {
				if idx := slices.Index(*participantsWhoseNodeIsExpanded, participant); idx != -1 {
					*participantsWhoseNodeIsExpanded = slices.Delete(*participantsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(participant, GetPointerToGongstructName[*Participant]())
		stage.Commit()
	}

}
