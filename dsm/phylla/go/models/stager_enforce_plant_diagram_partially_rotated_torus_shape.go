package models

func (stager *Stager) enforcePlantDiagramPartiallyRotatedTorusShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram.PartiallyRotatedTorusShape == nil {
			shape := (&PartiallyRotatedTorusShape{
				Name:    plantDiagram.Name + "-PartiallyRotatedTorusShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.PartiallyRotatedTorusShape = shape
			modified = true
		}
	}

	return modified
}
