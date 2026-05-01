package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceControlFlowRules() (needCommit bool) {
	stage := stager.stage

	// Build a reverse map from ControlFlow to its owning Participants
	controlFlowToParticipants := make(map[*ControlFlow][]*Participant)
	for _, participant := range GetGongstrucsSorted[*Participant](stage) {
		for _, controlFlow := range participant.ControlFlows {
			controlFlowToParticipants[controlFlow] = append(controlFlowToParticipants[controlFlow], participant)
		}
	}

	for _, controlFlow := range GetGongstrucsSorted[*ControlFlow](stage) {
		// Rule: A control flow must have exactly one owning participant.
		owners := controlFlowToParticipants[controlFlow]
		if len(owners) == 0 {
			controlFlow.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan control flow \"%s\" (no participant)", controlFlow.GetName()))
			}
			continue // No further checks needed for this orphan
		}
		if len(owners) > 1 {
			// To make the operation idempotent and clear, we remove the control flow from all its owners' lists
			// before unstaging it.
			for _, owner := range owners {
				if idx := slices.Index(owner.ControlFlows, controlFlow); idx != -1 {
					owner.ControlFlows = slices.Delete(owner.ControlFlows, idx, idx+1)
				}
			}
			controlFlow.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged control flow \"%s\" with multiple participants", controlFlow.GetName()))
			}
			continue // No further checks needed
		}

		// Rule: A control flow must have a start task.
		if controlFlow.Start == nil {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" has no start, unstaging",
					controlFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: A control flow must have an end task.
		if controlFlow.End == nil {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" has no end, unstaging",
					controlFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: Start and end tasks of a control flow must belong to the same participant.
		if controlFlow.Start.owningParticipant != controlFlow.End.owningParticipant {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" connects tasks from different participants, unstaging", controlFlow.GetName()))
			}
			needCommit = true
			continue
		}
	}

	return
}
