package models

import "fmt"

func (stager *Stager) enforcePlantHasTubeVaseAbstract() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
		if plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze {
			if plant.TubeVaseAbstract == nil {
				va := (&TubeVaseAbstract{
					Name:                plant.Name + "-TubeVaseAbstract",
					RibbonVerticalScale: 1.0,
				}).Stage(stager.stage)
				plant.TubeVaseAbstract = va
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s: created missing TubeVaseAbstract", plant.Name))
			}
		} else {
			if plant.TubeVaseAbstract != nil {
				plant.TubeVaseAbstract = nil
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s: removed TubeVaseAbstract because PlantType is %s", plant.Name, plant.PlantType))
			}
		}
	}

	// Unstage unreferenced TubeVaseAbstract
	for va := range *stager.stage.GetInstancesSet[*TubeVaseAbstract]() {
		hasOwner := false
		for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
			if plant.TubeVaseAbstract == va {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			va.Unstage(stager.stage)
			needCommit = true
			stager.logAndNotify(fmt.Sprintf("Removed orphaned TubeVaseAbstract %s", va.Name))
		}
	}

	return needCommit
}

func (stager *Stager) enforceTubeVaseAbstractName() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
		if (plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze) && plant.TubeVaseAbstract != nil {
			expectedName := plant.Name + "-TubeVaseAbstract"
			if plant.TubeVaseAbstract.Name != expectedName {
				oldName := plant.TubeVaseAbstract.Name
				plant.TubeVaseAbstract.Name = expectedName
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Renamed TubeVaseAbstract from '%s' to '%s'", oldName, expectedName))
			}
		}
	}
	return needCommit
}
