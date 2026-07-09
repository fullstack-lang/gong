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
		{"Enforce default values", stager.enforceDefaultValues},
		{"Enforce N <= M", stager.enforcePlantNM},
		{"Enforce duplicate remove", stager.enforceDuplicateRemove},
		{"Enforce single selected plant", stager.enforceSingleSelectedPlant},
		{"Enforce plant has diagram", stager.enforcePlantHasDiagram},
		{"Enforce plant diagram has axes", stager.enforcePlantDiagramHasAxes},
		{"Enforce axes shape name", stager.enforceAxesShapeName},
		{"Enforce plant diagram has reference rhombus", stager.enforcePlantDiagramHasReferenceRhombus},
		{"Enforce reference rhombus name", stager.enforceReferenceRhombusName},
		{"Enforce plant diagram has grid path shape", stager.enforcePlantDiagramHasGridPathShape},
		{"Enforce grid path shape name", stager.enforceGridPathShapeName},
		{"Enforce plant diagram has rhombus grid shape", stager.enforcePlantDiagramHasRhombusGridShape},
		{"Enforce rhombus grid shape name", stager.enforceRhombusGridShapeName},
		{"Enforce rhombus grid shape has rhombuses", stager.enforcePlantDiagramRhombusGridShapeHasRhombuses},
		{"Enforce plant diagram has explanation text shape", stager.enforcePlantDiagramHasExplanationTextShape},
		{"Enforce explanation text shape name", stager.enforceExplanationTextShapeName},
		{"Enforce plant diagram has rotated shapes", stager.enforcePlantDiagramHasRotatedShapes},
		{"Enforce rotated shapes names", stager.enforceRotatedShapesNames},
		{"Enforce plant diagram has growth vector shape", stager.enforcePlantDiagramHasPlantCircumferenceShape},
		{"Enforce compute growth vector shape", stager.enforceComputePlantCircumferenceShape},
		{"Enforce growth vector shape name", stager.enforcePlantCircumferenceShapeName},

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

	var selectedPlant *Plant
	for plant := range stager.stage.Plants {
		if plant.IsSelected {
			if selectedPlant == nil {
				selectedPlant = plant
			} else {
				// Only one plant can be selected, unselect the other one
				plant.IsSelected = false
				modified = true
			}
		}
	}

	if selectedPlant != nil {
		if stager.selectedPlant != selectedPlant {
			stager.selectedPlant = selectedPlant
		}
	} else if len(stager.stage.Plants) > 0 {
		// No plant selected, select the first one
		for plant := range stager.stage.Plants {
			plant.IsSelected = true
			stager.selectedPlant = plant
			modified = true
			break
		}
	} else {
		if stager.selectedPlant != nil {
			stager.selectedPlant = nil
		}
	}

	return modified
}
