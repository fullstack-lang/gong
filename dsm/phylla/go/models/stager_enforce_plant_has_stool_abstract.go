package models

import "fmt"

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
				stager.logAndNotify(fmt.Sprintf("Plant %s: created missing StoolAbstract", plant.Name))
			} else if plant.StoolAbstract.RadialRepetitions < 1 {
				plant.StoolAbstract.RadialRepetitions = 1
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RadialRepetitions set to 1", plant.Name))
			}
		} else {
			if plant.StoolAbstract != nil {
				plant.StoolAbstract = nil
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s: removed StoolAbstract because PlantType is %s", plant.Name, plant.PlantType))
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
			stager.logAndNotify(fmt.Sprintf("Removed orphaned StoolAbstract %s", sa.Name))
		}
	}

	return needCommit
}

func (stager *Stager) enforceStoolAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Stool && plant.StoolAbstract != nil {
			expectedName := plant.Name + "-StoolAbstract"
			if plant.StoolAbstract.Name != expectedName {
				oldName := plant.StoolAbstract.Name
				plant.StoolAbstract.Name = expectedName
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Renamed StoolAbstract from '%s' to '%s'", oldName, expectedName))
			}
		}
	}
	return needCommit
}
