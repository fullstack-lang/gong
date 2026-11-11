package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceDataFlowShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {

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
		dataFlowsInDiagram := make(map[*DataFlow]bool)
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

		// Ensure Protocol_Shapes match an active DataFlowShape and its current Protocol
		var validProtocolShapes []*ProtocolShape
		for _, protocolShape := range diagram.Protocol_Shapes {
			if protocolShape.DataFlow != nil &&
				protocolShape.Protocol != nil &&
				dataFlowsInDiagram[protocolShape.DataFlow] &&
				protocolShape.DataFlow.Protocol == protocolShape.Protocol {
				validProtocolShapes = append(validProtocolShapes, protocolShape)
			} else {
				protocolShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}
		if len(validProtocolShapes) != len(diagram.Protocol_Shapes) {
			diagram.Protocol_Shapes = validProtocolShapes
			needCommit = true
		}

		// Ensure ProtocolLayer_Shapes match an active DataFlowShape and its current Protocol's ProtocolLayers
		var validProtocolLayerShapes []*ProtocolLayerShape
		for _, layerShape := range diagram.ProtocolLayer_Shapes {
			if layerShape.DataFlow != nil &&
				layerShape.ProtocolLayer != nil &&
				dataFlowsInDiagram[layerShape.DataFlow] &&
				layerShape.DataFlow.Protocol != nil &&
				slices.Contains(layerShape.DataFlow.Protocol.ProtocolLayers, layerShape.ProtocolLayer) {
				validProtocolLayerShapes = append(validProtocolLayerShapes, layerShape)
			} else {
				layerShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}
		if len(validProtocolLayerShapes) != len(diagram.ProtocolLayer_Shapes) {
			diagram.ProtocolLayer_Shapes = validProtocolLayerShapes
			needCommit = true
		}
	}

	return
}
