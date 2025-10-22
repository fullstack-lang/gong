package models

import (
	"fmt"
	"strings"
)

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation
	ShapeConditions

	Animations []*Animate

	IsSelectable bool // alllow selected
	IsSelected   bool

	CanHaveLeftHandle bool
	HasLeftHandle     bool

	CanHaveRightHandle bool
	HasRightHandle     bool

	CanHaveTopHandle bool
	HasTopHandle     bool

	IsScalingProportionally bool

	CanHaveBottomHandle bool
	HasBottomHandle     bool

	CanMoveHorizontaly bool
	CanMoveVerticaly   bool

	RectAnchoredTexts []*RectAnchoredText
	RectAnchoredRects []*RectAnchoredRect
	RectAnchoredPaths []*RectAnchoredPath

	ChangeColorWhenHovered bool
	ColorWhenHovered       string
	OriginalColor          string
	FillOpacityWhenHovered float64
	OriginalFillOpacity    float64

	HasToolTip      bool
	ToolTipText     string
	ToolTipPosition ToolTipPositionEnum

	Impl RectImplInterface

	ImplWithMouseEvent RectImplWithMouseEventInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *Stage, _, frontRect *Rect) {

	if rect.Impl != nil {
		rect.Impl.RectUpdated(frontRect)
	}
}

func (rect *Rect) OnAfterUpdateWithMouseEvent(stage *Stage, frontRect *Rect, mouseEvent *Gong__MouseEvent) {

	if rect.ImplWithMouseEvent != nil {
		rect.ImplWithMouseEvent.RectUpdatedWithMouseEvent(frontRect, mouseEvent)
	}
}

func (rect *Rect) WriteSVG(sb *strings.Builder) (maxX, maxY float64) {

	sb.WriteString(
		fmt.Sprintf(
			`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
			formatFloat(rect.X),
			formatFloat(rect.Y),
			formatFloat(rect.Width),
			formatFloat(rect.Height),
			formatFloat(rect.RX),
			formatFloat(rect.RX)))

	maxX = rect.X + rect.Width
	maxY = rect.Y + rect.Height
	maxX += rect.Presentation.StrokeWidth / 2
	maxY += rect.Presentation.StrokeWidth / 2

	rect.Presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

	for _, rectAnchoredText := range rect.RectAnchoredTexts {
		x, y := getRectAnchorPoint(rect, rectAnchoredText.RectAnchorType)
		rectAnchoredText.WriteSVG(sb, x, y)

	}

	for _, rectAnchoredRect := range rect.RectAnchoredRects {
		x, y := getRectAnchorPoint(rect, rectAnchoredRect.RectAnchorType)
		maxX_, maxY_ := rectAnchoredRect.WriteSVG(sb, x, y)
		updateMaxx(maxX_, maxY_, &maxX, &maxY)
	}

	for _, rectAnchoredPath := range rect.RectAnchoredPaths {
		x, y := getRectAnchorPoint(rect, rectAnchoredPath.RectAnchorType)
		maxX_, maxY_ := rectAnchoredPath.WriteSVG(sb, x, y)
		updateMaxx(maxX_, maxY_, &maxX, &maxY)
	}

	return
}
