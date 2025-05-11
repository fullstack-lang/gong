package models

import (

	// For escaping text and attribute values
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateMaxx(maxX_, maxY_ float64, maxX, maxY *float64) {
	if maxX_ > *maxX {
		*maxX = maxX_
	}
	if maxY_ > *maxY {
		*maxY = maxY_
	}
}

// GenerateFile generates an SVG file that represents the content of the SVG object.
func (svg *SVG) GenerateFile(pathToFile string) (err error) {
	var sb strings.Builder

	maxX_string_to_be_replaced := "__gong__maxX_string_to_be_replaced"
	maxY_string_to_be_replaced := "__gong__maxY_string_to_be_replaced"

	sb.WriteString(
		fmt.Sprintf(`
<svg xmlns="http://www.w3.org/2000/svg" width="100%%" height="100%%" viewBox="-10 -10 %s %s">
`,
			maxX_string_to_be_replaced,
			maxY_string_to_be_replaced))
	sb.WriteString("<style>text { font-family: Roboto, Arial, sans-serif !important; }</style>")

	maxX, maxY := 0.0, 0.0

	for _, layer := range svg.Layers {
		// Rects
		for _, rect := range layer.Rects {

			maxX_, maxY_ := rect.WriteSVG(&sb)
			updateMaxx(maxX_, maxY_, &maxX, &maxY)
		}

		for _, link := range layer.Links {
			if link.Start == nil || link.End == nil {
				continue
			}

			segments := link.generateSegments()

			for idx, segment := range segments {
				maxX_, maxY_ := segment.WriteSVG(&sb, link)
				updateMaxx(maxX_, maxY_, &maxX, &maxY)

				// draw the end arrow
				if segment.Type == EndSegment && link.HasEndArrow {
					link.WriteSVGEndArrow(&sb, &segment)
				}

				// draw the start arrow (yes, that's strange, why start and end ?)
				if segment.Type == StartSegment && link.HasStartArrow {
					swapedSegment := swapSegment(segment)
					link.WriteSVGEndArrow(&sb, &swapedSegment)
				}

				// draw arcs between segment
				if idx < len(segments)-1 {
					link.WriteSVGArcPath(&sb, &segment, &segments[idx+1])
				}

				// draw the text at the end of the link
				if segment.Type == EndSegment {
					for _, linkAnchoredText := range link.TextAtArrowEnd {
						linkAnchoredText.WriteSVG(&sb, link, &segment)
					}
				}
			}

			_ = segments
		}

		/*
			// Circles
			for _, circle := range layer.Circles {
				circlePresentation := circle.Presentation
				sb.WriteString(fmt.Sprintf(`  <circle cx="%s" cy="%s" r="%s"`,
					formatFloat(circle.CX), formatFloat(circle.CY), formatFloat(circle.Radius)))
				appendPresentationAttributes(&sb, &circlePresentation, true)
				sb.WriteString(" />\n")
			}

			// Lines
			for _, line := range layer.Lines {
				linePresentation := line.Presentation
				if linePresentation.Stroke == "" {
					linePresentation.Stroke = "black"
				}
				if linePresentation.StrokeWidth == 0 && (linePresentation.Stroke != "" && strings.ToLower(linePresentation.Stroke) != "none") {
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
				sb.WriteString(fmt.Sprintf(`  <ellipse cx="%s" cy="%s" rx="%s" ry="%s"`,
					formatFloat(ellipse.CX), formatFloat(ellipse.CY), formatFloat(ellipse.RX), formatFloat(ellipse.RY)))
				appendPresentationAttributes(&sb, &ellipsePresentation, true)
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
				polylinePresentation := polyline.Presentation
				sb.WriteString(fmt.Sprintf(`  <polyline points="%s"`, html.EscapeString(polyline.Points)))
				appendPresentationAttributes(&sb, &polylinePresentation, true)
				sb.WriteString(" />\n")
			}

			// Texts (general texts, not anchored to rects or links)
			for _, text := range layer.Texts {
				// models.Text has FontSize (string) and FontWeight (string). No FontStyle, No TextAnchorType.
				fontSize := text.FontSize // This is a string
				if fontSize != "" {
					// Check if it's a plain number string, if so, append "px"
					if _, errConv := strconv.Atoi(fontSize); errConv == nil {
						fontSize += "px"
					}
				}
				generateTextElement(&sb, text.Content, text.X, text.Y, fontSize, text.FontWeight, "", "", &text.Presentation)
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
				if customLinkPresentation.StrokeWidth == 0 && (customLinkPresentation.Stroke != "" && strings.ToLower(customLinkPresentation.Stroke) != "none") {
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
					for _, anchoredText := range link.TextAtArrowStart { // anchoredText is *LinkAnchoredText
						textX := anchorXStart + anchoredText.X_Offset
						textY := anchorYStart + anchoredText.Y_Offset

						// LinkAnchoredText embeds TextAttributes: FontSize (string), FontWeight (string), FontStyle (string)
						fontSizeStrLAT := anchoredText.FontSize // Use string directly from TextAttributes
						if fontSizeStrLAT != "" {
							// Append "px" if it's a plain number string
							if _, errConv := strconv.Atoi(fontSizeStrLAT); errConv == nil {
								fontSizeStrLAT += "px"
							}
						}
						// LinkAnchoredText does not have TextAnchorType. Default to "middle".
						generateTextElement(&sb, anchoredText.Content, textX, textY,
							fontSizeStrLAT, anchoredText.FontWeight, anchoredText.FontStyle, // FontStyle is available via TextAttributes
							string(TEXT_ANCHOR_CENTER), // Default for link anchored text (models.TEXT_ANCHOR_MIDDLE)
							&anchoredText.Presentation)
					}

					lastSegment := segments[len(segments)-1]
					anchorXEnd := lastSegment.EndPoint.X
					anchorYEnd := lastSegment.EndPoint.Y
					for _, anchoredText := range link.TextAtArrowEnd { // anchoredText is *LinkAnchoredText
						textX := anchorXEnd + anchoredText.X_Offset
						textY := anchorYEnd + anchoredText.Y_Offset

						fontSizeStrLAT := anchoredText.FontSize // Use string directly from TextAttributes
						if fontSizeStrLAT != "" {
							// Append "px" if it's a plain number string
							if _, errConv := strconv.Atoi(fontSizeStrLAT); errConv == nil {
								fontSizeStrLAT += "px"
							}
						}
						generateTextElement(&sb, anchoredText.Content, textX, textY,
							fontSizeStrLAT, anchoredText.FontWeight, anchoredText.FontStyle, // FontStyle is available via TextAttributes
							string(TEXT_ANCHOR_CENTER), // Default for link anchored text (models.TEXT_ANCHOR_CENTER)
							&anchoredText.Presentation)
					}
				}
			}
		*/
	}

	sb.WriteString("</svg>\n")

	result := sb.String()

	result = strings.ReplaceAll(result, maxX_string_to_be_replaced, fmt.Sprintf("%f", maxX+10))
	result = strings.ReplaceAll(result, maxY_string_to_be_replaced, fmt.Sprintf("%f", maxY+10))

	return os.WriteFile(pathToFile, []byte(result), 0644)
}

func getRectAnchorPoint(rect *Rect, anchorType RectAnchorType) (x, y float64) {
	switch anchorType {
	case RECT_TOP_LEFT:
		return rect.X, rect.Y
	case RECT_TOP:
		return rect.X + rect.Width/2, rect.Y
	case RECT_TOP_RIGHT:
		return rect.X + rect.Width, rect.Y
	case RECT_LEFT:
		return rect.X, rect.Y + rect.Height/2
	case RECT_CENTER:
		return rect.X + rect.Width/2, rect.Y + rect.Height/2
	case RECT_RIGHT:
		return rect.X + rect.Width, rect.Y + rect.Height/2
	case RECT_BOTTOM_LEFT:
		return rect.X, rect.Y + rect.Height
	case RECT_BOTTOM:
		return rect.X + rect.Width/2, rect.Y + rect.Height
	case RECT_BOTTOM_RIGHT:
		return rect.X + rect.Width, rect.Y + rect.Height
	default:
		return rect.X + rect.Width/2, rect.Y + rect.Height/2
	}
}

func getRectAnchorPointXY(x_in, y_in, Width, Height float64, anchorType RectAnchorType) (x, y float64) {
	switch anchorType {
	case RECT_TOP_LEFT:
		return x_in, y_in
	case RECT_TOP:
		return x_in + Width/2, y_in
	case RECT_TOP_RIGHT:
		return x_in + Width, y_in
	case RECT_LEFT:
		return x_in, y_in + Height/2
	case RECT_CENTER:
		return x_in + Width/2, y_in + Height/2
	case RECT_RIGHT:
		return x_in + Width, y_in + Height/2
	case RECT_BOTTOM_LEFT:
		return x_in, y_in + Height
	case RECT_BOTTOM:
		return x_in + Width/2, y_in + Height
	case RECT_BOTTOM_RIGHT:
		return x_in + Width, y_in + Height
	default:
		return x_in + Width/2, y_in + Height/2
	}
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
