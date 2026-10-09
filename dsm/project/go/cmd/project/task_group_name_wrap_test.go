package main

import (
	"strings"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func TestTaskGroupNameWrapping(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_taskgroup_name_wrap", "", "", true, false, false)
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
		NbPixPerCharacter: 9.0,
	}).Stage(stage)

	// Create TaskGroup with a long repeated name like in schedule.go
	longName := "Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 Task 1 "
	tg := (&models.TaskGroup{Name: longName}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	task := (&models.Task{
		Name:     "Some Task",
		Start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		End:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
		IsAllDay: true,
	}).Stage(stage)
	tg.Tasks = []*models.Task{task}
	lib.RootTasks = []*models.Task{task}

	for d := range *stage.GetInstancesSet[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:          "Gantt",
		IsTimeDiagram: true,
		IsChecked:     true,
		IsEditable_:   true,
		TextHeight:    15.0,
		LaneHeight:    85.0,
		YTopMargin:    40.0,
		DateYOffset:   15.0,
		XLeftText:     10.0,
		XLeftLanes:    200.0,
		XRightMargin:  1250.0,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{Name: "Gantt-" + longName, TaskGroup: tg}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{Name: "Gantt-Some Task", Task: task}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	stage.Commit()

	svgObj := stager.GetSvgObject()
	if svgObj == nil {
		t.Fatalf("expected svgObject to be generated")
	}

	// Find the lane rect for the TaskGroup
	var laneRect *svg.Rect
	for _, layer := range svgObj.Layers {
		for _, rect := range layer.Rects {
			if rect.Name == longName {
				laneRect = rect
				break
			}
		}
	}

	if laneRect == nil {
		t.Fatalf("lane rectangle for %s not found", longName)
	}

	if len(laneRect.RectAnchoredTexts) == 0 {
		t.Fatalf("expected laneRect to have RectAnchoredTexts")
	}

	laneText := laneRect.RectAnchoredTexts[0]
	if !strings.Contains(laneText.Content, "\n") {
		t.Errorf("expected laneText content to be wrapped with newlines, got %q", laneText.Content)
	}

	// Verify that each line length in characters does not exceed cutoff (20)
	lines := strings.Split(laneText.Content, "\n")
	for i, line := range lines {
		if len(line) > 20 {
			t.Errorf("line %d has length %d > cutoff 20: %q", i, len(line), line)
		}
	}

	// Verify X_Offset places text at XLeftText
	expectedXOffset := diag.XLeftText - diag.XLeftLanes
	if laneText.X_Offset != expectedXOffset {
		t.Errorf("expected X_Offset %f, got %f", expectedXOffset, laneText.X_Offset)
	}
}

func TestTaskGroupSVGSelection(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_taskgroup_svg_selection", "", "", true, false, false)
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

	tgName := "Engineering Tasks"
	tg := (&models.TaskGroup{Name: tgName}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	task := (&models.Task{
		Name:     "Some Task",
		Start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		End:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
		IsAllDay: true,
	}).Stage(stage)
	tg.Tasks = []*models.Task{task}
	lib.RootTasks = []*models.Task{task}

	for d := range *stage.GetInstancesSet[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:          "Gantt",
		IsTimeDiagram: true,
		IsChecked:     true,
		IsEditable_:   true,
		LaneHeight:    85.0,
		YTopMargin:    40.0,
		DateYOffset:   15.0,
		XLeftText:     10.0,
		XLeftLanes:    200.0,
		XRightMargin:  1250.0,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{Name: "Gantt-" + tgName, TaskGroup: tg}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{Name: "Gantt-Some Task", Task: task}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	stage.Commit()

	svgObj := stager.GetSvgObject()
	if svgObj == nil {
		t.Fatalf("expected svgObject to be generated")
	}

	var laneRect *svg.Rect
	var headerRect *svg.Rect
	for _, layer := range svgObj.Layers {
		for _, rect := range layer.Rects {
			if rect.Name == tgName {
				laneRect = rect
			}
			if rect.Name == tgName+" Header" {
				headerRect = rect
			}
		}
	}

	if laneRect == nil {
		t.Fatalf("lane rectangle for %s not found", tgName)
	}
	if !laneRect.IsSelectable {
		t.Errorf("expected laneRect.IsSelectable to be true")
	}
	if laneRect.OnSelect == nil {
		t.Fatalf("expected laneRect.OnSelect to be non-nil")
	}

	if headerRect == nil {
		t.Fatalf("header rectangle for %s not found", tgName)
	}
	if !headerRect.IsSelectable {
		t.Errorf("expected headerRect.IsSelectable to be true")
	}
	if headerRect.OnSelect == nil {
		t.Fatalf("expected headerRect.OnSelect to be non-nil")
	}

	// Test 1: clicking header rect updates the form
	headerRect.OnSelect()
	formStage := stack.Probe.GetFormStage()
	fgSet := *formStage.GetInstancesSet[*form.FormGroup]()
	if len(fgSet) == 0 {
		t.Fatalf("expected form stage to have a FormGroup after header select")
	}
	var fg *form.FormGroup
	for item := range fgSet {
		fg = item
		break
	}
	if fg.TypeLabel != "TaskGroup" {
		t.Errorf("expected TypeLabel TaskGroup, got %s", fg.TypeLabel)
	}
	if fg.Label != tgName {
		t.Errorf("expected Label %s, got %s", tgName, fg.Label)
	}

	// Test 2: clicking lane rect also updates the form
	laneRect.OnSelect()
	fgSet = *formStage.GetInstancesSet[*form.FormGroup]()
	if len(fgSet) == 0 {
		t.Fatalf("expected form stage to have a FormGroup after lane select")
	}
	for item := range fgSet {
		fg = item
		break
	}
	if fg.TypeLabel != "TaskGroup" {
		t.Errorf("expected TypeLabel TaskGroup, got %s", fg.TypeLabel)
	}
	if fg.Label != tgName {
		t.Errorf("expected Label %s, got %s", tgName, fg.Label)
	}
}
