package models

func (stager *Stager) enforcePlantDiagramRendered3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.Rendered3DShape == nil {
			shape := (&Rendered3DShape{
				Name:    plantDiagram.Name + "-Rendered3DShape",
				ViewX:   300,
				ViewY:   300,
				ViewZ:   300,
				TargetX: 0,
				TargetY: 100,
				TargetZ: 0,
				Fov:     45,
			}).Stage(stager.stage)
			plantDiagram.Rendered3DShape = shape
			modified = true
		}
	}

	return modified
}
