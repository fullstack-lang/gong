package models

import (
	"fmt"
)

// enforcePlantHasDiagram ensures that each Plant has at least one Plant2DDiagram
func (stager *Stager) enforcePlantHasDiagram() (needCommit bool) {
	stage := stager.stage

	for plant := range *stage.GetInstancesSet[*PlantAbstract]() {
		if len(plant.Plant2DDiagrams) == 0 {
			plantDiagram := new(Plant2DDiagram).Stage(stage)
			plantDiagram.Name = plant.Name + " - Diagram"
			plantDiagram.OriginX = 280.0
			plantDiagram.OriginY = 950.0
			plantDiagram.Zoom = 1.0
			hasAnyChecked := false
			for d := range *stage.GetInstancesSet[*Plant2DDiagram]() {
				if d.IsChecked {
					hasAnyChecked = true
					break
				}
			}
			if !hasAnyChecked || plant.IsSelected {
				if plant.IsSelected {
					for plantDiagram_ := range *stager.stage.GetInstancesSet[*Plant2DDiagram]() {
						plantDiagram_.IsChecked = false
					}
				}
				plantDiagram.IsChecked = true
			}
			plant.Plant2DDiagrams = append(plant.Plant2DDiagrams, plantDiagram)

			stager.logAndNotify(fmt.Sprintf("Added default Plant2DDiagram for plant %s", plant.Name))

			needCommit = true
		}

		if len(plant.Plant3DDiagrams) == 0 {
			plant3DDiagram := new(Plant3DDiagram).Stage(stage)
			plant3DDiagram.Name = plant.Name + " - 3D Diagram"
			plant.Plant3DDiagrams = append(plant.Plant3DDiagrams, plant3DDiagram)

			stager.logAndNotify(fmt.Sprintf("Added default Plant3DDiagram for plant %s", plant.Name))

			needCommit = true
		}
	}

	return
}
