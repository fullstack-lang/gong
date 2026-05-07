package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDataFlowShapesRules() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {

		// Build a set of tasks that have a shape in this diagram
		tasksInDiagram := make(map[*Task]bool)
		for _, taskShape := range diagram.Task_Shapes {
			if taskShape.Task != nil {
				tasksInDiagram[taskShape.Task] = true
			}
		}

		// Build a set of external participants that have a shape in this diagram
		externalParticipantsInDiagram := make(map[*Participant]bool)
		for _, epShape := range diagram.ExternalParticipant_Shapes {
			if epShape.Participant != nil {
				externalParticipantsInDiagram[epShape.Participant] = true
			}
		}

		// Check DataFlow_Shapes
		var validDataFlowShapes []*DataFlowShape
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
					if df.StartExternalParticipant == nil || !externalParticipantsInDiagram[df.StartExternalParticipant] ||
						df.EndTask == nil || !tasksInDiagram[df.EndTask] {
						isValid = false
					}
				case DataFlow_Task2ExternalParticipant:
					if df.StartTask == nil || !tasksInDiagram[df.StartTask] ||
						df.EndExternalParticipant == nil || !externalParticipantsInDiagram[df.EndExternalParticipant] {
						isValid = false
					}
				}
			} else {
				isValid = false
			}

			if isValid {
				validDataFlowShapes = append(validDataFlowShapes, dataFlowShape)
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
	}

	return
}
