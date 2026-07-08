package models

import (
	"fmt"
	"time"
)

// enforceGridPathShapeName ensures that the name of the GridPathShape matches its owning PlantDiagram
func (stager *Stager) enforceGridPathShapeName() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.GridPathShape != nil {
			expectedName := plantDiagram.Name + "-GridPathShape"
			if plantDiagram.GridPathShape.Name != expectedName {
				plantDiagram.GridPathShape.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed GridPathShape to %s", expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}
