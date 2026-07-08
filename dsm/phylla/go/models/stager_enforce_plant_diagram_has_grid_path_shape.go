package models

import (
	"fmt"
	"time"
)

// enforcePlantDiagramHasGridPathShape ensures that each PlantDiagram has one and only one GridPathShape that belong to it
func (stager *Stager) enforcePlantDiagramHasGridPathShape() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each PlantDiagram has a GridPathShape
	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.GridPathShape == nil {
			gridPathShape := new(GridPathShape).Stage(stage)

			plantDiagram.GridPathShape = gridPathShape

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default GridPathShape for plant diagram %s", plantDiagram.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each GridPathShape belongs to exactly one PlantDiagram. If orphaned, remove it.
	for gridPathShape := range *GetGongstructInstancesSetFromPointerType[*GridPathShape](stage) {
		hasOwner := false
		for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
			if plantDiagram.GridPathShape == gridPathShape {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			gridPathShape.Unstage(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removed orphaned GridPathShape %s", gridPathShape.Name))
			}
			needCommit = true
		}
	}

	return
}
