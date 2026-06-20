package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceControlFlowShapesRules() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {

		// Build a set of ports that have a shape in this diagram
		portsInDiagram := make(map[*Port]bool)
		for _, portShape := range diagram.Port_Shapes {
			if portShape.Port != nil {
				portsInDiagram[portShape.Port] = true
			}
		}

		// 1. Check ControlFlow_Shapes
		var validControlFlowShapes []*ControlFlowShape
		for _, controlFlowShape := range diagram.ControlFlow_Shapes {
			isValid := true
			if controlFlowShape.ControlFlow != nil {
				if !portsInDiagram[controlFlowShape.ControlFlow.Start] || !portsInDiagram[controlFlowShape.ControlFlow.End] {
					isValid = false
				}
			} else {
				isValid = false
			}

			if isValid {
				validControlFlowShapes = append(validControlFlowShapes, controlFlowShape)
				if controlFlowShape.Name != controlFlowShape.ControlFlow.Name {
					controlFlowShape.Name = controlFlowShape.ControlFlow.Name
					needCommit = true
				}
			} else {
				controlFlowShape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged control flow shape %s (missing start or end port shape)", controlFlowShape.GetName()))
				}
			}
		}
		if len(validControlFlowShapes) != len(diagram.ControlFlow_Shapes) {
			diagram.ControlFlow_Shapes = validControlFlowShapes
			needCommit = true
		}
	}

	return
}
