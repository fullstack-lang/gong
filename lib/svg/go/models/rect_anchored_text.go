package models

import (
	"fmt"
	"html"
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

func (rectAnchoredText *RectAnchoredText) WriteString(sb *strings.Builder, x, y float64) {

	sb.WriteString(
		fmt.Sprintf(
			`  <text 
			x="%s" 
			y="%s" 
			text-anchor="%s"
			font-weight="%s"
			font-style="%s"
			font-size="%s"`,
			formatFloat(x),
			formatFloat(y),
			rectAnchoredText.TextAnchorType.ToString(),
			rectAnchoredText.FontWeight,
			rectAnchoredText.FontStyle,
			rectAnchoredText.FontSize,
		))

	rectAnchoredText.Presentation.WriteString(sb)
	sb.WriteString(" >")

	lines := strings.Split(rectAnchoredText.Content, "\n")
	for i, line := range lines {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("    <tspan>%s</tspan>\n", html.EscapeString(line)))
		} else {
			sb.WriteString(fmt.Sprintf("    <tspan x=\"%s\" dy=\"1.2em\">%s</tspan>\n", formatFloat(x), html.EscapeString(line)))
		}
	}
	sb.WriteString(rectAnchoredText.Content)
	sb.WriteString("</text>\n")
}
