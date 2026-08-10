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
			if plantDiagram.StoolDiagram.SeatTopCurveShape == nil {
				s := (&SeatTopCurveShape{
					Name: plantDiagram.Name + "-SeatTopCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.SeatTopCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.SeatTopCurveShape.Name != plantDiagram.Name+"-SeatTopCurveShape" {
				plantDiagram.StoolDiagram.SeatTopCurveShape.Name = plantDiagram.Name + "-SeatTopCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.RotatedSeatTopCurveShape == nil {
				s := (&PartiallyRotatedSeatTopCurveShape{
					Name: plantDiagram.Name + "-PartiallyRotatedSeatTopCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.RotatedSeatTopCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.RotatedSeatTopCurveShape.Name != plantDiagram.Name+"-PartiallyRotatedSeatTopCurveShape" {
				plantDiagram.StoolDiagram.RotatedSeatTopCurveShape.Name = plantDiagram.Name + "-PartiallyRotatedSeatTopCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.SeatBottomCurveShape == nil {
				s := (&SeatBottomCurveShape{
					Name: plantDiagram.Name + "-SeatBottomCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.SeatBottomCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.SeatBottomCurveShape.Name != plantDiagram.Name+"-SeatBottomCurveShape" {
				plantDiagram.StoolDiagram.SeatBottomCurveShape.Name = plantDiagram.Name + "-SeatBottomCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.RotatedSeatBottomCurveShape == nil {
				s := (&PartiallyRotatedSeatBottomCurveShape{
					Name: plantDiagram.Name + "-PartiallyRotatedSeatBottomCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.RotatedSeatBottomCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.RotatedSeatBottomCurveShape.Name != plantDiagram.Name+"-PartiallyRotatedSeatBottomCurveShape" {
				plantDiagram.StoolDiagram.RotatedSeatBottomCurveShape.Name = plantDiagram.Name + "-PartiallyRotatedSeatBottomCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.Torus3DShape == nil {
				s := (&Torus3DShape{
					Name: plantDiagram.Name + "-Torus3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.Torus3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.Torus3DShape.Name != plantDiagram.Name+"-Torus3DShape" {
				plantDiagram.StoolDiagram.Torus3DShape.Name = plantDiagram.Name + "-Torus3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.RotatedTorusShape == nil {
				ts := (&PartiallyRotatedTorusShape{
					Name: plantDiagram.Name + "-PartiallyRotatedTorusShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.RotatedTorusShape = ts
				modified = true
			} else if plantDiagram.StoolDiagram.RotatedTorusShape.Name != plantDiagram.Name+"-PartiallyRotatedTorusShape" {
				plantDiagram.StoolDiagram.RotatedTorusShape.Name = plantDiagram.Name + "-PartiallyRotatedTorusShape"
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
			if plantDiagram.StoolDiagram.RotatedSampledPoints3DShape == nil {
				s := (&RotatedSampledPoints3DShape{
					Name: plantDiagram.Name + "-RotatedSampledPoints3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.RotatedSampledPoints3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.RotatedSampledPoints3DShape.Name != plantDiagram.Name+"-RotatedSampledPoints3DShape" {
				plantDiagram.StoolDiagram.RotatedSampledPoints3DShape.Name = plantDiagram.Name + "-RotatedSampledPoints3DShape"
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
