package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage

	// computes fields that are not persisted
	stager.enforceProducersConsumers()
	stager.enforceLibrariesObject()
	stager.enforceDiagramMaps()
	stager.enforceParentAssociation()

	pass := 0
	for {
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
	stager.enforceProducersConsumers()
	stager.enforceLibrariesObject()
	stager.enforceDiagramMaps()
	stager.enforceParentAssociation()

	if needCommit {
		stager.probeForm.CommitNotificationTable()
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	methods := []struct {
		name string
		fn   func() bool
	}{
		// VERY important because the probe only unstages objects
		// this is the Clean that delete them from slices and pointers that reference
		// them. If the checkout is not performed, the stage might be dirty
		// with slices of pointer or pointer to unstaged instance
		{"stage.Clean", func() bool { return stage.Clean() }},
		{"enforceRoot", stager.enforceRoot},
		// Enforce that all libraries are appended to the [root]
		// if one library is not appended, append it
		{"enforceLibrariesAppendedToRoot", stager.enforceLibrariesAppendedToRoot},
		{"enforceDefaultValues", stager.enforceDefaultValues},
		{"enforceTreesAndDAG", stager.enforceTreesAndDAG},
		{"unstageAllOrphans", stager.unstageAllOrphans},
		{"enforceComputedPrefix", stager.enforceComputedPrefix},
		// enforce visibility will unstage shapes that are not visible
		// and indirectly shapes whose abstract element is not staged (because they are not visible)
		{"enforceVisibility", stager.enforceVisibility},
		{"enforceRelationDuplicates", stager.enforceRelationDuplicates},
		{"enforceNodeShapeDuplicates", stager.enforceNodeShapeDuplicates},
		{"enforceShapeOrphans", stager.enforceShapeOrphans},
		{"enforceShapesAbstractConsistency", func() bool { return stager.enforceShapesAbstractConsistency(stage, false) }},
		{"enforceDiagramSize", stager.enforceDiagramSize},
		{"enforceAssociationShapeConsistency", stager.enforceAssociationShapeConsistency},
		{"enforceTaskInputOutputLibraryConsistency", stager.enforceTaskInputOutputLibraryConsistency},
		{"enforceDuplicateRemove", stager.enforceDuplicateRemove},
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

func (stager *Stager) enforceLibrariesAppendedToRoot() (needCommit bool) {
	root := stager.root
	for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
		if slices.Index(root.Libraries, library) == -1 {
			root.Libraries = append(root.Libraries, library)
			needCommit = true
		}
	}
	return
}

func (stager *Stager) enforceRoot() (needCommit bool) {
	stage := stager.stage
	// Ensures that there is one and only one root
	// prune the other
	// check that there is at least one root
	// and that one can safely access stager.root
	roots := GetGongstrucsSorted[*Root](stage)
	if len(roots) == 0 {
		stager.root = (&Root{Name: "Root"}).Stage(stage)
		needCommit = true
	} else {
		stager.root = roots[0]
		if len(roots) > 1 {
			for _, root := range roots[1:] {
				root.Unstage(stage)
				needCommit = true
			}
		}
	}
	return
}
