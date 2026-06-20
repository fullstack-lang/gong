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

		// Enforce System_Shapes presence
		if len(diagram.System_Shapes) == 0 && diagram.owningSystem != nil {
			systemShape := (&SystemShape{Name: diagram.owningSystem.Name}).Stage(stage)
			systemShape.System = diagram.owningSystem
			systemShape.X = 50
			systemShape.Y = 50
			
			// Default sizes if diagram has bounds
			systemShape.Width = diagram.Width - 100
			if systemShape.Width < 200 {
				systemShape.Width = 800
			}
			systemShape.Height = diagram.Height - 100
			if systemShape.Height < 200 {
				systemShape.Height = 600
			}
			
			diagram.System_Shapes = append(diagram.System_Shapes, systemShape)
			needCommit = true
		}

		// Rebuild map_System_SystemShape and remove duplicates
		diagram.map_System_SystemShape = make(map[*System]*SystemShape)
		var validSystemShapes []*SystemShape
		for _, systemShape := range diagram.System_Shapes {
			if systemShape.System != nil {
				if _, exists := diagram.map_System_SystemShape[systemShape.System]; exists {
					// Duplicate shape! Remove it
					systemShape.UnstageVoid(stage)
					needCommit = true
				} else {
					diagram.map_System_SystemShape[systemShape.System] = systemShape
					validSystemShapes = append(validSystemShapes, systemShape)
				}
			} else {
				validSystemShapes = append(validSystemShapes, systemShape)
			}
		}
		if len(validSystemShapes) != len(diagram.System_Shapes) {
			diagram.System_Shapes = validSystemShapes
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
