package models

import (
	"fmt"
	"html" // For escaping text and attribute values
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
			rectPresentation := rect.Presentation
			if rectPresentation.Color == "" {
				rectPresentation.Color = "white"
			}
			if rectPresentation.Stroke == "" {
				rectPresentation.Stroke = "black"
			}
			if rectPresentation.StrokeWidth == 0 && (rectPresentation.Stroke != "" && strings.ToLower(rectPresentation.Stroke) != "none") {
				rectPresentation.StrokeWidth = 1
			}

			sb.WriteString(fmt.Sprintf(`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
				formatFloat(rect.X), formatFloat(rect.Y), formatFloat(rect.Width), formatFloat(rect.Height),
				formatFloat(rect.RX), formatFloat(rect.RX)))
			appendPresentationAttributes(&sb, &rectPresentation, true)
			sb.WriteString(" />\n")

			// Process RectAnchoredTexts for this rect
			for _, anchoredText := range rect.RectAnchoredTexts {
				anchorX, anchorY := getRectAnchorPoint(rect, anchoredText.RectAnchorType)
				textX := anchorX + anchoredText.X_Offset
				textY := anchorY + anchoredText.Y_Offset

				generateTextElement(
					&sb,
					anchoredText.Content,
					textX, textY,
					anchoredText.FontSize,
					anchoredText.FontWeight,
					anchoredText.FontStyle,
					string(anchoredText.TextAnchorType), // RectAnchoredText has TextAnchorType
					&anchoredText.Presentation,
				)
			}
		}

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
	}

	sb.WriteString("</svg>\n")
	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
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

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func appendPresentationAttributes(sb *strings.Builder, p *Presentation, isFillable bool) {
	if isFillable {
		if p.Color != "" {
			if strings.ToLower(p.Color) == "none" {
				sb.WriteString(` fill="none"`)
			} else {
				sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
			}
		}
		if p.FillOpacity != 1.0 && p.FillOpacity != 0.0 {
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		} else if p.FillOpacity == 0.0 && (p.Color != "" && strings.ToLower(p.Color) != "none") {
			sb.WriteString(` fill-opacity="0"`)
		}
	} else {
		sb.WriteString(` fill="none"`)
	}

	if p.Stroke != "" {
		if strings.ToLower(p.Stroke) == "none" {
			sb.WriteString(` stroke="none"`)
		} else {
			sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
		}
	}

	if p.Stroke != "" && strings.ToLower(p.Stroke) != "none" {
		if p.StrokeWidth != 1.0 {
			sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
		}
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
