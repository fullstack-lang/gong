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

	WritingMode WritingMode

	Presentation
	Animates []*Animate
}

/*
The writing-mode property in SVG and CSS specifies how lines of text are laid out, determining whether text
flows horizontally or vertically, and the direction in which lines stack. Here are the primary possible values:

horizontal-tb:

This is the default value for most languages.
Text flows horizontally from left to right (for left-to-right scripts like English) or right to left (for right-to-left scripts like Arabic or Hebrew).
Lines are stacked from top to bottom.
Think of a standard English book.

vertical-rl:

Text flows vertically from top to bottom.
Lines are stacked from right to left.
This is common for traditional East Asian scripts like Chinese, Japanese, and Korean.
Individual characters within the vertical lines are typically oriented according to their script's rules (Latin characters will usually appear rotated 90 degrees clockwise by default).

vertical-lr:

Text flows vertically from top to bottom.
Lines are stacked from left to right.
This is less common for natural language scripts but can be used for specific layout effects or some boustrophedon scripts.
Similar to vertical-rl, Latin characters will usually appear rotated 90 degrees clockwise by default.
*/
type WritingMode string

const (
	WritingModeHorizontalTB WritingMode = "horizontal-tb"
	WritingModeVerticalRL   WritingMode = "vertical-rl"
	WritingModeVertivcalLR  WritingMode = "vertical-lr"
)

func (rectAnchoredText *RectAnchoredText) WriteSVG(sb *strings.Builder, x, y float64) {

	sb.WriteString(
		fmt.Sprintf(
			`  <text xml:space="preserve"
			x="%s" 
			y="%s"
			writing-mode="%s"
			text-anchor="%s"
			font-weight="%s"
			font-style="%s"
			font-size="%s"`,
			formatFloat(x+rectAnchoredText.X_Offset),
			formatFloat(y+rectAnchoredText.Y_Offset),
			rectAnchoredText.WritingMode.ToString(),
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
