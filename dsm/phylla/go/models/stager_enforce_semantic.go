package models

import (
	"fmt"
	"log"
	"time"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage
	needCommit = stager.enforceThereIsARootLibrary() || needCommit

	// computes fields that are not persisted
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()

	pass := 0
	for {
		if pass > 10 {
			log.Println("enforceSemantic reached 10 passes. Breaking loop.")
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), "Semantic enforcement reached maximum number of passes (10). Breaking loop.")
			}
			break
		}
		if stager.enforceSemanticOnePass(false, stage) {
			needCommit = true
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
		{"Enforce single selected plant", stager.enforceSingleSelectedPlant, false},
		{"Enforce plant has diagram", stager.enforcePlantHasDiagram, true},

		// concrete / omit shape generation (normal runtime behavior, not notified)
		{"Enforce plant has axes", stager.enforcePlantHasAxes, false},
		{"Enforce axes shape name", stager.enforceAxesShapeName, false},
		{"Enforce plant has rhombus stuff", stager.enforcePlantHasRhombusStuff, false},
		{"Enforce rhombus stuff name", stager.enforceRhombusStuffName, false},
		{"Enforce plant has vase abstract", stager.enforcePlantHasVaseAbstract, false},
		{"Enforce vase abstract name", stager.enforceVaseAbstractName, false},
		{"Enforce plant has stool abstract", stager.enforcePlantHasStoolAbstract, false},
		{"Enforce stool abstract name", stager.enforceStoolAbstractName, false},
		{"Enforce plant has clock abstract", stager.enforcePlantHasClockAbstract, false},
		{"Enforce clock abstract name", stager.enforceClockAbstractName, false},
		{"Enforce plant has reference rhombus", stager.enforcePlantHasReferenceRhombus, false},
		{"Enforce reference rhombus name", stager.enforceReferenceRhombusName, false},
		{"Enforce plant has grid path shape", stager.enforcePlantHasGridPathShape, false},
		{"Enforce grid path shape name", stager.enforceGridPathShapeName, false},
		{"Enforce plant has initial rhombus grid shape", stager.enforcePlantHasInitialRhombusGridShape, false},
		{"Enforce initial rhombus grid shape name", stager.enforceInitialRhombusGridShapeName, false},
		{"Enforce plant has explanation text shape", stager.enforcePlantHasExplanationTextShape, false},
		{"Enforce explanation text shape name", stager.enforceExplanationTextShapeName, false},
		{"Enforce plant has rotated shapes", stager.enforcePlantHasRotatedShapes, false},
		{"Enforce rotated shapes names", stager.enforceRotatedShapesNames, false},
		{"Enforce vase has shapes", stager.enforceVaseHasShapes, false},
		{"Enforce vase shape names", stager.enforceVaseShapeNames, false},
		{"Enforce plant has growth vector shape", stager.enforcePlantHasPlantCircumferenceShape, false},
		{"Enforce compute growth vector shape", stager.enforceComputePlantCircumferenceShape, false},
		{"Enforce growth vector shape name", stager.enforcePlantCircumferenceShapeName, false},
		{"Enforce rhombus grid shape has rhombuses", stager.enforcePlantRhombusGridShapeHasRhombuses, false},
		{"Enforce diagram shapes", stager.enforceDiagramShapes, false},
		{"Enforce plant rotation ratio heights", stager.enforcePlantRotationRatioHeights, false},
	}

	for _, method := range methods {
		modified := method.fn()
		if modified {
			if method.notify {
				log.Printf("Semantic check '%s' generated a stage modification", method.name)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Semantic check '%s' generated a stage modification", method.name))
				}
			}
			needCommit = true
		}
	}

	return needCommit
}


func (stager *Stager) enforceSingleSelectedPlant() bool {
	modified := false

	// Try to find if any diagram across all plants is checked
	var plantWithCheckedDiagram *PlantAbstract
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		checked := false
		for _, d := range plant.Plant2DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Plant3DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Vase2DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Vase3DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Stool2DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Stool3DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Clock2DDiagrams { if d.IsChecked { checked = true } }
		for _, d := range plant.Clock3DDiagrams { if d.IsChecked { checked = true } }

		if checked {
			plantWithCheckedDiagram = plant
			break
		}
	}

	if plantWithCheckedDiagram != nil {
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			shouldBeSelected := (plant == plantWithCheckedDiagram)
			if plant.IsSelected != shouldBeSelected {
				plant.IsSelected = shouldBeSelected
				modified = true
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
				}
			}
		}

		checkDefaultDiagramForPlant := func(p *PlantAbstract) {
			if p.PlantType == Vase {
				if len(p.Vase3DDiagrams) > 0 {
					p.Vase3DDiagrams[0].IsChecked = true
					p.Vase3DDiagrams[0].IsExpanded = true
					p.IsVase3DDiagramsNodeExpanded = true
				} else if len(p.Vase2DDiagrams) > 0 {
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
		} else if len(plants) > 0 {
			for plant := range plants {
				plant.IsSelected = true
				stager.selectedPlant = plant
				checkDefaultDiagramForPlant(plant)
				modified = true
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
		} else if stager.selectedPlant.PlantType == Stool && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_STOOL_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
		} else if stager.selectedPlant.PlantType == Clock && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_CLOCK_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
		} else if stager.selectedPlant.PlantType == Vase && stager.selectedPlant.CurrentView != VIEW_PLANT_2D && stager.selectedPlant.CurrentView != VIEW_PLANT_3D && stager.selectedPlant.CurrentView != VIEW_VASE_FORM && stager.selectedPlant.CurrentView != VIEW_VASE_2D && stager.selectedPlant.CurrentView != VIEW_VASE_3D && stager.selectedPlant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			stager.selectedPlant.CurrentView = VIEW_PLANT_2D
			modified = true
		}
	}

	return modified
}
func (stager *Stager) enforcePlantRotationRatioHeights() bool {
	modified := false
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType != Vase {
			continue
		}
		h0 := ComputeStackHeightForRotationRatio(plant, 0.0)
		if plant.VaseAbstract.heightAtRotRatio0 != h0 {
			plant.VaseAbstract.heightAtRotRatio0 = h0
			modified = true
		}
		h1 := ComputeStackHeightForRotationRatio(plant, 1.0)
		if plant.VaseAbstract.heightAtRotRatio1 != h1 {
			plant.VaseAbstract.heightAtRotRatio1 = h1
			modified = true
		}
	}
	return modified
}
