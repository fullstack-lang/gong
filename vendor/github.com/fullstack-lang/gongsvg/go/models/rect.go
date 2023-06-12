package models

import "log"

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool // alllow selected
	IsSelected   bool

	CanHaveLeftHandle bool
	HasLeftHandle     bool

	CanHaveRightHandle bool
	HasRightHandle     bool

	CanHaveTopHandle bool
	HasTopHandle     bool

	CanHaveBottomHandle bool
	HasBottomHandle     bool

	CanMoveHorizontaly bool
	CanMoveVerticaly   bool

	RectAnchoredTexts []*RectAnchoredText
	RectAnchoredRects []*RectAnchoredRect

	Impl RectImplInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *StageStruct, _, frontRect *Rect) {

	log.Println("Rect, OnAfterUpdate", rect.Name)

	if rect.Impl != nil {
		rect.Impl.RectUpdated(frontRect)
	}
}
