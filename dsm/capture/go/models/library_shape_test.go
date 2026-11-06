package models

import (
	"os"
	"testing"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func TestLibraryShapeDoubleClickNonEditable(t *testing.T) {
	stage := NewStage("test")
	svgStage := svg.NewStage("svg")

	rootLib := (&Library{Name: "Root Library", IsRootLibrary: true}).Stage(stage)
	rootDiag := (&Diagram{Name: "Root Diagram", IsChecked: true, IsEditable_: false}).Stage(stage)
	rootLib.Diagrams = append(rootLib.Diagrams, rootDiag)

	subLib := (&Library{Name: "Sub Library"}).Stage(stage)
	subDiag1 := (&Diagram{Name: "Sub Diagram 1", IsChecked: false, IsEditable_: false}).Stage(stage)
	subDiag2 := (&Diagram{Name: "Sub Diagram 2", IsChecked: false, IsEditable_: false}).Stage(stage)
	subLib.Diagrams = append(subLib.Diagrams, subDiag1, subDiag2)
	rootLib.SubLibraries = append(rootLib.SubLibraries, subLib)

	libShape := (&LibraryShape{
		Name:    "Sub Library-Root Diagram",
		Library: subLib,
		RectShape: RectShape{
			X:      100,
			Y:      100,
			Width:  200,
			Height: 80,
		},
	}).Stage(stage)
	rootDiag.Library_Shapes = append(rootDiag.Library_Shapes, libShape)

	stager := &Stager{
		stage:    stage,
		svgStage: svgStage,
	}

	svgObj := stager.generateSvgObject(rootDiag)
	if len(svgObj.Layers) != 1 || len(svgObj.Layers[0].Rects) != 1 {
		t.Fatalf("expected 1 rect in layer, got %d", len(svgObj.Layers[0].Rects))
	}
	rect := svgObj.Layers[0].Rects[0]

	// First click inside rect bounds
	frontRect := *rect
	frontRect.MouseX = 150
	frontRect.MouseY = 120
	rect.OnAfterUpdate(svgStage, rect, &frontRect)

	if !rootDiag.IsChecked || subDiag1.IsChecked {
		t.Fatalf("single click should not switch diagram")
	}

	// Second click inside rect bounds within 500ms -> should open subDiag1
	rect.OnAfterUpdate(svgStage, rect, &frontRect)

	if rootDiag.IsChecked {
		t.Errorf("expected rootDiag.IsChecked to be false after double click")
	}
	if !subDiag1.IsChecked {
		t.Errorf("expected subDiag1.IsChecked to be true after double click")
	}
	if subDiag2.IsChecked {
		t.Errorf("expected subDiag2.IsChecked to be false")
	}
}

func TestLibraryShapeDoubleClickEditable(t *testing.T) {
	stage := NewStage("test")
	svgStage := svg.NewStage("svg")

	rootLib := (&Library{Name: "Root Library", IsRootLibrary: true}).Stage(stage)
	rootDiag := (&Diagram{Name: "Root Diagram", IsChecked: true, IsEditable_: true}).Stage(stage)
	rootLib.Diagrams = append(rootLib.Diagrams, rootDiag)

	subLib := (&Library{Name: "Sub Library"}).Stage(stage)
	subDiag1 := (&Diagram{Name: "Sub Diagram 1", IsChecked: false, IsEditable_: true}).Stage(stage)
	subLib.Diagrams = append(subLib.Diagrams, subDiag1)
	rootLib.SubLibraries = append(rootLib.SubLibraries, subLib)

	libShape := (&LibraryShape{
		Name:    "Sub Library-Root Diagram",
		Library: subLib,
		RectShape: RectShape{
			X:      100,
			Y:      100,
			Width:  200,
			Height: 80,
		},
	}).Stage(stage)
	rootDiag.Library_Shapes = append(rootDiag.Library_Shapes, libShape)

	stager := &Stager{
		stage:    stage,
		svgStage: svgStage,
	}

	svgObj := stager.generateSvgObject(rootDiag)
	rect := svgObj.Layers[0].Rects[0]

	// Click 1 in editable mode: frontend selects rect (1 update with IsSelected=true)
	frontRect := *rect
	frontRect.IsSelected = true
	rect.OnUpdate(&frontRect)

	if !rootDiag.IsChecked || subDiag1.IsChecked {
		t.Fatalf("single click should not switch diagram")
	}

	// Click 2 in editable mode within 500ms: frontend calls unselectAllRects() + unselectRect() (2 updates with IsSelected=false)
	frontRect.IsSelected = false
	rect.OnUpdate(&frontRect)
	if !rootDiag.IsChecked || subDiag1.IsChecked {
		t.Fatalf("single unselect (e.g. clicking background) should not switch diagram")
	}
	rect.OnUpdate(&frontRect)

	if rootDiag.IsChecked || !subDiag1.IsChecked {
		t.Errorf("expected double click in editable mode to switch to subDiag1")
	}
}

func TestLibraryShapeDoubleClickEditableWhenAlreadySelected(t *testing.T) {
	stage := NewStage("test")
	svgStage := svg.NewStage("svg")

	rootLib := (&Library{Name: "Root Library", IsRootLibrary: true}).Stage(stage)
	rootDiag := (&Diagram{Name: "Root Diagram", IsChecked: true, IsEditable_: true}).Stage(stage)
	rootLib.Diagrams = append(rootLib.Diagrams, rootDiag)

	subLib := (&Library{Name: "Sub Library"}).Stage(stage)
	subDiag1 := (&Diagram{Name: "Sub Diagram 1", IsChecked: false, IsEditable_: true}).Stage(stage)
	subLib.Diagrams = append(subLib.Diagrams, subDiag1)
	rootLib.SubLibraries = append(rootLib.SubLibraries, subLib)

	libShape := (&LibraryShape{
		Name:    "Sub Library-Root Diagram",
		Library: subLib,
		RectShape: RectShape{
			X:      100,
			Y:      100,
			Width:  200,
			Height: 80,
		},
	}).Stage(stage)
	rootDiag.Library_Shapes = append(rootDiag.Library_Shapes, libShape)

	stager := &Stager{
		stage:    stage,
		svgStage: svgStage,
	}

	svgObj := stager.generateSvgObject(rootDiag)
	rect := svgObj.Layers[0].Rects[0]

	// Initial single click to select the shape
	frontRect := *rect
	frontRect.IsSelected = true
	rect.OnUpdate(&frontRect)

	// Wait > 500ms so the next two clicks form a fresh double-click on an already-selected shape
	stager.lastClickedLibraryShapeTime = time.Now().Add(-600 * time.Millisecond)

	// First click of the double-click: unselects the already-selected rect (2 updates with IsSelected=false)
	frontRect.IsSelected = false
	rect.OnUpdate(&frontRect)
	rect.OnUpdate(&frontRect)

	if !rootDiag.IsChecked || subDiag1.IsChecked {
		t.Fatalf("first click of double-click on selected rect should only unselect it")
	}

	// Second click of the double-click: selects the rect again (1 update with IsSelected=true)
	frontRect.IsSelected = true
	rect.OnUpdate(&frontRect)

	if rootDiag.IsChecked || !subDiag1.IsChecked {
		t.Errorf("expected double click on already-selected rect to switch to subDiag1")
	}
}

func TestParse15288Data(t *testing.T) {
	data, err := os.ReadFile("../cmd/capture/data/15288.go")
	if err != nil {
		t.Fatalf("failed to read 15288.go: %v", err)
	}
	stage := NewStage("test15288")
	if err := ParseAstFromBytes(stage, data); err != nil {
		t.Fatalf("failed to parse 15288.go: %v", err)
	}
	libShapes := stage.GetInstancesSorted[*LibraryShape]()
	if len(libShapes) != 64 {
		t.Fatalf("expected 64 LibraryShapes in 15288.go, got %d", len(libShapes))
	}
	for _, s := range libShapes {
		if s.Library == nil {
			t.Errorf("LibraryShape %q has nil Library", s.Name)
		} else if len(s.Library.Diagrams) == 0 {
			t.Errorf("LibraryShape %q references Library %q with 0 Diagrams", s.Name, s.Library.Name)
		}
	}
}

