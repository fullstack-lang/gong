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
		{"Enforce duplicate remove", stager.enforceDuplicateRemove},
		{"Enforce plant has diagram", stager.enforcePlantHasDiagram},
		{"Enforce plant diagram has axes", stager.enforcePlantDiagramHasAxes},
		{"Enforce axes shape name", stager.enforceAxesShapeName},
		{"Enforce plant diagram has reference rhombus", stager.enforcePlantDiagramHasReferenceRhombus},
		{"Enforce reference rhombus name", stager.enforceReferenceRhombusName},
		{"Enforce plant diagram has grid path shape", stager.enforcePlantDiagramHasGridPathShape},
		{"Enforce grid path shape name", stager.enforceGridPathShapeName},
		{"Enforce plant diagram has rotated shapes", stager.enforcePlantDiagramHasRotatedShapes},
		{"Enforce rotated shapes names", stager.enforceRotatedShapesNames},
		{"Enforce plant diagram has growth vector shape", stager.enforcePlantDiagramHasGrowthVectorShape},
		{"Enforce compute growth vector shape", stager.enforceComputeGrowthVectorShape},
		{"Enforce growth vector shape name", stager.enforceGrowthVectorShapeName},

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
