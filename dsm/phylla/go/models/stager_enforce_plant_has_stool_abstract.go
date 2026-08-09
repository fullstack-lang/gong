package models

func (stager *Stager) enforcePlantHasStoolAbstract() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Stool {
			if plant.StoolAbstract == nil {
				sa := (&StoolAbstract{
					Name:              plant.Name + "-StoolAbstract",
					RadialRepetitions: 1,
				}).Stage(stager.stage)
				plant.StoolAbstract = sa
				needCommit = true
			} else if plant.StoolAbstract.RadialRepetitions < 1 {
				plant.StoolAbstract.RadialRepetitions = 1
				needCommit = true
			}
		} else {
			if plant.StoolAbstract != nil {
				plant.StoolAbstract = nil
				needCommit = true
			}
		}
	}

	// Unstage unreferenced StoolAbstract
	for sa := range *GetGongstructInstancesSetFromPointerType[*StoolAbstract](stager.stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.StoolAbstract == sa {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			sa.Unstage(stager.stage)
			needCommit = true
		}
	}

	return needCommit
}

func (stager *Stager) enforceStoolAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Stool && plant.StoolAbstract != nil {
			expectedName := plant.Name + "-StoolAbstract"
			if plant.StoolAbstract.Name != expectedName {
				plant.StoolAbstract.Name = expectedName
				needCommit = true
			}
		}
	}
	return needCommit
}
