package models

import (
	"fmt"
	"math"
	"strings"
)

type Link struct {
	Name string

	Type LinkType

	// IsBezierCurve, if true will draw the link as a bezier curve instead of segments
	IsBezierCurve bool

	Start           *Rect
	StartAnchorType AnchorType

	End           *Rect
	EndAnchorType AnchorType

	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64

	// corner radius
	CornerRadius float64

	// End Arrows
	HasEndArrow  bool
	EndArrowSize float64

	// Start Arrows
	HasStartArrow  bool
	StartArrowSize float64

	// to be displayed at the start
	TextAtArrowStart []*LinkAnchoredText

	// to be displayed at the end
	TextAtArrowEnd []*LinkAnchoredText

	// for non floating orthogonal anchors
	ControlPoints []*Point

	Presentation

	Impl LinkImplInterface
}

func (link *Link) OnAfterUpdate(stage *Stage, _, frontLink *Link) {

	if link.Impl != nil {
		link.Impl.LinkUpdated(frontLink)
	}
}

func (link *Link) generateSegments() []Segment {
	if link.Start == nil || link.End == nil {
		return []Segment{}
	}
	startRect, endRect := link.Start, link.End
	startDirection, endDirection := link.StartOrientation, link.EndOrientation
	startRatio, endRatio := link.StartRatio, link.EndRatio
	cornerOffsetRatio, cornerRadius := link.CornerOffsetRatio, link.CornerRadius
	segments := []Segment{}

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_VERTICAL {
		c1Y := startRect.Y + startRatio*startRect.Height
		c1X := endRect.X + endRatio*endRect.Width
		c1 := createPoint(c1X, c1Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true, StartSegment)
		s2 := generatePointRectSegment(c1, endRect, link, endDirection, cornerRadius, 1, false, EndSegment)
		segments = append(segments, s1, s2)

	} else if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X := startRect.X + startRatio*startRect.Width
		c1Y := endRect.Y + endRatio*endRect.Height
		c1 := createPoint(c1X, c1Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true, StartSegment)
		s2 := generatePointRectSegment(c1, endRect, link, endDirection, cornerRadius, 1, false, EndSegment)
		segments = append(segments, s1, s2)

	} else if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X := startRect.X + cornerOffsetRatio*startRect.Width
		c1Y := startRect.Y + startRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)
		c2X := c1X
		c2Y := endRect.Y + endRatio*endRect.Height
		c2 := createPoint(c2X, c2Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true, StartSegment)
		s2 := generatePointPointSegment(c1, c2, ORIENTATION_VERTICAL, cornerRadius, 1)
		s3 := generatePointRectSegment(c2, endRect, link, endDirection, cornerRadius, 2, false, EndSegment)
		if math.Abs(c1Y-c2Y) <= 2*cornerRadius && cornerRadius > 0 {
			c2a := createPoint(c2X, c1Y)
			s1 = generatePointRectSegment(c1, startRect, link, startDirection, 0, 0, true, StartSegment)
			s2 = generatePointPointSegment(c1, c2a, ORIENTATION_HORIZONTAL, 0, 1)
			s3 = generatePointRectSegment(c2a, endRect, link, endDirection, 0, 2, false, EndSegment)
		}
		segments = append(segments, s1, s2, s3)

	} else if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_VERTICAL {
		c1X := startRect.X + startRatio*startRect.Width
		c1Y := startRect.Y + cornerOffsetRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)
		c2X := endRect.X + endRatio*endRect.Width
		c2Y := c1Y
		c2 := createPoint(c2X, c2Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true, StartSegment)
		s2 := generatePointPointSegment(c1, c2, ORIENTATION_HORIZONTAL, cornerRadius, 1)
		s3 := generatePointRectSegment(c2, endRect, link, endDirection, cornerRadius, 2, false, EndSegment)
		if math.Abs(c1X-c2X) <= 2*cornerRadius && cornerRadius > 0 {
			c2a := createPoint(c1X, c2Y)
			s1 = generatePointRectSegment(c1, startRect, link, startDirection, 0, 0, true, StartSegment)
			s2 = generatePointPointSegment(c1, c2a, ORIENTATION_VERTICAL, 0, 1)
			s3 = generatePointRectSegment(c2a, endRect, link, endDirection, 0, 2, false, EndSegment)
		}
		segments = append(segments, s1, s2, s3)
	}
	return segments
}

