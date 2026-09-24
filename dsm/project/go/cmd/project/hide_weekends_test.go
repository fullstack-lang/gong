package main

import (
	"math"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func TestHideWeekendsPeriodDateToX(t *testing.T) {
	// Diagram spanning 2 weeks: Monday Jan 5, 2026 to Monday Jan 19, 2026 (14 days, 10 work days)
	// XLeftLanes = 200, XRightMargin = 1200 => width = 1000
	start := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)  // Monday
	end := time.Date(2026, 1, 19, 0, 0, 0, 0, time.UTC)   // Monday (2 weeks later)

	diagram := &models.Diagram{
		Name:             "test_gantt",
		IsTimeDiagram:    true,
		ComputedStart:    start,
		ComputedEnd:      end,
		ComputedDuration: end.Sub(start),
		XLeftLanes:       200.0,
		XRightMargin:     1200.0,
	}

	// 1. Without HideWeekendsPeriod (standard calendar time)
	diagram.HideWeekendsPeriod = false

	// Friday Jan 9 end of day (5 days out of 14)
	friEnd := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC) // Sat 00:00 = Fri 24:00
	xNormalFri := diagram.DateToX(friEnd)
	expectedNormalFri := 200.0 + 1000.0*(5.0/14.0)
	if math.Abs(xNormalFri-expectedNormalFri) > 0.001 {
		t.Errorf("Normal Fri: got %f, expected %f", xNormalFri, expectedNormalFri)
	}

	// Saturday Jan 10 noon
	satNoon := time.Date(2026, 1, 10, 12, 0, 0, 0, time.UTC)
	xNormalSat := diagram.DateToX(satNoon)
	expectedNormalSat := 200.0 + 1000.0*(5.5/14.0)
	if math.Abs(xNormalSat-expectedNormalSat) > 0.001 {
		t.Errorf("Normal Sat: got %f, expected %f", xNormalSat, expectedNormalSat)
	}

	// 2. With HideWeekendsPeriod = true
	diagram.HideWeekendsPeriod = true

	// Friday end of day: 5 work days out of 10 => exactly 50% => x = 700
	xFri := diagram.DateToX(friEnd)
	expectedFri := 700.0
	if math.Abs(xFri-expectedFri) > 0.001 {
		t.Errorf("HideWeekends Fri: got %f, expected %f", xFri, expectedFri)
	}

	// Saturday noon: should have the EXACT SAME x coordinate as Friday end of day!
	xSat := diagram.DateToX(satNoon)
	if math.Abs(xSat-expectedFri) > 0.001 {
		t.Errorf("HideWeekends Sat noon: got %f, expected %f (same as Fri)", xSat, expectedFri)
	}

	// Sunday 18:00: should have the EXACT SAME x coordinate as Friday end of day!
	sunEve := time.Date(2026, 1, 11, 18, 0, 0, 0, time.UTC)
	xSun := diagram.DateToX(sunEve)
	if math.Abs(xSun-expectedFri) > 0.001 {
		t.Errorf("HideWeekends Sun eve: got %f, expected %f (same as Fri)", xSun, expectedFri)
	}

	// Monday Jan 12 00:00: should ALSO have the exact same x coordinate as Friday end of day!
	monStart := time.Date(2026, 1, 12, 0, 0, 0, 0, time.UTC)
	xMon := diagram.DateToX(monStart)
	if math.Abs(xMon-expectedFri) > 0.001 {
		t.Errorf("HideWeekends Mon 00:00: got %f, expected %f (same as Fri)", xMon, expectedFri)
	}

	// Monday Jan 12 12:00 (halfway through Monday): 5.5 work days out of 10 => 55% => x = 750
	monNoon := time.Date(2026, 1, 12, 12, 0, 0, 0, time.UTC)
	xMonNoon := diagram.DateToX(monNoon)
	expectedMonNoon := 200.0 + 1000.0*0.55
	if math.Abs(xMonNoon-expectedMonNoon) > 0.001 {
		t.Errorf("HideWeekends Mon noon: got %f, expected %f", xMonNoon, expectedMonNoon)
	}

	// Task spanning across the weekend: Friday 12:00 to Monday 12:00
	// Working time = 12h (Fri) + 0h (Sat) + 0h (Sun) + 12h (Mon) = 24h = 1 work day
	// 1 work day out of 10 = 10% width = 100px
	tFriNoon := time.Date(2026, 1, 9, 12, 0, 0, 0, time.UTC)
	xStart := diagram.DateToX(tFriNoon)
	xEnd := diagram.DateToX(monNoon)
	taskWidth := xEnd - xStart
	expectedWidth := 100.0
	if math.Abs(taskWidth-expectedWidth) > 0.001 {
		t.Errorf("Task spanning weekend width: got %f, expected %f", taskWidth, expectedWidth)
	}
}

func TestHideWeekendsPeriodSVGIntegration(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_hw_svg", "", "", true, false, false)
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

	// Task from Friday Jan 9 12:00 to Monday Jan 12 12:00
	task := (&models.Task{
		Name:  "WeekendTask",
		Start: time.Date(2026, 1, 9, 12, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 1, 12, 12, 0, 0, 0, time.UTC),
	}).Stage(stage)
	lib.RootTasks = []*models.Task{task}

	tg := (&models.TaskGroup{
		Name:  "Group1",
		Tasks: []*models.Task{task},
	}).Stage(stage)

	diag := (&models.Diagram{
		Name:               "GanttDiag",
		IsTimeDiagram:      true,
		IsChecked:          true,
		XLeftLanes:         200.0,
		XRightMargin:       1200.0,
		LaneHeight:         60.0,
		HideWeekendsPeriod: true,
		UseManualStartAndEndDates: true,
		ManualStart:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
		ManualEnd:          time.Date(2026, 1, 19, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{
		Name:      "GanttDiag-Group1",
		TaskGroup: tg,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{
		Name: "GanttDiag-WeekendTask",
		Task: task,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	stage.Commit()

	// Verify the SVG rect for task
	rect := diag.GetTaskRect(task)
	if rect == nil {
		t.Fatal("Task rect was not generated in SVG")
	}

	// 10 work days total (Jan 5 to Jan 19)
	// Friday noon is day 4.5 -> x = 200 + 1000 * 0.45 = 650
	// Monday noon is day 5.5 -> x = 200 + 1000 * 0.55 = 750
	// Width = 100
	expectedX := 650.0
	expectedW := 100.0
	if math.Abs(rect.X-expectedX) > 0.001 {
		t.Errorf("rect.X = %f, expected %f", rect.X, expectedX)
	}
	if math.Abs(rect.Width-expectedW) > 0.001 {
		t.Errorf("rect.Width = %f, expected %f", rect.Width, expectedW)
	}
}

