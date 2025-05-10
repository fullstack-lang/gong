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

	// Define marker for arrowhead if any link uses it
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
		if needsEndArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadend" markerWidth="10" markerHeight="7" refX="0" refY="3.5" orient="auto" fill="%s">
      <polygon points="0 0, 10 3.5, 0 7" />
    </marker>
`, "black")) // Default fill, can be customized or inherit from link stroke
		}
		if needsStartArrowMarker {
			sb.WriteString(fmt.Sprintf(`    <marker id="arrowheadstart" markerWidth="10" markerHeight="7" refX="10" refY="3.5" orient="auto" fill="%s">
      <polygon points="10 0, 0 3.5, 10 7" />
    </marker>
`, "black")) // Default fill
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
			appendPresentationAttributes(&sb, &polyline.Presentation, true)
			sb.WriteString(" />\n")
		}

		// Texts
		for _, text := range layer.Texts {
			sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(text.X), formatFloat(text.Y)))
			if text.FontSize != "" {
				sb.WriteString(fmt.Sprintf(` font-size="%s"`, text.FontSize))
			}
			if text.FontWeight != "" {
				sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(text.FontWeight)))
			}
			appendPresentationAttributes(&sb, &text.Presentation, true)
			sb.WriteString(">")
			sb.WriteString(html.EscapeString(text.Content))
			sb.WriteString("</text>\n")
		}

		// Links
		for _, link := range layer.Links {
			if link.Start == nil || link.End == nil {
				continue
			}

			var pathData strings.Builder

			if link.IsBezierCurve {
				startX, startY := getAnchorPoint(link.Start, link.StartAnchorType, link.StartOrientation, link.StartRatio)
				endX, endY := getAnchorPoint(link.End, link.EndAnchorType, link.EndOrientation, link.EndRatio)

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
				segments := generateSegments(link)
				if len(segments) == 0 {
					continue
				}

				pathData.WriteString(fmt.Sprintf("M %s %s", formatFloat(segments[0].StartPoint.X), formatFloat(segments[0].StartPoint.Y)))
				for i, seg := range segments {
					if i > 0 && link.CornerRadius > 0 {
						prevSeg := segments[i-1]
						if prevSeg.Orientation != seg.Orientation {
							// Arc rendering logic is complex.
							// The current segment points (seg.StartPoint, prevSeg.EndPoint) are already adjusted for radius.
							// An arc command would go from prevSeg.EndPoint to seg.StartPoint.
							// For simplicity, we are currently drawing lines to these adjusted points.
							// To implement arcs, careful calculation of large-arc-flag and sweep-flag is needed.
							// Example: pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(prevSeg.EndPoint.X), formatFloat(prevSeg.EndPoint.Y)))
							// arcStartX := prevSeg.EndPoint.X
							// arcStartY := prevSeg.EndPoint.Y
							arcEndX := seg.StartPoint.X
							arcEndY := seg.StartPoint.Y
							// sweepFlag := 1 // Placeholder, requires complex logic

							// Fallback to L connecting the pre-adjusted points of the segments
							pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(arcEndX), formatFloat(arcEndY)))
						}
					}
					pathData.WriteString(fmt.Sprintf(" L %s %s", formatFloat(seg.EndPoint.X), formatFloat(seg.EndPoint.Y)))
				}
			}

			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, pathData.String()))
			// Pass link.Presentation for styling, false for isFillable
			customLinkPresentation := link.Presentation
			if customLinkPresentation.Stroke == "" { // Ensure links have a stroke color if not specified
				customLinkPresentation.Stroke = "black" // Default stroke color for links
			}
			if customLinkPresentation.StrokeWidth == 0 { // Ensure links have a stroke width if not specified
				customLinkPresentation.StrokeWidth = 1 // Default stroke width
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
		}
	}

	sb.WriteString("</svg>\n")

	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func appendPresentationAttributes(sb *strings.Builder, p *Presentation, isFillable bool) {
	if isFillable {
		if p.Color != "" {
			sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
		}
		if p.FillOpacity != 1.0 && p.FillOpacity != 0 { // Default is 1, 0 might mean unset or transparent
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		}
	} else {
		sb.WriteString(` fill="none"`)
	}

	if p.Stroke != "" {
		sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
	}
	if p.StrokeWidth != 1.0 && p.StrokeWidth != 0 { // Default is 1, 0 might mean unset or thinnest
		sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
	}
	if p.StrokeDashArray != "" {
		sb.WriteString(fmt.Sprintf(` stroke-dasharray="%s"`, html.EscapeString(p.StrokeDashArray)))
	}
	if p.StrokeOpacity != 1.0 && p.StrokeOpacity != 0 { // Default is 1
		sb.WriteString(fmt.Sprintf(` stroke-opacity="%s"`, formatFloat(p.StrokeOpacity)))
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
	// This is a simplified interpretation based on typical orthogonal link start/end points.
	// It assumes ratio applies along the dimension perpendicular to the orientation.
	// E.g., Horizontal orientation exits from Left/Right face, ratio applies along Height.
	// A more robust version would use AnchorType for specific points (Center, N, S, E, W, etc.).
	switch orientation {
	case ORIENTATION_HORIZONTAL:
		y = rect.Y + ratio*rect.Height
		// Decide left or right edge based on which side the EndRect is, or a convention.
		// For simplicity, this needs a more concrete rule if not using explicit AnchorType values for edges.
		// Assuming a default: if no other info, connect from right edge if ratio > 0.5, else left. This is arbitrary.
		// A common approach: If StartRect is to the left of EndRect, exit from Right of StartRect.
		// This function is primarily for Bezier curves in this context, where precise edge choice might be simpler.
		// For orthogonal, generateSegments computes these implicitly.
		x = rect.X // Placeholder, should be rect.X or rect.X + rect.Width
	case ORIENTATION_VERTICAL:
		x = rect.X + ratio*rect.Width
		y = rect.Y // Placeholder, should be rect.Y or rect.Y + rect.Height
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
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0, true)
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1, false)
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_VERTICAL && endDirection == ORIENTATION_HORIZONTAL {
		startX := startRect.X + startRatio*startRect.Width
		c1Y := endRect.Y + endRatio*endRect.Height
		c1 := createPoint(startX, c1Y)
		firstSegment := generatePointRectSegment(c1, startRect, startDirection, cornerRadius, 0, true)
		secondSegment := generatePointRectSegment(c1, endRect, endDirection, cornerRadius, 1, false)
		segments = append(segments, firstSegment, secondSegment)
	}

	if startDirection == ORIENTATION_HORIZONTAL && endDirection == ORIENTATION_HORIZONTAL {
		c1X_ts := startRect.X + cornerOffsetRatio*startRect.Width
		c1Y_ts := startRect.Y + startRatio*startRect.Height
		c1_ts := createPoint(c1X_ts, c1Y_ts)

		c2X_ts := c1X_ts
		c2Y_ts := endRect.Y + endRatio*endRect.Height
		c2_ts := createPoint(c2X_ts, c2Y_ts)

		firstSegment := generatePointRectSegment(c1_ts, startRect, startDirection, cornerRadius, 0, true)
		secondSegment := generatePointPointSegment(c1_ts, c2_ts, ORIENTATION_VERTICAL, cornerRadius, 1)
		thirdSegment := generatePointRectSegment(c2_ts, endRect, endDirection, cornerRadius, 2, false)

		if math.Abs(c1Y_ts-c2Y_ts) <= 2*cornerRadius && cornerRadius > 0 {
			// Reduction logic for aligned segments (simplified: current implementation does not change segments yet)
			// For full correctness, segments should be recalculated with 0 radius and different orientation for middle segment.
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

		firstSegment := generatePointRectSegment(c1_ts, startRect, startDirection, cornerRadius, 0, true)
		secondSegment := generatePointPointSegment(c1_ts, c2_ts, ORIENTATION_HORIZONTAL, cornerRadius, 1)
		thirdSegment := generatePointRectSegment(c2_ts, endRect, endDirection, cornerRadius, 2, false)

		if math.Abs(c1X_ts-c2X_ts) <= 2*cornerRadius && cornerRadius > 0 {
			// Reduction logic (simplified)
		}
		segments = append(segments, firstSegment, secondSegment, thirdSegment)
	}
	return segments
}

func generatePointRectSegment(point Point, rect *Rect, direction OrientationType, cornerRadius float64, number int, isStartSegment bool) Segment {
	seg := Segment{
		StartPointWithoutRadius: point,
		Orientation:             direction,
		Number:                  number,
	}

	if direction == ORIENTATION_HORIZONTAL {
		seg.EndPointWithoutRadius.Y = point.Y
		// Determine connection to left or right edge of 'rect'
		// This logic needs to be robust for all StartRect/EndRect relative positions.
		// The TS code implies: if point.X <= rect.X + rect.Width / 2, connect to rect.X, else rect.X + rect.Width.
		// This is ambiguous. A clearer rule:
		// If isStartSegment, we are drawing from 'rect' to 'point'.
		// If !isStartSegment, we are drawing from 'point' to 'rect'.
		// The 'point' parameter is the corner of the bend.

		// Let's assume standard connection: if link starts horizontally, it exits from a vertical face.
		// If point.X is to the left of rect center, exit from right face of rect.
		// If point.X is to the right of rect center, exit from left face of rect.
		if isStartSegment { // Segment from rect to point
			seg.StartPointWithoutRadius.Y = point.Y
			if point.X < rect.X+rect.Width/2 { // Corner is to the left of rect's center, so exit rect from right.
				seg.StartPointWithoutRadius.X = rect.X + rect.Width
			} else { // Corner is to the right (or center) of rect's center, so exit rect from left.
				seg.StartPointWithoutRadius.X = rect.X
			}
			seg.EndPointWithoutRadius = point // The corner point is the end of this segment from the rect
		} else { // Segment from point to rect
			seg.StartPointWithoutRadius = point // The corner point is the start
			seg.EndPointWithoutRadius.Y = point.Y
			if point.X < rect.X+rect.Width/2 { // Corner is to the left of rect's center, so enter rect from right.
				seg.EndPointWithoutRadius.X = rect.X + rect.Width
			} else {
				seg.EndPointWithoutRadius.X = rect.X
			}
		}

		// Apply corner radius
		if seg.StartPointWithoutRadius.X < seg.EndPointWithoutRadius.X {
			seg.StartPoint.X = seg.StartPointWithoutRadius.X + cornerRadius
			seg.EndPoint.X = seg.EndPointWithoutRadius.X - cornerRadius
		} else {
			seg.StartPoint.X = seg.StartPointWithoutRadius.X - cornerRadius
			seg.EndPoint.X = seg.EndPointWithoutRadius.X + cornerRadius
		}
		seg.StartPoint.Y = seg.StartPointWithoutRadius.Y
		seg.EndPoint.Y = seg.EndPointWithoutRadius.Y

	} else { // ORIENTATION_VERTICAL
		// Similar logic for vertical segments
		if isStartSegment { // Segment from rect to point
			seg.StartPointWithoutRadius.X = point.X
			if point.Y < rect.Y+rect.Height/2 { // Corner is above rect's center, so exit rect from bottom.
				seg.StartPointWithoutRadius.Y = rect.Y + rect.Height
			} else { // Corner is below (or center) rect's center, so exit rect from top.
				seg.StartPointWithoutRadius.Y = rect.Y
			}
			seg.EndPointWithoutRadius = point
		} else { // Segment from point to rect
			seg.StartPointWithoutRadius = point
			seg.EndPointWithoutRadius.X = point.X
			if point.Y < rect.Y+rect.Height/2 { // Corner is above rect's center, so enter rect from bottom.
				seg.EndPointWithoutRadius.Y = rect.Y + rect.Height
			} else {
				seg.EndPointWithoutRadius.Y = rect.Y
			}
		}

		if seg.StartPointWithoutRadius.Y < seg.EndPointWithoutRadius.Y {
			seg.StartPoint.Y = seg.StartPointWithoutRadius.Y + cornerRadius
			seg.EndPoint.Y = seg.EndPointWithoutRadius.Y - cornerRadius
		} else {
			seg.StartPoint.Y = seg.StartPointWithoutRadius.Y - cornerRadius
			seg.EndPoint.Y = seg.EndPointWithoutRadius.Y + cornerRadius
		}
		seg.StartPoint.X = seg.StartPointWithoutRadius.X
		seg.EndPoint.X = seg.EndPointWithoutRadius.X
	}

	// The original TS `drawPointRectSegment` with `swapSegment` for number == 0:
	// If it's the first segment (number==0), it means it connects FROM a rect TO the corner `point`.
	// The logic above with `isStartSegment` should correctly set StartPointWithoutRadius on the rect
	// and EndPointWithoutRadius as the `point` (the corner).
	// The TS `swapSegment` inverted StartPoint and EndPoint if number == 0.
	// The current `isStartSegment` logic aims to establish the correct initial direction (rect to point).
	// The `generatePointRectSegment` in TS seems to assume drawing from `point` (corner) to `rect` (anchor face)
	// and then swaps if it's the first segment.
	// My `isStartSegment` logic directly tries to set Start on rect and End at corner for the first segment.
	// Let's re-verify the TS `drawPointRectSegment` and `swapSegment` intent.
	// TS: `drawPointRectSegment(c1, StartRect, ...)` --> c1 is corner, StartRect is target. Produces segment from c1 to StartRect.
	// Then `swapSegment` if number==0 makes it from StartRect to c1.
	// So, my `isStartSegment` is intended to achieve this directly:
	// If `isStartSegment` is true: segment is from `rect` to `point`.
	//   So, `seg.StartPointWithoutRadius` should be on `rect`.
	//   And `seg.EndPointWithoutRadius` should be `point`.
	// This is what the `if isStartSegment` block above does.

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
