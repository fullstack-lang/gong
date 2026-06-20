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

		// Rule : StartPort, EndPortEnd, StartExternalPart & EndExternalPart
		// must be non nil according to the dataFlow type
		switch dataFlow.Type {
		case DataFlow_Port2Port:
			// Rule: A data flow must have a start port.
			if dataFlow.StartPort == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndPort == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no end, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
		case DataFlow_ExternalPart2Port:
			if dataFlow.StartExternalPart == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndPort == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no end, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
		case DataFlow_Port2ExternalPart:
			// Rule: A data flow must have a start port.
			if dataFlow.StartPort == nil {
				dataFlow.UnstageVoid(stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" has no start, unstaging",
						dataFlow.GetName()))
				}
				needCommit = true
				continue
			}
			if dataFlow.EndExternalPart == nil {
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

		// Rule: A data flow must have an end port.

		// Rule: A data flow cannot connect a start or an end port.
		if dataFlow.Type == DataFlow_Port2Port &&
			(dataFlow.StartPort.IsStartPort || dataFlow.StartPort.IsEndPort || dataFlow.EndPort.IsStartPort || dataFlow.EndPort.IsEndPort) {
			dataFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" connects to a start or end port, unstaging", dataFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// Rule: Start and end port cannot belong to the same part.
		if dataFlow.Type == DataFlow_Port2Port &&
			(dataFlow.StartPort.owningPart == dataFlow.EndPort.owningPart) {
			dataFlow.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Data flow \"%s\" connects ports from the same part, unstaging", dataFlow.GetName()))
			}
			needCommit = true
			continue
		}

		// enforce name to be "StartPortName to EndPortName" for Port2Port data flow
		if dataFlow.Type == DataFlow_Port2Port {
			expectedName := "\"" + dataFlow.StartPort.GetName() + "\"" + " to " + "\"" + dataFlow.EndPort.GetName() + "\""
			if dataFlow.Name != expectedName {
				dataFlow.Name = expectedName
				needCommit = true
			}
		}

	}

	return
}
