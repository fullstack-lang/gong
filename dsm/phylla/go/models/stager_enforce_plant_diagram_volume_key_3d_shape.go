package models

func (stager *Stager) enforcePlantDiagramVolumeKey3DShape() bool {
	modified := false

	for plantDiagram := range stager.stage.PlantDiagrams {
		if plantDiagram.VolumeKey3DShape == nil {
			shape := (&VolumeKey3DShape{
				Name: plantDiagram.Name + "-VolumeKey3DShape",
			}).Stage(stager.stage)
			plantDiagram.VolumeKey3DShape = shape
			modified = true
		}
	}

	return modified
}
