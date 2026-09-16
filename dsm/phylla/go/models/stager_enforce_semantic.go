package models

import (
	"fmt"
	"log"
	"time"
)

func (stager *Stager) logAndNotify(message string) {
	log.Println(message)
	if stager.probeForm != nil {
		stager.probeForm.AddNotification(time.Now(), message)
	}
}

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage
	needCommit = stager.enforceThereIsARootLibrary() || needCommit

	// computes fields that are not persisted
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()

	pass := 0
	for {
		if pass > 10 {
			stager.logAndNotify("Semantic enforcement reached maximum number of passes (10). Breaking loop.")
			break
		}
		if stager.enforceSemanticOnePass(false, stage) {
			needCommit = true
			stager.logAndNotify(fmt.Sprint("Stage was modified to enforce semantic, pass ", pass))
			pass++
		} else {
			break
		}
	}

	// computes fields that are not persisted
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()

	if needCommit {
		if stager.probeForm != nil {
			stager.probeForm.CommitNotificationTable()
		}
		stage.CommitWithSuspendedCallbacks()
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	methods := []struct {
		name   string
		fn     func() bool
		notify bool
	}{
		// abstract semantic checks
		{"Clean the stage", func() bool { return stage.Clean() }, true},
		{"Enforce orphans abstract element", stager.enforceOrphansAbstractElement, true},
		{"Enforce orphan shape remove", stager.enforceOrphanShapeRemove, true},
		{"Enforce default values", stager.enforceDefaultValues, true},
		{"Enforce N <= M", stager.enforcePlantNM, true},
		{"Enforce duplicate remove", stager.enforceDuplicateRemove, true},
		{"Enforce single selected plant", stager.enforceSingleSelectedPlant, true},
		{"Enforce plant has diagram", stager.enforcePlantHasDiagram, true},

		// concrete semantic checks
		{"Enforce plant has axes", stager.enforcePlantHasAxes, true},
		{"Enforce axes shape name", stager.enforceAxesShapeName, true},
		{"Enforce plant has rhombus stuff", stager.enforcePlantHasRhombusStuff, true},
		{"Enforce rhombus stuff name", stager.enforceRhombusStuffName, true},
		{"Enforce plant has tube vase abstract", stager.enforcePlantHasTubeVaseAbstract, true},
		{"Enforce tube vase abstract name", stager.enforceTubeVaseAbstractName, true},
		{"Enforce plant has stool abstract", stager.enforcePlantHasStoolAbstract, true},
		{"Enforce stool abstract name", stager.enforceStoolAbstractName, true},
		{"Enforce plant has clock abstract", stager.enforcePlantHasClockAbstract, true},
		{"Enforce clock abstract name", stager.enforceClockAbstractName, true},
		{"Enforce plant has music abstract", stager.enforcePlantHasMusicAbstract, true},
		{"Enforce music abstract name", stager.enforceMusicAbstractName, true},
		{"Enforce plant has reference rhombus", stager.enforcePlantHasReferenceRhombus, true},
		{"Enforce reference rhombus name", stager.enforceReferenceRhombusName, true},
		{"Enforce plant has grid path shape", stager.enforcePlantHasGridPathShape, true},
		{"Enforce grid path shape name", stager.enforceGridPathShapeName, true},
		{"Enforce plant has initial rhombus grid shape", stager.enforcePlantHasInitialRhombusGridShape, true},
		{"Enforce initial rhombus grid shape name", stager.enforceInitialRhombusGridShapeName, true},
		{"Enforce plant has explanation text shape", stager.enforcePlantHasExplanationTextShape, true},
		{"Enforce explanation text shape name", stager.enforceExplanationTextShapeName, true},
		{"Enforce plant has rotated shapes", stager.enforcePlantHasRotatedShapes, true},
		{"Enforce rotated shapes names", stager.enforceRotatedShapesNames, true},
		{"Enforce tube vase has shapes", stager.enforceTubeVaseHasShapes, true},
		{"Enforce tube vase shape names", stager.enforceTubeVaseShapeNames, true},
		{"Enforce plant has growth vector shape", stager.enforcePlantHasPlantCircumferenceShape, true},
		{"Enforce growth vector shape name", stager.enforcePlantCircumferenceShapeName, true},
		{"Enforce diagram shapes", stager.enforceDiagramShapes, true},

		// continuous shape geometry recalculations (normal runtime behavior, not notified)
		{"Enforce compute growth vector shape", stager.enforceComputePlantCircumferenceShape, false},
		{"Enforce rhombus grid shape has rhombuses", stager.enforcePlantRhombusGridShapeHasRhombuses, false},
		{"Enforce plant rotation ratio heights", stager.enforcePlantRotationRatioHeights, false},
	}

	for _, method := range methods {
		modified := method.fn()
		if modified {
			if method.notify {
				stager.logAndNotify(fmt.Sprintf("Semantic check '%s' generated a stage modification", method.name))
			}
			needCommit = true
		}
	}

	return needCommit
}


