package models

import (
	"fmt"
	"time"
)

// enforcePlantHasAxes ensures that each Plant has one and only one Axes that belong to it
func (stager *Stager) enforcePlantHasAxes() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each Plant has an Axes and its name matches plant.Name + "-Axes"
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		expectedName := plant.Name + "-Axes"

		if plant.Axes == nil {
			axes := new(Axes).Stage(stage)
			axes.Name = expectedName
			
			// Default values
			axes.LengthX = 1.0
			axes.LengthY = 1.0
			
			plant.Axes = axes

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default Axes for plant %s", plant.Name))
			}

			needCommit = true
		} else if plant.Axes.Name != expectedName {
			plant.Axes.Name = expectedName
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed Axes to %s for plant %s", expectedName, plant.Name))
			}
			needCommit = true
		}
	}

	// 2. Ensure each Axes belongs to exactly one Plant. If orphaned, remove it.
	for axes := range *GetGongstructInstancesSetFromPointerType[*Axes](stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
			if plant.Axes == axes {
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
