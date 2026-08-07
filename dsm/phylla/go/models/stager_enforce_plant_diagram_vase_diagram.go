package models

func (stager *Stager) enforcePlantDiagramVaseDiagram() bool {
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
		if owner != nil && owner.PlantType == PLANT_TYPE_VASE {
			if plantDiagram.VaseDiagram == nil {
				vd := (&VaseDiagram{
					Name: plantDiagram.Name + "-VaseDiagram",
				}).Stage(stager.stage)
				plantDiagram.VaseDiagram = vd
				modified = true
			} else if plantDiagram.VaseDiagram.Name != plantDiagram.Name+"-VaseDiagram" {
				plantDiagram.VaseDiagram.Name = plantDiagram.Name + "-VaseDiagram"
				modified = true
			}
		} else {
			if plantDiagram.VaseDiagram != nil {
				plantDiagram.VaseDiagram = nil
				modified = true
			}
		}
	}

	// Clean up unreferenced VaseDiagrams
	for vd := range *GetGongstructInstancesSetFromPointerType[*VaseDiagram](stager.stage) {
		hasOwner := false
		for plantDiagram := range stager.stage.PlantDiagrams {
			if plantDiagram.VaseDiagram == vd {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			vd.Unstage(stager.stage)
			modified = true
		}
	}

	return modified
}

