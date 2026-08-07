package models

func (stager *Stager) enforcePlantDiagramKeyHole3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram != nil && plantDiagram.VaseDiagram.KeyHole3DShape == nil {
			shape := (&KeyHole3DShape{
				Name: plantDiagram.Name + "-KeyHole3DShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.KeyHole3DShape = shape
			modified = true
		}
	}

	return modified
}
