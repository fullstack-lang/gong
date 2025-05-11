package models

import (
	"fmt"
	"math"
	"strings"
)

type SegmentType string

const (
	StartSegment  SegmentType = "Start Segment"
	MiddleSegment SegmentType = "Middle Segment"
	EndSegment    SegmentType = "End Segment"
)

type Segment struct {
	StartPoint              Point
	EndPoint                Point
	StartPointWithoutRadius Point
	EndPointWithoutRadius   Point
	Orientation             OrientationType
	Number                  int
	Type                    SegmentType
}

func (segment *Segment) WriteSVG(sb *strings.Builder, link *Link) (maxX, maxY float64) {

	var str string

	switch segment.Type {
	case StartSegment:
		str = fmt.Sprintf(
			`	<line x1="%s" y1="%s" x2="%s" y2="%s"`,
			formatFloat(segment.StartPointWithoutRadius.X),
			formatFloat(segment.StartPointWithoutRadius.Y),
			formatFloat(segment.EndPoint.X),
			formatFloat(segment.EndPoint.Y))
	case MiddleSegment:
		str = fmt.Sprintf(
			`	<line x1="%s" y1="%s" x2="%s" y2="%s"`,
			formatFloat(segment.StartPoint.X),
			formatFloat(segment.StartPoint.Y),
			formatFloat(segment.EndPoint.X),
			formatFloat(segment.EndPoint.Y))
	case EndSegment:
		str = fmt.Sprintf(
			`	<line x1="%s" y1="%s" x2="%s" y2="%s"`,
			formatFloat(segment.StartPoint.X),
			formatFloat(segment.StartPoint.Y),
			formatFloat(segment.EndPointWithoutRadius.X),
			formatFloat(segment.EndPointWithoutRadius.Y))
	}

	sb.WriteString(str)

	maxX = math.Max(segment.StartPoint.X, segment.EndPoint.X)
	maxY = math.Max(segment.StartPoint.Y, segment.EndPoint.Y)

	link.Presentation.WriteSVG(sb)
	sb.WriteString(" />\n")

	return
}

func createPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func swapSegment(segment Segment) Segment {
	res := segment
	res.StartPoint, res.EndPoint = segment.EndPoint, segment.StartPoint
	res.StartPointWithoutRadius, res.EndPointWithoutRadius = segment.EndPointWithoutRadius, segment.StartPointWithoutRadius
	return res
}

// rotateToSegmentDirection rotates a point (x, y) based on the segment's direction.
// This is a Go translation of the TypeScript helper function.
func rotateToSegmentDirection(segment *Segment, x, y float64) (float64, float64) {
	xRes := 0.0
	yRes := 0.0

	if segment.Orientation == ORIENTATION_HORIZONTAL { // Assumes ORIENTATION_HORIZONTAL is defined
		if segment.EndPoint.X > segment.StartPoint.X { // 0 degrees
			xRes = x
			yRes = y
		} else { // 180 degrees (pi radians)
			xRes = -x
			yRes = y
		}
	}
	if segment.Orientation == ORIENTATION_VERTICAL { // Assumes ORIENTATION_VERTICAL is defined
		if segment.EndPoint.Y > segment.StartPoint.Y { // 90 degrees (pi/2 radians)
			xRes = y
			yRes = x
		} else { // 270 degrees (3*pi/2 radians)
			xRes = -y
			yRes = -x
		}
	}
	return xRes, yRes
}

func generatePointPointSegment(start Point, end Point, orientation OrientationType, cornerRadius float64, number int) Segment {
	newStart, newEnd := start, end
	if orientation == ORIENTATION_HORIZONTAL {
		if start.X < end.X {
			newStart.X += cornerRadius
			newEnd.X -= cornerRadius
		} else {
			newStart.X -= cornerRadius
			newEnd.X += cornerRadius
		}
	} else {
		if start.Y < end.Y {
			newStart.Y += cornerRadius
			newEnd.Y -= cornerRadius
		} else {
			newStart.Y -= cornerRadius
			newEnd.Y += cornerRadius
		}
	}
	return Segment{
		StartPoint:              newStart,
		EndPoint:                newEnd,
		StartPointWithoutRadius: start,
		EndPointWithoutRadius:   end,
		Orientation:             orientation,
		Number:                  number,
		Type:                    MiddleSegment,
	}
}

func generatePointRectSegment(
	point Point,
	rect *Rect,
	link *Link,
	direction OrientationType,
	cornerRadius float64,
	number int,
	isStartSegment bool,
	segmentType SegmentType) Segment {

	var ratio float64
	if isStartSegment {
		ratio = link.StartRatio
	} else {
		ratio = link.EndRatio
	}

	var rectConnectionPoint Point
	if direction == ORIENTATION_HORIZONTAL {
		rectConnectionPoint.Y = rect.Y + ratio*rect.Height
		if point.X <= rect.X+rect.Width/2 {
			rectConnectionPoint.X = rect.X
		} else {
			rectConnectionPoint.X = rect.X + rect.Width
		}
	} else {
		rectConnectionPoint.X = rect.X + ratio*rect.Width
		if point.Y <= rect.Y+rect.Height/2 {
			rectConnectionPoint.Y = rect.Y
		} else {
			rectConnectionPoint.Y = rect.Y + rect.Height
		}
	}

	segment := generatePointPointSegment(point, rectConnectionPoint, direction, cornerRadius, number)
	segment.Type = segmentType

	if isStartSegment {
		segment = swapSegment(segment)
	}
	return segment
}

func getAnchorPoint(rect *Rect, _ AnchorType, orientation OrientationType, ratio float64) (x, y float64) {
	if orientation == ORIENTATION_HORIZONTAL {
		y = rect.Y + ratio*rect.Height
		x = rect.X
	} else if orientation == ORIENTATION_VERTICAL {
		x = rect.X + ratio*rect.Width
		y = rect.Y
	} else {
		x, y = rect.X+rect.Width/2, rect.Y+rect.Height/2
	}
	return
}
