package models

func (stager *Stager) enforcePlantDiagramRendered3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.Rendered3DShape == nil {
			shape := (&Rendered3DShape{
				Name:  plantDiagram.Name + "-Rendered3DShape",
				ViewX: 0,
				ViewY: 300,
				ViewZ: 0,
			}).Stage(stager.stage)
			plantDiagram.Rendered3DShape = shape
			modified = true
		}
	}

	return modified
}
