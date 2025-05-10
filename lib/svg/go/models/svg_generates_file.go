package models

import (
	"fmt"
	"html" // For escaping text and attribute values
	"math"
	"os"
	"strconv"
	"strings"
)

// GenerateFile generates an SVG file that represents the content of the SVG object.
func (svg *SVG) GenerateFile(pathToFile string) (err error) {
	var sb strings.Builder

	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"100%\" height=\"100%\">\n")

	for _, layer := range svg.Layers {
		// Rects
		for _, rect := range layer.Rects {
			// Apply default styling for rectangles to match frontend (white fill, black stroke)
			rectPresentation := rect.Presentation // Make a copy to modify
			if rectPresentation.Color == "" {
				rectPresentation.Color = "white" // Default fill to white
			}
			if rectPresentation.Stroke == "" {
				rectPresentation.Stroke = "black" // Default stroke to black
			}
			if rectPresentation.StrokeWidth == 0 { // If stroke width is not set, default to 1
				rectPresentation.StrokeWidth = 1
			}

			sb.WriteString(fmt.Sprintf(`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
				formatFloat(rect.X), formatFloat(rect.Y), formatFloat(rect.Width), formatFloat(rect.Height),
				formatFloat(rect.RX), formatFloat(rect.RX)))
			appendPresentationAttributes(&sb, &rectPresentation, true) // Pass modified presentation
			sb.WriteString(" />\n")
		}

		// Circles
		for _, circle := range layer.Circles {
			// Apply similar default styling for circles if needed, or keep as is
			circlePresentation := circle.Presentation
			if circlePresentation.Color == "" { // Example: default circles to white fill too
				// circlePresentation.Color = "white"
			}
			if circlePresentation.Stroke == "" {
				// circlePresentation.Stroke = "black"
			}
			if circlePresentation.StrokeWidth == 0 && circlePresentation.Stroke != "" && circlePresentation.Stroke != "none" {
				// circlePresentation.StrokeWidth = 1
			}
			sb.WriteString(fmt.Sprintf(`  <circle cx="%s" cy="%s" r="%s"`,
				formatFloat(circle.CX), formatFloat(circle.CY), formatFloat(circle.Radius)))
			appendPresentationAttributes(&sb, &circlePresentation, true)
			sb.WriteString(" />\n")
		}

		// Lines
		for _, line := range layer.Lines {
			linePresentation := line.Presentation
			if linePresentation.Stroke == "" {
				linePresentation.Stroke = "black" // Default lines to black
			}
			if linePresentation.StrokeWidth == 0 {
				linePresentation.StrokeWidth = 1
			}
			sb.WriteString(fmt.Sprintf(`  <line x1="%s" y1="%s" x2="%s" y2="%s"`,
				formatFloat(line.X1), formatFloat(line.Y1), formatFloat(line.X2), formatFloat(line.Y2)))
			appendPresentationAttributes(&sb, &linePresentation, false)
			sb.WriteString(" />\n")
		}

		// Ellipses
		for _, ellipse := range layer.Ellipses {
			ellipsePresentation := ellipse.Presentation
			// Apply defaults if needed
			// if ellipsePresentation.Color == "" { ellipsePresentation.Color = "white" }
			// if ellipsePresentation.Stroke == "" { ellipsePresentation.Stroke = "black" }
			// if ellipsePresentation.StrokeWidth == 0 { ellipsePresentation.StrokeWidth = 1 }
			sb.WriteString(fmt.Sprintf(`  <ellipse cx="%s" cy="%s" rx="%s" ry="%s"`,
				formatFloat(ellipse.CX), formatFloat(ellipse.CY), formatFloat(ellipse.RX), formatFloat(ellipse.RY)))
			appendPresentationAttributes(&sb, &ellipsePresentation, true)
			sb.WriteString(" />\n")
		}

		// Paths (general paths, not links or arrows)
		for _, pathElement := range layer.Paths {
			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, html.EscapeString(pathElement.Definition)))
			appendPresentationAttributes(&sb, &pathElement.Presentation, true) // Path can be fillable
			sb.WriteString(" />\n")
		}

		// Polygones
		for _, polygone := range layer.Polygones {
			sb.WriteString(fmt.Sprintf(`  <polygon points="%s"`, html.EscapeString(polygone.Points)))
			appendPresentationAttributes(&sb, &polygone.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Polylines
		for _, polyline := range layer.Polylines {
			// Polylines are often not filled by default in some renderers, but can be.
			// SVG default fill is black. If you want them open with no fill unless specified:
			polylinePresentation := polyline.Presentation
			// if polylinePresentation.Color == "" { // If Color means fill for polyline
			// 	// To make it behave like an open path, ensure fill is "none" if not specified
			//  // However, appendPresentationAttributes with isFillable=true and empty Color will lead to SVG default black fill.
			//  // So, if default for polyline should be no fill:
			//  // Option 1: Set polylinePresentation.Color = "none" if empty
			//  // Option 2: Call appendPresentationAttributes with isFillable=false if Color is empty
			// }
			sb.WriteString(fmt.Sprintf(`  <polyline points="%s"`, html.EscapeString(polyline.Points)))
			appendPresentationAttributes(&sb, &polylinePresentation, true) // isFillable=true, so fill will be black if Color is empty
			sb.WriteString(" />\n")
		}

		// Texts
		for _, text := range layer.Texts {
			generateTextElement(&sb, text.Content, text.X, text.Y, text.FontSize, text.FontWeight, &text.Presentation, true)
		}

		// Links
		for _, link := range layer.Links {
			if link.Start == nil || link.End == nil {
				continue
			}

			var segments []Segment
			var pathData strings.Builder

			if link.IsBezierCurve {
				startX, startY := getAnchorPoint(link.Start, link.StartAnchorType, link.StartOrientation, link.StartRatio)
				endX, endY := getAnchorPoint(link.End, link.EndAnchorType, link.EndOrientation, link.EndRatio)
				segments = []Segment{
					{StartPoint: createPoint(startX, startY), EndPoint: createPoint(endX, endY)},
				}

				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(startX), formatFloat(startY)))
				if len(link.ControlPoints) == 2 {
					cp1, cp2 := link.ControlPoints[0], link.ControlPoints[1]
					pathData.WriteString(fmt.Sprintf(" C %s %s, %s %s, %s %s",
						formatFloat(cp1.X), formatFloat(cp1.Y), formatFloat(cp2.X), formatFloat(cp2.Y), formatFloat(endX), formatFloat(endY)))
				} else if len(link.ControlPoints) == 1 {
					cp1 := link.ControlPoints[0]
					pathData.WriteString(fmt.Sprintf(" Q %s %s, %s %s", formatFloat(cp1.X), formatFloat(cp1.Y), formatFloat(endX), formatFloat(endY)))
				} else {
					pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(endX), formatFloat(endY)))
				}
			} else {
				segments = generateSegments(link)
				if len(segments) == 0 {
					continue
				}
				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(segments[0].StartPoint.X), formatFloat(segments[0].StartPoint.Y)))
				for _, seg := range segments {
					pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(seg.EndPoint.X), formatFloat(seg.EndPoint.Y)))
				}
			}

			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, pathData.String()))
			customLinkPresentation := link.Presentation
			if customLinkPresentation.Stroke == "" {
				customLinkPresentation.Stroke = "black"
			}
			if customLinkPresentation.StrokeWidth == 0 {
				customLinkPresentation.StrokeWidth = 1
			}
			appendPresentationAttributes(&sb, &customLinkPresentation, false)
			sb.WriteString(" />\n")

			if len(segments) > 0 {
				if link.HasStartArrow && link.StartArrowSize > 0 {
					arrowPathDataStart := generateArrowPath(link, segments[0], link.StartArrowSize, true)
					if arrowPathDataStart != "" {
						sb.WriteString(fmt.Sprintf(`  <path d="%s"`, arrowPathDataStart))
						appendPresentationAttributes(&sb, &customLinkPresentation, false)
						sb.WriteString(" />\n")
					}
				}

				if link.HasEndArrow && link.EndArrowSize > 0 {
					arrowPathDataEnd := generateArrowPath(link, segments[len(segments)-1], link.EndArrowSize, false)
					if arrowPathDataEnd != "" {
						sb.WriteString(fmt.Sprintf(`  <path d="%s"`, arrowPathDataEnd))
						appendPresentationAttributes(&sb, &customLinkPresentation, false)
						sb.WriteString(" />\n")
					}
				}
			}

			if len(segments) > 0 {
				firstSegment := segments[0]
				anchorXStart := firstSegment.StartPoint.X
				anchorYStart := firstSegment.StartPoint.Y
				for _, anchoredText := range link.TextAtArrowStart {
					textX := anchorXStart + anchoredText.X_Offset
					textY := anchorYStart + anchoredText.Y_Offset
					generateTextElement(&sb, anchoredText.Content, textX, textY, "", "", &anchoredText.Presentation, false)
				}

				lastSegment := segments[len(segments)-1]
				anchorXEnd := lastSegment.EndPoint.X
				anchorYEnd := lastSegment.EndPoint.Y
				for _, anchoredText := range link.TextAtArrowEnd {
					textX := anchorXEnd + anchoredText.X_Offset
					textY := anchorYEnd + anchoredText.Y_Offset
					generateTextElement(&sb, anchoredText.Content, textX, textY, "", "", &anchoredText.Presentation, false)
				}
			}
		}
	}

	sb.WriteString("</svg>\n")
	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
}

func generateTextElement(sb *strings.Builder, content string, x, y float64, explicitFontSize, explicitFontWeight string, p *Presentation, useDefaultAnchors bool) {
	sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(x), formatFloat(y)))
	if !useDefaultAnchors {
		sb.WriteString(` text-anchor="middle" dominant-baseline="middle"`)
	}

	clonedPres := *p
	if clonedPres.Color == "" {
		clonedPres.Color = "black"
	}

	finalFontSize := explicitFontSize
	if finalFontSize == "" {
		finalFontSize = "10px"
	}
	sb.WriteString(fmt.Sprintf(` font-size="%s"`, html.EscapeString(finalFontSize)))

	finalFontWeight := explicitFontWeight
	if finalFontWeight != "" {
		sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(finalFontWeight)))
	}

	appendPresentationAttributes(sb, &clonedPres, true)

	sb.WriteString(">")
	sb.WriteString(html.EscapeString(content))
	sb.WriteString("</text>\n")
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func appendPresentationAttributes(sb *strings.Builder, p *Presentation, isFillable bool) {
	if isFillable {
		if p.Color != "" {
			// If color is "none", treat it as such, otherwise escape it.
			if strings.ToLower(p.Color) == "none" {
				sb.WriteString(` fill="none"`)
			} else {
				sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
			}
		} else {
			// If p.Color is empty for a fillable element, SVG defaults to black.
			// This was the issue for rectangles. Now handled by defaulting rectPresentation.Color to "white".
			// For other fillable elements, if Color is empty, they will get SVG default black fill.
		}

		if p.FillOpacity != 1.0 && p.FillOpacity != 0.0 {
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		} else if p.FillOpacity == 0.0 && (p.Color != "" && strings.ToLower(p.Color) != "none") {
			sb.WriteString(` fill-opacity="0"`)
		}
	} else { // Not fillable (e.g., lines, link paths)
		sb.WriteString(` fill="none"`)
	}

	if p.Stroke != "" {
		if strings.ToLower(p.Stroke) == "none" {
			sb.WriteString(` stroke="none"`)
		} else {
			sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
		}
	} else {
		// If stroke is not defined, SVG default is "none".
		// For elements like rects where we want a default border, this is handled by setting p.Stroke to "black" before this call.
	}

	// Handle stroke-width carefully:
	// SVG default stroke-width is 1.
	// If p.StrokeWidth is 0, it means "hairline" or thinnest possible, not "no stroke".
	// To have "no stroke", p.Stroke should be "none" or empty.
	if p.Stroke != "" && strings.ToLower(p.Stroke) != "none" { // Only apply stroke-width if there's a stroke color
		if p.StrokeWidth != 1.0 { // Write only if not the default 1.0
			sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
		}
		// If p.StrokeWidth is 1.0 (default), we don't need to write it.
		// If p.StrokeWidth is 0, it will be written as "0".
	}

	if p.StrokeDashArray != "" {
		sb.WriteString(fmt.Sprintf(` stroke-dasharray="%s"`, html.EscapeString(p.StrokeDashArray)))
	}
	if p.StrokeOpacity != 1.0 && p.StrokeOpacity != 0.0 {
		sb.WriteString(fmt.Sprintf(` stroke-opacity="%s"`, formatFloat(p.StrokeOpacity)))
	} else if p.StrokeOpacity == 0.0 && p.Stroke != "" && strings.ToLower(p.Stroke) != "none" {
		sb.WriteString(` stroke-opacity="0"`)
	}

	if p.Transform != "" {
		sb.WriteString(fmt.Sprintf(` transform="%s"`, html.EscapeString(p.Transform)))
	}
}

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
		if deltaX > 0 { // Right
			return localX, localY
		} else { // Left
			return -localX, localY
		}
	} else if segment.Orientation == ORIENTATION_VERTICAL {
		if deltaY > 0 { // Down
			return localY, localX
		} else { // Up
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
