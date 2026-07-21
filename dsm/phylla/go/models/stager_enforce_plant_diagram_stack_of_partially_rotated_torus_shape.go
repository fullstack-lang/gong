package models

func (stager *Stager) enforcePlantDiagramStackOfPartiallyRotatedTorusShape() bool {
	modified := false

	for _, plantDiagram := range GetGongstrucsSorted[*PlantDiagram](stager.stage) {
		if plantDiagram.StackOfPartiallyRotatedTorusShape == nil {
			shape := (&StackOfPartiallyRotatedTorusShape{
				Name:    plantDiagram.Name + "-StackOfPartiallyRotatedTorusShape",
			}).Stage(stager.stage)
			plantDiagram.StackOfPartiallyRotatedTorusShape = shape
			modified = true
		}
	}

	return modified
}
