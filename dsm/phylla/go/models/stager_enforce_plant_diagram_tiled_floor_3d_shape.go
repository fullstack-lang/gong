package models

func (stager *Stager) enforcePlantDiagramTiledFloor3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram != nil {
			if plantDiagram.VaseDiagram.TiledFloor3DShape == nil {
				shape := (&TiledFloor3DShape{
					Name: plantDiagram.Name + "-TiledFloor3DShape",
				}).Stage(stager.stage)
				plantDiagram.VaseDiagram.TiledFloor3DShape = shape
				modified = true
			} else if plantDiagram.VaseDiagram.TiledFloor3DShape.Name != plantDiagram.Name+"-TiledFloor3DShape" {
				plantDiagram.VaseDiagram.TiledFloor3DShape.Name = plantDiagram.Name + "-TiledFloor3DShape"
				modified = true
			}
		}
		if plantDiagram.StoolDiagram != nil {
			if plantDiagram.StoolDiagram.TiledFloor3DShape == nil {
				shape := (&TiledFloor3DShape{
					Name: plantDiagram.Name + "-TiledFloor3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.TiledFloor3DShape = shape
				modified = true
			} else if plantDiagram.StoolDiagram.TiledFloor3DShape.Name != plantDiagram.Name+"-TiledFloor3DShape" {
				plantDiagram.StoolDiagram.TiledFloor3DShape.Name = plantDiagram.Name + "-TiledFloor3DShape"
				modified = true
			}
		}
		if plantDiagram.ClockDiagram != nil {
			if plantDiagram.ClockDiagram.TiledFloor3DShape == nil {
				shape := (&TiledFloor3DShape{
					Name: plantDiagram.Name + "-TiledFloor3DShape",
				}).Stage(stager.stage)
				plantDiagram.ClockDiagram.TiledFloor3DShape = shape
				modified = true
			} else if plantDiagram.ClockDiagram.TiledFloor3DShape.Name != plantDiagram.Name+"-TiledFloor3DShape" {
				plantDiagram.ClockDiagram.TiledFloor3DShape.Name = plantDiagram.Name + "-TiledFloor3DShape"
				modified = true
			}
		}
	}

	return modified
}
