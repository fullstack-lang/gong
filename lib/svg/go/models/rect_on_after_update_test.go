package models

import "testing"

func TestRectOnAfterUpdateSelectAfterMove(t *testing.T) {
	stage := NewStage("svg_test")
	rect := (&Rect{
		Name:   "P1",
		X:      100,
		Y:      150,
		Width:  250,
		Height: 70,
	}).Stage(stage)

	selectCount := 0
	moveCount := 0
	var movedX, movedY float64

	rect.OnSelect = func() {
		selectCount++
	}
	rect.OnMove = func(x, y float64) {
		moveCount++
		movedX = x
		movedY = y
	}

	// 1. Move rect without regenerating stage (mirrors CommitWithSuspendedCallbacks on move)
	frontMove := *rect
	frontMove.X = 220
	frontMove.Y = 340
	rect.OnAfterUpdate(stage, rect, &frontMove)

	if moveCount != 1 || movedX != 220 || movedY != 340 {
		t.Fatalf("expected 1 OnMove to (220, 340), got count=%d (%v, %v)", moveCount, movedX, movedY)
	}
	if selectCount != 0 {
		t.Fatalf("expected 0 OnSelect during move, got %d", selectCount)
	}

	// 2. Click moved rect to select it at its new position
	frontSelect := frontMove
	frontSelect.IsSelected = true
	rect.OnAfterUpdate(stage, rect, &frontSelect)

	if selectCount != 1 {
		t.Fatalf("expected OnSelect to be called once after moving and clicking rect, got %d (moveCount=%d)", selectCount, moveCount)
	}
	if moveCount != 1 {
		t.Fatalf("expected OnMove not to be called on click after move, got %d", moveCount)
	}

	// 3. Click again to unselect, then click to select again
	frontUnselect := frontSelect
	frontUnselect.IsSelected = false
	rect.OnAfterUpdate(stage, rect, &frontUnselect)

	if selectCount != 1 {
		t.Fatalf("expected OnSelect not to fire on unselect, got %d", selectCount)
	}

	rect.OnAfterUpdate(stage, rect, &frontSelect)
	if selectCount != 2 {
		t.Fatalf("expected OnSelect to fire on re-selection, got %d", selectCount)
	}
}

func TestLinkOnAfterUpdateSelectAfterChange(t *testing.T) {
	stage := NewStage("svg_link_test")
	link := (&Link{
		Name:              "L1",
		StartRatio:        0.5,
		EndRatio:          0.5,
		StartOrientation:  ORIENTATION_HORIZONTAL,
		EndOrientation:    ORIENTATION_HORIZONTAL,
		CornerOffsetRatio: 1.5,
	}).Stage(stage)

	selectCount := 0
	changeCount := 0
	link.OnSelect = func() { selectCount++ }
	link.OnChange = func(_ *Link) { changeCount++ }

	frontChange := *link
	frontChange.StartRatio = 0.75
	link.OnAfterUpdate(stage, link, &frontChange)

	if changeCount != 1 || selectCount != 0 {
		t.Fatalf("expected 1 OnChange and 0 OnSelect, got change=%d select=%d", changeCount, selectCount)
	}

	// Clicking link after change should fire OnSelect, not OnChange again
	link.OnAfterUpdate(stage, link, &frontChange)
	if selectCount != 1 || changeCount != 1 {
		t.Fatalf("expected 1 OnSelect and 1 OnChange, got select=%d change=%d", selectCount, changeCount)
	}
}
