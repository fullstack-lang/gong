package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
)

func (stager *Stager) enforcePlantHasStoolAbstract() (needCommit bool) {
	if stager.stageSet == nil || stager.stageSet.StoolStage == nil {
		return false
	}
	stoolStage := stager.stageSet.StoolStage

	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
		if plant.PlantType == Stool {
			if plant.StoolAbstract == nil {
				sa := (&stool.StoolAbstract{
					Name:              plant.Name + "-StoolAbstract",
					RadialRepetitions: 1,
				}).Stage(stoolStage)
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
	for sa := range *stoolStage.GetInstancesSet[*stool.StoolAbstract]() {
		hasOwner := false
		for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
			if plant.StoolAbstract == sa {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			sa.Unstage(stoolStage)
			needCommit = true
			stager.logAndNotify(fmt.Sprintf("Removed orphaned StoolAbstract %s", sa.Name))
		}
	}

	return needCommit
}

func (stager *Stager) enforceStoolAbstractName() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
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
