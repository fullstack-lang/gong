package models

func (stager *Stager) enforcePlantDiagramVerticalTorusStackShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram.VerticalTorusStackShape == nil {
			shape := (&VerticalTorusStackShape{
				Name: plantDiagram.Name + " - Vertical 3D Torus Stack",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.VerticalTorusStackShape = shape
			modified = true
		} else if plantDiagram.VaseDiagram.VerticalTorusStackShape.Name != plantDiagram.Name+" - Vertical 3D Torus Stack" {
			plantDiagram.VaseDiagram.VerticalTorusStackShape.Name = plantDiagram.Name + " - Vertical 3D Torus Stack"
			modified = true
		}
	}

	return modified
}
