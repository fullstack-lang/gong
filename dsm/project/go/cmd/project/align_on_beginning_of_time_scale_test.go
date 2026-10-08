package main

import (
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func TestAlignOnBeginningOfTimeScale(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_align_beginning", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)
	_ = stager
	stage := stack.Stage

	lib := (&models.Library{
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	task := (&models.Task{
		Name:  "TaskMidMonth",
		Start: time.Date(2026, 3, 15, 10, 30, 0, 0, time.UTC),
		End:   time.Date(2026, 4, 10, 18, 0, 0, 0, time.UTC),
	}).Stage(stage)

	tg := (&models.TaskGroup{
		Name:  "TG1",
		Tasks: []*models.Task{task},
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                        "GanttDiag",
		IsTimeDiagram:               true,
		IsChecked:                   true,
		IsEditable_:                 true,
		AlignOnBeginningOfTimeScale: false,
		TimeStep:                    1,
		TimeStepScale:               models.MONTHS,
	}).Stage(stage)

	// Attach task to diagram
	taskShape := (&models.TaskShape{
		Name: "GanttDiag-TaskMidMonth",
		Task: task,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{taskShape}

	tgShape := (&models.TaskGroupShape{
		Name:      "GanttDiag-TG1",
		TaskGroup: tg,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgShape}

	stage.Commit()

	// When AlignOnBeginningOfTimeScale is false, ComputedStart is task.Start
	if !diag.ComputedStart.Equal(task.Start) {
		t.Errorf("expected ComputedStart to be %v, got %v", task.Start, diag.ComputedStart)
	}

	// Test 1: MONTHS scale
	diag.AlignOnBeginningOfTimeScale = true
	diag.TimeStepScale = models.MONTHS
	stage.Commit()

	expectedMonthStart := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
	if !diag.ComputedStart.Equal(expectedMonthStart) {
		t.Errorf("MONTHS: expected ComputedStart %v, got %v", expectedMonthStart, diag.ComputedStart)
	}

	// Test 2: YEARS scale
	diag.TimeStepScale = models.YEARS
	stage.Commit()

	expectedYearStart := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	if !diag.ComputedStart.Equal(expectedYearStart) {
		t.Errorf("YEARS: expected ComputedStart %v, got %v", expectedYearStart, diag.ComputedStart)
	}

	// Test 3: WEEKS scale (2026-03-15 is Sunday, previous Monday was 2026-03-09)
	diag.TimeStepScale = models.WEEKS
	stage.Commit()

	expectedWeekStart := time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC)
	if !diag.ComputedStart.Equal(expectedWeekStart) {
		t.Errorf("WEEKS: expected ComputedStart %v, got %v", expectedWeekStart, diag.ComputedStart)
	}

	// Test 4: DAYS scale
	diag.TimeStepScale = models.DAYS
	stage.Commit()

	expectedDayStart := time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC)
	if !diag.ComputedStart.Equal(expectedDayStart) {
		t.Errorf("DAYS: expected ComputedStart %v, got %v", expectedDayStart, diag.ComputedStart)
	}
}
