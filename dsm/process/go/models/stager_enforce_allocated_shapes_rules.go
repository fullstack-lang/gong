package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceAllocatedShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		seenResourceKeys := make(map[allocatedResourceShapeKey]bool)
		var validAllocatedResourceShapes []*AllocatedResourceShape
		for _, shape := range diagram.AllocatedResourceShapes {
			key := allocatedResourceShapeKey{
				participant:         shape.Participant,
				externalParticipant: shape.ExternalParticipant,
				resource:            shape.Resource,
			}
			hasValidParticipantOwner := shape.Participant != nil &&
				shape.ExternalParticipant == nil &&
				shape.Resource != nil &&
				slices.Contains(shape.Participant.Resources, shape.Resource)
			hasValidExternalParticipantOwner := shape.ExternalParticipant != nil &&
				shape.Participant == nil &&
				shape.Resource != nil &&
				slices.Contains(shape.ExternalParticipant.Resources, shape.Resource)

			isValid := (hasValidParticipantOwner || hasValidExternalParticipantOwner) && !seenResourceKeys[key]
			if isValid {
				seenResourceKeys[key] = true
				validAllocatedResourceShapes = append(validAllocatedResourceShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged allocated resource shape %s (invalid or duplicate allocation)", shape.GetName()))
				}
			}
		}
		if len(validAllocatedResourceShapes) != len(diagram.AllocatedResourceShapes) {
			diagram.AllocatedResourceShapes = validAllocatedResourceShapes
			needCommit = true
		}

		seenProcessKeys := make(map[allocatedProcessShapeKey]bool)
		var validAllocatedProcessShapes []*AllocatedProcessShape
		for _, shape := range diagram.AllocatedProcessShapes {
			key := allocatedProcessShapeKey{
				externalParticipant: shape.ExternalParticipant,
				process:             shape.Process,
			}
			isValid := shape.ExternalParticipant != nil &&
				shape.Process != nil &&
				slices.Contains(shape.ExternalParticipant.Processes, shape.Process) &&
				!seenProcessKeys[key]
			if isValid {
				seenProcessKeys[key] = true
				validAllocatedProcessShapes = append(validAllocatedProcessShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged allocated process shape %s (invalid or duplicate allocation)", shape.GetName()))
				}
			}
		}
		if len(validAllocatedProcessShapes) != len(diagram.AllocatedProcessShapes) {
			diagram.AllocatedProcessShapes = validAllocatedProcessShapes
			needCommit = true
		}
	}

	return
}
