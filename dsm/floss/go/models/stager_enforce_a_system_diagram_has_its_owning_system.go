package models

import "time"

func (stager *Stager) enforceASystemDiagramHasItsOwningSystem() (needCommit bool) {

	rm := GetSliceOfPointersReverseMap[System, DiagramFloss](GetAssociationName[System]().DiagramFlosses[0].Name, stager.stage)

	for diagramFloss := range *GetGongstructInstancesSetFromPointerType[*DiagramFloss](stager.stage) {

		owningSystemes := rm[diagramFloss]
		if len(owningSystemes) > 1 {
			stager.probeForm.AddNotification(time.Now(), "DiagramFloss "+diagramFloss.Name+" has more than one owning system")
			needCommit = true
			diagramFloss.UnstageVoid(stager.stage)
			continue
		}

		if len(owningSystemes) == 0 {
			stager.probeForm.AddNotification(time.Now(), "DiagramFloss "+diagramFloss.Name+" has no owning system")
			diagramFloss.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		diagramFloss.owningSystem = owningSystemes[0]

		// check that that there is a systemShape that have the system as its system
		isSystemShapeOfOwningSystemFound := false
		for _, systemShape := range diagramFloss.System_Shapes {
			if systemShape.System == diagramFloss.owningSystem {
				isSystemShapeOfOwningSystemFound = true
				break
			}
		}

		if !isSystemShapeOfOwningSystemFound {
			stager.probeForm.AddNotification(time.Now(), "DiagramFloss "+diagramFloss.Name+" has no system shape for its owning system "+diagramFloss.owningSystem.Name)
			systemShape := (&SystemShape{
				Name:   "SystemShape",
				System: diagramFloss.owningSystem,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stager.stage)
			diagramFloss.System_Shapes = append(diagramFloss.System_Shapes, systemShape)
			needCommit = true
			continue
		}

	}
	return
}
