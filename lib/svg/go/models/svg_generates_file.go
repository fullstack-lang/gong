package models

import (
	"fmt"
	"html" // For escaping text and attribute values
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

	// SVG header
	// Note: The gongsvg.SVG model does not currently have Width and Height fields for the root <svg> tag.
	// If specific dimensions are required for the SVG viewport, these fields could be added to the SVG struct
	// (e.g., Width float64, Height float64) and then used here:
	// if svg.Width > 0 && svg.Height > 0 {
	// 	sb.WriteString(fmt.Sprintf("<svg width=\"%s\" height=\"%s\" xmlns=\"http://www.w3.org/2000/svg\">\n",
	// 		formatFloat(svg.Width), formatFloat(svg.Height)))
	// } else {
	// 	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\">\n")
	// }
	// For now, omitting width and height attributes from the <svg> root tag.
	// This makes the generated SVG scalable by default, its size determined by the container it's viewed in.
	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\">\n")

	for _, layer := range svg.Layers {

		// Rects
		for _, rect := range layer.Rects {
			sb.WriteString(fmt.Sprintf(`  <rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s"`,
				formatFloat(rect.X), formatFloat(rect.Y), formatFloat(rect.Width), formatFloat(rect.Height),
				formatFloat(rect.RX), formatFloat(rect.RX))) // SVG spec allows ry to be omitted if same as rx
			appendPresentationAttributes(&sb, &rect.Presentation, true) // true: isFillable
			sb.WriteString(" />\n")
		}

		// Circles
		for _, circle := range layer.Circles {
			sb.WriteString(fmt.Sprintf(`  <circle cx="%s" cy="%s" r="%s"`,
				formatFloat(circle.CX), formatFloat(circle.CY), formatFloat(circle.Radius)))
			appendPresentationAttributes(&sb, &circle.Presentation, true)
			sb.WriteString(" />\n") // Assuming Circle does not have Animates field
		}

		// Lines
		for _, line := range layer.Lines {
			sb.WriteString(fmt.Sprintf(`  <line x1="%s" y1="%s" x2="%s" y2="%s"`,
				formatFloat(line.X1), formatFloat(line.Y1), formatFloat(line.X2), formatFloat(line.Y2)))
			appendPresentationAttributes(&sb, &line.Presentation, false) // false: not fillable by default
			sb.WriteString(" />\n")                                      // Assuming Line does not have Animates field
		}

		// Ellipses
		for _, ellipse := range layer.Ellipses {
			sb.WriteString(fmt.Sprintf(`  <ellipse cx="%s" cy="%s" rx="%s" ry="%s"`,
				formatFloat(ellipse.CX), formatFloat(ellipse.CY), formatFloat(ellipse.RX), formatFloat(ellipse.RY)))
			appendPresentationAttributes(&sb, &ellipse.Presentation, true)
			sb.WriteString(" />\n") // Assuming Ellipse does not have Animates field
		}

		// Paths
		for _, path := range layer.Paths {
			sb.WriteString(fmt.Sprintf(`  <path d="%s"`, html.EscapeString(path.Definition))) // Path data 'd' can be complex
			appendPresentationAttributes(&sb, &path.Presentation, true)
			sb.WriteString(" />\n") // Assuming Path does not have Animates field
		}

		// Polygones (Note: model name is Polygone, SVG tag is <polygon>)
		for _, polygone := range layer.Polygones {
			sb.WriteString(fmt.Sprintf(`  <polygon points="%s"`, html.EscapeString(polygone.Points)))
			appendPresentationAttributes(&sb, &polygone.Presentation, true)
			sb.WriteString(" />\n") // Assuming Polygone does not have Animates field
		}

		// Polylines
		for _, polyline := range layer.Polylines {
			sb.WriteString(fmt.Sprintf(`  <polyline points="%s"`, html.EscapeString(polyline.Points)))
			appendPresentationAttributes(&sb, &polyline.Presentation, true) // Polylines can be filled
			sb.WriteString(" />\n")                                         // Assuming Polyline does not have Animates field
		}

		// Texts
		// Note: The gongsvg.Text model currently does not have an Animates field.
		for _, text := range layer.Texts {
			sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, formatFloat(text.X), formatFloat(text.Y)))
			if text.FontSize != "" { // Assuming 0 means not set or use CSS default
				sb.WriteString(fmt.Sprintf(` font-size="%s"`, text.FontSize))
			}
			if text.FontWeight != "" {
				sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(text.FontWeight)))
			}
			// Other text-specific attributes like text-anchor, dominant-baseline are not in the current model.
			appendPresentationAttributes(&sb, &text.Presentation, true)
			sb.WriteString(">")
			sb.WriteString(html.EscapeString(text.Content))
			sb.WriteString("</text>\n")
		}
	}

	sb.WriteString("</svg>\n")

	return os.WriteFile(pathToFile, []byte(sb.String()), 0644)
}

// formatFloat converts a float64 to its string representation using 'g' format.
// This format uses the shortest representation (decimal or scientific) and removes trailing zeros.
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

// appendPresentationAttributes appends SVG presentation attributes (fill, stroke, transform, etc.)
// to the strings.Builder.
// `isFillable` indicates if the element type typically supports fill attributes.
func appendPresentationAttributes(sb *strings.Builder, p *Presentation, isFillable bool) {
	// Fill attributes
	if isFillable {
		if p.Color != "" {
			sb.WriteString(fmt.Sprintf(` fill="%s"`, html.EscapeString(p.Color)))
		}
		// fill-opacity: SVG default is 1 (opaque). Only write the attribute if the value is different.
		if p.FillOpacity != 1.0 {
			sb.WriteString(fmt.Sprintf(` fill-opacity="%s"`, formatFloat(p.FillOpacity)))
		}
	} else {
		// For elements not typically filled (e.g., <line>), explicitly set fill to "none".
		sb.WriteString(` fill="none"`)
	}

	// Stroke attributes
	if p.Stroke != "" {
		// stroke: SVG default is "none". Explicitly writing it if set.
		sb.WriteString(fmt.Sprintf(` stroke="%s"`, html.EscapeString(p.Stroke)))
	}
	// stroke-width: SVG default is 1. Only write if different from 1, or if explicitly set to 0.
	if p.StrokeWidth != 1.0 {
		sb.WriteString(fmt.Sprintf(` stroke-width="%s"`, formatFloat(p.StrokeWidth)))
	}

	if p.StrokeDashArray != "" {
		sb.WriteString(fmt.Sprintf(` stroke-dasharray="%s"`, html.EscapeString(p.StrokeDashArray)))
	}

	// Transform attribute
	if p.Transform != "" {
		sb.WriteString(fmt.Sprintf(` transform="%s"`, html.EscapeString(p.Transform)))
	}
}

// appendAnimates appends SVG <animate> tags to the strings.Builder.
// `indent` provides spacing for pretty-printing.
func appendAnimates(sb *strings.Builder, animates []*Animate, indent string) {
	for _, animate := range animates {
		sb.WriteString(fmt.Sprintf("%s<animate attributeName=\"%s\"", indent, html.EscapeString(animate.AttributeName)))
		if animate.Values != "" {
			sb.WriteString(fmt.Sprintf(" values=\"%s\"", html.EscapeString(animate.Values)))
		}
		if animate.Dur != "" {
			sb.WriteString(fmt.Sprintf(" dur=\"%s\"", html.EscapeString(animate.Dur)))
		}
		if animate.RepeatCount != "" {
			sb.WriteString(fmt.Sprintf(" repeatCount=\"%s\"", html.EscapeString(animate.RepeatCount)))
		}
		sb.WriteString(" />\n")
	}
}
