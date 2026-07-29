package models

func (stager *Stager) enforcePlantDiagramPointsAndLines3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.PointsAndLines3DShape == nil {
			shape := (&PointsAndLines3DShape{
				Name: plantDiagram.Name + "-PointsAndLines3DShape",
			}).Stage(stager.stage)
			plantDiagram.PointsAndLines3DShape = shape
			modified = true
		}
	}

	return modified
}
