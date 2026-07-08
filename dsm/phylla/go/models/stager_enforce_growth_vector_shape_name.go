package models

import (
	"fmt"
	"time"
)

// enforceGrowthVectorShapeName ensures that the name of the GrowthVectorShape matches its owning PlantDiagram
func (stager *Stager) enforceGrowthVectorShapeName() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.GrowthVectorShape != nil {
			expectedName := plantDiagram.Name + "-GrowthVectorShape"
			if plantDiagram.GrowthVectorShape.Name != expectedName {
				plantDiagram.GrowthVectorShape.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed GrowthVectorShape to %s", expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}
