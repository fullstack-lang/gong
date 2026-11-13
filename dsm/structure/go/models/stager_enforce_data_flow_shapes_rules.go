package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceDataFlowShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {

		// Build a set of external parts that have a shape in this diagram
		externalPartsInDiagram := make(map[*Part]bool)
		for _, epShape := range diagram.ExternalPart_Shapes {
			if epShape.Part != nil {
				externalPartsInDiagram[epShape.Part] = true
			}
		}

		// Check DataFlow_Shapes
		var validDataFlowShapes []*DataFlowShape
		validDataFlowShapeSet := make(map[*DataFlowShape]bool)
		dataFlowsInDiagram := make(map[*DataFlow]bool)

		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			isValid := true
			if dataFlowShape.DataFlow != nil {
				df := dataFlowShape.DataFlow
				if dataFlowShape.ParentPartShape != nil {
					if !slices.Contains(diagram.Part_Shapes, dataFlowShape.ParentPartShape) ||
						dataFlowShape.ParentPartShape.Part == nil ||
						dataFlowShape.ParentPartShape.Part.Type == nil ||
						!slices.Contains(dataFlowShape.ParentPartShape.Part.Type.DataFlows, df) {
						isValid = false
					}
				}
				if isValid {
					_, startOK, _, endOK := resolveDataFlowEndpointPortShapes(diagram, dataFlowShape)
					switch df.Type {
					case DataFlow_Port2Port:
						if df.StartPort == nil || !startOK || df.EndPort == nil || !endOK {
							isValid = false
						}
					case DataFlow_ExternalPart2Port:
						if df.StartExternalPart == nil || !externalPartsInDiagram[df.StartExternalPart] ||
							df.EndPort == nil || !endOK {
							isValid = false
						}
					case DataFlow_Port2ExternalPart:
						if df.StartPort == nil || !startOK ||
							df.EndExternalPart == nil || !externalPartsInDiagram[df.EndExternalPart] {
							isValid = false
						}
					}
				}
			} else {
				isValid = false
			}

			if isValid {
				validDataFlowShapes = append(validDataFlowShapes, dataFlowShape)
				validDataFlowShapeSet[dataFlowShape] = true
				dataFlowsInDiagram[dataFlowShape.DataFlow] = true
				if dataFlowShape.Name != dataFlowShape.DataFlow.Name {
					dataFlowShape.Name = dataFlowShape.DataFlow.Name
					needCommit = true
				}
			} else {
				// Cascade unstage any DataShape / ProtocolShape / ProtocolLayerShape explicitly scoped to this DataFlowShape
				// before stage.Clean() zeroes their DataFlowShape pointer.
				for _, ds := range diagram.Data_Shapes {
					if ds != nil && ds.DataFlowShape == dataFlowShape {
						ds.UnstageVoid(stager.stage)
					}
				}
				for _, ps := range diagram.Protocol_Shapes {
					if ps != nil && ps.DataFlowShape == dataFlowShape {
						ps.UnstageVoid(stager.stage)
					}
				}
				for _, ls := range diagram.ProtocolLayer_Shapes {
					if ls != nil && ls.DataFlowShape == dataFlowShape {
						ls.UnstageVoid(stager.stage)
					}
				}
				for _, cp := range dataFlowShape.ControlPointShapes {
					if cp != nil {
						cp.UnstageVoid(stager.stage)
					}
				}

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

		// Ensure Data_Shapes with explicit DataFlowShape match an active DataFlowShape
		var validDataShapes []*DataShape
		for _, dataShape := range diagram.Data_Shapes {
			if dataShape.DataFlowShape != nil && !validDataFlowShapeSet[dataShape.DataFlowShape] {
				dataShape.UnstageVoid(stager.stage)
				needCommit = true
			} else {
				validDataShapes = append(validDataShapes, dataShape)
			}
		}
		if len(validDataShapes) != len(diagram.Data_Shapes) {
			diagram.Data_Shapes = validDataShapes
			needCommit = true
		}

		// Ensure Protocol_Shapes match an active DataFlowShape and its current Protocol
		var validProtocolShapes []*ProtocolShape
		for _, protocolShape := range diagram.Protocol_Shapes {
			effDFS := effectiveDataFlowShape(diagram, protocolShape.DataFlow, protocolShape.DataFlowShape)
			isActive := (protocolShape.DataFlowShape != nil && validDataFlowShapeSet[protocolShape.DataFlowShape]) ||
				(protocolShape.DataFlowShape == nil && ((effDFS != nil && validDataFlowShapeSet[effDFS]) || dataFlowsInDiagram[protocolShape.DataFlow]))
			if protocolShape.DataFlow != nil &&
				protocolShape.Protocol != nil &&
				isActive &&
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
			effDFS := effectiveDataFlowShape(diagram, layerShape.DataFlow, layerShape.DataFlowShape)
			isActive := (layerShape.DataFlowShape != nil && validDataFlowShapeSet[layerShape.DataFlowShape]) ||
				(layerShape.DataFlowShape == nil && ((effDFS != nil && validDataFlowShapeSet[effDFS]) || dataFlowsInDiagram[layerShape.DataFlow]))
			if layerShape.DataFlow != nil &&
				layerShape.ProtocolLayer != nil &&
				isActive &&
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
