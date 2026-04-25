package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeParticipants(
	participant *Participant,
	parentNode *tree.Node,
	participantsWhoseNodeIsExpanded *[]*Participant) {

	participantNode := &tree.Node{
		Name:            participant.Name,
		IsExpanded:      slices.Contains(*participantsWhoseNodeIsExpanded, participant),
		IsNodeClickable: true,
		IsInEditMode:    participant.isInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, participantNode)
	addRenameButton(participant, participantNode, stager)

	participantNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			participant.Name = frontNode.Name
			participant.isInRenameMode = false
			stager.stage.Commit()
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
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(participant, GetPointerToGongstructName[*Participant]())
		stager.stage.Commit()
	}

}
