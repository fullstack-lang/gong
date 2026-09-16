package models

func (stager *Stager) computePortShapeFields() {
	stage := stager.stage

	// reset field
	for _, portShape := range stage.GetInstancesSorted[*PortShape]() {
		portShape.owningPartShape = nil
	}

	for _, diagramStructure := range stage.GetInstancesSorted[*DiagramStructure]() {
		for _, portShape := range diagramStructure.Port_Shapes {
			if portShape.Port != nil && portShape.Port.owningPart != nil {
				if partShape, ok := diagramStructure.map_Part_PartShape[portShape.Port.owningPart]; ok {
					portShape.owningPartShape = partShape
				}
			}
		}
	}
}
