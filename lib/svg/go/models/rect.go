package models

import (
	"fmt"
	"strings"
)

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

	CheckboxHasToolTip  bool
	CheckboxToolTipText string

	Impl RectImplInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *Stage, _, frontRect *Rect) {

	if rect.Impl != nil {
		rect.Impl.RectUpdated(frontRect)
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

	rect.Presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

	for _, anchoredText := range rect.RectAnchoredTexts {

		x, y := getRectAnchorPoint(rect, anchoredText.RectAnchorType)

		anchoredText.WriteSVG(sb, x, y)

	}

	return
}
