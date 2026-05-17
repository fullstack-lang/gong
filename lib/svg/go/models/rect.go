package models

import (
	"fmt"
	"strings"
)

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64

	// Peers are rects that move with this rect
	// notice that they are not updated on the mouse up.
	Peers []*Rect

	// EnclosingRect is the rect that encloses this rect
	// the rect cannot move outside of the enclosing rect
	EnclosingRect *Rect

	// Obstacles are rects that this rect cannot overlap with
	Obstacles []*Rect

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

	RectAnchoredTexts     []*RectAnchoredText
	RectAnchoredRects     []*RectAnchoredRect
	RectAnchoredPaths     []*RectAnchoredPath
	RectAnchoredPngImages []*RectAnchoredPngImage

	ChangeColorWhenHovered bool
	ColorWhenHovered       string
	OriginalColor          string
	FillOpacityWhenHovered float64
	OriginalFillOpacity    float64

	HasToolTip      bool
	ToolTipText     string
	ToolTipPosition ToolTipPositionEnum

	// deprecated, uses OnUpdate instead
	Impl RectImplInterface

	OnUpdate func(updatedRect *Rect)

	MouseEvent

	// URLPath is the path, if not empty, to navigate to when the rect is clicked, if not empty
	// only works with backend svg generation
	// <a href="https://example.com" target="_blank">
	//   <rect x="10" y="10" width="100" height="50" />
	// </a>
	URLPath string

	// URLTarget specifies how to open the link if URLPath is present
	URLTarget LinkTargetType
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *Stage, _, frontRect *Rect) {

	if rect.Impl != nil {
		rect.Impl.RectUpdated(frontRect)
	}
	if rect.OnUpdate != nil {
		rect.OnUpdate(frontRect)
	}
}

func (rect *Rect) WriteSVG(sb *strings.Builder) (maxX, maxY float64) {
	if rect.URLPath != "" {
		targetAttr := ""
		if rect.URLTarget != "" {
			targetAttr = fmt.Sprintf(` target="%s"`, rect.URLTarget)
		}
		sb.WriteString(fmt.Sprintf(`  <a href="%s"%s>`+"\n", rect.URLPath, targetAttr))
	}

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

	if rect.URLPath != "" {
		sb.WriteString("  </a>\n")
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

	// if there are after, they are more clickable than the rect itself
	for _, rectAnchoredText := range rect.RectAnchoredTexts {
		x, y := getRectAnchorPoint(rect, rectAnchoredText.RectAnchorType)
		rectAnchoredText.WriteSVG(sb, x, y)

	}

	return
}
