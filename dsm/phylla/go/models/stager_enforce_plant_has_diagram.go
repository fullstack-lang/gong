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

		if (plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze) && len(plant.Vase2DDiagrams) == 0 {
			vase2DDiagram := new(Vase2DDiagram).Stage(stage)
			vase2DDiagram.Name = plant.Name + "-Vase2DDiagram"
			vase2DDiagram.Zoom = 0.45
			vase2DDiagram.IsExpanded = true
			if plant.PlantType == VaseTrapeze {
				vase2DDiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 = true
				vase2DDiagram.IsHiddenPxShape = true
				vase2DDiagram.IsHiddenChosenP1P2PairShape = true
				vase2DDiagram.IsHiddenKeyHoleShape = true
			}
			plant.Vase2DDiagrams = append(plant.Vase2DDiagrams, vase2DDiagram)
			stager.logAndNotify(fmt.Sprintf("Added default Vase2DDiagram for plant %s", plant.Name))
			needCommit = true
		}

		if (plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze) && len(plant.TubeVase3DDiagrams) == 0 {
			vase3DDiagram := new(TubeVase3DDiagram).Stage(stage)
			vase3DDiagram.Name = plant.Name + "-TubeVase3DDiagram"
			vase3DDiagram.IsExpanded = true
			if plant.PlantType == VaseTrapeze {
				vase3DDiagram.IsHiddenTorusStackShape = false
				vase3DDiagram.IsHiddenVerticalTorusStackShape = true
				vase3DDiagram.IsHiddenPartiallyRotatedTorusShape = true
				vase3DDiagram.IsHiddenStackOfPartiallyRotatedTorusShape = true
				vase3DDiagram.IsHiddenKeyHole3DShape = true
				vase3DDiagram.IsHiddenKey3DShape = true
				vase3DDiagram.IsHiddenVolumeKey3DShape = true
				vase3DDiagram.IsHiddenTorusEdge3DShape = true
				vase3DDiagram.IsHiddenSampledPoints3DShape = true
				vase3DDiagram.IsHiddenOriginalPoints3DShape = true
				vase3DDiagram.IsHiddenAngle0Shape = true
				vase3DDiagram.IsHiddenTiledFloor3DShape = false
				vase3DDiagram.IsHiddenTopCurvePlane1Shape = false
				vase3DDiagram.IsHiddenBottomCurvePlane1Shape = false
				vase3DDiagram.IsHiddenTopCurvePlane2Shape = false
				vase3DDiagram.IsHiddenBottomCurvePlane2Shape = false
				vase3DDiagram.IsHiddenVaseTrapezeRingShape = false
				vase3DDiagram.IsHiddenStackOfVaseTrapezeRingsShape = false
				vase3DDiagram.IsHiddenStackOfRotatedVaseTrapezeRingsShape = false
			} else {
				vase3DDiagram.IsHiddenTopCurvePlane1Shape = true
				vase3DDiagram.IsHiddenBottomCurvePlane1Shape = true
				vase3DDiagram.IsHiddenTopCurvePlane2Shape = true
				vase3DDiagram.IsHiddenBottomCurvePlane2Shape = true
				vase3DDiagram.IsHiddenVaseTrapezeRingShape = true
				vase3DDiagram.IsHiddenStackOfVaseTrapezeRingsShape = true
				vase3DDiagram.IsHiddenStackOfRotatedVaseTrapezeRingsShape = true
			}
			plant.TubeVase3DDiagrams = append(plant.TubeVase3DDiagrams, vase3DDiagram)
			stager.logAndNotify(fmt.Sprintf("Added default TubeVase3DDiagram for plant %s", plant.Name))
			needCommit = true
		}
	}

	return
}
