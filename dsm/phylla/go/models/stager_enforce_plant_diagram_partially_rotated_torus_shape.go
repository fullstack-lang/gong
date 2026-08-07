package models

func (stager *Stager) enforcePlantDiagramPartiallyRotatedTorusShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram != nil && plantDiagram.VaseDiagram.PartiallyRotatedTorusShape == nil {
			shape := (&PartiallyRotatedTorusShape{
				Name: plantDiagram.Name + " - Partially Rotated 3D Torus",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.PartiallyRotatedTorusShape = shape
			modified = true
		}
	}

	return modified
}