func (stager *Stager) enforceSingleSelectedPlant() bool {
	modified := false

	hasChecked := func(p *PlantAbstract) bool {
		if p == nil {
			return false
		}
		for _, d := range p.Plant2DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Plant3DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Vase2DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.TubeVase3DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Stool2DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Stool3DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Clock2DDiagrams { if d.IsChecked { return true } }
		for _, d := range p.Clock3DDiagrams { if d.IsChecked { return true } }
		if p.MusicAbstract != nil && p.MusicAbstract.IsChecked { return true }
		return false
	}

	// Try to find if any diagram across all plants is checked, prioritizing selectedPlant or IsSelected
	var plantWithCheckedDiagram *PlantAbstract
	if hasChecked(stager.selectedPlant) {
		plantWithCheckedDiagram = stager.selectedPlant
	} else {
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.IsSelected && hasChecked(plant) {
				plantWithCheckedDiagram = plant
				break
			}
		}
		if plantWithCheckedDiagram == nil {
			for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
				if hasChecked(plant) {
					plantWithCheckedDiagram = plant
					break
				}
			}
		}
	}

	if plantWithCheckedDiagram != nil {
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			shouldBeSelected := (plant == plantWithCheckedDiagram)
			if plant.IsSelected != shouldBeSelected {
				plant.IsSelected = shouldBeSelected
				modified = true
				if shouldBeSelected {
					stager.logAndNotify(fmt.Sprintf("Selected plant set to %s", plant.Name))
				}
			}
			if !shouldBeSelected {
				for _, d := range plant.Plant2DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Plant3DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Vase2DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.TubeVase3DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Stool2DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Stool3DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Clock2DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				for _, d := range plant.Clock3DDiagrams { if d.IsChecked { d.IsChecked = false; modified = true } }
				if plant.MusicAbstract != nil && plant.MusicAbstract.IsChecked { plant.MusicAbstract.IsChecked = false; modified = true }
			}
		}
		if stager.selectedPlant != plantWithCheckedDiagram {
			stager.selectedPlant = plantWithCheckedDiagram
		}
	} else {
		var selectedPlant *PlantAbstract
		plants := *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage)
		for plant := range plants {
			if plant.IsSelected {
				if selectedPlant == nil {
					selectedPlant = plant
				} else {
					plant.IsSelected = false
					modified = true
					stager.logAndNotify(fmt.Sprintf("Deselected plant %s (only single selected plant allowed)", plant.Name))
				}
			}
		}

		checkDefaultDiagramForPlant := func(p *PlantAbstract) {
			if p.PlantType == TubeVase {
				if len(p.TubeVase3DDiagrams) > 0 {
					p.TubeVase3DDiagrams[0].IsChecked = true
					p.TubeVase3DDiagrams[0].IsExpanded = true
					p.IsTubeVase3DDiagramsNodeExpanded = true
				} else if len(p.Vase2DDiagrams) > 0 {
					p.Vase2DDiagrams[0].IsChecked = true
				}
			} else if p.PlantType == TrapezeVase {
				if len(p.Vase2DDiagrams) > 0 {
					p.Vase2DDiagrams[0].IsChecked = true
				}
			} else if p.PlantType == Stool {
				if len(p.Stool3DDiagrams) > 0 {
					p.Stool3DDiagrams[0].IsChecked = true
				} else if len(p.Stool2DDiagrams) > 0 {
					p.Stool2DDiagrams[0].IsChecked = true
				}
			} else if p.PlantType == Clock {
				if len(p.Clock3DDiagrams) > 0 {
					p.Clock3DDiagrams[0].IsChecked = true
				} else if len(p.Clock2DDiagrams) > 0 {
					p.Clock2DDiagrams[0].IsChecked = true
				}
			} else {
				if len(p.Plant2DDiagrams) > 0 {
					p.Plant2DDiagrams[0].IsChecked = true
				}
			}
		}

		if selectedPlant != nil {
			if stager.selectedPlant != selectedPlant {
				stager.selectedPlant = selectedPlant
			}
			checkDefaultDiagramForPlant(selectedPlant)
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: checked default diagram", selectedPlant.Name))
		} else if len(plants) > 0 {
			for plant := range plants {
				plant.IsSelected = true
				stager.selectedPlant = plant
				checkDefaultDiagramForPlant(plant)
				modified = true
				stager.logAndNotify(fmt.Sprintf("Selected default plant %s and checked default diagram", plant.Name))
				break
			}
		} else {
			if stager.selectedPlant != nil {
				stager.selectedPlant = nil
			}
		}
	}

	if stager.selectedPlant != nil {
		if stager.selectedPlant.PlantType == Plant && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: reset invalid CurrentView to VIEW_PLANT_2D", stager.selectedPlant.Name))
		} else if stager.selectedPlant.PlantType == Stool && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_STOOL_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: reset invalid CurrentView to VIEW_PLANT_2D", stager.selectedPlant.Name))
		} else if stager.selectedPlant.PlantType == Clock && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_CLOCK_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: reset invalid CurrentView to VIEW_PLANT_2D", stager.selectedPlant.Name))
		} else if stager.selectedPlant.PlantType == TubeVase && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_VASE_FORM && stager.selectedPlant.CurrentView != VIEW_VASE_2D && stager.selectedPlant.CurrentView != VIEW_TUBE_VASE_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: reset invalid CurrentView to VIEW_PLANT_2D", stager.selectedPlant.Name))
		} else if stager.selectedPlant.PlantType == TrapezeVase && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_VASE_FORM && stager.selectedPlant.CurrentView != VIEW_VASE_2D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
			stager.logAndNotify(fmt.Sprintf("Plant %s: reset invalid CurrentView to VIEW_PLANT_2D", stager.selectedPlant.Name))
		}
	}

	return modified
}
func (stager *Stager) enforcePlantRotationRatioHeights() bool {
	modified := false
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType != TubeVase {
			continue
		}
		h0 := ComputeStackHeightForRotationRatio(plant, 0.0)
		if plant.TubeVaseAbstract.heightAtRotRatio0 != h0 {
			plant.TubeVaseAbstract.heightAtRotRatio0 = h0
			modified = true
		}
		h1 := ComputeStackHeightForRotationRatio(plant, 1.0)
		if plant.TubeVaseAbstract.heightAtRotRatio1 != h1 {
			plant.TubeVaseAbstract.heightAtRotRatio1 = h1
			modified = true
		}
	}
	return modified
}
