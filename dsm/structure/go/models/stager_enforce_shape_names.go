package models

func (stager *Stager) enforceShapeNames() (needCommit bool) {
	stage := stager.stage

	// 1. PartShapes
	for diagram := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stage) {
		for _, partShape := range diagram.Part_Shapes {
			if partShape.Part != nil {
				expectedName := partShape.Part.GetName() + "-" + diagram.GetName()
				if partShape.Name != expectedName {
					partShape.Name = expectedName
					needCommit = true
				}
			}
		}

		// 2. LinkShapes
		for _, linkShape := range diagram.Link_Shapes {
			if linkShape.Link != nil {
				var startName, endName string
				if linkShape.Link.Source != nil {
					startName = linkShape.Link.Source.GetName()
				} else {
					startName = "undefined"
				}
				if linkShape.Link.Target != nil {
					endName = linkShape.Link.Target.GetName()
				} else {
					endName = "undefined"
				}
				expectedName := startName + " to " + endName
				if linkShape.Name != expectedName {
					linkShape.Name = expectedName
					needCommit = true
				}
			}
		}
	}

	return needCommit
}
