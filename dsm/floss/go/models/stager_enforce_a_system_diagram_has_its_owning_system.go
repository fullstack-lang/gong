package models

import "time"

func (stager *Stager) enforceASystemDiagramHasItsOwningSystem() (needCommit bool) {

	rm := GetSliceOfPointersReverseMap[System, DiagramFlossEquation](GetAssociationName[System]().DiagramFlossEquations[0].Name, stager.stage)

	for diagram := range *GetGongstructInstancesSetFromPointerType[*DiagramFlossEquation](stager.stage) {

		owningSystems := rm[diagram]
		if len(owningSystems) > 1 {
			stager.probeForm.AddNotification(time.Now(), "DiagramFlossEquation "+diagram.Name+" has more than one owning system")
			needCommit = true
			diagram.UnstageVoid(stager.stage)
			continue
		}

		if len(owningSystems) == 1 {
			diagram.owningSystem = owningSystems[0]
		}
	}
	return
}
