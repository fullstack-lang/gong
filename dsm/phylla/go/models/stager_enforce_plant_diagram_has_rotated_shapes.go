package models

import (
	"fmt"
	"time"
)

// enforcePlantDiagramHasRotatedShapes ensures that each PlantDiagram has its Rotated shapes
func (stager *Stager) enforcePlantDiagramHasRotatedShapes() (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if plantDiagram.RotatedReferenceRhombus == nil {
			shape := new(ReferenceRhombus).Stage(stage)
			plantDiagram.RotatedReferenceRhombus = shape
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default RotatedReferenceRhombus for %s", plantDiagram.Name))
			}
			needCommit = true
		}
		if plantDiagram.RotatedGrowthVectorShape == nil {
			shape := new(GrowthVectorShape).Stage(stage)
			plantDiagram.RotatedGrowthVectorShape = shape
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default RotatedGrowthVectorShape for %s", plantDiagram.Name))
			}
			needCommit = true
		}
		if plantDiagram.RotatedGridPathShape == nil {
			shape := new(GridPathShape).Stage(stage)
			plantDiagram.RotatedGridPathShape = shape
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default RotatedGridPathShape for %s", plantDiagram.Name))
			}
			needCommit = true
		}
	}

	return
}
