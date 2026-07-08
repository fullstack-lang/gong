package models

import (
	"fmt"
	"time"
)

// enforcePlantDiagramHasGrowthVectorShape ensures that each PlantDiagram has one and only one GrowthVectorShape that belong to it
func (stager *Stager) enforcePlantDiagramHasGrowthVectorShape() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each PlantDiagram has a GrowthVectorShape
	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.GrowthVectorShape == nil {
			growthVectorShape := new(GrowthVectorShape).Stage(stage)

			plantDiagram.GrowthVectorShape = growthVectorShape

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default GrowthVectorShape for plant diagram %s", plantDiagram.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each GrowthVectorShape belongs to exactly one PlantDiagram. If orphaned, remove it.
	for growthVectorShape := range *GetGongstructInstancesSetFromPointerType[*GrowthVectorShape](stage) {
		hasOwner := false
		for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
			if plantDiagram.GrowthVectorShape == growthVectorShape || plantDiagram.RotatedGrowthVectorShape == growthVectorShape {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			growthVectorShape.Unstage(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removed orphaned GrowthVectorShape %s", growthVectorShape.Name))
			}
			needCommit = true
		}
	}

	return
}
