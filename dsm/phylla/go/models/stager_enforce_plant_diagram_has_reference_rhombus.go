package models

import (
	"fmt"
	"time"
)

// enforcePlantDiagramHasReferenceRhombus ensures that each PlantDiagram has one and only one ReferenceRhombus that belong to it
func (stager *Stager) enforcePlantDiagramHasReferenceRhombus() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each PlantDiagram has a ReferenceRhombus
	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.ReferenceRhombus == nil {
			referenceRhombus := new(ReferenceRhombus).Stage(stage)

			plantDiagram.ReferenceRhombus = referenceRhombus

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default ReferenceRhombus for plant diagram %s", plantDiagram.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each ReferenceRhombus belongs to exactly one PlantDiagram. If orphaned, remove it.
	for referenceRhombus := range *GetGongstructInstancesSetFromPointerType[*ReferenceRhombus](stage) {
		hasOwner := false
		for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
			if plantDiagram.ReferenceRhombus == referenceRhombus {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			referenceRhombus.Unstage(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removed orphaned ReferenceRhombus %s", referenceRhombus.Name))
			}
			needCommit = true
		}
	}

	return
}
