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
		IsExpanded:              slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, externalParticipant),
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

	node.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			externalParticipant.Name = frontNode.Name
			externalParticipant.isInRenameMode = false
			stage.Commit()
			return
		}
		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(externalParticipant, diagramProcess, &diagramProcess.ExternalParticipant_Shapes, stage)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramProcess.ExternalParticipant_Shapes, shape)
			diagramProcess.ExternalParticipant_Shapes = slices.Delete(diagramProcess.ExternalParticipant_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if !slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, externalParticipant) {
					diagramProcess.ExternalParticipantWhoseNodeIsExpanded = append(diagramProcess.ExternalParticipantWhoseNodeIsExpanded, externalParticipant)
				}
			} else {
				if idx := slices.Index(diagramProcess.ExternalParticipantWhoseNodeIsExpanded, externalParticipant); idx != -1 {
					diagramProcess.ExternalParticipantWhoseNodeIsExpanded = slices.Delete(diagramProcess.ParticipantWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(externalParticipant, GetPointerToGongstructName[*Participant]())
		stage.Commit()
	}

}
