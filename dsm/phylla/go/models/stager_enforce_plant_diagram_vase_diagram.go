package models

func (stager *Stager) enforcePlantDiagramVaseDiagram() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram == nil {
			vd := (&VaseDiagram{
				Name: plantDiagram.Name + "-VaseDiagram",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram = vd
			modified = true
		}
	}

	return modified
}
