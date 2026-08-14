package models

import (
	"fmt"
	"time"
)

// enforcePlantHasDiagram ensures that each Plant has at least one Plant2DDiagram
func (stager *Stager) enforcePlantHasDiagram() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		if len(plant.Plant2DDiagrams) == 0 {
			plantDiagram := new(Plant2DDiagram).Stage(stage)
			plantDiagram.Name = plant.Name + " - Diagram"
			hasAnyChecked := false
			for d := range *GetGongstructInstancesSetFromPointerType[*Plant2DDiagram](stage) {
				if d.IsChecked {
					hasAnyChecked = true
					break
				}
			}
			if !hasAnyChecked || plant.IsSelected {
				if plant.IsSelected {
					for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*Plant2DDiagram](stager.stage) {
						plantDiagram_.IsChecked = false
					}
				}
				plantDiagram.IsChecked = true
			}
			plant.Plant2DDiagrams = append(plant.Plant2DDiagrams, plantDiagram)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default Plant2DDiagram for plant %s", plant.Name))
			}

			needCommit = true
		}
	}

	return
}
