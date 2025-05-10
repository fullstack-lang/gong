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
// It aims to produce an output similar to what would be displayed in a frontend
// rendering this SVG model.
// The pathToFile is the desired path for the output SVG file.
func (svg *SVG) GenerateFile(pathToFile string) (err error) {
	var sb strings.Builder

	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"100%\" height=\"100%\">\n") // Added width/height for better viewing

	needsStartArrowMarker := false
	needsEndArrowMarker := false
	for _, layer := range svg.Layers {
		for _, link := range layer.Links {
			if link.HasStartArrow {
				needsStartArrowMarker = true
			}
			if link.HasEndArrow {
				needsEndArrowMarker = true
			}
			if needsStartArrowMarker && needsEndArrowMarker {
				break
			}
		}
		if needsStartArrowMarker && needsEndArrowMarker {
			break
		}
	}

	if needsStartArrowMarker || needsEndArrowMarker {
		sb.WriteString("  <defs>\n")
		// Arrowhead marker color should ideally match the link's stroke.
		// For now, using black. The 'fill' attribute on marker can be set.
		if needsEndArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadend" markerWidth="10" markerHeight="7" refX="9" refY="3.5" orient="auto" fill="black">
      <polygon points="0 0, 10 3.5, 0 7" />
    </marker>
`)) // refX="9" to pull arrow back slightly onto the line
		}
		if needsStartArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadstart" markerWidth="10" markerHeight="7" refX="1" refY="3.5" orient="auto" fill="black">
      <polygon points="10 0, 0 3.5, 10 7" />
    </marker>
`)) // refX="1" for start arrow
		}
		sb.WriteString("  </defs>\n")
	}

	for _, layer := range svg.Layers {
		// Rects
		for _, rect := range layer.Rects {
			sb.WriteString(fmt.Sprintf(`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
				formatFloat(rect.X), formatFloat(rect.Y), formatFloat(rect.Width), formatFloat(rect.Height),
				formatFloat(rect.RX), formatFloat(rect.RX)))
			appendPresentationAttributes(&sb, &rect.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Circles, Lines, Ellipses, Paths, Polygones, Polylines (omitted for brevity, assumed unchanged)
		for _, circle := range layer.Circles {
			sb.WriteString(fmt.Sprintf(`  <circle cx="%s" cy="%s" r="%s"`,
				formatFloat(circle.CX), formatFloat(circle.CY), formatFloat(circle.Radius)))
			appendPresentationAttributes(&sb, &circle.Presentation, true)
			sb.WriteString(" />\n")
		}

		for _, line := range layer.Lines {
			sb.WriteString(fmt.Sprintf(`  <line x1="%s" y1="%s" x2="%s" y2="%s"`,
				formatFloat(line.X1), formatFloat(line.Y1), formatFloat(line.X2), formatFloat(line.Y2)))
			appendPresentationAttributes(&sb, &line.Presentation, false)
			sb.WriteString(" />\n")
		}

		for _, ellipse := range layer.Ellipses {
			sb.WriteString(fmt.Sprintf(`  <ellipse cx="%s" cy="%s" rx="%s" ry="%s"`,
				formatFloat(ellipse.CX), formatFloat(ellipse.CY), formatFloat(ellipse.RX), formatFloat(ellipse.RY)))
			appendPresentationAttributes(&sb, &ellipse.Presentation, true)
			sb.WriteString(" />\n")
		}

		for _, path := range layer.Paths {
			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, html.EscapeString(path.Definition)))
			appendPresentationAttributes(&sb, &path.Presentation, true)
			sb.WriteString(" />\n")
		}

		for _, polygone := range layer.Polygones {
			sb.WriteString(fmt.Sprintf(`  <polygon points="%s"`, html.EscapeString(polygone.Points)))
			appendPresentationAttributes(&sb, &polygone.Presentation, true)
			sb.WriteString(" />\n")
		}

		for _, polyline := range layer.Polylines {
			sb.WriteString(fmt.Sprintf(`  <polyline points="%s"`, html.EscapeString(polyline.Points)))
			appendPresentationAttributes(&sb, &polyline.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Texts
		for _, text := range layer.Texts {
			generateTextElement(&sb, text.Content, text.X, text.Y, &text.Presentation, true) // true for general text anchor
		}

		// Links
		for _, link := range layer.Links {
			if link.Start == nil || link.End == nil {
				continue
			}

			var segments []Segment
			var pathData strings.Builder

			if link.IsBezierCurve {
				// Bezier curve logic (simplified)
				startX, startY := getAnchorPoint(link.Start, link.StartAnchorType, link.StartOrientation, link.StartRatio) // TODO: This might need StartRect/EndRect context
				endX, endY := getAnchorPoint(link.End, link.EndAnchorType, link.EndOrientation, link.EndRatio)             // TODO: Same here

				segments = []Segment{ // Simplified segment representation for Bezier text anchoring
					{StartPoint: createPoint(startX, startY), EndPoint: createPoint(endX, endY)},
				}

				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(startX), formatFloat(startY)))
				if len(link.ControlPoints) == 2 {
					cp1 := link.ControlPoints[0]
					cp2 := link.ControlPoints[1]
					pathData.WriteString(fmt.Sprintf(" C %s %s, %s %s, %s %s",
						formatFloat(cp1.X), formatFloat(cp1.Y),
						formatFloat(cp2.X), formatFloat(cp2.Y),
						formatFloat(endX), formatFloat(endY)))
				} else if len(link.ControlPoints) == 1 {
					cp1 := link.ControlPoints[0]
					pathData.WriteString(fmt.Sprintf(" Q %s %s, %s %s",
						formatFloat(cp1.X), formatFloat(cp1.Y),
						formatFloat(endX), formatFloat(endY)))
				} else {
					pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(endX), formatFloat(endY)))
				}
			} else {
				// Orthogonal line segments - using the revised logic
				segments = generateSegments(link) // Corrected function name
				if len(segments) == 0 {
					continue
				}

				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(segments[0].StartPoint.X), formatFloat(segments[0].StartPoint.Y)))
				for _, seg := range segments {
					// The points in segments are already adjusted for corner radius.
					// So, just draw lines between them. Arcs would require more complex path commands.
					// The visual output suggests sharp corners, so L commands should be fine.
					// If seg.StartPoint is where the previous segment ended (or M for the first),
					// we just need to L to seg.EndPoint.
					// The loop structure was slightly off, let's draw from start of first segment to its end,
					// then from start of second to its end, etc.
					// Path should be: M seg[0].Start -> L seg[0].End -> L seg[1].Start (if different) -> L seg[1].End ...
					// However, the segment generation ensures seg[i].StartPoint is where seg[i-1].EndPoint was, after cornering.
					// So, just L to each EndPoint is correct.
					pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(seg.EndPoint.X), formatFloat(seg.EndPoint.Y)))
				}
			}

			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, pathData.String()))
			customLinkPresentation := link.Presentation
			if customLinkPresentation.Stroke == "" {
				customLinkPresentation.Stroke = "black"
			}
			if customLinkPresentation.StrokeWidth == 0 {
				customLinkPresentation.StrokeWidth = 2 // Made links slightly thicker like in example
			}
			appendPresentationAttributes(&sb, &customLinkPresentation, false)

			markerEnd := ""
			if link.HasEndArrow && needsEndArrowMarker {
				markerEnd = ` marker-end="url(#arrowheadend)"`
			}
			markerStart := ""
			if link.HasStartArrow && needsStartArrowMarker {
				markerStart = ` marker-start="url(#arrowheadstart)"`
			}
			sb.WriteString(markerStart)
			sb.WriteString(markerEnd)
			sb.WriteString(" />\n")

			// Process TextAtArrowStart
			if len(segments) > 0 {
				firstSegment := segments[0]
				// Anchor point for text is the StartPoint of the link path
				anchorX := firstSegment.StartPoint.X
				anchorY := firstSegment.StartPoint.Y

				for _, anchoredText := range link.TextAtArrowStart {
					textX := anchorX + anchoredText.X_Offset
					textY := anchorY + anchoredText.Y_Offset
					generateTextElement(&sb, anchoredText.Content, textX, textY, &anchoredText.Presentation, false)
				}
			}

			// Process TextAtArrowEnd
			if len(segments) > 0 {
				lastSegment := segments[len(segments)-1]
				// Anchor point for text is the EndPoint of the link path
				anchorX := lastSegment.EndPoint.X
				anchorY := lastSegment.EndPoint.Y

				for _, anchoredText := range link.TextAtArrowEnd {
					textX := anchorX + anchoredText.X_Offset
					textY := anchorY + anchoredText.Y_Offset
					generateTextElement(&sb, anchoredText.Content, textX, textY, &anchoredText.Presentation, false)
				}
			}
		}
	}

	sb.WriteString("</svg>\n")
	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
}

// generateTextElement helper
// set useDefaultAnchors to true for general text elements, false for link-anchored text
// where specific x,y are usually enough and default SVG text-anchor="start" is often fine.
func generateTextElement(sb *strings.Builder, content string, x, y float64, p *Presentation, useDefaultAnchors bool) {
	sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(x), formatFloat(y)))

	if useDefaultAnchors { // For general texts from layer.Texts
		// Consider if text-anchor or dominant-baseline should be set by default for layer.Texts
		// For link anchored text, X_Offset, Y_Offset are more direct.
	} else { // For link-anchored text, to mimic frontend better:
		sb.WriteString(` text-anchor="middle" dominant-baseline="middle"`)
	}

	clonedPres := *p
	if clonedPres.Color == "" {
		clonedPres.Color = "black"
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
			sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
		}
		if p.FillOpacity != 1.0 && p.FillOpacity != 0.0 { // 0.0 is fully transparent
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
	if p.StrokeWidth != 1.0 && p.StrokeWidth != 0.0 { // 0.0 is no stroke
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

	// Text anchor and dominant baseline are now handled in generateTextElement for finer control

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
	// ArrowEndAnchoredText: Array<Offset> // This was in TS, not directly used in Go path gen
}

func createPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// Go equivalent of swapSegment
func swapSegment(segment Segment) Segment {
	res := segment // Shallow copy for simple fields
	res.StartPoint = segment.EndPoint
	res.StartPointWithoutRadius = segment.EndPointWithoutRadius
	res.EndPoint = segment.StartPoint
	res.EndPointWithoutRadius = segment.StartPointWithoutRadius
	return res
}

// Go equivalent of drawPointPointSegment from draw.point.point.segment.ts
func generatePointPointSegment(start Point, end Point, orientation OrientationType, cornerRadius float64, number int) Segment {
	newStart := start
	newEnd := end

	if orientation == ORIENTATION_HORIZONTAL {
		if start.X < end.X { // Left to Right
			newStart.X = start.X + cornerRadius
			newEnd.X = end.X - cornerRadius
		} else { // Right to Left
			newStart.X = start.X - cornerRadius
			newEnd.X = end.X + cornerRadius
		}
	} else { // ORIENTATION_VERTICAL
		if start.Y < end.Y { // Top to Bottom
			newStart.Y = start.Y + cornerRadius
			newEnd.Y = end.Y - cornerRadius
		} else { // Bottom to Top
			newStart.Y = start.Y - cornerRadius
			newEnd.Y = end.Y + cornerRadius
		}
	}

	return Segment{
		StartPoint:              newStart,
		EndPoint:                newEnd,
		StartPointWithoutRadius: start,
		EndPointWithoutRadius:   end,
		Orientation:             orientation,
		Number:                  number,
	}
}

// Go equivalent of drawPointRectSegment from draw.point.rect.segment.ts
func generatePointRectSegment(point Point, rect *Rect, direction OrientationType, cornerRadius float64, number int) Segment {
	// Initial segment is from `point` (corner) to a dummy end point.
	// Its StartPoint will be adjusted by corner radius, EndPoint will be on rect edge.
	segment := generatePointPointSegment(point, createPoint(0, 0), direction, cornerRadius, number) // Dummy end, will be replaced

	// Overwrite StartPoint based on point and cornerRadius (as per TS)
	// The logic in TS drawPointRectSegment for c1_X_withCornerOffset / c1_Y_withCornerOffset
	// adjusts the segment's start point away from the raw 'point' (corner) by the cornerRadius.
	if direction == ORIENTATION_HORIZONTAL {
		// TS: segment.StartPoint.X calculation based on point.X relative to rect and cornerRadius
		tempStartX := point.X
		if point.X < rect.X+0.5*rect.Width { // Point is to the left of or near middle of rect
			if point.X < rect.X { // Point is fully to the left of rect
				tempStartX = point.X + cornerRadius
			} else { // Point is inside left half of rect
				tempStartX = point.X - cornerRadius // This seems to pull it back if inside
			}
		} else { // Point is to the right of or near middle of rect
			if point.X > rect.X+rect.Width { // Point is fully to the right of rect
				tempStartX = point.X - cornerRadius
			} else { // Point is inside right half of rect
				tempStartX = point.X + cornerRadius // This seems to push it further if inside
			}
		}
		segment.StartPoint = createPoint(tempStartX, point.Y)

		// Set EndPoint on the correct edge of the rect
		if point.X <= rect.X+rect.Width/2 {
			segment.EndPoint = createPoint(rect.X, point.Y) // Connect to left edge
		} else {
			segment.EndPoint = createPoint(rect.X+rect.Width, point.Y) // Connect to right edge
		}
	} else { // ORIENTATION_VERTICAL
		tempStartY := point.Y
		if point.Y < rect.Y+0.5*rect.Height {
			if point.Y < rect.Y {
				tempStartY = point.Y + cornerRadius
			} else {
				tempStartY = point.Y - cornerRadius
			}
		} else {
			if point.Y < rect.Y+rect.Height { // TS had point.Y < rect.Y + rect.Height (not >)
				tempStartY = point.Y + cornerRadius
			} else { // point.Y >= rect.Y + rect.Height
				tempStartY = point.Y - cornerRadius
			}
		}
		segment.StartPoint = createPoint(point.X, tempStartY)

		if point.Y <= rect.Y+rect.Height/2 {
			segment.EndPoint = createPoint(point.X, rect.Y) // Connect to top edge
		} else {
			segment.EndPoint = createPoint(point.X, rect.Y+rect.Height) // Connect to bottom edge
		}
	}

	// Set EndPointWithoutRadius to be the same as the calculated EndPoint (on the rect edge)
	segment.EndPointWithoutRadius = segment.EndPoint
	// StartPointWithoutRadius was already set by generatePointPointSegment to be the original 'point' (corner)

	if number == 0 { // If it's the first segment, swap its start and end
		segment = swapSegment(segment)
	}
	return segment
}

// Go equivalent of drawSegments from draw.segments.ts
func generateSegments(link *Link) []Segment { // Renamed with Go convention
	if link.Start == nil || link.End == nil {
		return []Segment{}
	}

	startRect := link.Start
	endRect := link.End
	startDirection := link.StartOrientation
	endDirection := link.EndOrientation
	startRatio := link.StartRatio
	endRatio := link.EndRatio
	cornerOffsetRatio := link.CornerOffsetRatio
	cornerRadius := link.CornerRadius

	segments := []Segment{}

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_VERTICAL {
		startY := startRect.Y + startRatio*startRect.Height
		c1X := endRect.X + endRatio*endRect.Width
		c1Y := startY
		c1 := createPoint(c1X, c1Y)
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0)
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1)
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_HORIZONTAL {
		startX := startRect.X + startRatio*startRect.Width
		c1X := startX
		c1Y := endRect.Y + endRatio*endRect.Height
		c1 := createPoint(c1X, c1Y)
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0)
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1)
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X := startRect.X + cornerOffsetRatio*startRect.Width
		c1Y := startRect.Y + startRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)

		c2X := c1X
		c2Y := endRect.Y + endRatio*endRect.Height
		c2 := createPoint(c2X, c2Y)

		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0)
		secondSegment := generatePointPointSegment(c1, c2, ORIENTATION_VERTICAL, cornerRadius, 1)
		thirdSegment := generatePointRectSegment(c2, endRect, endDirection, cornerRadius, 2)

		// Reduce second segment if start and end are aligned vertically
		if math.Abs(c1Y-c2Y) <= 2*cornerRadius && cornerRadius > 0 { // Added CR > 0 check
			c2Adjusted := createPoint(c2X, c1Y) // Align Y of c2 with c1
			firstSegment = generatePointRectSegment(c1, startRect, startDirection, 0, 0)
			secondSegment = generatePointPointSegment(c1, c2Adjusted, ORIENTATION_HORIZONTAL, 0, 1) // Middle segment becomes horizontal
			thirdSegment = generatePointRectSegment(c2Adjusted, endRect, endDirection, 0, 2)
		}
		segments = append(segments, firstSegment, secondSegment, thirdSegment)
	}

	if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_VERTICAL {
		c1X := startRect.X + startRatio*startRect.Width
		c1Y := startRect.Y + cornerOffsetRatio*startRect.Height
		c1 := createPoint(c1X, c1Y)

		c2X := endRect.X + endRatio*endRect.Width
		c2Y := c1Y
		c2 := createPoint(c2X, c2Y)

		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0)
		secondSegment := generatePointPointSegment(c1, c2, ORIENTATION_HORIZONTAL, cornerRadius, 1) // Middle is horizontal
		thirdSegment := generatePointRectSegment(c2, endRect, endDirection, cornerRadius, 2)

		// Reduce second segment if start and end are aligned horizontally
		if math.Abs(c1X-c2X) <= 2*cornerRadius && cornerRadius > 0 { // Added CR > 0 check
			c2Adjusted := createPoint(c1X, c2Y) // Align X of c2 with c1
			firstSegment = generatePointRectSegment(c1, startRect, startDirection, 0, 0)
			// TS uses ORIENTATION_HORIZONTAL for the middle segment in V-V reduction in draw.segments.ts line 99
			// but then it uses it again in line 115 for V-V.
			// The code at line 99 seems to be for H-H where middle is Vertical.
			// For V-V reduction, the middle segment should become Vertical if it's too short horizontally.
			// Let's check the TS code carefully:
			// In H-H: middle is Vertical. If c1.Y and c2.Y are close, middle becomes Horizontal.
			// In V-V: middle is Horizontal. If c1.X and c2.X are close, middle becomes Vertical.
			secondSegment = generatePointPointSegment(c1, c2Adjusted, ORIENTATION_VERTICAL, 0, 1) // Middle segment becomes vertical
			thirdSegment = generatePointRectSegment(c2Adjusted, endRect, endDirection, 0, 2)
		}
		segments = append(segments, firstSegment, secondSegment, thirdSegment)
	}
	return segments
}

// getAnchorPoint is primarily for Bezier curves or if a link doesn't use complex orthogonal routing.
// For orthogonal links, the detailed segment generation functions handle anchor points.
func getAnchorPoint(rect *Rect, anchorType AnchorType, orientation OrientationType, ratio float64) (x, y float64) {
	// This is a simplified interpretation for general anchor calculation.
	// The detailed logic in `generateSegments` and `generatePointRectSegment` is more specific for orthogonal links.
	switch orientation {
	case ORIENTATION_HORIZONTAL: // Exits/enters left or right face. Ratio applies along height.
		y = rect.Y + ratio*rect.Height
		// Default to center of the relevant edge if not further specified by AnchorType
		// For now, let's assume if it's HORIZONTAL, it primarily means the *direction* of the link segment
		// and the actual X connection point is chosen by other logic (e.g. mid-point of rect edge or based on relative position).
		// This function might be too simplistic for precise orthogonal start/end.
		// The `generateSegments` calculates precise start/end based on ratios for the *first bend point*.
		x = rect.X + rect.Width/2 // Fallback: center of rect (needs to be edge)
	case ORIENTATION_VERTICAL: // Exits/enters top or bottom face. Ratio applies along width.
		x = rect.X + ratio*rect.Width
		y = rect.Y + rect.Height/2 // Fallback: center of rect (needs to be edge)
	default:
		return rect.X + rect.Width/2, rect.Y + rect.Height/2
	}
	return x, y
}
