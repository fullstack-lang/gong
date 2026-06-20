package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceDataFlowRules() (needCommit bool) {
	stage := stager.stage

	// Build a reverse map from DataFlow to its owning Libraries
	dataFlowToLibraries := make(map[*DataFlow][]*Library)
	for _, library := range GetGongstrucsSorted[*Library](stage) {
		for _, dataFlow := range library.RootDataFlows {
			dataFlowToLibraries[dataFlow] = append(dataFlowToLibraries[dataFlow], library)
		}
	}

	for _, dataFlow := range GetGongstrucsSorted[*DataFlow](stage) {
		// Rule: A data flow belongs to one and only one library.
		owners := dataFlowToLibraries[dataFlow]
		if len(owners) == 0 {
			dataFlow.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan data flow \"%s\" (no library)", dataFlow.GetName()))
			}
			continue // No further checks needed for this orphan
		}
		if len(owners) > 1 {
			// To make the operation idempotent and clear, we remove the data flow from all its owners' lists
			// before unstaging it.
			for _, owner := range owners {
				if idx := slices.Index(owner.RootDataFlows, dataFlow); idx != -1 {
					owner.RootDataFlows = slices.Delete(owner.RootDataFlows, idx, idx+1)
				}
			}
			dataFlow.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged data flow \"%s\" with multiple libraries", dataFlow.GetName()))
			}
			continue // No further checks needed
		}

		// Rule : StartTask, EndTaskEnd, StartExternalParticipant & EndExternalParticipant
		// must be non nil according to the dataFlow type
		switch dataFlow.Type {
		case DataFlow_Task2Task:
			// Rule: A data flow must have a start task.
			if dataFlow.StartTask == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndTask == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no end, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
		case DataFlow_ExternalParticipant2Task:
			if dataFlow.StartExternalParticipant == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndTask == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no end, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
		case DataFlow_Task2ExternalParticipant:
			// Rule: A data flow must have a start task.
			if dataFlow.StartTask == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndExternalParticipant == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no end, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
		default:
			dataFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no type, unstaging",
					dataFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: A data flow must have an end task.

		// Rule: A data flow cannot connect a start or an end task.
		if dataFlow.Type == DataFlow_Task2Task &&
			(dataFlow.StartTask.IsStartTask || dataFlow.StartTask.IsEndTask || dataFlow.EndTask.IsStartTask || dataFlow.EndTask.IsEndTask) {
			dataFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" connects to a start or end task, unstaging", dataFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: Start and end task cannot belong to the same participant.
		if dataFlow.Type == DataFlow_Task2Task &&
			(dataFlow.StartTask.owningParticipant == dataFlow.EndTask.owningParticipant) {
			dataFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" connects tasks from the same participant, unstaging", dataFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// enforce name to be "StartTaskName to EndTaskName" for Task2Task data flow
		if dataFlow.Type == DataFlow_Task2Task {
			expectedName := "\"" + dataFlow.StartTask.GetName() + "\"" + " to " + "\"" + dataFlow.EndTask.GetName() + "\""
			if dataFlow.Name != expectedName {
				dataFlow.Name = expectedName
				needCommit = true
			}
		}

	}

	return
}
