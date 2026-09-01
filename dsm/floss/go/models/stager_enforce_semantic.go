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
	stager.enforceStagerMaps()

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
	stager.enforceStagerMaps()

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
		{"Assign Probe Callback", stager.assignProbeCallback},

		// VERY important because the probe only unstages objects
		// this is the Clean that delete them from slices and pointers that reference
		// them. If the checkout is not performed, the stage might be dirty
		// with slices of pointer or pointer to unstaged instance
		{"Clean stage", func() bool { return stage.Clean() }},
		{"Enforce default values", stager.enforceDefaultValues},
		{"Enforce orphans abstract element", stager.enforceOrphansAbstractElement},

		// concrete semantic check
		{"Enforce at least one diagram per system", stager.enforceAtLeastOneDiagramPerSystem},
		{"Enforce at least one diagram per compare analysis", stager.enforceAtLeastOneDiagramPerCompareAnalysis},
		{"Enforce diagram floss equation exclusive owner", stager.enforceDiagramFlossEquationExclusiveOwner},
		{"Enforce node shape duplicates", stager.enforceNodeShapeDuplicates},
		{"Enforce shape orphans", stager.enforceShapeOrphans},
		{"Enforce shapes not attached to system", stager.enforceShapesNotAttachedToSystem},
		{"Enforce diagram size", stager.enforceDiagramSize},
		{"Enforce floss equation", stager.enforceFlossEquation},
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
