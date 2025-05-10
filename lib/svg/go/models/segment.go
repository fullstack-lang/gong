package models

import (
	"fmt"
	"math"
)

type Segment struct {
	StartPoint              Point
	EndPoint                Point
	StartPointWithoutRadius Point
	EndPointWithoutRadius   Point
	Orientation             OrientationType
	Number                  int
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

func rotateToSegmentDirection(segment Segment, localX, localY float64) (float64, float64) {
	deltaX := segment.EndPoint.X - segment.StartPoint.X
	deltaY := segment.EndPoint.Y - segment.StartPoint.Y

	if segment.Orientation == ORIENTATION_HORIZONTAL {
		if deltaX > 0 {
			return localX, localY
		} else {
			return -localX, localY
		}
	} else if segment.Orientation == ORIENTATION_VERTICAL {
		if deltaY > 0 {
			return localY, localX
		} else {
			return -localY, -localX
		}
	}
	return localX, localY
}

func generateArrowPath(link *Link, segment Segment, arrowSize float64, isStartArrow bool) string {
	if arrowSize <= 0 {
		return ""
	}
	const tsRatio = 0.3535533905

	refX, refY := 0.0, 0.0
	var localTip1X, localTip1Y, localBase1X, localBase1Y float64
	var localTip2X, localTip2Y, localBase2X, localBase2Y float64

	if isStartArrow {
		refX, refY = segment.StartPoint.X, segment.StartPoint.Y
		localTip1X = arrowSize
		localTip1Y = -arrowSize
		localTip2X = arrowSize
		localTip2Y = arrowSize

		localBase1X = -link.StrokeWidth * tsRatio
		localBase1Y = link.StrokeWidth * tsRatio
		localBase2X = -link.StrokeWidth * tsRatio
		localBase2Y = -link.StrokeWidth * tsRatio

	} else {
		refX, refY = segment.EndPoint.X, segment.EndPoint.Y
		localTip1X = -arrowSize
		localTip1Y = -arrowSize
		localTip2X = -arrowSize
		localTip2Y = arrowSize

		localBase1X = link.StrokeWidth * tsRatio
		localBase1Y = link.StrokeWidth * tsRatio
		localBase2X = link.StrokeWidth * tsRatio
		localBase2Y = -link.StrokeWidth * tsRatio
	}

	rotatedTip1X, rotatedTip1Y := rotateToSegmentDirection(segment, localTip1X, localTip1Y)
	rotatedBase1X, rotatedBase1Y := rotateToSegmentDirection(segment, localBase1X, localBase1Y)
	rotatedTip2X, rotatedTip2Y := rotateToSegmentDirection(segment, localTip2X, localTip2Y)
	rotatedBase2X, rotatedBase2Y := rotateToSegmentDirection(segment, localBase2X, localBase2Y)

	finalTip1X := refX + rotatedTip1X
	finalTip1Y := refY + rotatedTip1Y
	finalBase1X := refX + rotatedBase1X
	finalBase1Y := refY + rotatedBase1Y

	finalTip2X := refX + rotatedTip2X
	finalTip2Y := refY + rotatedTip2Y
	finalBase2X := refX + rotatedBase2X
	finalBase2Y := refY + rotatedBase2Y

	return fmt.Sprintf("M %s %s L %s %s M %s %s L %s %s",
		formatFloat(finalBase1X), formatFloat(finalBase1Y),
		formatFloat(finalTip1X), formatFloat(finalTip1Y),
		formatFloat(finalBase2X), formatFloat(finalBase2Y),
		formatFloat(finalTip2X), formatFloat(finalTip2Y))
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
		StartPoint: newStart, EndPoint: newEnd,
		StartPointWithoutRadius: start, EndPointWithoutRadius: end,
		Orientation: orientation, Number: number,
	}
}

func generatePointRectSegment(point Point, rect *Rect, link *Link, direction OrientationType, cornerRadius float64, number int, isStartSegment bool) Segment {
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

	if isStartSegment {
		segment = swapSegment(segment)
	}
	return segment
}

func generateSegments(link *Link) []Segment {
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
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true)
		s2 := generatePointRectSegment(c1, endRect, link, endDirection, cornerRadius, 1, false)
		segments = append(segments, s1, s2)

	} else if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X := startRect.X + startRatio*startRect.Width
		c1Y := endRect.Y + endRatio*endRect.Height
		c1 := createPoint(c1X, c1Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true)
		s2 := generatePointRectSegment(c1, endRect, link, endDirection, cornerRadius, 1, false)
		segments = append(segments, s1, s2)

	} else if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X := startRect.X + cornerOffsetRatio*startRect.Width
		c1Y := startRect.Y + startRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)
		c2X := c1X
		c2Y := endRect.Y + endRatio*endRect.Height
		c2 := createPoint(c2X, c2Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true)
		s2 := generatePointPointSegment(c1, c2, ORIENTATION_VERTICAL, cornerRadius, 1)
		s3 := generatePointRectSegment(c2, endRect, link, endDirection, cornerRadius, 2, false)
		if math.Abs(c1Y-c2Y) <= 2*cornerRadius && cornerRadius > 0 {
			c2a := createPoint(c2X, c1Y)
			s1 = generatePointRectSegment(c1, startRect, link, startDirection, 0, 0, true)
			s2 = generatePointPointSegment(c1, c2a, ORIENTATION_HORIZONTAL, 0, 1)
			s3 = generatePointRectSegment(c2a, endRect, link, endDirection, 0, 2, false)
		}
		segments = append(segments, s1, s2, s3)

	} else if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_VERTICAL {
		c1X := startRect.X + startRatio*startRect.Width
		c1Y := startRect.Y + cornerOffsetRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)
		c2X := endRect.X + endRatio*endRect.Width
		c2Y := c1Y
		c2 := createPoint(c2X, c2Y)
		s1 := generatePointRectSegment(c1, startRect, link, startDirection, cornerRadius, 0, true)
		s2 := generatePointPointSegment(c1, c2, ORIENTATION_HORIZONTAL, cornerRadius, 1)
		s3 := generatePointRectSegment(c2, endRect, link, endDirection, cornerRadius, 2, false)
		if math.Abs(c1X-c2X) <= 2*cornerRadius && cornerRadius > 0 {
			c2a := createPoint(c1X, c2Y)
			s1 = generatePointRectSegment(c1, startRect, link, startDirection, 0, 0, true)
			s2 = generatePointPointSegment(c1, c2a, ORIENTATION_VERTICAL, 0, 1)
			s3 = generatePointRectSegment(c2a, endRect, link, endDirection, 0, 2, false)
		}
		segments = append(segments, s1, s2, s3)
	}
	return segments
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
