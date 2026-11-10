package main

import (
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func TestPartClickUpdatesFormInDiagram(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_part_select", "data/stage.go", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected active diagram SVG object with at least one layer")
	}

	// Find P1 and P2 svg.Rects on the checked diagram ("sdqsdqsd" in data/stage.go)
	var p1Rect, p2Rect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		switch r.Name {
		case "P1-sdqsdqsd":
			p1Rect = r
		case "P2-sdqsdqsd":
			p2Rect = r
		}
	}
	if p1Rect == nil || p2Rect == nil {
		t.Fatalf("expected P1-sdqsdqsd and P2-sdqsdqsd rects, got p1=%v p2=%v", p1Rect, p2Rect)
	}

	assertFormShowsPart := func(expectedPartName, expectedPartShapeName string) {
		t.Helper()
		formStage := stack.Probe.GetFormStage()
		formGroups := formStage.GetInstancesSorted[*form.FormGroup]()
		if len(formGroups) != 1 {
			t.Fatalf("expected 1 FormGroup in probe formStage, got %d", len(formGroups))
		}
		fg := formGroups[0]
		if fg.TypeLabel != "Part" || fg.Label != expectedPartName {
			t.Fatalf("expected probe form to display Part %q, got TypeLabel=%q Label=%q", expectedPartName, fg.TypeLabel, fg.Label)
		}

		shapeFormStage := stager.GetShapeFormStage()
		shapeGroups := shapeFormStage.GetInstancesSorted[*form.FormGroup]()
		if len(shapeGroups) != 1 {
			t.Fatalf("expected 1 FormGroup in shapeFormStage, got %d", len(shapeGroups))
		}
		if shapeGroups[0].TypeLabel != "PartShape" || shapeGroups[0].Label != expectedPartShapeName {
			t.Fatalf("expected shape form to display PartShape %q, got TypeLabel=%q Label=%q", expectedPartShapeName, shapeGroups[0].TypeLabel, shapeGroups[0].Label)
		}
	}

	// 1. Direct click on P1 updates probe form to P1 and appearance form to P1-sdqsdqsd
	frontP1Select := *p1Rect
	frontP1Select.IsSelected = true
	p1Rect.OnAfterUpdate(nil, p1Rect, &frontP1Select)
	assertFormShowsPart("P1", "P1-sdqsdqsd")

	// 2. Drag P2 to a new position (triggers CommitWithSuspendedCallbacks without rebuilding SVG)
	//    and then click P2 -> probe form must switch from P1 to P2!
	frontP2Move := *p2Rect
	frontP2Move.X += 45
	frontP2Move.Y -= 30
	p2Rect.OnAfterUpdate(nil, p2Rect, &frontP2Move)

	frontP2Select := frontP2Move
	frontP2Select.IsSelected = true
	p2Rect.OnAfterUpdate(nil, p2Rect, &frontP2Select)
	assertFormShowsPart("P2", "P2-sdqsdqsd")
}
