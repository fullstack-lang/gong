package models

func (stager *Stager) enforcePlantDiagramPartiallyRotatedTorusShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.PartiallyRotatedTorusShape == nil {
			shape := (&PartiallyRotatedTorusShape{
				Name:    plantDiagram.Name + "-PartiallyRotatedTorusShape",
			}).Stage(stager.stage)
			plantDiagram.PartiallyRotatedTorusShape = shape
			modified = true
		}
	}

	return modified
}
