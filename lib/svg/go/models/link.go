package models

import "math"

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
