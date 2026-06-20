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
	stager.computePortFields()
	stager.computePartFields()
	stager.computeDataFlowShapeFields()
	stager.computePortShapeFields()

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
	stager.computePortFields()
	stager.computePartFields()
	stager.computeDataFlowShapeFields()
	stager.computePortShapeFields()

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
		{"Enforce port part consistency", stager.enforcePortPartConsistency},
		{"Enforce port semantic rules", stager.enforcePortSemanticRules},
		{"Enforce part semantic rules", stager.enforcePartSemanticRules},
		{"Enforce control flow semantic rules", stager.enforceControlFlowRules},
		{"Enforce data flow semantic rules", stager.enforceDataFlowRules},

		// concrete semantic check
		{"Enforce at least one diagram per system", stager.enforceAtLeastOneDiagramPerSystem},
		{"Enforce a system diagram has its owning system", stager.enforceASystemDiagramHasItsOwningSystem},
		{"Enforce node shape duplicates", stager.enforceNodeShapeDuplicates},
		{"Enforce relation duplicates", stager.enforceRelationDuplicates},
		{"Enforce shape orphans", stager.enforceShapeOrphans},
		{"Enforce control flow shapes rules", stager.enforceControlFlowShapesRules},
		{"Enforce data flow shapes rules", stager.enforceDataFlowShapesRules},
		{"Enforce part shape semantic", stager.enforcePartShapeSemantic},
		{"Enforce port shape within part", stager.enforcePortShapeWithinPart},
		{"Enforce port shape part presence", stager.enforcePortShapePartPresence},
		{"Enforce shapes abstract consistency", stager.enforceShapesAbstractConsistency},
		{"Enforce diagram size", stager.enforceDiagramSize},
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
