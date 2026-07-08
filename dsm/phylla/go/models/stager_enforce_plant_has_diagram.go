package models

import (
	"fmt"
	"time"
)

// enforcePlantHasDiagram ensures that each Plant has at least one PlantDiagram
func (stager *Stager) enforcePlantHasDiagram() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		if len(plant.PlantDiagrams) == 0 {
			plantDiagram := new(PlantDiagram).Stage(stage)
			plantDiagram.Name = plant.Name + " - Diagram"
			plant.PlantDiagrams = append(plant.PlantDiagrams, plantDiagram)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default PlantDiagram for plant %s", plant.Name))
			}

			needCommit = true
		}
	}

	return
}
