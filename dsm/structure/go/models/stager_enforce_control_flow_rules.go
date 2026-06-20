package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceControlFlowRules() (needCommit bool) {
	stage := stager.stage

	// Build a reverse map from ControlFlow to its owning Parts
	controlFlowToParts := make(map[*ControlFlow][]*Part)
	for _, part := range GetGongstrucsSorted[*Part](stage) {
		for _, controlFlow := range part.ControlFlows {
			controlFlowToParts[controlFlow] = append(controlFlowToParts[controlFlow], part)
		}
	}

	for _, controlFlow := range GetGongstrucsSorted[*ControlFlow](stage) {
		// Rule: A control flow must have exactly one owning part.
		owners := controlFlowToParts[controlFlow]
		if len(owners) == 0 {
			controlFlow.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan control flow \"%s\" (no part)", controlFlow.GetName()))
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
					fmt.Sprintf("Unstaged control flow \"%s\" with multiple parts", controlFlow.GetName()))
			}
			continue // No further checks needed
		}

		// Rule: A control flow must have a start port.
		if controlFlow.Start == nil {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" has no start, unstaging",
					controlFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: A control flow must have an end port.
		if controlFlow.End == nil {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" has no end, unstaging",
					controlFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: Start and end ports of a control flow must belong to the same part.
		if controlFlow.Start.owningPart != controlFlow.End.owningPart {
			controlFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" connects ports from different parts, unstaging", controlFlow.GetName()))
			}
			needCommit = true
			continue
		}

		expectedName := "\"" + controlFlow.Start.GetName() + "\"" + " to " + "\"" + controlFlow.End.GetName() + "\""
		if controlFlow.Name != expectedName {
			controlFlow.Name = expectedName
			needCommit = true
		}
	}

	return
}
