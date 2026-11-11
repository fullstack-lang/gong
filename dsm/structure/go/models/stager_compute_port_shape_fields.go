package models

import "slices"

func (stager *Stager) computePortShapeFields() {
	stage := stager.stage

	// reset field
	for _, portShape := range stage.GetInstancesSorted[*PortShape]() {
		portShape.owningPartShape = nil
	}

	for _, diagramStructure := range stage.GetInstancesSorted[*DiagramStructure]() {
		for _, portShape := range diagramStructure.Port_Shapes {
			if portShape.Port != nil && portShape.Port.owningPart != nil {
				parentPS := effectivePortParentPartShape(diagramStructure, portShape)
				if parentPS != nil && parentPS.Part == portShape.Port.owningPart && slices.Contains(diagramStructure.Part_Shapes, parentPS) {
					portShape.owningPartShape = parentPS
				}
			}
		}
	}
}
