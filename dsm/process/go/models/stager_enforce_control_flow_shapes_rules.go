package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceControlFlowShapesRules() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {

		// Build a set of tasks that have a shape in this diagram
		tasksInDiagram := make(map[*Task]bool)
		for _, taskShape := range diagram.Task_Shapes {
			if taskShape.Task != nil {
				tasksInDiagram[taskShape.Task] = true
			}
		}

		// 1. Check ControlFlow_Shapes
		var validControlFlowShapes []*ControlFlowShape
		for _, controlFlowShape := range diagram.ControlFlow_Shapes {
			isValid := true
			if controlFlowShape.ControlFlow != nil {
				if !tasksInDiagram[controlFlowShape.ControlFlow.Start] || !tasksInDiagram[controlFlowShape.ControlFlow.End] {
					isValid = false
				}
			} else {
				isValid = false
			}

			if isValid {
				validControlFlowShapes = append(validControlFlowShapes, controlFlowShape)
			} else {
				controlFlowShape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged control flow shape %s (missing start or end task shape)", controlFlowShape.GetName()))
				}
			}
		}
		if len(validControlFlowShapes) != len(diagram.ControlFlow_Shapes) {
			diagram.ControlFlow_Shapes = validControlFlowShapes
			needCommit = true
		}

		// 2. Check DataFlow_Shapes
		var validDataFlowShapes []*DataFlowShape
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			isValid := true
			if dataFlowShape.DataFlow != nil {
				if !tasksInDiagram[dataFlowShape.DataFlow.Start] || !tasksInDiagram[dataFlowShape.DataFlow.End] {
					isValid = false
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
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged data flow shape %s (missing start or end task shape)", dataFlowShape.GetName()))
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
