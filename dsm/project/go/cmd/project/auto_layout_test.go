package main

import (
	"testing"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func TestAutoLayoutOnSubItemsReorder(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test", "", "", true, false, false)
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

	pA := (&models.Product{
		Name: "A",
	}).Stage(stage)

	pA1 := (&models.Product{
		Name: "A.1",
	}).Stage(stage)

	pA2 := (&models.Product{
		Name: "A.2",
	}).Stage(stage)

	pA.SubProducts = []*models.Product{pA1, pA2}
	lib.RootProducts = []*models.Product{pA}

	diag := (&models.Diagram{
		Name:               "TestDiagram",
		IsInAutoLayoutMode: true,
		IsChecked:          true,
		IsEditable_:        true,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	pAShape := (&models.ProductShape{
		Name:    "Shape-A",
		Product: pA,
		Width:   250,
		Height:  70,
	}).Stage(stage)

	pA1Shape := (&models.ProductShape{
		Name:    "Shape-A1",
		Product: pA1,
		Width:   250,
		Height:  70,
	}).Stage(stage)

	pA2Shape := (&models.ProductShape{
		Name:    "Shape-A2",
		Product: pA2,
		Width:   250,
		Height:  70,
	}).Stage(stage)

	diag.Product_Shapes = []*models.ProductShape{pAShape, pA1Shape, pA2Shape}

	comp1 := (&models.ProductCompositionShape{
		Name:    "Comp-A-A1",
		Product: pA1,
	}).Stage(stage)
	comp2 := (&models.ProductCompositionShape{
		Name:    "Comp-A-A2",
		Product: pA2,
	}).Stage(stage)
	diag.ProductComposition_Shapes = []*models.ProductCompositionShape{comp1, comp2}

	// First pass commit (triggers enforceSemantic including enforceAutoLayout)
	stage.Commit()

	initialA1X, initialA1Y := pA1Shape.X, pA1Shape.Y
	initialA2X, initialA2Y := pA2Shape.X, pA2Shape.Y

	t.Logf("Initial: A1=(%f, %f), A2=(%f, %f)", initialA1X, initialA1Y, initialA2X, initialA2Y)

	if initialA1X == initialA2X && initialA1Y == initialA2Y {
		t.Fatalf("Shapes should have distinct positions")
	}

	// Now reorder sub items: swap A1 and A2
	pA.SubProducts = []*models.Product{pA2, pA1}

	// Re-commit stage (as happens on probe form save or code update)
	stage.Commit()

	newA1X, newA1Y := pA1Shape.X, pA1Shape.Y
	newA2X, newA2Y := pA2Shape.X, pA2Shape.Y

	t.Logf("Reordered: A1=(%f, %f), A2=(%f, %f)", newA1X, newA1Y, newA2X, newA2Y)

	// Since A2 is now the first child and A1 is the second child,
	// A2 should take A1's old position and A1 should take A2's old position!
	if newA2X != initialA1X || newA2Y != initialA1Y {
		t.Errorf("Expected A2 to take A1's initial position (%f, %f), got (%f, %f)",
			initialA1X, initialA1Y, newA2X, newA2Y)
	}
	if newA1X != initialA2X || newA1Y != initialA2Y {
		t.Errorf("Expected A1 to take A2's initial position (%f, %f), got (%f, %f)",
			initialA2X, initialA2Y, newA1X, newA1Y)
	}

	// Now disable auto layout mode
	diag.IsInAutoLayoutMode = false
	stage.Commit()

	// Swap back
	pA.SubProducts = []*models.Product{pA1, pA2}
	stage.Commit()

	// Since auto layout mode is false, positions should NOT change
	if pA1Shape.X != newA1X || pA1Shape.Y != newA1Y {
		t.Errorf("When IsInAutoLayoutMode is false, A1 position should not change")
	}
	if pA2Shape.X != newA2X || pA2Shape.Y != newA2Y {
		t.Errorf("When IsInAutoLayoutMode is false, A2 position should not change")
	}

	// Test Task SubTasks reordering
	diag.IsInAutoLayoutMode = true

	tT := (&models.Task{Name: "T"}).Stage(stage)
	tT1 := (&models.Task{Name: "T.1"}).Stage(stage)
	tT2 := (&models.Task{Name: "T.2"}).Stage(stage)
	tT.SubTasks = []*models.Task{tT1, tT2}
	lib.RootTasks = []*models.Task{tT}

	tTShape := (&models.TaskShape{Name: "Shape-T", Task: tT, Width: 250, Height: 70}).Stage(stage)
	tT1Shape := (&models.TaskShape{Name: "Shape-T1", Task: tT1, Width: 250, Height: 70}).Stage(stage)
	tT2Shape := (&models.TaskShape{Name: "Shape-T2", Task: tT2, Width: 250, Height: 70}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{tTShape, tT1Shape, tT2Shape}

	compT1 := (&models.TaskCompositionShape{Name: "Comp-T-T1", Task: tT1}).Stage(stage)
	compT2 := (&models.TaskCompositionShape{Name: "Comp-T-T2", Task: tT2}).Stage(stage)
	diag.TaskComposition_Shapes = []*models.TaskCompositionShape{compT1, compT2}

	stage.Commit()

	initialT1X := tT1Shape.X
	initialT2X := tT2Shape.X

	// Swap SubTasks
	tT.SubTasks = []*models.Task{tT2, tT1}
	stage.Commit()

	if tT2Shape.X != initialT1X || tT1Shape.X != initialT2X {
		t.Errorf("Expected T2 to take T1's initial X position and vice-versa")
	}
}

