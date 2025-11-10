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
	// auto: (Default)
	// If this property occurs on a <text> element, the computed value depends on the writing-mode.
	// - If horizontal, the dominant-baseline component is 'alphabetic'.
	// - If vertical, the dominant-baseline component is 'central'.
	//
	// If on a <tspan> or <textPath>, it remains the same as the parent.
	// If there is no parent text content element, it's computed as for <text>.
	DominantBaselineAuto DominantBaselineType = "auto"

	// alphabetic:
	// The baseline-identifier for the dominant-baseline is set to be 'alphabetic'.
	// The derived baseline-table is constructed using the 'alphabetic' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineAlphabetic DominantBaselineType = "alphabetic"

	// ideographic:
	// The baseline-identifier for the dominant-baseline is set to be 'ideographic'.
	// The derived baseline-table is constructed using the 'ideographic' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineIdeographic DominantBaselineType = "ideographic"

	// hanging:
	// The baseline-identifier for the dominant-baseline is set to be 'hanging'.
	// The derived baseline-table is constructed using the 'hanging' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineHanging DominantBaselineType = "hanging"

	// mathematical:
	// The baseline-identifier for the dominant-baseline is set to be 'mathematical'.
	// The derived baseline-table is constructed using the 'mathematical' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineMathematical DominantBaselineType = "mathematical"

	// central:
	// The baseline-identifier for the dominant-baseline is set to be 'central'.
	// The derived baseline-table is constructed from a font baseline-table chosen using the
	// following priority order: 'ideographic', 'alphabetic', 'hanging', 'mathematical'.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineCentral DominantBaselineType = "central"

	// middle:
	// The baseline-identifier for the dominant-baseline is set to be 'middle'.
	// The derived baseline-table is constructed from a font baseline-table chosen using the
	// following priority order: 'alphabetic', 'ideographic', 'hanging', 'mathematical'.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineMiddle DominantBaselineType = "middle"

	// text-after-edge:
	// The baseline-identifier for the dominant-baseline is set to be 'text-after-edge'.
	// The choice of which font baseline-table to use is browser dependent.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineTextAfterEdge DominantBaselineType = "text-after-edge"

	// text-before-edge:
	// The baseline-identifier for the dominant-baseline is set to be 'text-before-edge'.
	// The choice of which font baseline-table to use is browser dependent.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineTextBeforeEdge DominantBaselineType = "text-before-edge"

	// text-top:
	// This value uses the top of the em box as the baseline.
	DominantBaselineTextTop DominantBaselineType = "text-top"

	// -- Deprecated Values --

	// use-script: [Deprecated]
	// The dominant-baseline and baseline-table are set by determining the predominant script.
	// The writing-mode selects the appropriate baseline-tables, and the dominant baseline
	// selects the table that corresponds to that baseline.
	DominantBaselineUseScript DominantBaselineType = "use-script"

	// no-change: [Deprecated]
	// The dominant-baseline, the baseline-table, and the baseline-table font-size
	// remain the same as that of the parent text content element.
	DominantBaselineNoChange DominantBaselineType = "no-change"

	// reset-size: [Deprecated]
	// The dominant-baseline and the baseline-table remain the same, but the
	// baseline-table font-size is changed to the value of the font-size attribute on this element.
	DominantBaselineResetSize DominantBaselineType = "reset-size"
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
			dominant-baseline="%s"
			text-anchor="%s"
			font-weight="%s"
			font-style="%s"
			font-size="%s"`,
			formatFloat(x+rectAnchoredText.X_Offset),
			formatFloat(y+rectAnchoredText.Y_Offset),
			rectAnchoredText.WritingMode.ToString(),
			rectAnchoredText.DominantBaseline.ToString(),
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
