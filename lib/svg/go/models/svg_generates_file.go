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

	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\">\n")

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
		// SVG 1.1 markers cannot directly inherit color from the referencing element's stroke.
		// A common workaround is to set marker's fill to "context-stroke" or "currentColor",
		// but browser support varies. For simplicity, we use a fixed color or pass it.
		// Another option is to create multiple markers for different colors if needed.
		if needsEndArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadend" markerWidth="10" markerHeight="7" refX="0" refY="3.5" orient="auto" fill="black">
      <polygon points="0 0, 10 3.5, 0 7" />
    </marker>
`))
		}
		if needsStartArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadstart" markerWidth="10" markerHeight="7" refX="10" refY="3.5" orient="auto" fill="black">
      <polygon points="10 0, 0 3.5, 10 7" />
    </marker>
`))
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

		// Paths
		for _, path := range layer.Paths {
			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, html.EscapeString(path.Definition)))
			appendPresentationAttributes(&sb, &path.Presentation, true)
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
			appendPresentationAttributes(&sb, &polyline.Presentation, true) // Polylines can be filled if closed or path-like
			sb.WriteString(" />\n")
		}

		// Texts
		for _, text := range layer.Texts {
			// Use a separate function to generate text elements for DRY principle
			generateTextElement(&sb, text.Content, text.X, text.Y, "", "", &text.Presentation)
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

				// For Bezier, "segments" aren't explicitly generated in the same way,
				// but we can define start and end points for text anchoring.
				// This is a simplification; text on Bezier curves might need path-based offsetting.
				segments = []Segment{ // Simplified segment representation for Bezier
					{StartPoint: createPoint(startX, startY), EndPoint: createPoint(endX, endY)},
				}
				if len(link.ControlPoints) > 0 {
					// More accurately, the final point of the Bezier is EndPoint.
					// If ControlPoints define the curve, the visual end might be different from just (endX, endY) if those are only anchors.
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
				segments = generateSegments(link)
				if len(segments) == 0 {
					continue
				}

				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(segments[0].StartPoint.X), formatFloat(segments[0].StartPoint.Y)))
				for i, seg := range segments {
					if i > 0 && link.CornerRadius > 0 {
						prevSeg := segments[i-1]
						if prevSeg.Orientation != seg.Orientation {
							arcEndX := seg.StartPoint.X
							arcEndY := seg.StartPoint.Y
							pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(arcEndX), formatFloat(arcEndY)))
						}
					}
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
			if len(segments) > 0 { // Ensure segments exist for anchoring
				firstSegment := segments[0]
				anchorX := firstSegment.StartPoint.X
				anchorY := firstSegment.StartPoint.Y
				// TODO: Adjust anchorX, anchorY based on PositionOnArrow if HasStartArrow is true
				// and the arrow marker's size/offset. This requires knowing arrow dimensions.
				// For now, simple offset from segment start.

				for _, anchoredText := range link.TextAtArrowStart {
					textX := anchorX + anchoredText.X_Offset
					textY := anchorY + anchoredText.Y_Offset
					// Use anchoredText's own presentation for styling
					generateTextElement(&sb, anchoredText.Content, textX, textY, "", "", &anchoredText.Presentation)
				}
			}

			// Process TextAtArrowEnd
			if len(segments) > 0 { // Ensure segments exist for anchoring
				lastSegment := segments[len(segments)-1]
				anchorX := lastSegment.EndPoint.X
				anchorY := lastSegment.EndPoint.Y
				// TODO: Adjust anchorX, anchorY based on PositionOnArrow if HasEndArrow is true
				// and the arrow marker's size/offset.

				for _, anchoredText := range link.TextAtArrowEnd {
					textX := anchorX + anchoredText.X_Offset
					textY := anchorY + anchoredText.Y_Offset
					generateTextElement(&sb, anchoredText.Content, textX, textY, "", "", &anchoredText.Presentation)
				}
			}
		}
	}

	sb.WriteString("</svg>\n")

	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
}

