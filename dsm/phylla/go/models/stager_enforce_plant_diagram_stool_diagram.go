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
			if plantDiagram.StoolDiagram.EyeSampledPoints3DShape == nil {
				s := (&EyeSampledPoints3DShape{
					Name: plantDiagram.Name + "-EyeSampledPoints3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.EyeSampledPoints3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.EyeSampledPoints3DShape.Name != plantDiagram.Name+"-EyeSampledPoints3DShape" {
				plantDiagram.StoolDiagram.EyeSampledPoints3DShape.Name = plantDiagram.Name + "-EyeSampledPoints3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.EyeCornersSampledPoints3DShape == nil {
				s := (&EyeCornersSampledPoints3DShape{
					Name: plantDiagram.Name + "-EyeCornersSampledPoints3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.EyeCornersSampledPoints3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.EyeCornersSampledPoints3DShape.Name != plantDiagram.Name+"-EyeCornersSampledPoints3DShape" {
				plantDiagram.StoolDiagram.EyeCornersSampledPoints3DShape.Name = plantDiagram.Name + "-EyeCornersSampledPoints3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.Eye3DShape == nil {
				s := (&Eye3DShape{
					Name: plantDiagram.Name + "-Eye3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.Eye3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.Eye3DShape.Name != plantDiagram.Name+"-Eye3DShape" {
				plantDiagram.StoolDiagram.Eye3DShape.Name = plantDiagram.Name + "-Eye3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.EyeSeatBottomCurveShape == nil {
				s := (&EyeSeatBottomCurveShape{
					Name: plantDiagram.Name + "-EyeSeatBottomCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.EyeSeatBottomCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.EyeSeatBottomCurveShape.Name != plantDiagram.Name+"-EyeSeatBottomCurveShape" {
				plantDiagram.StoolDiagram.EyeSeatBottomCurveShape.Name = plantDiagram.Name + "-EyeSeatBottomCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.EyeStoolBottomCurveShape == nil {
				s := (&EyeStoolBottomCurveShape{
					Name: plantDiagram.Name + "-EyeStoolBottomCurveShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.EyeStoolBottomCurveShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.EyeStoolBottomCurveShape.Name != plantDiagram.Name+"-EyeStoolBottomCurveShape" {
				plantDiagram.StoolDiagram.EyeStoolBottomCurveShape.Name = plantDiagram.Name + "-EyeStoolBottomCurveShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.Seat3DShape == nil {
				s := (&Seat3DShape{
					Name: plantDiagram.Name + "-Seat3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.Seat3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.Seat3DShape.Name != plantDiagram.Name+"-Seat3DShape" {
				plantDiagram.StoolDiagram.Seat3DShape.Name = plantDiagram.Name + "-Seat3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.EyeVolume3DShape == nil {
				s := (&EyeVolume3DShape{
					Name: plantDiagram.Name + "-EyeVolume3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.EyeVolume3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.EyeVolume3DShape.Name != plantDiagram.Name+"-EyeVolume3DShape" {
				plantDiagram.StoolDiagram.EyeVolume3DShape.Name = plantDiagram.Name + "-EyeVolume3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.SeatAndLegs3DShape == nil {
				s := (&SeatAndLegs3DShape{
					Name: plantDiagram.Name + "-SeatAndLegs3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.SeatAndLegs3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.SeatAndLegs3DShape.Name != plantDiagram.Name+"-SeatAndLegs3DShape" {
				plantDiagram.StoolDiagram.SeatAndLegs3DShape.Name = plantDiagram.Name + "-SeatAndLegs3DShape"
				modified = true
			}
			if plantDiagram.StoolDiagram.RotatedSeatAndLegs3DShape == nil {
				s := (&RotatedSeatAndLegs3DShape{
					Name: plantDiagram.Name + "-RotatedSeatAndLegs3DShape",
				}).Stage(stager.stage)
				plantDiagram.StoolDiagram.RotatedSeatAndLegs3DShape = s
				modified = true
			} else if plantDiagram.StoolDiagram.RotatedSeatAndLegs3DShape.Name != plantDiagram.Name+"-RotatedSeatAndLegs3DShape" {
				plantDiagram.StoolDiagram.RotatedSeatAndLegs3DShape.Name = plantDiagram.Name + "-RotatedSeatAndLegs3DShape"
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
