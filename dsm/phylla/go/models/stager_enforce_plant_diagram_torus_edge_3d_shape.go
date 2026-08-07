package models

func (stager *Stager) enforcePlantDiagramTorusEdge3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram.TorusEdge3DShape == nil {
			shape := (&TorusEdge3DShape{
				Name: plantDiagram.Name + "-TorusEdge3DShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.TorusEdge3DShape = shape
			modified = true
		}
	}

	return modified
}
