package models

import (
	"fmt"
	"strings"

	"github.com/fullstack-lang/gong/lib/svg/go/models/path"
)

type RectAnchoredPath struct {
	Name string

	Definition string

	X_Offset float64
	Y_Offset float64

	// gong:width 400
	RectAnchorType RectAnchorType

	// if true, rect has the same Dimension that the rect it is anchored to
	// rect must scale proportionnaly
	ScalePropotionnally bool

	// AppliedScaling is the scale that is applied is ScalePropotionnally is set
	// The value is initialized to 1 then scales with the scaling action
	// of the end user
	AppliedScaling float64

	Presentation
}

func (rectAnchoredPath *RectAnchoredPath) WriteSVG(sb *strings.Builder, x, y float64) (maxX, maxY float64) {

	sb.WriteString(
		fmt.Sprintf(
			`<path
			d="%s"`,
			rectAnchoredPath.Definition,
		))

	// add a translation to the presentation of the path
	// (path does not treat x,y)
	var presentation Presentation
	presentation = rectAnchoredPath.Presentation

	presentation.Transform =
		fmt.Sprintf("translate(%s %s) ",
			formatFloat(x+rectAnchoredPath.X_Offset),
			formatFloat(y+rectAnchoredPath.Y_Offset),
		) + presentation.Transform
	presentation.WriteSVG(sb)
	sb.WriteString(" >\n")

	sb.WriteString("</path>\n")

	pathBound := path.ProcessSVGPath(rectAnchoredPath.Definition)
	maxX = x + rectAnchoredPath.X_Offset + pathBound.MaxX
	maxY = y + rectAnchoredPath.Y_Offset + pathBound.MaxY

	return
}
