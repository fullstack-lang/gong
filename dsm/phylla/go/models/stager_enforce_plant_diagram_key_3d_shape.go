package models

func (stager *Stager) enforcePlantDiagramKey3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram.Key3DShape == nil {
			shape := (&Key3DShape{
				Name: plantDiagram.Name + "-Key3DShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.Key3DShape = shape
			modified = true
		}
	}

	return modified
}