func (link *Link) WriteSVGEndArrow(sb *strings.Builder, segment *Segment) {
	const ratio = 0.707106781 / 2 // (1/sqrt(2)) / 2

	arrowSize := link.EndArrowSize

	firstStartX := segment.EndPointWithoutRadius.X
	firstStartY := segment.EndPointWithoutRadius.Y

	secondStartX := segment.EndPointWithoutRadius.X
	secondStartY := segment.EndPointWithoutRadius.Y

	firstTipX := segment.EndPointWithoutRadius.X
	firstTipY := segment.EndPointWithoutRadius.Y
	secondTipX := segment.EndPointWithoutRadius.X
	secondTipY := segment.EndPointWithoutRadius.Y

	// Calculate first tip and start points
	dx, dy := rotateToSegmentDirection(segment, -arrowSize, -arrowSize)
	firstTipX += dx
	firstTipY += dy

	dx, dy = rotateToSegmentDirection(segment, link.StrokeWidth*ratio, link.StrokeWidth*ratio)
	firstStartX += dx
	firstStartY += dy

	// Calculate second tip and start points
	dx, dy = rotateToSegmentDirection(segment, -arrowSize, arrowSize)
	secondTipX += dx
	secondTipY += dy

	dx, dy = rotateToSegmentDirection(segment, link.StrokeWidth*ratio, -link.StrokeWidth*ratio)
	secondStartX += dx
	secondStartY += dy

	sb.WriteString(fmt.Sprintf("	<path d=\"M %f %f L %f %f M %f %f L %f %f\"",
		firstStartX,
		firstStartY,
		firstTipX,
		firstTipY,

		secondStartX,
		secondStartY,
		secondTipX,
		secondTipY))

	link.Presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

}

func (link *Link) WriteSVGArcPath(sb *strings.Builder, segment, nextSegment *Segment) {

	// The TypeScript version uses constant startDegree and endDegree,
	// which means startRadians and endRadians are also constant.
	// However, largeArcFlag depends on these, so we keep the calculation
	// for clarity, even though it's currently fixed.
	const startDegree = 180.0
	const endDegree = 270.0
	// const startRadians = (startDegree * math.Pi) / 180.0 // Not directly used in path string
	// const endRadians = (endDegree * math.Pi) / 180.0 // Not directly used in path string

	startX := segment.EndPoint.X
	startY := segment.EndPoint.Y
	endX := nextSegment.StartPoint.X
	endY := nextSegment.StartPoint.Y

	largeArcFlag := 0
	if (endDegree - startDegree) > 180.0 { // Or endDegree - startDegree <= 180 ? 0 : 1
		largeArcFlag = 1
	}

	sweepFlag := 0

	if segment.Orientation == ORIENTATION_HORIZONTAL {
		segmentDirection := 0
		if segment.EndPoint.X > segment.StartPoint.X {
			segmentDirection = 1
		} else {
			segmentDirection = -1
		}

		cornerDirection := 0
		if segment.EndPoint.Y < nextSegment.StartPoint.Y {
			cornerDirection = 1
		} else {
			cornerDirection = -1
		}

		if segmentDirection*cornerDirection == 1 {
			sweepFlag = 1
		} else {
			sweepFlag = 0
		}

	} else if segment.Orientation == ORIENTATION_VERTICAL {
		segmentDirection := 0
		if segment.EndPoint.Y > segment.StartPoint.Y {
			segmentDirection = 1
		} else {
			segmentDirection = -1
		}

		cornerDirection := 0
		if segment.EndPoint.X < nextSegment.StartPoint.X { // Corrected based on typical SVG arc logic (clockwise/counter-clockwise)
			cornerDirection = 1
		} else {
			cornerDirection = -1
		}

		// This logic for sweepFlag in vertical orientation seems to be swapped compared to horizontal
		// in the original TS. If the visual output is correct with the TS code,
		// then this translation maintains that logic.
		if segmentDirection*cornerDirection == 1 {
			sweepFlag = 0
		} else {
			sweepFlag = 1
		}
	}

	// Ensure link.CornerRadius is available and has a sensible value.
	// If link.CornerRadius could be zero, SVG arc with 0 radius is problematic.
	// The original TS code doesn't show a default or handling for 0.
	cornerRadius := link.CornerRadius
	if cornerRadius <= 0 {
		// Default to a small radius if not set or invalid, or handle as an error
		// For now, let's assume it's always positive as in the TS
		// For a straight line instead of an arc if radius is 0, one might use L command
		// return fmt.Sprintf("M %f %f L %f %f", startX, startY, endX, endY)
	}

	sb.WriteString(fmt.Sprintf("	<path d=\"M %f %f A %f %f 0 %d %d %f %f\"",
		startX, startY,
		cornerRadius, cornerRadius, /* rx, ry */
		largeArcFlag, sweepFlag,
		endX, endY))

	link.Presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

}
