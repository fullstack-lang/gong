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
		{"Enforce task predecessor dates", stager.enforceTaskPredecessorDates},
		{"Enforce task duration dates", stager.enforceTaskDurationDates},
		{"Enforce task milestone dates", stager.enforceTaskMilestoneDates},

		// concrete semantic check

		{"Enforce visibility", stager.enforceVisibility},
		{"Enforce relation duplicates", stager.enforceRelationDuplicates},
		{"Enforce node shape duplicates", stager.enforceNodeShapeDuplicates},
		{"Enforce shape orphans", stager.enforceShapeOrphans},
		{"Enforce shapes abstract consistency", stager.enforceShapesAbstractConsistency},
		{"Enforce auto layout", stager.enforceAutoLayout},
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
	for _, task := range stager.stage.GetInstancesSorted[*Task]() {
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
	for _, task := range stager.stage.GetInstancesSorted[*Task]() {
		if task.IsEndDateComputedFromDuration {
			// If End date is computed from predecessors (FF or SF), shift Start backwards from End
			if len(task.Predecessors) > 0 && (task.DependencyType == FINISH_TO_FINISH || task.DependencyType == START_TO_FINISH) {
				days := task.DurationWeeks*7 + task.DurationDays
				fractionalDays := days - float64(int(days))
				hours := task.DurationHours + fractionalDays*24

				effectiveEnd := task.End
				if task.IsAllDay && !task.IsMilestone {
					effectiveEnd = effectiveEnd.AddDate(0, 0, 1)
				}

				expectedStart := effectiveEnd.AddDate(
					-int(task.DurationYears),
					-int(task.DurationMonths),
					-int(days),
				).Add(-time.Duration(hours * float64(time.Hour)))

				if !task.Start.Equal(expectedStart) {
					task.Start = expectedStart
					needCommit = true
					if stager.probeForm != nil {
						stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted start date backwards from duration and predecessor", task.Name))
					}
				}
				continue
			}

			days := task.DurationWeeks*7 + task.DurationDays
			fractionalDays := days - float64(int(days))
			hours := task.DurationHours + fractionalDays*24

			expectedEnd := task.Start.AddDate(
				int(task.DurationYears),
				int(task.DurationMonths),
				int(days),
			).Add(time.Duration(hours * float64(time.Hour)))

			if task.IsAllDay && !task.IsMilestone && (task.DurationYears > 0 || task.DurationMonths > 0 || days > 0 || hours > 0) {
				expectedEnd = expectedEnd.AddDate(0, 0, -1)
			}

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
	for _, task := range stager.stage.GetInstancesSorted[*Task]() {
		if task.DependencyType == "" || task.DependencyType == NO_DEPENDENCY || len(task.Predecessors) == 0 {
			continue
		}

		var maxDate time.Time
		first := true
		for _, predecessor := range task.Predecessors {
			if predecessor == nil {
				continue
			}
			var predDate time.Time
			switch task.DependencyType {
			case FINISH_TO_START, FINISH_TO_FINISH:
				predDate = predecessor.End
				if predecessor.IsAllDay && !predecessor.IsMilestone {
					predDate = predDate.AddDate(0, 0, 1)
				}
			case START_TO_START, START_TO_FINISH:
				predDate = predecessor.Start
			}

			if first || predDate.After(maxDate) {
				maxDate = predDate
				first = false
			}
		}

		if first {
			continue
		}

		depDays := task.DependencyDurationWeeks*7 + task.DependencyDurationDays
		fractionalDays := depDays - float64(int(depDays))
		depHours := task.DependencyDurationHours + fractionalDays*24

		targetDate := maxDate.AddDate(
			int(task.DependencyDurationYears),
			int(task.DependencyDurationMonths),
			int(depDays),
		).Add(time.Duration(depHours * float64(time.Hour)))

		switch task.DependencyType {
		case FINISH_TO_START: // Start from End (FS)
			if !task.Start.Equal(targetDate) {
				task.Start = targetDate
				if task.IsMilestone {
					task.End = targetDate
				}
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted start date from predecessor end date and dependency duration (FS)", task.Name))
				}
			} else if task.IsMilestone && !task.End.Equal(targetDate) {
				task.End = targetDate
				needCommit = true
			}
		case START_TO_START: // Start from Start (SS)
			if !task.Start.Equal(targetDate) {
				task.Start = targetDate
				if task.IsMilestone {
					task.End = targetDate
				}
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted start date from predecessor start date and dependency duration (SS)", task.Name))
				}
			} else if task.IsMilestone && !task.End.Equal(targetDate) {
				task.End = targetDate
				needCommit = true
			}
		case FINISH_TO_FINISH: // End from End (FF)
			expectedEnd := targetDate
			if task.IsAllDay && !task.IsMilestone {
				expectedEnd = expectedEnd.AddDate(0, 0, -1)
			}
			if !task.End.Equal(expectedEnd) {
				task.End = expectedEnd
				if task.IsMilestone {
					task.Start = expectedEnd
				}
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted end date from predecessor end date and dependency duration (FF)", task.Name))
				}
			} else if task.IsMilestone && !task.Start.Equal(expectedEnd) {
				task.Start = expectedEnd
				needCommit = true
			}
		case START_TO_FINISH: // End from Start (SF)
			expectedEnd := targetDate
			if task.IsAllDay && !task.IsMilestone {
				expectedEnd = expectedEnd.AddDate(0, 0, -1)
			}
			if !task.End.Equal(expectedEnd) {
				task.End = expectedEnd
				if task.IsMilestone {
					task.Start = expectedEnd
				}
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task %s: shifted end date from predecessor start date and dependency duration (SF)", task.Name))
				}
			} else if task.IsMilestone && !task.Start.Equal(expectedEnd) {
				task.Start = expectedEnd
				needCommit = true
			}
		}
	}
	return
}
