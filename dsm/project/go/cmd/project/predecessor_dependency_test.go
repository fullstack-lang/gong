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
