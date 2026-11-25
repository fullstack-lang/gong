package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDataFlowShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {

		// Build a set of tasks that have a shape in this diagram
		tasksInDiagram := make(map[*Task]bool)
		for _, taskShape := range diagram.Task_Shapes {
			if taskShape.Task != nil {
				tasksInDiagram[taskShape.Task] = true
			}
		}

		// Build a set of external participants that have a shape in this diagram
		externalParticipantsInDiagram := make(map[*ExternalParticipant]bool)
		for _, epShape := range diagram.ExternalParticipant_Shapes {
			if epShape.ExternalParticipant != nil {
				externalParticipantsInDiagram[epShape.ExternalParticipant] = true
			}
		}

		// Build a set of data stores that have a shape in this diagram
		dataStoresInDiagram := make(map[*DataStore]bool)
		for _, dsShape := range diagram.DataStore_Shapes {
			if dsShape.DataStore != nil {
				dataStoresInDiagram[dsShape.DataStore] = true
			}
		}

		// Check DataFlow_Shapes
		var validDataFlowShapes []*DataFlowShape
		dataFlowsInDiagram := make(map[*DataFlow]bool)
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			isValid := true
			if dataFlowShape.DataFlow != nil {
				df := dataFlowShape.DataFlow
				switch df.Type {
				case DataFlow_Task2Task:
					if df.StartTask == nil || !tasksInDiagram[df.StartTask] ||
						df.EndTask == nil || !tasksInDiagram[df.EndTask] {
						isValid = false
					}
				case DataFlow_ExternalParticipant2Task:
					if df.StartParticipant == nil || !externalParticipantsInDiagram[df.StartParticipant] ||
						df.EndTask == nil || !tasksInDiagram[df.EndTask] {
						isValid = false
					}
				case DataFlow_Task2ExternalParticipant:
					if df.StartTask == nil || !tasksInDiagram[df.StartTask] ||
						df.EndParticipant == nil || !externalParticipantsInDiagram[df.EndParticipant] {
						isValid = false
					}
				case DataFlow_DataStore2Task:
					if df.StartDataStore == nil || !dataStoresInDiagram[df.StartDataStore] ||
						df.EndTask == nil || !tasksInDiagram[df.EndTask] {
						isValid = false
					}
				case DataFlow_Task2DataStore:
					if df.StartTask == nil || !tasksInDiagram[df.StartTask] ||
						df.EndDataStore == nil || !dataStoresInDiagram[df.EndDataStore] {
						isValid = false
					}
				default:
					isValid = false
				}
			} else {
				isValid = false
			}

			if isValid {
				validDataFlowShapes = append(validDataFlowShapes, dataFlowShape)
				dataFlowsInDiagram[dataFlowShape.DataFlow] = true
				if dataFlowShape.Name != dataFlowShape.DataFlow.Name {
					dataFlowShape.Name = dataFlowShape.DataFlow.Name
					needCommit = true
				}
			} else {
				dataFlowShape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged data flow shape %s (missing start or end shape)", dataFlowShape.GetName()))
				}
			}
		}
		if len(validDataFlowShapes) != len(diagram.DataFlow_Shapes) {
			diagram.DataFlow_Shapes = validDataFlowShapes
			needCommit = true
		}

		controlFlowsInDiagram := make(map[*ControlFlow]bool)
		for _, cfShape := range diagram.ControlFlow_Shapes {
			if cfShape != nil && cfShape.ControlFlow != nil {
				controlFlowsInDiagram[cfShape.ControlFlow] = true
			}
		}

		// Ensure Data_Shapes match either an active DataFlowShape or an active ControlFlowShape
		// (exclusively) and a Data reachable from that flow's Datas.
		var validDataShapes []*DataShape
		for _, dataShape := range diagram.Data_Shapes {
			if dataShape == nil {
				needCommit = true
				continue
			}
			hasValidDataFlow := dataShape.Data != nil &&
				dataShape.DataFlow != nil &&
				dataShape.ControlFlow == nil &&
				dataFlowsInDiagram[dataShape.DataFlow] &&
				isDataReachableFromDataFlow(dataShape.DataFlow, dataShape.Data)
			hasValidControlFlow := dataShape.Data != nil &&
				dataShape.ControlFlow != nil &&
				dataShape.DataFlow == nil &&
				controlFlowsInDiagram[dataShape.ControlFlow] &&
				isDataReachableFromControlFlow(dataShape.ControlFlow, dataShape.Data)

			if hasValidDataFlow || hasValidControlFlow {
				validDataShapes = append(validDataShapes, dataShape)
			} else {
				dataShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}
		if len(validDataShapes) != len(diagram.Data_Shapes) {
			diagram.Data_Shapes = validDataShapes
			needCommit = true
		}
	}

	return
}
