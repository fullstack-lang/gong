package models

func (stager *Stager) enforcePlantDiagramClockDiagram() bool {
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
		if owner != nil && owner.PlantType == Clock {
			if plantDiagram.ClockDiagram == nil {
				cd := (&ClockDiagram{
					Name: plantDiagram.Name + "-ClockDiagram",
				}).Stage(stager.stage)
				plantDiagram.ClockDiagram = cd
				modified = true
			} else if plantDiagram.ClockDiagram.Name != plantDiagram.Name+"-ClockDiagram" {
				plantDiagram.ClockDiagram.Name = plantDiagram.Name + "-ClockDiagram"
				modified = true
			}
			if plantDiagram.ClockDiagram.ClockTopCurveShape == nil {
				s := (&ClockTopCurveShape{
					Name: plantDiagram.Name + "-ClockTopCurveShape",
				}).Stage(stager.stage)
				plantDiagram.ClockDiagram.ClockTopCurveShape = s
				modified = true
			} else if plantDiagram.ClockDiagram.ClockTopCurveShape.Name != plantDiagram.Name+"-ClockTopCurveShape" {
				plantDiagram.ClockDiagram.ClockTopCurveShape.Name = plantDiagram.Name + "-ClockTopCurveShape"
				modified = true
			}
			if plantDiagram.ClockDiagram.Torus3DShape == nil {
				s := (&Torus3DShape{
					Name: plantDiagram.Name + "-Torus3DShape",
				}).Stage(stager.stage)
				plantDiagram.ClockDiagram.Torus3DShape = s
				modified = true
			} else if plantDiagram.ClockDiagram.Torus3DShape.Name != plantDiagram.Name+"-Torus3DShape" {
				plantDiagram.ClockDiagram.Torus3DShape.Name = plantDiagram.Name + "-Torus3DShape"
				modified = true
			}
			if plantDiagram.ClockDiagram.SampledPoints3DShape == nil {
				s := (&SampledPoints3DShape{
					Name: plantDiagram.Name + "-SampledPoints3DShape",
				}).Stage(stager.stage)
				plantDiagram.ClockDiagram.SampledPoints3DShape = s
				modified = true
			} else if plantDiagram.ClockDiagram.SampledPoints3DShape.Name != plantDiagram.Name+"-SampledPoints3DShape" {
				plantDiagram.ClockDiagram.SampledPoints3DShape.Name = plantDiagram.Name + "-SampledPoints3DShape"
				modified = true
			}
		} else {
			if plantDiagram.ClockDiagram != nil {
				plantDiagram.ClockDiagram = nil
				modified = true
			}
		}
	}

	// Clean up unreferenced ClockDiagrams
	for cd := range *GetGongstructInstancesSetFromPointerType[*ClockDiagram](stager.stage) {
		hasOwner := false
		for plantDiagram := range stager.stage.PlantDiagrams {
			if plantDiagram.ClockDiagram == cd {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			cd.Unstage(stager.stage)
			modified = true
		}
	}

	return modified
}
