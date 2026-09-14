package models

func (stager *Stager) enforcePlantHasTubeVaseAbstract() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == TubeVase {
			if plant.TubeVaseAbstract == nil {
				va := (&TubeVaseAbstract{
					Name: plant.Name + "-TubeVaseAbstract",
				}).Stage(stager.stage)
				plant.TubeVaseAbstract = va
				needCommit = true
			}
		} else {
			if plant.TubeVaseAbstract != nil {
				plant.TubeVaseAbstract = nil
				needCommit = true
			}
		}
	}

	// Unstage unreferenced TubeVaseAbstract
	for va := range *GetGongstructInstancesSetFromPointerType[*TubeVaseAbstract](stager.stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.TubeVaseAbstract == va {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			va.Unstage(stager.stage)
			needCommit = true
		}
	}

	return needCommit
}

func (stager *Stager) enforceTubeVaseAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == TubeVase {
			expectedName := plant.Name + "-TubeVaseAbstract"
			if plant.TubeVaseAbstract.Name != expectedName {
				plant.TubeVaseAbstract.Name = expectedName
				needCommit = true
			}
		}
	}
	return needCommit
}
