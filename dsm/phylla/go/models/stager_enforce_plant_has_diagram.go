package models

import (
	"fmt"
	"time"
)

// enforcePlantHasDiagram ensures that each Plant has at least one PlantDiagram
func (stager *Stager) enforcePlantHasDiagram() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		if len(plant.PlantDiagrams) == 0 {
			plantDiagram := new(PlantDiagram).Stage(stage)
			plantDiagram.Name = plant.Name + " - Diagram"
			hasAnyChecked := false
			for d := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
				if d.IsChecked {
					hasAnyChecked = true
					break
				}
			}
			if !hasAnyChecked || plant.IsSelected {
				if plant.IsSelected {
					for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stager.stage) {
						plantDiagram_.IsChecked = false
					}
				}
				plantDiagram.IsChecked = true
			}
			plant.PlantDiagrams = append(plant.PlantDiagrams, plantDiagram)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default PlantDiagram for plant %s", plant.Name))
			}

			needCommit = true
		}
	}

	return
}
