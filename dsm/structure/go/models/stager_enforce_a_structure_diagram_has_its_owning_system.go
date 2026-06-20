package models

import "time"

func (stager *Stager) enforceAStructureDiagramHasItsOwningSystem() (needCommit bool) {

	rm := GetSliceOfPointersReverseMap[System, DiagramStructure](GetAssociationName[System]().DiagramStructures[0].Name, stager.stage)

	for diagramStructure := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stager.stage) {

		owningSystems := rm[diagramStructure]
		if len(owningSystems) > 1 {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), "DiagramStructure "+diagramStructure.Name+" has more than one owning system")
			}
			needCommit = true
			diagramStructure.UnstageVoid(stager.stage)
			continue
		}

		if len(owningSystems) == 0 {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), "DiagramStructure "+diagramStructure.Name+" has no owning system")
			}
			diagramStructure.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		diagramStructure.owningSystem = owningSystems[0]
	}
	return
}
