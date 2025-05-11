package models

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkAnchoredText struct {
	Name string

	//gong:text width:300 height:300
	Content string

	// AutomaticLayout is true will have the front compute the
	// X_Offset / Y_Offset of the anchored text
	AutomaticLayout bool
	LinkAnchorType  LinkAnchorType

	// values if AutomaticLayout is false
	X_Offset float64
	Y_Offset float64

	TextAttributes

	Presentation
	Animates []*Animate

	Impl LinkAnchoredTextImplInterface
}

func (linkAnchoredText *LinkAnchoredText) OnAfterUpdate(stage *Stage, _, frontLinkAnchoredText *LinkAnchoredText) {

	if linkAnchoredText.Impl != nil {
		linkAnchoredText.Impl.AnchoredTextUpdated(frontLinkAnchoredText)
	}
}

func (linkAnchoredText *LinkAnchoredText) WriteSVG(sb *strings.Builder, link *Link, segment *Segment) {

	lines := strings.Split(linkAnchoredText.Content, "\n")
	x, y := segment.EndPointWithoutRadius.X, segment.EndPointWithoutRadius.Y

	offset := 10.0

	offsetPerLine := offset

	s := linkAnchoredText.FontSize
	if strings.HasSuffix(s, "px") {
		numericStr := strings.TrimSuffix(s, "px")

		// Try to convert the numeric string part to an integer
		value, err := strconv.Atoi(numericStr)
		if err != nil {

		} else {
			offsetPerLine = float64(value)
		}
	}

	var anchorType string
	switch segment.Orientation {
	case ORIENTATION_HORIZONTAL:

		if segment.EndPoint.X > segment.StartPoint.X {
			x -= offset / 2
			if link.HasEndArrow {
				x -= link.EndArrowSize
			}
			anchorType = TEXT_ANCHOR_END.ToString()
		} else {
			x += offset / 2
			if link.HasEndArrow {
				x += link.EndArrowSize
			}
			anchorType = TEXT_ANCHOR_START.ToString()
		}

		switch linkAnchoredText.LinkAnchorType {
		case LINK_LEFT_OR_TOP:
			y = y - offsetPerLine*float64(len(lines))
		case LINK_RIGHT_OR_BOTTOM:
			y += offset
			if link.HasEndArrow {
				y += link.EndArrowSize
			}
		}
	case ORIENTATION_VERTICAL:
		y -= offset

		if linkAnchoredText.LinkAnchorType == LINK_LEFT_OR_TOP {
			x -= offset / 2

			if link.HasEndArrow {
				x -= link.EndArrowSize
			}
			anchorType = TEXT_ANCHOR_END.ToString()
		} else {
			anchorType = TEXT_ANCHOR_START.ToString()
		}
	}
	sb.WriteString(
		fmt.Sprintf(
			`  <text xml:space="preserve"
			x="%s" 
			y="%s" 
			text-anchor="%s"
			font-weight="%s"
			font-style="%s"
			font-size="%s"`,
			formatFloat(x),
			formatFloat(y),
			anchorType,
			linkAnchoredText.FontWeight,
			linkAnchoredText.FontStyle,
			linkAnchoredText.FontSize,
		))

	linkAnchoredText.Presentation.WriteSVG(sb)
	sb.WriteString(" >\n")

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