// generateTextElement is a helper to create SVG <text>...</text> string.
// Pass empty string for fontSize or fontWeight if not needed or to use defaults from presentation.
func generateTextElement(sb *strings.Builder, content string, x, y float64, fontSize, fontWeight string, p *Presentation) {
	sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(x), formatFloat(y)))

	// Font attributes from Presentation if not overridden by direct params
	// (Assuming TextSpecificAttributes like FontSize and FontWeight might be part of Presentation or a sub-struct)
	// For now, direct params take precedence if provided.
	// The `Text` struct has FontSize and FontWeight directly. `LinkAnchoredText` uses Presentation.
	// We need to decide if LinkAnchoredText.Presentation can carry font info or if we need dedicated fields.
	// Let's assume for now that LinkAnchoredText.Presentation might contain fill (for color),
	// but font size/weight might need to be sourced differently or have defaults.
	// If Presentation is expected to have all font info, this logic would change.

	// Example: if p has font size/weight fields (e.g., p.FontSize)
	// if fontSize == "" && p.FontSize != "" { fontSize = p.FontSize }
	// if fontWeight == "" && p.FontWeight != "" { fontWeight = p.FontWeight }

	if fontSize != "" { // This implies direct font size for LinkAnchoredText is not in its Presentation
		sb.WriteString(fmt.Sprintf(` font-size="%s"`, fontSize))
	} else if fontSize != "" { // Check if Presentation struct has it
		sb.WriteString(fmt.Sprintf(` font-size="%s"`, html.EscapeString(fontSize)))
	}

	if fontWeight != "" { // Same for font weight
		sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(fontWeight)))
	} else if fontWeight != "" {
		sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(fontWeight)))
	}

	// Text anchor and dominant baseline can be important for precise positioning.
	// Default "text-anchor" is "start". Default "dominant-baseline" is "auto".
	// These could be added to LinkAnchoredText or its Presentation if needed.
	// sb.WriteString(` text-anchor="middle" dominant-baseline="middle"`) // Example for centering

	clonedPres := *p            // Clone presentation to avoid modifying original link presentation
	if clonedPres.Color == "" { // Default text color
		clonedPres.Color = "black"
	}
	appendPresentationAttributes(sb, &clonedPres, true) // true: text is fillable

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
		// SVG default fill-opacity is 1. Only write if different from 1, and not 0 (unless 0 means fully transparent and not "unset")
		if p.FillOpacity != 1.0 && p.FillOpacity != 0.0 {
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		} else if p.FillOpacity == 0.0 && p.Color != "" { // If color is set and opacity is 0, it's explicitly transparent
			sb.WriteString(` fill-opacity="0"`)
		}
	} else {
		sb.WriteString(` fill="none"`)
	}

	if p.Stroke != "" {
		sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
	}
	// SVG default stroke-width is 1.
	if p.StrokeWidth != 1.0 && p.StrokeWidth != 0.0 {
		sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
	} else if p.StrokeWidth == 0.0 && p.Stroke != "" { // If stroke is set and width is 0, it's explicitly no width
		sb.WriteString(` stroke-width="0"`)
	}

	if p.StrokeDashArray != "" {
		sb.WriteString(fmt.Sprintf(` stroke-dasharray="%s"`, html.EscapeString(p.StrokeDashArray)))
	}
	// SVG default stroke-opacity is 1.
	if p.StrokeOpacity != 1.0 && p.StrokeOpacity != 0.0 {
		sb.WriteString(fmt.Sprintf(` stroke-opacity="%s"`, formatFloat(p.StrokeOpacity)))
	} else if p.StrokeOpacity == 0.0 && p.Stroke != "" { // If stroke is set and opacity is 0, explicitly transparent
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

func getAnchorPoint(rect *Rect, anchorType AnchorType, orientation OrientationType, ratio float64) (x, y float64) {
	switch orientation {
	case ORIENTATION_HORIZONTAL:
		y = rect.Y + ratio*rect.Height
		// This is a placeholder. Actual X should be rect.X or rect.X + rect.Width
		// depending on which side the link connects (e.g., based on relative position of EndRect).
		// For a link starting horizontally from the right face: rect.X + rect.Width
		// For a link starting horizontally from the left face: rect.X
		// A common default might be to pick the side "facing" the other linked element.
		// If StartRect is to the left of EndRect, pick StartRect's right face.
		// If StartRect is to the right of EndRect, pick StartRect's left face.
		// This detail is crucial for correct Bezier anchor points.
		// For now, defaulting to left edge.
		x = rect.X
		// A more robust approach might be:
		// if link.End != nil && rect.X + rect.Width/2 < link.End.X + link.End.Width/2 {
		// 	x = rect.X + rect.Width // StartRect is to the left, exit right
		// } else {
		// 	x = rect.X // StartRect is to the right (or aligned), exit left
		// }

	case ORIENTATION_VERTICAL:
		x = rect.X + ratio*rect.Width
		y = rect.Y // Placeholder, should be rect.Y or rect.Y + rect.Height
		// Similar logic for top/bottom face based on relative Y position of EndRect
		// if link.End != nil && rect.Y + rect.Height/2 < link.End.Y + link.End.Height/2 {
		// 	y = rect.Y + rect.Height // StartRect is above, exit bottom
		// } else {
		// 	y = rect.Y // StartRect is below (or aligned), exit top
		// }
	default:
		return rect.X + rect.Width/2, rect.Y + rect.Height/2
	}
	return x, y
}

func generateSegments(link *Link) []Segment {
	if link.Start == nil || link.End == nil {
		return []Segment{}
	}

	segments := []Segment{}
	startRect := link.Start
	endRect := link.End
	startDirection := link.StartOrientation
	endDirection := link.EndOrientation
	startRatio := link.StartRatio
	endRatio := link.EndRatio
	cornerRadius := link.CornerRadius
	cornerOffsetRatio := link.CornerOffsetRatio

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_VERTICAL {
		startY := startRect.Y + startRatio*startRect.Height
		c1X := endRect.X + endRatio*endRect.Width
		c1 := createPoint(c1X, startY)
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0, true, endRect) // Pass endRect for context
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1, false, startRect) // Pass startRect for context
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_HORIZONTAL {
		startX := startRect.X + startRatio*startRect.Width
		c1Y := endRect.Y + endRatio*endRect.Height
		c1 := createPoint(startX, c1Y)
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0, true, endRect)
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1, false, startRect)
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X_ts := startRect.X + cornerOffsetRatio*startRect.Width
		c1Y_ts := startRect.Y + startRatio*startRect.Height
		c1_ts := createPoint(c1X_ts, c1Y_ts)

		c2X_ts := c1X_ts
		c2Y_ts := endRect.Y + endRatio*endRect.Height
		c2_ts := createPoint(c2X_ts, c2Y_ts)

		firstSegment := generatePointRectSegment(c1_ts, startRect, startDirection, cornerRadius, 0, true, endRect)
		secondSegment := generatePointPointSegment(c1_ts, c2_ts, ORIENTATION_VERTICAL, cornerRadius, 1)
		thirdSegment := generatePointRectSegment(c2_ts, endRect, endDirection, cornerRadius, 2, false, startRect)

		if math.Abs(c1Y_ts-c2Y_ts) <= 2*cornerRadius && cornerRadius > 0 {
			// Reduction logic could be implemented here by re-calculating segments
		}
		segments = append(segments, firstSegment, secondSegment, thirdSegment)
	}

	if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_VERTICAL {
		c1Y_ts := startRect.Y + cornerOffsetRatio*startRect.Height
		c1X_ts := startRect.X + startRatio*startRect.Width
		c1_ts := createPoint(c1X_ts, c1Y_ts)

		c2Y_ts := c1Y_ts
		c2X_ts := endRect.X + endRatio*endRect.Width
		c2_ts := createPoint(c2X_ts, c2Y_ts)

		firstSegment := generatePointRectSegment(c1_ts, startRect, startDirection, cornerRadius, 0, true, endRect)
		secondSegment := generatePointPointSegment(c1_ts, c2_ts, ORIENTATION_HORIZONTAL, cornerRadius, 1)
		thirdSegment := generatePointRectSegment(c2_ts, endRect, endDirection, cornerRadius, 2, false, startRect)

		if math.Abs(c1X_ts-c2X_ts) <= 2*cornerRadius && cornerRadius > 0 {
			// Reduction logic
		}
		segments = append(segments, firstSegment, secondSegment, thirdSegment)
	}
	return segments
}

