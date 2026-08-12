package models

func (stager *Stager) enforcePlantHasClockAbstract() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Clock {
			if plant.ClockAbstract == nil {
				ca := (&ClockAbstract{
					Name:                    plant.Name + "-ClockAbstract",
					RadialRepetitions:       1,
					RelativeTubeDiameter:    0.01,
					ClockTorusVerticalScale: 1.0,
					RelativeHeight:          1.0,
				}).Stage(stager.stage)
				plant.ClockAbstract = ca
				needCommit = true
			} else if plant.ClockAbstract.RadialRepetitions < 1 {
				plant.ClockAbstract.RadialRepetitions = 1
				needCommit = true
			}
		} else {
			if plant.ClockAbstract != nil {
				plant.ClockAbstract = nil
				needCommit = true
			}
		}
	}

	// Unstage unreferenced ClockAbstract
	for ca := range *GetGongstructInstancesSetFromPointerType[*ClockAbstract](stager.stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.ClockAbstract == ca {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			ca.Unstage(stager.stage)
			needCommit = true
		}
	}

	return needCommit
}

func (stager *Stager) enforceClockAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Clock && plant.ClockAbstract != nil {
			expectedName := plant.Name + "-ClockAbstract"
			if plant.ClockAbstract.Name != expectedName {
				plant.ClockAbstract.Name = expectedName
				needCommit = true
			}
		}
	}
	return needCommit
}
