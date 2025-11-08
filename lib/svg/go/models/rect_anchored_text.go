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

	RectAnchorType   RectAnchorType
	TextAnchorType   TextAnchorType
	DominantBaseline DominantBaselineType

	WritingMode WritingMode

	Presentation
	Animates []*Animate
}

type DominantBaselineType string

const (
	//	These are the most frequently used values for precise text alignment.
	//
	// auto (Default): The browser determines the baseline automatically. For horizontal text, this is typically the same as alphabetic. For vertical text, it's typically central.
	DominantBaselineAuto DominantBaselineType = "auto"

	// middle: This aligns the vertical midpoint of the text to the y coordinate. This is the value you'd use to perfectly center text, as you asked in your previous question.
	DominantBaselineMiddle DominantBaselineType = "middle"

	// alphabetic: The standard baseline used by most Latin, Greek, and Cyrillic scripts (the line the bottom of letters like 'a', 'b', 'c' sit on). This is often the default behavior if auto is set.
	DominantBaselineAlphabetic DominantBaselineType = "alphabetic"

	// hanging: The baseline used by scripts like Devanagari (Hindi) and Tibetan, where the characters "hang" from the line.
	DominantBaselineHanging DominantBaselineType = "hanging"

	// ideographic: The baseline used for East Asian scripts (Chinese, Japanese, Korean) that aligns with the bottom of the ideographic characters.
	DominantBaselineIdeographic DominantBaselineType = "ideographic"
)

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
			`  <text
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
	sb.WriteString(" >")

	lines := strings.Split(rectAnchoredText.Content, "\n")
	for i, line := range lines {

		// if line is empty, it is not displayed by SVG
		if line == "" {
			line = " "
		}

		if i == 0 {
			sb.WriteString(fmt.Sprintf("<tspan x=\"%s\" text-anchor=\"%s\">%s</tspan>",
				formatFloat(x),
				rectAnchoredText.TextAnchorType.ToString(),
				line,
			))
		} else {
			sb.WriteString(fmt.Sprintf("<tspan x=\"%s\" dy=\"1.2em\" text-anchor=\"%s\">%s</tspan>",
				formatFloat(x),
				rectAnchoredText.TextAnchorType.ToString(),
				line,
			))
		}
	}
	sb.WriteString("</text>\n")
}
