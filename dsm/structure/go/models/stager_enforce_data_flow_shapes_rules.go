package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDataFlowShapesRules() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {

		// Build a set of ports that have a shape in this diagram
		portsInDiagram := make(map[*Port]bool)
		for _, portShape := range diagram.Port_Shapes {
			if portShape.Port != nil {
				portsInDiagram[portShape.Port] = true
			}
		}

		// Build a set of external parts that have a shape in this diagram
		externalPartsInDiagram := make(map[*Part]bool)
		for _, epShape := range diagram.ExternalPart_Shapes {
			if epShape.Part != nil {
				externalPartsInDiagram[epShape.Part] = true
			}
		}

		// Check DataFlow_Shapes
		var validDataFlowShapes []*DataFlowShape
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			isValid := true
			if dataFlowShape.DataFlow != nil {
				df := dataFlowShape.DataFlow
				switch df.Type {
				case DataFlow_Port2Port:
					if df.StartPort == nil || !portsInDiagram[df.StartPort] ||
						df.EndPort == nil || !portsInDiagram[df.EndPort] {
						isValid = false
					}
				case DataFlow_ExternalPart2Port:
					if df.StartExternalPart == nil || !externalPartsInDiagram[df.StartExternalPart] ||
						df.EndPort == nil || !portsInDiagram[df.EndPort] {
						isValid = false
					}
				case DataFlow_Port2ExternalPart:
					if df.StartPort == nil || !portsInDiagram[df.StartPort] ||
						df.EndExternalPart == nil || !externalPartsInDiagram[df.EndExternalPart] {
						isValid = false
					}
				}
			} else {
				isValid = false
			}

			if isValid {
				validDataFlowShapes = append(validDataFlowShapes, dataFlowShape)
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
	}

	return
}
