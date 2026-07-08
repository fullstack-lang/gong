package models

import (
	"fmt"
	"time"
)

// enforceReferenceRhombusName ensures that the name of the ReferenceRhombus matches its owning PlantDiagram
func (stager *Stager) enforceReferenceRhombusName() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.ReferenceRhombus != nil {
			expectedName := plantDiagram.Name + "-ReferenceRhombus"
			if plantDiagram.ReferenceRhombus.Name != expectedName {
				plantDiagram.ReferenceRhombus.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed ReferenceRhombus to %s", expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}
