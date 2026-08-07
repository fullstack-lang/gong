package models

func (stager *Stager) enforcePlantDiagramVolumeKey3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VaseDiagram != nil && plantDiagram.VaseDiagram.VolumeKey3DShape == nil {
			shape := (&VolumeKey3DShape{
				Name: plantDiagram.Name + "-VolumeKey3DShape",
			}).Stage(stager.stage)
			plantDiagram.VaseDiagram.VolumeKey3DShape = shape
			modified = true
		}
	}

	return modified
}
