package models

import "fmt"

func (stager *Stager) enforcePlantHasClockAbstract() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
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
				stager.logAndNotify(fmt.Sprintf("Plant %s: created missing ClockAbstract", plant.Name))
			} else if plant.ClockAbstract.RadialRepetitions < 1 {
				plant.ClockAbstract.RadialRepetitions = 1
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s Clock: default RadialRepetitions set to 1", plant.Name))
			}
		} else {
			if plant.ClockAbstract != nil {
				plant.ClockAbstract = nil
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s: removed ClockAbstract because PlantType is %s", plant.Name, plant.PlantType))
			}
		}
	}

	// Unstage unreferenced ClockAbstract
	for ca := range *stager.stage.GetInstancesSet[*ClockAbstract]() {
		hasOwner := false
		for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
			if plant.ClockAbstract == ca {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			ca.Unstage(stager.stage)
			needCommit = true
			stager.logAndNotify(fmt.Sprintf("Removed orphaned ClockAbstract %s", ca.Name))
		}
	}

	return needCommit
}

func (stager *Stager) enforceClockAbstractName() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
		if plant.PlantType == Clock && plant.ClockAbstract != nil {
			expectedName := plant.Name + "-ClockAbstract"
			if plant.ClockAbstract.Name != expectedName {
				oldName := plant.ClockAbstract.Name
				plant.ClockAbstract.Name = expectedName
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Renamed ClockAbstract from '%s' to '%s'", oldName, expectedName))
			}
		}
	}
	return needCommit
}
