package main

import (
	"strings"
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

func TestDataShapeSelectAndResizeOnDataFlow(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_data_select_resize", "data/stage.go", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected active diagram SVG object with at least one layer")
	}

	// Find the svg.Rect for Data "D1" on DataFlow `"Port 1" to "P1"`
	var d1Rect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "D1" {
			d1Rect = r
			break
		}
	}
	if d1Rect == nil {
		t.Fatalf("expected to find an svg.Rect for Data 'D1'")
	}
	if !d1Rect.CanHaveLeftHandle || !d1Rect.CanHaveRightHandle {
		t.Errorf("expected Data svg.Rect to have left and right resize handles enabled")
	}
	if d1Rect.CanHaveTopHandle || d1Rect.CanHaveBottomHandle {
		t.Errorf("expected Data svg.Rect to disable top and bottom resize handles")
	}
	if d1Rect.OnSelect == nil || d1Rect.OnResize == nil {
		t.Fatalf("expected Data svg.Rect to have OnSelect and OnResize callbacks")
	}

	// 1. Clicking Data rect updates probe form to Data "D1" and shape form to DataShape (with Width field)
	frontSelect := *d1Rect
	frontSelect.IsSelected = true
	d1Rect.OnAfterUpdate(nil, d1Rect, &frontSelect)

	formGroups := stack.Probe.GetFormStage().GetInstancesSorted[*form.FormGroup]()
	if len(formGroups) != 1 || formGroups[0].TypeLabel != "Data" || formGroups[0].Label != "D1" {
		t.Fatalf("expected probe form to display Data 'D1', got %+v", formGroups)
	}

	shapeGroups := stager.GetShapeFormStage().GetInstancesSorted[*form.FormGroup]()
	if len(shapeGroups) != 1 || shapeGroups[0].TypeLabel != "DataShape" {
		t.Fatalf("expected shape appearance form to display DataShape, got %+v", shapeGroups)
	}
	var hasWidthField bool
	for _, fd := range shapeGroups[0].FormDivs {
		if fd.Name == "Width" {
			hasWidthField = true
			break
		}
	}
	if !hasWidthField {
		t.Errorf("expected DataShape appearance form to expose 'Width' field")
	}

	// 2. Rename Data to a multi-word name and resize narrower -> wraps across lines and increases height
	var d1Data *models.Data
	for _, d := range stack.Stage.GetInstancesSorted[*models.Data]() {
		if d.Name == "D1" {
			d1Data = d
			break
		}
	}
	if d1Data == nil {
		t.Fatalf("expected Data 'D1' in stage")
	}
	d1Data.Name = "Flight Plan Request Message"
	stack.Stage.Commit()

	svgObj = stager.GetSvgObject()
	var renamedRect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "Flight Plan Request Message" {
			renamedRect = r
			break
		}
	}
	if renamedRect == nil {
		t.Fatalf("expected svg.Rect for renamed Data")
	}
	singleLineHeight := renamedRect.Height

	// Resize narrower via OnAfterUpdate (simulating dragging right/left handle to Width = 90)
	frontResize := *renamedRect
	frontResize.Width = 90.0
	renamedRect.OnAfterUpdate(nil, renamedRect, &frontResize)

	svgObjAfter := stager.GetSvgObject()
	var resizedRect *svg.Rect
	for _, r := range svgObjAfter.Layers[0].Rects {
		if r.Name == "Flight Plan Request Message" && r.Width == 90.0 {
			resizedRect = r
			break
		}
	}
	if resizedRect == nil {
		t.Fatalf("expected resized svg.Rect with Width=90.0")
	}
	if len(resizedRect.RectAnchoredTexts) == 0 {
		t.Fatalf("expected RectAnchoredText inside Data svg.Rect")
	}
	if !strings.Contains(resizedRect.RectAnchoredTexts[0].Content, "\n") {
		t.Errorf("expected text to wrap across multiple lines at Width=90.0, got %q", resizedRect.RectAnchoredTexts[0].Content)
	}
	if resizedRect.Height <= singleLineHeight {
		t.Errorf("expected wrapped Data rect height (%v) > single-line height (%v)", resizedRect.Height, singleLineHeight)
	}

	// 3. Resizing below minimum width clamps to 30.0
	frontTooSmall := *resizedRect
	frontTooSmall.Width = 10.0
	resizedRect.OnAfterUpdate(nil, resizedRect, &frontTooSmall)

	svgObjClamped := stager.GetSvgObject()
	var clampedRect *svg.Rect
	for _, r := range svgObjClamped.Layers[0].Rects {
		if r.Name == "Flight Plan Request Message" {
			clampedRect = r
			break
		}
	}
	if clampedRect == nil || clampedRect.Width != 30.0 {
		t.Fatalf("expected Data svg.Rect width to clamp to 30.0, got %v", clampedRect)
	}
}

