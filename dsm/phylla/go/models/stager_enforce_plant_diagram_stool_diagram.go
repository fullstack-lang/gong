package models

func (stager *Stager) enforcePlantDiagramStoolDiagram() bool {
	modified := false

	// Map diagrams to their owning plant
	diagramOwner := make(map[*PlantDiagram]*PlantAbstract)
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		for _, d := range plant.PlantDiagrams {
			diagramOwner[d] = plant
		}
	}

	for plantDiagram := range stager.stage.PlantDiagrams {
		owner := diagramOwner[plantDiagram]
		if owner != nil && owner.PlantType == Stool {
			if plantDiagram.StoolDiagram == nil {
				sd := (&StoolDiagram{
					Name: plantDiagram.Name + "-StoolDiagram",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram = sd
				modified = true
			} else if plantDiagram.StoolDiagram.Name != plantDiagram.Name+"-StoolDiagram" {
				plantDiagram.StoolDiagram.Name = plantDiagram.Name + "-StoolDiagram"
				modified = true
			}
			if plantDiagram.StoolDiagram.SampledPoints3DShape == nil {
				s := (&SampledPoints3DShape{
					Name: plantDiagram.Name + "-SampledPoints3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.SampledPoints3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.SampledPoints3DShape.Name != plantDiagram.Name+"-SampledPoints3DShape" {
				plantDiagram.StoolDiagram.SampledPoints3DShape.Name = plantDiagram.Name + "-SampledPoints3DShape"
				modified = true
			}
		} else {
			if plantDiagram.StoolDiagram != nil {
				plantDiagram.StoolDiagram = nil
				modified = true
			}
		}
	}

	// Clean up unreferenced StoolDiagrams
	for sd := range *GetGongstructInstancesSetFromPointerType[*StoolDiagram](stager.stage) {
		hasOwner := false
		for plantDiagram := range stager.stage.PlantDiagrams {
			if plantDiagram.StoolDiagram == sd {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			sd.Unstage(stager.stage)
			modified = true
		}
	}

	return modified
}
