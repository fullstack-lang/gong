package models

import (
	"fmt"
	"strings"

	"github.com/fullstack-lang/gong/lib/svg/go/models/path"
)

type LinkAnchoredPath struct {
	Name string

	Definition string

	X_Offset float64
	Y_Offset float64

	// if true, path has the same Dimension that the rect it is anchored to
	// path must scale proportionnaly
	ScalePropotionnally bool

	// AppliedScaling is the scale that is applied is ScalePropotionnally is set
	// The value is initialized to 1 then scales with the scaling action
	// of the end user
	AppliedScaling float64

	Presentation
}

func (linkAnchoredPath *LinkAnchoredPath) WriteSVG(sb *strings.Builder, link *Link, segment *Segment) (maxX, maxY float64) {
	segmentOfInterest := *segment
	if segmentOfInterest.Type == StartSegment {
		segmentOfInterest = swapSegment(segmentOfInterest)
	}

	x, y := segmentOfInterest.EndPointWithoutRadius.X, segmentOfInterest.EndPointWithoutRadius.Y

	return linkAnchoredPath.writeSVG(sb, x, y)
}

func (linkAnchoredPath *LinkAnchoredPath) WriteSVGCorner(sb *strings.Builder, link *Link, segments []Segment) (maxX, maxY float64) {
	if len(segments) < 2 {
		return
	}

	segment0 := segments[0]
	segment1 := segments[1]

	var x, y float64

	isSeg0Horizontal := segment0.Orientation == ORIENTATION_HORIZONTAL
	isSeg0Vertical := segment0.Orientation == ORIENTATION_VERTICAL

	if isSeg0Horizontal {
		// Segment 1 is vertical: Center the text exactly on the corner's X coordinate
		x = segment0.EndPoint.X
	} else {
		// Segment 1 is horizontal: Place text opposite to the horizontal direction
		dirX := segment1.EndPoint.X - segment1.StartPoint.X
		signX := 1.0
		if dirX < 0 {
			signX = -1.0
		}
		paddingX := 12.0
		x = segment0.EndPoint.X - signX*paddingX
	}

	var dirY float64
	if isSeg0Vertical {
		dirY = segment0.EndPoint.Y - segment0.StartPoint.Y
	} else {
		dirY = segment1.EndPoint.Y - segment1.StartPoint.Y
	}

	signY := 1.0
	if dirY < 0 {
		signY = -1.0
	}
	paddingY := 8.0

	var yOffset float64
	if isSeg0Vertical {
		yOffset = signY * paddingY
	} else {
		yOffset = -signY * paddingY
	}

	y = segment0.EndPoint.Y + yOffset

	return linkAnchoredPath.writeSVG(sb, x, y)
}

func (linkAnchoredPath *LinkAnchoredPath) writeSVG(sb *strings.Builder, x, y float64) (maxX, maxY float64) {
	sb.WriteString(
		fmt.Sprintf(
			`	<path
			d="%s"`,
			linkAnchoredPath.Definition,
		))

	var presentation Presentation
	presentation = linkAnchoredPath.Presentation

	var scaling string

	if linkAnchoredPath.ScalePropotionnally {
		scaling = fmt.Sprintf("scale(%s %s)",
			formatFloat(linkAnchoredPath.AppliedScaling),
			formatFloat(linkAnchoredPath.AppliedScaling),
		)
	}

	presentation.Transform =
		fmt.Sprintf("translate(%s %s) %s",
			formatFloat(x+linkAnchoredPath.X_Offset),
			formatFloat(y+linkAnchoredPath.Y_Offset),
			scaling,
		) + presentation.Transform
	presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

	pathBound := path.ProcessSVGPath(linkAnchoredPath.Definition)
	maxX = x + linkAnchoredPath.X_Offset + pathBound.MaxX
	maxY = y + linkAnchoredPath.Y_Offset + pathBound.MaxY

	return
}
