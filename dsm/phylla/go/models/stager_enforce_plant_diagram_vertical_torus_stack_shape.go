package models

func (stager *Stager) enforcePlantDiagramVerticalTorusStackShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VerticalTorusStackShape == nil {
			shape := (&VerticalTorusStackShape{
				Name: plantDiagram.Name + " - Vertical 3D Torus Stack",
			}).Stage(stager.stage)
			plantDiagram.VerticalTorusStackShape = shape
			modified = true
		} else if plantDiagram.VerticalTorusStackShape.Name != plantDiagram.Name+" - Vertical 3D Torus Stack" {
			plantDiagram.VerticalTorusStackShape.Name = plantDiagram.Name + " - Vertical 3D Torus Stack"
			modified = true
		}
	}

	return modified
}
