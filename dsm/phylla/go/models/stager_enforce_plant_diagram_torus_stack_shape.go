package models

func (stager *Stager) enforcePlantDiagramTorusStackShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram.TorusStackShape == nil {
			shape := (&TorusStackShape{
				Name:    plantDiagram.Name + "-TorusStackShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.TorusStackShape = shape
			modified = true
		}
	}

	return modified
}
