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
			sb.WriteString(fmt.Sprintf(`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
				formatFloat(rect.X), formatFloat(rect.Y), formatFloat(rect.Width), formatFloat(rect.Height),
				formatFloat(rect.RX), formatFloat(rect.RX)))
			appendPresentationAttributes(&sb, &rect.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Circles
		for _, circle := range layer.Circles {
			sb.WriteString(fmt.Sprintf(`  <circle cx="%s" cy="%s" r="%s"`,
				formatFloat(circle.CX), formatFloat(circle.CY), formatFloat(circle.Radius)))
			appendPresentationAttributes(&sb, &circle.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Lines
		for _, line := range layer.Lines {
			sb.WriteString(fmt.Sprintf(`  <line x1="%s" y1="%s" x2="%s" y2="%s"`,
				formatFloat(line.X1), formatFloat(line.Y1), formatFloat(line.X2), formatFloat(line.Y2)))
			appendPresentationAttributes(&sb, &line.Presentation, false)
			sb.WriteString(" />\n")
		}

		// Ellipses
		for _, ellipse := range layer.Ellipses {
			sb.WriteString(fmt.Sprintf(`  <ellipse cx="%s" cy="%s" rx="%s" ry="%s"`,
				formatFloat(ellipse.CX), formatFloat(ellipse.CY), formatFloat(ellipse.RX), formatFloat(ellipse.RY)))
			appendPresentationAttributes(&sb, &ellipse.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Paths (general paths, not links or arrows)
		for _, pathElement := range layer.Paths {
			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, html.EscapeString(pathElement.Definition)))
			appendPresentationAttributes(&sb, &pathElement.Presentation, true)
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
			sb.WriteString(fmt.Sprintf(`  <polyline points="%s"`, html.EscapeString(polyline.Points)))
			appendPresentationAttributes(&sb, &polyline.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Texts
		for _, text := range layer.Texts {
			// For general texts, pass their own font attributes if they exist, otherwise use defaults from generateTextElement
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
				segments = []Segment{ // Simplified for text anchoring
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
					// LinkAnchoredText uses its own Presentation. Font attributes are not in Presentation struct.
					// So, pass empty strings for explicit font size/weight, relying on defaults or what appendPresentationAttributes might do.
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

// generateTextElement helper
// explicitFontSize and explicitFontWeight are for Text struct fields.
// For LinkAnchoredText, these will be empty, and we rely on defaults or Presentation if it ever gets font fields.
func generateTextElement(sb *strings.Builder, content string, x, y float64, explicitFontSize, explicitFontWeight string, p *Presentation, useDefaultAnchors bool) {
	sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(x), formatFloat(y)))
	if !useDefaultAnchors { // For link-anchored text
		sb.WriteString(` text-anchor="middle" dominant-baseline="middle"`)
	}

	clonedPres := *p // Make a copy to set defaults without modifying the original
	if clonedPres.Color == "" {
		clonedPres.Color = "black"
	}

	// Handle font size and weight
	// If explicit values are provided (from Text struct), use them.
	// Otherwise, check Presentation (though it doesn't have them now) or use a default.
	finalFontSize := explicitFontSize
	if finalFontSize == "" { // Not provided by Text struct, or called for LinkAnchoredText
		// if p.FontSize != "" { finalFontSize = p.FontSize } // If Presentation had font size
		// else {
		finalFontSize = "10px" // Default if no other source
		// }
	}
	sb.WriteString(fmt.Sprintf(` font-size="%s"`, html.EscapeString(finalFontSize)))

	finalFontWeight := explicitFontWeight
	if finalFontWeight == "" {
		// if p.FontWeight != "" { finalFontWeight = p.FontWeight } // If Presentation had font weight
		// No default font weight, browser will handle.
	}
	if finalFontWeight != "" {
		sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(finalFontWeight)))
	}

	appendPresentationAttributes(sb, &clonedPres, true) // Pass the copy

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
			sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
		}
		if p.FillOpacity != 1.0 && p.FillOpacity != 0.0 {
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		} else if p.FillOpacity == 0.0 && (p.Color != "" && p.Color != "none") {
			sb.WriteString(` fill-opacity="0"`)
		}
	} else {
		sb.WriteString(` fill="none"`)
	}

	if p.Stroke != "" {
		sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
	}
	if p.StrokeWidth != 1.0 && p.StrokeWidth != 0.0 {
		sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
	} else if p.StrokeWidth == 0.0 && p.Stroke != "" && p.Stroke != "none" {
		sb.WriteString(` stroke-width="0"`)
	}

	if p.StrokeDashArray != "" {
		sb.WriteString(fmt.Sprintf(` stroke-dasharray="%s"`, html.EscapeString(p.StrokeDashArray)))
	}
	if p.StrokeOpacity != 1.0 && p.StrokeOpacity != 0.0 {
		sb.WriteString(fmt.Sprintf(` stroke-opacity="%s"`, formatFloat(p.StrokeOpacity)))
	} else if p.StrokeOpacity == 0.0 && p.Stroke != "" && p.Stroke != "none" {
		sb.WriteString(` stroke-opacity="0"`)
	}

	// Fontsize and fontweight are NOT part of the generic Presentation struct
	// They are handled by generateTextElement directly or via Text struct fields.
	// So, no p.FontSize or p.FontWeight access here.

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
	return localX, localY // Fallback
}

func generateArrowPath(link *Link, segment Segment, arrowSize float64, isStartArrow bool) string {
	if arrowSize <= 0 {
		return ""
	}
	const tsRatio = 0.3535533905 // (1/sqrt(2))/2, from TS for base point offsetting

	refX, refY := 0.0, 0.0
	var localTip1X, localTip1Y, localBase1X, localBase1Y float64
	var localTip2X, localTip2Y, localBase2X, localBase2Y float64

	if isStartArrow { // Arrow at segment.StartPoint, pointing towards segment.EndPoint
		refX, refY = segment.StartPoint.X, segment.StartPoint.Y
		// Local coords for arrow pointing right (+X direction) relative to refX, refY
		// Tip components define the ends of the two lines forming the ">"
		localTip1X = arrowSize
		localTip1Y = -arrowSize // Upper leg of ">"
		localTip2X = arrowSize
		localTip2Y = arrowSize // Lower leg of ">"

		// Base components from TS logic, adjusted for arrow pointing right
		// Original TS for end arrow (pointing left): (sw*ratio, sw*ratio) and (sw*ratio, -sw*ratio)
		// For start arrow (pointing right), X base offset might be negative to be "behind" the tip.
		// Or, if base is where the lines *start* at the anchor point, localBase can be 0,0.
		// Let's use the TS logic for base points for consistency:
		localBase1X = -link.StrokeWidth * tsRatio // Offset "behind" the tip along X
		localBase1Y = link.StrokeWidth * tsRatio  // Offset perpendicular for one leg
		localBase2X = -link.StrokeWidth * tsRatio // Same X offset
		localBase2Y = -link.StrokeWidth * tsRatio // Offset perpendicular for the other leg

	} else { // End Arrow: at segment.EndPoint, pointing "backwards" (like "<")
		refX, refY = segment.EndPoint.X, segment.EndPoint.Y
		// Local coords for arrow pointing left (-X direction) relative to refX, refY
		// Tip components
		localTip1X = -arrowSize
		localTip1Y = -arrowSize // Upper leg of "<"
		localTip2X = -arrowSize
		localTip2Y = arrowSize // Lower leg of "<"

		// Base components from TS logic (getEndArrowPath)
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

	// Path: M Base1 L Tip1 M Base2 L Tip2
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
	} else { // ORIENTATION_VERTICAL
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

// generatePointRectSegment now uses link.StartRatio/EndRatio correctly
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
		// Original TS logic for choosing X connection point based on corner's relative position:
		if point.X <= rect.X+rect.Width/2 {
			rectConnectionPoint.X = rect.X
		} else {
			rectConnectionPoint.X = rect.X + rect.Width
		}
	} else { // ORIENTATION_VERTICAL
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
		// For Bezier, deciding which X (rect.X or rect.X + rect.Width) can be complex.
		// Defaulting to rect.X for now. A better heuristic might be needed.
		x = rect.X
	} else if orientation == ORIENTATION_VERTICAL {
		x = rect.X + ratio*rect.Width
		y = rect.Y // Defaulting to rect.Y
	} else {
		x, y = rect.X+rect.Width/2, rect.Y+rect.Height/2
	}
	return
}
