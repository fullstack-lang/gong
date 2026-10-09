package main

import (
	"slices"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestTaskTaskGroupLink(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_task_taskgroup_link", "", "", true, false, false)
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

	tg1 := (&models.TaskGroup{Name: "Backend"}).Stage(stage)
	tg2 := (&models.TaskGroup{Name: "Frontend"}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg1, tg2}

	// Task 1: linked via tg1.Tasks
	task1 := (&models.Task{
		Name:     "DB Migration",
		Start:    time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		End:      time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC),
		IsAllDay: true,
	}).Stage(stage)
	tg1.Tasks = []*models.Task{task1}

	// Task 2: linked via task2.TaskGroups
	task2 := (&models.Task{
		Name:       "UI Components",
		Start:      time.Date(2026, 4, 3, 0, 0, 0, 0, time.UTC),
		End:        time.Date(2026, 4, 10, 0, 0, 0, 0, time.UTC),
		IsAllDay:   true,
		TaskGroups: []*models.TaskGroup{tg2},
	}).Stage(stage)

	// Task 3: linked via both tg1.Tasks and task3.TaskGroups
	task3 := (&models.Task{
		Name:       "API Endpoints",
		Start:      time.Date(2026, 4, 2, 0, 0, 0, 0, time.UTC),
		End:        time.Date(2026, 4, 7, 0, 0, 0, 0, time.UTC),
		IsAllDay:   true,
		TaskGroups: []*models.TaskGroup{tg1},
	}).Stage(stage)
	tg1.Tasks = append(tg1.Tasks, task3)

	lib.RootTasks = []*models.Task{task1, task2, task3}

	for d := range *stage.GetInstancesSet[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:          "ProjectGantt",
		IsTimeDiagram: true,
		IsChecked:     true,
		IsEditable_:   true,
		TextHeight:    15.0,
		LaneHeight:    85.0,
		YTopMargin:    40.0,
		DateYOffset:   15.0,
		XLeftLanes:    240.0,
		XRightMargin:  1250.0,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs1 := (&models.TaskGroupShape{Name: "ProjectGantt-Backend", TaskGroup: tg1}).Stage(stage)
	tgs2 := (&models.TaskGroupShape{Name: "ProjectGantt-Frontend", TaskGroup: tg2}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs1, tgs2}

	ts1 := (&models.TaskShape{Name: "ProjectGantt-DB Migration", Task: task1}).Stage(stage)
	ts2 := (&models.TaskShape{Name: "ProjectGantt-UI Components", Task: task2}).Stage(stage)
	ts3 := (&models.TaskShape{Name: "ProjectGantt-API Endpoints", Task: task3}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts1, ts2, ts3}

	stage.Commit()

	// 1. Verify GetTasks on TaskGroups
	tg1Tasks := tg1.GetTasks(stage)
	if len(tg1Tasks) != 2 {
		t.Fatalf("expected 2 tasks in tg1, got %d", len(tg1Tasks))
	}
	if !slices.Contains(tg1Tasks, task1) || !slices.Contains(tg1Tasks, task3) {
		t.Errorf("tg1Tasks should contain task1 and task3, got %v", tg1Tasks)
	}

	tg2Tasks := tg2.GetTasks(stage)
	if len(tg2Tasks) != 1 {
		t.Fatalf("expected 1 task in tg2, got %d", len(tg2Tasks))
	}
	if !slices.Contains(tg2Tasks, task2) {
		t.Errorf("tg2Tasks should contain task2, got %v", tg2Tasks)
	}

	// 2. Verify Diagram dates include task2 (which ends April 10 + 1 day = April 11)
	expectedEnd := time.Date(2026, 4, 11, 0, 0, 0, 0, time.UTC)
	if !diag.ComputedEnd.Equal(expectedEnd) {
		t.Errorf("expected diagram ComputedEnd %v, got %v", expectedEnd, diag.ComputedEnd)
	}

	// 3. Verify SVG contains rectangles for all 3 tasks in their lanes
	svgObj := stager.GetSvgObject()
	if svgObj == nil {
		t.Fatalf("expected svgObject to be generated")
	}

	foundTask1 := false
	foundTask2 := false
	foundTask3 := false
	for _, layer := range svgObj.Layers {
		for _, rect := range layer.Rects {
			if rect.Name == task1.Name {
				foundTask1 = true
			}
			if rect.Name == task2.Name {
				foundTask2 = true
			}
			if rect.Name == task3.Name {
				foundTask3 = true
			}
		}
	}
	if !foundTask1 {
		t.Errorf("task1 (%s) not found in SVG rects", task1.Name)
	}
	if !foundTask2 {
		t.Errorf("task2 (%s) not found in SVG rects", task2.Name)
	}
	if !foundTask3 {
		t.Errorf("task3 (%s) not found in SVG rects", task3.Name)
	}

	treeStage := stager.GetTreeStage()
	if treeStage == nil {
		t.Fatalf("expected treeStage to be present")
	}
	frontendChildFound := false
	for node := range *treeStage.GetInstancesSet[*tree_models.Node]() {
		if node.Name == "Frontend" {
			for _, child := range node.Children {
				if child.Name == "UI Components" {
					frontendChildFound = true
				}
			}
		}
	}
	if !frontendChildFound {
		t.Errorf("expected 'UI Components' to be a child of 'Frontend' in tree")
	}
}
