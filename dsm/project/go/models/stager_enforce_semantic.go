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
	stager.enforceProducersConsumers()
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()
	stager.enforceParentAssociation()

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
	stager.enforceProducersConsumers()
	stager.enforceOwningLibraryAndObjects()
	stager.enforceDiagramMaps()
	stager.enforceParentAssociation()

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
		{"Enforce trees and DAG", stager.enforceTreesAndDAG},
		{"Enforce task input output library consistency", stager.enforceTaskInputOutputLibraryConsistency},
		{"Enforce duplicate remove", stager.enforceDuplicateRemove},
		{"Enforce library has at least one diagram", stager.enforceLibraryHasAtLeastOneDiagram},
		{"Enforce task milestone dates", stager.enforceTaskMilestoneDates},
		{"Enforce task duration dates", stager.enforceTaskDurationDates},
		{"Enforce task predecessor dates", stager.enforceTaskPredecessorDates},

		// concrete semantic check

		{"Enforce visibility", stager.enforceVisibility},
		{"Enforce relation duplicates", stager.enforceRelationDuplicates},
		{"Enforce node shape duplicates", stager.enforceNodeShapeDuplicates},
		{"Enforce shape orphans", stager.enforceShapeOrphans},
		{"Enforce shapes abstract consistency", stager.enforceShapesAbstractConsistency},
		{"Enforce diagram size", stager.enforceDiagramSize},
		{"Enforce association shape consistency", stager.enforceAssociationShapeConsistency},
		{"Enforce shape names", stager.enforceShapeNames},
		{"Enforce diagram dates", stager.enforceDiagramDates},

		// to be performed at the end
		{"Enforce computed prefix", stager.enforceComputedPrefix},
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

func (stager *Stager) enforceTaskMilestoneDates() (needCommit bool) {
	for _, task := range GetGongstrucsSorted[*Task](stager.stage) {
		if task.IsMilestone && task.End != task.Start {
			task.End = task.Start
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted end date to start date since it is a milestone", task.Name))
			}
		}
	}
	return
}

func (stager *Stager) enforceTaskDurationDates() (needCommit bool) {
	for _, task := range GetGongstrucsSorted[*Task](stager.stage) {
		if task.IsEndDateComputedFromDuration {
			days := task.DurationWeeks*7 + task.DurationDays
			fractionalDays := days - float64(int(days))
			hours := task.DurationHours + fractionalDays*24

			expectedEnd := task.Start.AddDate(
				int(task.DurationYears),
				int(task.DurationMonths),
				int(days),
			).Add(time.Duration(hours * float64(time.Hour)))
			if !task.End.Equal(expectedEnd) {
				task.End = expectedEnd
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted end date from duration", task.Name))
				}
			}
		}
	}
	return
}

func (stager *Stager) enforceTaskPredecessorDates() (needCommit bool) {
	for _, task := range GetGongstrucsSorted[*Task](stager.stage) {
		if task.IsStartDateComputedFromPredecessors && len(task.Predecessors) > 0 {
			var maxEnd time.Time
			for i, predecessor := range task.Predecessors {
				if i == 0 || predecessor.End.After(maxEnd) {
					maxEnd = predecessor.End
				}
			}
			if !task.Start.Equal(maxEnd) {
				task.Start = maxEnd
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted start date from predecessors", task.Name))
				}
			}
		}
	}
	return
}
