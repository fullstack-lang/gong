package models

import (
	"fmt"
	"strings"
)

type RectAnchoredRect struct {
	Name string

	X, Y, Width, Height, RX float64

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType

	// if true, rect has the same Dimension that the rect it is anchored to
	WidthFollowRect  bool
	HeightFollowRect bool

	HasToolTip  bool
	ToolTipText string

	Presentation
}

func (rectAnchoredRect *RectAnchoredRect) WriteSVG(sb *strings.Builder, x, y float64) (maxX, maxY float64) {

	sb.WriteString(
		fmt.Sprintf(
			`  <rect
			x="%s" 
			y="%s"
			width="%s"
			height="%s"`,
			formatFloat(x+rectAnchoredRect.X_Offset),
			formatFloat(y+rectAnchoredRect.Y_Offset),
			formatFloat(rectAnchoredRect.Width),
			formatFloat(rectAnchoredRect.Height),
		))

	rectAnchoredRect.Presentation.WriteSVG(sb)
	sb.WriteString(" >\n")

	sb.WriteString("</rect>\n")

	maxX = rectAnchoredRect.X + rectAnchoredRect.Width + rectAnchoredRect.X_Offset
	maxY = rectAnchoredRect.Y + rectAnchoredRect.Height + rectAnchoredRect.Y_Offset

	return
}
