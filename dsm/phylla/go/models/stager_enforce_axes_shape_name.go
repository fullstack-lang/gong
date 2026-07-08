package models

import (
	"fmt"
	"time"
)

// enforceAxesShapeName ensures that the name of the AxesShape matches its owning PlantDiagram
func (stager *Stager) enforceAxesShapeName() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.AxesShape != nil {
			expectedName := plantDiagram.Name + "-AxesShape"
			if plantDiagram.AxesShape.Name != expectedName {
				plantDiagram.AxesShape.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed AxesShape to %s", expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}
