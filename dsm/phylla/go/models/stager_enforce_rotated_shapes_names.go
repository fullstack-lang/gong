package models

import (
	"fmt"
	"time"
)

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning PlantDiagram
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.RotatedReferenceRhombus != nil {
			expectedName := plantDiagram.Name + "-RotatedReferenceRhombus"
			if plantDiagram.RotatedReferenceRhombus.Name != expectedName {
				plantDiagram.RotatedReferenceRhombus.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed RotatedReferenceRhombus to %s", expectedName))
				}
				needCommit = true
			}
		}

		if plantDiagram.RotatedGrowthVectorShape != nil {
			expectedName := plantDiagram.Name + "-RotatedGrowthVectorShape"
			if plantDiagram.RotatedGrowthVectorShape.Name != expectedName {
				plantDiagram.RotatedGrowthVectorShape.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed RotatedGrowthVectorShape to %s", expectedName))
				}
				needCommit = true
			}
		}

		if plantDiagram.RotatedGridPathShape != nil {
			expectedName := plantDiagram.Name + "-RotatedGridPathShape"
			if plantDiagram.RotatedGridPathShape.Name != expectedName {
				plantDiagram.RotatedGridPathShape.Name = expectedName
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed RotatedGridPathShape to %s", expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}
