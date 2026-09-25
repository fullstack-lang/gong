package main

import (
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func TestEnforceDiagramSizeTimeDiagram(t *testing.T) {
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
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	tg1 := (&models.TaskGroup{Name: "TG1"}).Stage(stage)
	tg2 := (&models.TaskGroup{Name: "TG2"}).Stage(stage)
	tg3 := (&models.TaskGroup{Name: "TG3"}).Stage(stage)
	tg4 := (&models.TaskGroup{Name: "TG4"}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg1, tg2, tg3, tg4}

	diag := (&models.Diagram{
		Name:                      "Gantt",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		IsEditable_:               true,
		TextHeight:                15.0,
		LaneHeight:                85.0,
		YTopMargin:                40.0,
		DateYOffset:               15.0,
		XLeftLanes:                240.0,
		XRightMargin:              1250.0,
		UseManualStartAndEndDates: true,
		ManualStart:               time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
		ManualEnd:                 time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs1 := (&models.TaskGroupShape{Name: "Gantt-TG1", TaskGroup: tg1}).Stage(stage)
	tgs2 := (&models.TaskGroupShape{Name: "Gantt-TG2", TaskGroup: tg2}).Stage(stage)
	tgs3 := (&models.TaskGroupShape{Name: "Gantt-TG3", TaskGroup: tg3}).Stage(stage)
	tgs4 := (&models.TaskGroupShape{Name: "Gantt-TG4", TaskGroup: tg4}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs1, tgs2, tgs3, tgs4}

	// Commit will trigger enforceSemantic and enforceDiagramSize
	stage.Commit()

	// yTimeLine = 40.0 + 85.0 * 4 = 380.0
	// dateMargin = DateYOffset(15.0) + TextHeight(15.0) + 5.0 = 35.0
	// expectedHeight = 380.0 + 35.0 = 415.0
	expectedHeight := 415.0
	if diag.Height != expectedHeight {
		t.Errorf("expected diagram height %f, got %f (old buggy value was 495.0)", expectedHeight, diag.Height)
	}

	// Now add a hidden ProductShape placed far down at Y=600.
	// It should NOT increase the diagram height because it is hidden.
	prod := (&models.Product{Name: "HiddenProd"}).Stage(stage)
	lib.RootProducts = []*models.Product{prod}
	prodShape := (&models.ProductShape{
		Name:     "Gantt-HiddenProd",
		Product:  prod,
		X:        50.0,
		Y:        600.0,
		Width:    250.0,
		Height:   70.0,
		IsHidden: true,
	}).Stage(stage)
	diag.Product_Shapes = []*models.ProductShape{prodShape}

	stage.Commit()

	if diag.Height != expectedHeight {
		t.Errorf("hidden shape at Y=600 should not enlarge diagram; expected %f, got %f", expectedHeight, diag.Height)
	}

	// Now unhide the shape: it should now enlarge the diagram height
	prodShape.IsHidden = false
	stage.Commit()

	// With visible shape at Y=600, Height=70, margin=100 => 600 + 70 + 100 = 770.0
	expectedExpandedHeight := 770.0
	if diag.Height != expectedExpandedHeight {
		t.Errorf("visible shape at Y=600 should enlarge diagram; expected %f, got %f", expectedExpandedHeight, diag.Height)
	}
}