// Added otherRect context to help decide connection side more robustly
func generatePointRectSegment(point Point, rect *Rect, direction OrientationType, cornerRadius float64, number int, isStartSegment bool, otherRect *Rect) Segment {
	seg := Segment{
		// StartPointWithoutRadius will be set based on isStartSegment
		Orientation: direction,
		Number:      number,
	}

	if direction == ORIENTATION_HORIZONTAL {
		if isStartSegment { // Segment from rect to point (corner)
			seg.StartPointWithoutRadius.Y = rect.Y + (rect.Height * 0.5) // Default to middle of rect face, can be link.StartRatio
			// Determine exit side of 'rect' based on 'otherRect' (which is link.End Rect)
			if rect.X+rect.Width < otherRect.X { // rect is to the left of otherRect, exit right
				seg.StartPointWithoutRadius.X = rect.X + rect.Width
			} else if rect.X > otherRect.X+otherRect.Width { // rect is to the right of otherRect, exit left
				seg.StartPointWithoutRadius.X = rect.X
			} else { // Rects are somewhat aligned vertically, or one contains the other X-wise
				// Default: if point (corner) is to the right of rect center, exit left, else exit right.
				if point.X > rect.X+rect.Width/2 {
					seg.StartPointWithoutRadius.X = rect.X
				} else {
					seg.StartPointWithoutRadius.X = rect.X + rect.Width
				}
			}
			seg.EndPointWithoutRadius = point // The corner point
		} else { // Segment from point (corner) to rect (target)
			seg.StartPointWithoutRadius = point   // The corner point
			seg.EndPointWithoutRadius.Y = point.Y // Y is same as corner for horizontal segment
			// Determine entry side of 'rect' based on 'otherRect' (which is link.Start Rect)
			if rect.X > otherRect.X+otherRect.Width { // rect is to the right of otherRect, enter left
				seg.EndPointWithoutRadius.X = rect.X
			} else if rect.X+rect.Width < otherRect.X { // rect is to the left of otherRect, enter right
				seg.EndPointWithoutRadius.X = rect.X + rect.Width
			} else {
				if point.X > rect.X+rect.Width/2 {
					seg.EndPointWithoutRadius.X = rect.X // Enter left if corner is to the right
				} else {
					seg.EndPointWithoutRadius.X = rect.X + rect.Width // Enter right if corner is to the left
				}
			}
		}

		// Apply corner radius adjustments
		seg.StartPoint.Y = seg.StartPointWithoutRadius.Y
		seg.EndPoint.Y = seg.EndPointWithoutRadius.Y
		if seg.StartPointWithoutRadius.X < seg.EndPointWithoutRadius.X { // Left to Right
			seg.StartPoint.X = seg.StartPointWithoutRadius.X + cornerRadius
			seg.EndPoint.X = seg.EndPointWithoutRadius.X - cornerRadius
		} else { // Right to Left
			seg.StartPoint.X = seg.StartPointWithoutRadius.X - cornerRadius
			seg.EndPoint.X = seg.EndPointWithoutRadius.X + cornerRadius
		}

	} else { // ORIENTATION_VERTICAL
		if isStartSegment { // Segment from rect to point (corner)
			seg.StartPointWithoutRadius.X = rect.X + (rect.Width * 0.5) // Default to middle of rect face
			if rect.Y+rect.Height < otherRect.Y {                       // rect is above otherRect, exit bottom
				seg.StartPointWithoutRadius.Y = rect.Y + rect.Height
			} else if rect.Y > otherRect.Y+otherRect.Height { // rect is below otherRect, exit top
				seg.StartPointWithoutRadius.Y = rect.Y
			} else {
				if point.Y > rect.Y+rect.Height/2 {
					seg.StartPointWithoutRadius.Y = rect.Y
				} else {
					seg.StartPointWithoutRadius.Y = rect.Y + rect.Height
				}
			}
			seg.EndPointWithoutRadius = point
		} else { // Segment from point (corner) to rect (target)
			seg.StartPointWithoutRadius = point
			seg.EndPointWithoutRadius.X = point.X
			if rect.Y > otherRect.Y+otherRect.Height { // rect is below otherRect, enter top
				seg.EndPointWithoutRadius.Y = rect.Y
			} else if rect.Y+rect.Height < otherRect.Y { // rect is above otherRect, enter bottom
				seg.EndPointWithoutRadius.Y = rect.Y + rect.Height
			} else {
				if point.Y > rect.Y+rect.Height/2 {
					seg.EndPointWithoutRadius.Y = rect.Y // Enter top if corner is below
				} else {
					seg.EndPointWithoutRadius.Y = rect.Y + rect.Height // Enter bottom if corner is above
				}
			}
		}

		seg.StartPoint.X = seg.StartPointWithoutRadius.X
		seg.EndPoint.X = seg.EndPointWithoutRadius.X
		if seg.StartPointWithoutRadius.Y < seg.EndPointWithoutRadius.Y { // Top to Bottom
			seg.StartPoint.Y = seg.StartPointWithoutRadius.Y + cornerRadius
			seg.EndPoint.Y = seg.EndPointWithoutRadius.Y - cornerRadius
		} else { // Bottom to Top
			seg.StartPoint.Y = seg.StartPointWithoutRadius.Y - cornerRadius
			seg.EndPoint.Y = seg.EndPointWithoutRadius.Y + cornerRadius
		}
	}
	return seg
}

func generatePointPointSegment(start Point, end Point, orientation OrientationType, cornerRadius float64, number int) Segment {
	newStart := start
	newEnd := end

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
		StartPoint:              newStart,
		EndPoint:                newEnd,
		StartPointWithoutRadius: start,
		EndPointWithoutRadius:   end,
		Orientation:             orientation,
		Number:                  number,
	}
}
