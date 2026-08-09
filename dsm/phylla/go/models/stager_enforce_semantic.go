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
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprint("Stage was modified to enforce semantic, pass ", pass))
			}
			pass++
		} else {
			break
		}
	}

	// computes fields that are not persisted
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()

	if needCommit {
		stager.probeForm.CommitNotificationTable()
		stage.CommitWithSuspendedCallbacks()
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	methods := []struct {
		name string
		fn   func() bool
	}{
		// abstract semantic check

		// VERY important because the probe only unstages objects
		// this is the Clean that delete them from slices and pointers that reference
		// them. If the checkout is not performed, the stage might be dirty
		// with slices of pointer or pointer to unstaged instance
		{"Clean the stage", func() bool { return stage.Clean() }},
		{"Enforce orphans abstract element", stager.enforceOrphansAbstractElement},
		{"Enforce orphan shape remove", stager.enforceOrphanShapeRemove},
		{"Enforce default values", stager.enforceDefaultValues},
		{"Enforce N <= M", stager.enforcePlantNM},
		{"Enforce duplicate remove", stager.enforceDuplicateRemove},
		{"Enforce single selected plant", stager.enforceSingleSelectedPlant},
		{"Enforce plant has diagram", stager.enforcePlantHasDiagram},
		{"Enforce plant has axes", stager.enforcePlantHasAxes},
		{"Enforce axes shape name", stager.enforceAxesShapeName},
		{"Enforce plant has rhombus stuff", stager.enforcePlantHasRhombusStuff},
		{"Enforce rhombus stuff name", stager.enforceRhombusStuffName},
		{"Enforce plant has vase abstract", stager.enforcePlantHasVaseAbstract},
		{"Enforce vase abstract name", stager.enforceVaseAbstractName},
		{"Enforce plant has reference rhombus", stager.enforcePlantHasReferenceRhombus},
		{"Enforce reference rhombus name", stager.enforceReferenceRhombusName},
		{"Enforce plant has grid path shape", stager.enforcePlantHasGridPathShape},
		{"Enforce grid path shape name", stager.enforceGridPathShapeName},
		{"Enforce plant has initial rhombus grid shape", stager.enforcePlantHasInitialRhombusGridShape},
		{"Enforce initial rhombus grid shape name", stager.enforceInitialRhombusGridShapeName},
		{"Enforce plant has explanation text shape", stager.enforcePlantHasExplanationTextShape},
		{"Enforce explanation text shape name", stager.enforceExplanationTextShapeName},
		{"Enforce plant has rotated shapes", stager.enforcePlantHasRotatedShapes},
		{"Enforce rotated shapes names", stager.enforceRotatedShapesNames},
		{"Enforce plant has growth vector shape", stager.enforcePlantHasPlantCircumferenceShape},
		{"Enforce compute growth vector shape", stager.enforceComputePlantCircumferenceShape},
		{"Enforce growth vector shape name", stager.enforcePlantCircumferenceShapeName},
		{"Enforce rhombus grid shape has rhombuses", stager.enforcePlantRhombusGridShapeHasRhombuses},
		{"Enforce plant diagram vase diagram", stager.enforcePlantDiagramVaseDiagram},
		{"Enforce plant diagram rendered 3d shape", stager.enforcePlantDiagramRendered3DShape},
		{"Enforce plant diagram torus stack shape", stager.enforcePlantDiagramTorusStackShape},
		{"Enforce plant diagram vertical torus stack shape", stager.enforcePlantDiagramVerticalTorusStackShape},
		{"Enforce plant diagram partially rotated torus shape", stager.enforcePlantDiagramPartiallyRotatedTorusShape},
		{"Enforce plant diagram stack of partially rotated torus shape", stager.enforcePlantDiagramStackOfPartiallyRotatedTorusShape},
		{"Enforce plant diagram points and lines 3d shape", stager.enforcePlantDiagramPointsAndLines3DShape},
		{"Enforce plant diagram key hole 3d shape", stager.enforcePlantDiagramKeyHole3DShape},
		{"Enforce plant diagram key 3d shape", stager.enforcePlantDiagramKey3DShape},
		{"Enforce plant diagram volume key 3d shape", stager.enforcePlantDiagramVolumeKey3DShape},
		{"Enforce plant diagram torus edge 3d shape", stager.enforcePlantDiagramTorusEdge3DShape},
		{"Enforce plant rotation ratio heights", stager.enforcePlantRotationRatioHeights},

		// concrete semantic check

	}

	for _, method := range methods {
		modified := method.fn()
		if modified {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Semantic check '%s' generated a stage modification", method.name))
			}
			needCommit = true
		}
	}

	return needCommit
}

func (stager *Stager) enforceSingleSelectedPlant() bool {
	modified := false

	// If there is a checked PlantDiagram, the plant owning that diagram must be the selected plant
	var plantWithCheckedDiagram *PlantAbstract
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		for _, diagram := range plant.PlantDiagrams {
			if diagram.IsChecked {
				plantWithCheckedDiagram = plant
				break
			}
		}
		if plantWithCheckedDiagram != nil {
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

		if selectedPlant != nil {
			if stager.selectedPlant != selectedPlant {
				stager.selectedPlant = selectedPlant
			}
			if len(selectedPlant.PlantDiagrams) > 0 && !selectedPlant.PlantDiagrams[0].IsChecked {
				for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stager.stage) {
					plantDiagram_.IsChecked = false
				}
				selectedPlant.PlantDiagrams[0].IsChecked = true
				modified = true
			}
		} else if len(plants) > 0 {
			for plant := range plants {
				plant.IsSelected = true
				stager.selectedPlant = plant
				if len(plant.PlantDiagrams) > 0 && !plant.PlantDiagrams[0].IsChecked {
					for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stager.stage) {
						plantDiagram_.IsChecked = false
					}
					plant.PlantDiagrams[0].IsChecked = true
				}
				modified = true
				break
			}
		} else {
			if stager.selectedPlant != nil {
				stager.selectedPlant = nil
			}
		}
	}

	if stager.selectedPlant != nil && stager.selectedPlant.PlantType != Vase && stager.selectedPlant.CurrentView != VIEW_PLANT_2D {
		stager.selectedPlant.CurrentView = VIEW_PLANT_2D
		modified = true
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
