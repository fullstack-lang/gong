package models

import (
	"fmt"
	"strings"
)

type RectAnchoredText struct {
	Name string

	//gong:text width:300 height:300
	Content string

	TextAttributes

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType

	Presentation
	Animates []*Animate
}

func (rectAnchoredText *RectAnchoredText) WriteSVG(sb *strings.Builder, x, y float64) {

	sb.WriteString(
		fmt.Sprintf(
			`  <text xml:space="preserve"
			x="%s" 
			y="%s" 
			text-anchor="%s"
			font-weight="%s"
			font-style="%s"
			font-size="%s"`,
			formatFloat(x+rectAnchoredText.X_Offset),
			formatFloat(y+rectAnchoredText.Y_Offset),
			rectAnchoredText.TextAnchorType.ToString(),
			rectAnchoredText.FontWeight,
			rectAnchoredText.FontStyle,
			rectAnchoredText.FontSize,
		))

	rectAnchoredText.Presentation.WriteSVG(sb)
	sb.WriteString(" >\n")

	lines := strings.Split(rectAnchoredText.Content, "\n")
	for i, line := range lines {

		// if line is empty, it is not displayed by SVG
		if line == "" {
			line = " "
		}

		if i == 0 {
			sb.WriteString(fmt.Sprintf("    <tspan >%s</tspan>\n", line))
		} else {
			sb.WriteString(fmt.Sprintf("    <tspan x=\"%s\" dy=\"1.2em\">%s</tspan>\n", formatFloat(x), line))
		}
	}
	sb.WriteString("</text>\n")
}
