package models

import "time"

func (stager *Stager) enforceASystemDiagramHasItsOwningSystem() (needCommit bool) {

	rm := GetSliceOfPointersReverseMap[System, DiagramStructure](GetAssociationName[System]().DiagramStructures[0].Name, stager.stage)

	for diagramStructure := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stager.stage) {

		owningSystemes := rm[diagramStructure]
		if len(owningSystemes) > 1 {
			stager.probeForm.AddNotification(time.Now(), "DiagramStructure "+diagramStructure.Name+" has more than one owning system")
			needCommit = true
			diagramStructure.UnstageVoid(stager.stage)
			continue
		}

		if len(owningSystemes) == 0 {
			stager.probeForm.AddNotification(time.Now(), "DiagramStructure "+diagramStructure.Name+" has no owning system")
			diagramStructure.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		diagramStructure.owningSystem = owningSystemes[0]

		// check that that there is a systemShape that have the system as its system
		isSystemShapeOfOwningSystemFound := false
		for _, systemShape := range diagramStructure.System_Shapes {
			if systemShape.System == diagramStructure.owningSystem {
				isSystemShapeOfOwningSystemFound = true
				break
			}
		}

		if !isSystemShapeOfOwningSystemFound {
			stager.probeForm.AddNotification(time.Now(), "DiagramStructure "+diagramStructure.Name+" has no system shape for its owning system "+diagramStructure.owningSystem.Name)
			systemShape := (&SystemShape{
				Name:   "SystemShape",
				System: diagramStructure.owningSystem,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stager.stage)
			diagramStructure.System_Shapes = append(diagramStructure.System_Shapes, systemShape)
			needCommit = true
			continue
		}

	}
	return
}
