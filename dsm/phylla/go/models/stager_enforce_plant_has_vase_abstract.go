package models

func (stager *Stager) enforcePlantHasVaseAbstract() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == PLANT_TYPE_VASE {
			if plant.VaseAbstract == nil {
				va := (&VaseAbstract{
					Name: plant.Name + "-VaseAbstract",
				}).Stage(stager.stage)
				plant.VaseAbstract = va
				needCommit = true
			}
		} else {
			if plant.VaseAbstract != nil {
				plant.VaseAbstract = nil
				needCommit = true
			}
		}
	}

	// Unstage unreferenced VaseAbstract
	for va := range *GetGongstructInstancesSetFromPointerType[*VaseAbstract](stager.stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.VaseAbstract == va {
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

func (stager *Stager) enforceVaseAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.VaseAbstract != nil {
			expectedName := plant.Name + "-VaseAbstract"
			if plant.VaseAbstract.Name != expectedName {
				plant.VaseAbstract.Name = expectedName
				needCommit = true
			}
		}
	}
	return needCommit
}

