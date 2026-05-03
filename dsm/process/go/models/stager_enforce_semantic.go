package models

import (
	"fmt"
	"log"
	"time"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage
	needCommit = stager.enforceThereIsADefaultLibrary() || needCommit

	// computes fields that are not persisted
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()
	stager.computeTaskFields()
	stager.computeParticipantFields()

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
	stager.computeTaskFields()
	stager.computeParticipantFields()

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
		{"Clean stage", func() bool { return stage.Clean() }},
		{"Enforce default values", stager.enforceDefaultValues},
		{"Enforce orphans abstract element", stager.enforceOrphansAbstractElement},
		{"Enforce task participant consistency", stager.enforceTaskParticipantConsistency},
		{"Enforce task semantic rules", stager.enforceTaskSemanticRules},
		{"Enforce control flow semantic rules", stager.enforceControlFlowRules},
		{"Enforce data flow semantic rules", stager.enforceDataFlowRules},

		// concrete semantic check
		{"Enforce at least one diagram per process", stager.enforceAtLeastOneDiagramPerProcess},
		{"Enforce a process diagram has its owning process", stager.enforceAProcessDiagramHasItsOwningProcess},
		{"Enforce node shape duplicates", stager.enforceNodeShapeDuplicates},
		{"Enforce relation duplicates", stager.enforceRelationDuplicates},
		{"Enforce shape orphans", stager.enforceShapeOrphans},
		{"Enforce control flow shapes rules", stager.enforceControlFlowShapesRules},
		{"Enforce participant shape order", stager.enforceParticipantShapeOrder},
		{"Enforce task shape within process", stager.enforceTaskShapeWithinProcess},
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

func (stager *Stager) enforceThereIsADefaultLibrary() (needCommit bool) {
	stage := stager.stage
	libraries := GetStructInstancesByOrderAuto[*Library](stage)
	if len(libraries) == 0 {
		stager.rootLibrary = (&Library{Name: ""}).Stage(stage)
		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				"Created root library")
		}
		needCommit = true
	}
	if stager.rootLibrary == nil {
		stager.rootLibrary = libraries[0] // The first library is the root library
	}

	return
}
