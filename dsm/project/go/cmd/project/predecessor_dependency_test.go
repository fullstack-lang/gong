package main

import (
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func TestEnforceTaskPredecessorDates(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)
	_ = stager
	stage := stack.Stage

	(&models.Library{
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	// Base dates for predecessor T1: Jan 1 to Feb 1
	t1Start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1End := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)

	t1 := (&models.Task{
		Name:  "T1",
		Start: t1Start,
		End:   t1End,
	}).Stage(stage)

	// FS: Successor Start from Predecessor End
	taskFS := (&models.Task{
		Name:           "TaskFS",
		Predecessors:   []*models.Task{t1},
		DependencyType: models.FINISH_TO_START,
		Start:          time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
		End:            time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	// SS: Successor Start from Predecessor Start
	taskSS := (&models.Task{
		Name:           "TaskSS",
		Predecessors:   []*models.Task{t1},
		DependencyType: models.START_TO_START,
		Start:          time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
		End:            time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	// FF: Successor End from Predecessor End
	taskFF := (&models.Task{
		Name:           "TaskFF",
		Predecessors:   []*models.Task{t1},
		DependencyType: models.FINISH_TO_FINISH,
		Start:          time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
		End:            time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	// SF: Successor End from Predecessor Start
	taskSF := (&models.Task{
		Name:           "TaskSF",
		Predecessors:   []*models.Task{t1},
		DependencyType: models.START_TO_FINISH,
		Start:          time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
		End:            time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	// FF with Duration: End is fixed by predecessor, Start should shift backwards
	taskFFDuration := (&models.Task{
		Name:                          "TaskFFDuration",
		Predecessors:                  []*models.Task{t1},
		DependencyType:                models.FINISH_TO_FINISH,
		DurationDays:                  10,
		IsEndDateComputedFromDuration: true,
		Start:                         time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		End:                           time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	stage.Commit()

	if !taskFS.Start.Equal(t1End) {
		t.Errorf("taskFS.Start = %v, expected %v", taskFS.Start, t1End)
	}

	if !taskSS.Start.Equal(t1Start) {
		t.Errorf("taskSS.Start = %v, expected %v", taskSS.Start, t1Start)
	}

	if !taskFF.End.Equal(t1End) {
		t.Errorf("taskFF.End = %v, expected %v", taskFF.End, t1End)
	}

	if !taskSF.End.Equal(t1Start) {
		t.Errorf("taskSF.End = %v, expected %v", taskSF.End, t1Start)
	}

	// For taskFFDuration: End = t1End (Feb 1), Duration = 10 days, so Start should be Jan 22
	expectedStart := t1End.AddDate(0, 0, -10)
	if !taskFFDuration.End.Equal(t1End) {
		t.Errorf("taskFFDuration.End = %v, expected %v", taskFFDuration.End, t1End)
	}
	if !taskFFDuration.Start.Equal(expectedStart) {
		t.Errorf("taskFFDuration.Start = %v, expected %v", taskFFDuration.Start, expectedStart)
	}
}

func TestEnforceTaskDependencyDuration(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_dep_duration", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)
	_ = stager
	stage := stack.Stage

	(&models.Library{
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	// Base dates for predecessor T1: Jan 1 to Feb 1
	t1Start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1End := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)

	t1 := (&models.Task{
		Name:  "T1",
		Start: t1Start,
		End:   t1End,
	}).Stage(stage)

	// FS with lag: 1 week and 3 days (10 days lag from T1 End)
	taskFSLag := (&models.Task{
		Name:                    "TaskFSLag",
		Predecessors:            []*models.Task{t1},
		DependencyType:          models.FINISH_TO_START,
		DependencyDurationWeeks: 1,
		DependencyDurationDays:  3,
		DurationDays:            5,
		IsEndDateComputedFromDuration: true,
	}).Stage(stage)

	// SS with lag: 2 months lag from T1 Start
	taskSSLag := (&models.Task{
		Name:                     "TaskSSLag",
		Predecessors:             []*models.Task{t1},
		DependencyType:           models.START_TO_START,
		DependencyDurationMonths: 2,
	}).Stage(stage)

	// FF with lag: 4 days lag from T1 End, Duration = 6 days, so Start should be shifted backwards
	taskFFLag := (&models.Task{
		Name:                          "TaskFFLag",
		Predecessors:                  []*models.Task{t1},
		DependencyType:                models.FINISH_TO_FINISH,
		DependencyDurationDays:        4,
		DurationDays:                  6,
		IsEndDateComputedFromDuration: true,
	}).Stage(stage)

	// Milestone with lag: 2 days after T1 End
	milestoneLag := (&models.Task{
		Name:                   "MilestoneLag",
		Predecessors:           []*models.Task{t1},
		DependencyType:         models.FINISH_TO_START,
		DependencyDurationDays: 2,
		IsMilestone:            true,
	}).Stage(stage)

	stage.Commit()

	// taskFSLag: Start = Feb 1 + 10 days = Feb 11; End = Feb 11 + 5 days = Feb 16
	expectedFSStart := t1End.AddDate(0, 0, 10)
	expectedFSEnd := expectedFSStart.AddDate(0, 0, 5)
	if !taskFSLag.Start.Equal(expectedFSStart) {
		t.Errorf("taskFSLag.Start = %v, expected %v", taskFSLag.Start, expectedFSStart)
	}
	if !taskFSLag.End.Equal(expectedFSEnd) {
		t.Errorf("taskFSLag.End = %v, expected %v", taskFSLag.End, expectedFSEnd)
	}

	// taskSSLag: Start = Jan 1 + 2 months = Mar 1
	expectedSSStart := t1Start.AddDate(0, 2, 0)
	if !taskSSLag.Start.Equal(expectedSSStart) {
		t.Errorf("taskSSLag.Start = %v, expected %v", taskSSLag.Start, expectedSSStart)
	}

	// taskFFLag: End = Feb 1 + 4 days = Feb 5; Start = Feb 5 - 6 days = Jan 30
	expectedFFEnd := t1End.AddDate(0, 0, 4)
	expectedFFStart := expectedFFEnd.AddDate(0, 0, -6)
	if !taskFFLag.End.Equal(expectedFFEnd) {
		t.Errorf("taskFFLag.End = %v, expected %v", taskFFLag.End, expectedFFEnd)
	}
	if !taskFFLag.Start.Equal(expectedFFStart) {
		t.Errorf("taskFFLag.Start = %v, expected %v", taskFFLag.Start, expectedFFStart)
	}

	// milestoneLag: Start = End = Feb 1 + 2 days = Feb 3
	expectedMilestoneDate := t1End.AddDate(0, 0, 2)
	if !milestoneLag.Start.Equal(expectedMilestoneDate) {
		t.Errorf("milestoneLag.Start = %v, expected %v", milestoneLag.Start, expectedMilestoneDate)
	}
	if !milestoneLag.End.Equal(expectedMilestoneDate) {
		t.Errorf("milestoneLag.End = %v, expected %v", milestoneLag.End, expectedMilestoneDate)
	}
}

