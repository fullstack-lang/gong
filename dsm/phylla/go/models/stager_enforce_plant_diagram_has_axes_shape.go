package models

import (
	"fmt"
	"time"
)

// enforcePlantDiagramHasAxes ensures that each PlantDiagram has one and only one Axes that belong to it
func (stager *Stager) enforcePlantDiagramHasAxes() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each PlantDiagram has an AxesShape
	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.AxesShape == nil {
			axes := new(AxesShape).Stage(stage)

			plantDiagram.AxesShape = axes

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default Axes for plant diagram %s", plantDiagram.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each Axes belongs to exactly one PlantDiagram. If orphaned, remove it.
	for axes := range *GetGongstructInstancesSetFromPointerType[*AxesShape](stage) {
		hasOwner := false
		for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
			if plantDiagram.AxesShape == axes {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			axes.Unstage(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removed orphaned Axes %s", axes.Name))
			}
			needCommit = true
		}
	}

	return
}
