package models

func (stager *Stager) enforceDiagramMaps() bool {
	needCommit := false
	stage := stager.stage

	for diagram := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stage) {
		// Rebuild map_Part_PartShape and remove duplicates
		diagram.map_Part_PartShape = make(map[*Part]*PartShape)
		var validPartShapes []*PartShape
		for _, partShape := range diagram.Part_Shapes {
			if partShape.Part != nil {
				if _, exists := diagram.map_Part_PartShape[partShape.Part]; exists {
					// Duplicate shape! Remove it
					partShape.UnstageVoid(stage)
					needCommit = true
				} else {
					diagram.map_Part_PartShape[partShape.Part] = partShape
					validPartShapes = append(validPartShapes, partShape)
				}
			} else {
				// Orphan shape will be removed by enforceShapesAbstractConsistency
				validPartShapes = append(validPartShapes, partShape)
			}
		}
		if len(validPartShapes) != len(diagram.Part_Shapes) {
			diagram.Part_Shapes = validPartShapes
			needCommit = true
		}

		// Rebuild map_Link_LinkAssociationShape and remove duplicates
		diagram.map_Link_LinkAssociationShape = make(map[*Link]*LinkAssociationShape)
		var validLinkShapes []*LinkAssociationShape
		for _, linkShape := range diagram.Link_Shapes {
			if linkShape.Link != nil {
				if _, exists := diagram.map_Link_LinkAssociationShape[linkShape.Link]; exists {
					// Duplicate link shape! Remove it
					linkShape.UnstageVoid(stage)
					needCommit = true
				} else {
					diagram.map_Link_LinkAssociationShape[linkShape.Link] = linkShape
					validLinkShapes = append(validLinkShapes, linkShape)
				}
			} else {
				// Orphan shape will be removed by enforceShapesAbstractConsistency
				validLinkShapes = append(validLinkShapes, linkShape)
			}
		}
		if len(validLinkShapes) != len(diagram.Link_Shapes) {
			diagram.Link_Shapes = validLinkShapes
			needCommit = true
		}
	}

	return needCommit
}
