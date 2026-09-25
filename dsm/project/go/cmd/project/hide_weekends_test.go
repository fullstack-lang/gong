package main

import (
	"math"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
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
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

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

func TestTimeDiagramTaskHandlesAndMove(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_handles", "", "", true, false, false)
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

	// Task from Monday Jan 5 to Tuesday Jan 6 (2 days, all day)
	task := (&models.Task{
		Name:                          "Task1",
		Start:                         time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
		End:                           time.Date(2026, 1, 6, 0, 0, 0, 0, time.UTC),
		IsAllDay:                      true,
		DurationDays:                  2.0,
		IsEndDateComputedFromDuration: true,
	}).Stage(stage)
	lib.RootTasks = []*models.Task{task}

	tg := (&models.TaskGroup{
		Name:  "Group1",
		Tasks: []*models.Task{task},
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                      "GanttDiag",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		XLeftLanes:                200.0,
		XRightMargin:              1200.0, // 1000px total
		LaneHeight:                60.0,
		RatioBarToLaneHeight:      0.8,
		HideWeekendsPeriod:        true,
		UseManualStartAndEndDates: true,
		ManualStart:               time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),  // Monday
		ManualEnd:                 time.Date(2026, 1, 19, 0, 0, 0, 0, time.UTC), // Monday 2 weeks later (10 work days)
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{
		Name:      "GanttDiag-Group1",
		TaskGroup: tg,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{
		Name: "GanttDiag-Task1",
		Task: task,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	stage.Commit()

	rect := diag.GetTaskRect(task)
	if rect == nil {
		t.Fatal("Task rect was not generated in SVG")
	}

	// 10 work days, width = 1000 => 100px per work day
	// Jan 5 00:00 is x = 200
	// Jan 7 00:00 (end of Tuesday) is x = 400 => width = 200
	if math.Abs(rect.X-200.0) > 0.001 || math.Abs(rect.Width-200.0) > 0.001 {
		t.Fatalf("Initial rect: got X=%f, Width=%f, expected X=200, Width=200", rect.X, rect.Width)
	}

	// Test 1: Move right handle (extend to Wednesday Jan 7 end of day, +100px width => 300px)
	rect.OnResize(rect.X, rect.Y, 300.0, rect.Height)

	expectedEnd1 := time.Date(2026, 1, 7, 0, 0, 0, 0, time.UTC)
	if !task.End.Equal(expectedEnd1) {
		t.Errorf("After right handle resize: task.End = %v, expected %v", task.End, expectedEnd1)
	}
	if task.DurationDays != 3.0 {
		t.Errorf("After right handle resize: task.DurationDays = %f, expected 3.0", task.DurationDays)
	}

	// Verify the SVG rect updated
	rect = diag.GetTaskRect(task)
	if math.Abs(rect.Width-300.0) > 0.001 {
		t.Errorf("After right handle resize: rect.Width = %f, expected 300.0", rect.Width)
	}

	// Test 2: Move left handle (start at Tuesday Jan 6, x = 300px, width = 200px)
	rect.OnResize(300.0, rect.Y, 200.0, rect.Height)

	expectedStart2 := time.Date(2026, 1, 6, 0, 0, 0, 0, time.UTC)
	if !task.Start.Equal(expectedStart2) {
		t.Errorf("After left handle resize: task.Start = %v, expected %v", task.Start, expectedStart2)
	}
	if !task.End.Equal(expectedEnd1) {
		t.Errorf("After left handle resize: task.End = %v, expected %v", task.End, expectedEnd1)
	}
	if task.DurationDays != 2.0 {
		t.Errorf("After left handle resize: task.DurationDays = %f, expected 2.0", task.DurationDays)
	}

	// Test 3: Move task horizontally (shift by 1 work day to Wednesday Jan 7, x = 400px)
	rect = diag.GetTaskRect(task)
	rect.OnMove(400.0, rect.Y)

	expectedStart3 := time.Date(2026, 1, 7, 0, 0, 0, 0, time.UTC)
	expectedEnd3 := time.Date(2026, 1, 8, 0, 0, 0, 0, time.UTC)
	if !task.Start.Equal(expectedStart3) {
		t.Errorf("After OnMove: task.Start = %v, expected %v", task.Start, expectedStart3)
	}
	if !task.End.Equal(expectedEnd3) {
		t.Errorf("After OnMove: task.End = %v, expected %v", task.End, expectedEnd3)
	}
	if task.DurationDays != 2.0 {
		t.Errorf("After OnMove: task.DurationDays = %f, expected 2.0", task.DurationDays)
	}

	// Test 4: Verify task form was updated in probeForm
	formStage := stack.Probe.GetFormStage()
	if len(formStage.FormGroups) == 0 {
		t.Fatal("Expected task form in probe formStage, but none found")
	}
	var fg *form.FormGroup
	for formGroup := range formStage.FormGroups {
		fg = formGroup
		break
	}
	if fg.TypeLabel != "Task" || fg.Label != task.Name {
		t.Errorf("FormGroup TypeLabel=%s, Label=%s, expected Task / %s", fg.TypeLabel, fg.Label, task.Name)
	}
	for _, div := range fg.FormDivs {
		for _, field := range div.FormFields {
			if field.Name == "StartDate" && field.FormFieldDate != nil {
				if !field.FormFieldDate.Value.Equal(expectedStart3) {
					t.Errorf("Form Start = %v, expected %v", field.FormFieldDate.Value, expectedStart3)
				}
			}
			if field.Name == "EndDate" && field.FormFieldDate != nil {
				if !field.FormFieldDate.Value.Equal(expectedEnd3) {
					t.Errorf("Form End = %v, expected %v", field.FormFieldDate.Value, expectedEnd3)
				}
			}
		}
	}
}

func TestTimeDiagramSizeTakesIntoAccountTextOnRight(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_diagram_size", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)
	_ = stager
	stage := stack.Stage

	lib := (&models.Library{
		Name:              "RootLib",
		IsRootLibrary:     true,
		NbPixPerCharacter: 8.0,
	}).Stage(stage)

	// Diagram spanning 2 weeks: Monday Jan 5 to Monday Jan 19
	start := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 19, 0, 0, 0, 0, time.UTC)

	// Milestone on the last day: Monday Jan 19 (which is at XRightMargin = 1200)
	milestoneName := "General Availability Milestone" // 30 chars -> 240px
	milestone := (&models.Task{
		Name:        milestoneName,
		Start:       end,
		End:         end,
		IsMilestone: true,
	}).Stage(stage)
	lib.RootTasks = []*models.Task{milestone}

	tg := (&models.TaskGroup{
		Name:  "MilestonesGroup",
		Tasks: []*models.Task{milestone},
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                      "GanttDiag",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		XLeftText:                 15.0,
		XLeftLanes:                200.0,
		XRightMargin:              1200.0,
		LaneHeight:                60.0,
		RatioBarToLaneHeight:      0.8,
		HideWeekendsPeriod:        true,
		UseManualStartAndEndDates: true,
		ManualStart:               start,
		ManualEnd:                 end,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{
		Name:      "GanttDiag-Group",
		TaskGroup: tg,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{
		Name: "GanttDiag-Milestone",
		Task: milestone,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	stage.Commit()

	// lineX for milestone is 1200.0
	// diamondWidth is 18.0 => dummyX is 1209.0
	// milestoneText offset is 15.0 => text starts at 1224.0
	// textWidth is 30 * 8 = 240.0 => text ends at 1464.0
	// margin is 100.0 => expected width is at least 1464 + 100 = 1564.0
	expectedMinRight := 1209.0 + 15.0 + float64(len(milestoneName))*8.0 // 1464.0
	expectedMinWidth := expectedMinRight + 100.0                       // 1564.0

	if diag.Width < expectedMinWidth {
		t.Errorf("diag.Width = %f, expected at least %f to accommodate right-side milestone text", diag.Width, expectedMinWidth)
	}

	// Also verify that SVG object has the updated OverriddenWidth
	svgObj := stager.GetSvgObject()
	if svgObj == nil {
		t.Fatal("SVG object was nil")
	}
	if svgObj.OverriddenWidth != diag.Width {
		t.Errorf("svgObj.OverriddenWidth = %f, expected %f", svgObj.OverriddenWidth, diag.Width)
	}
}

func TestTaskDefaultDependencyTypeFinishToStart(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_default_dependency_type", "", "", true, false, false)
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
		Name: "TaskWithoutDependencyType",
	}).Stage(stage)
	lib.RootTasks = []*models.Task{task}

	stage.Commit()

	if task.DependencyType != models.FINISH_TO_START {
		t.Errorf("task.DependencyType = %q, expected %q", task.DependencyType, models.FINISH_TO_START)
	}
}


