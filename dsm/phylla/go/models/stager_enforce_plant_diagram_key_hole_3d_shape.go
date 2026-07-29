package models

func (stager *Stager) enforcePlantDiagramKeyHole3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.KeyHole3DShape == nil {
			shape := (&KeyHole3DShape{
				Name: plantDiagram.Name + "-KeyHole3DShape",
			}).Stage(stager.stage)
			plantDiagram.KeyHole3DShape = shape
			modified = true
		}
	}

	return modified
}
