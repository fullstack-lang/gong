package models

import (
	"fmt"
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

	segmentOfInterest := *segment
	var positionArrowType PositionOnArrowType

	if segmentOfInterest.Type == StartSegment {
		positionArrowType = POSITION_ON_ARROW_START
		segmentOfInterest = swapSegment(segmentOfInterest)
	} else {
		positionArrowType = POSITION_ON_ARROW_END
	}

	lines := strings.Split(linkAnchoredText.Content, "\n")
	x, y := segmentOfInterest.EndPointWithoutRadius.X, segmentOfInterest.EndPointWithoutRadius.Y

	y += linkAnchoredText.auto_Y_offset(link, segment, positionArrowType)
	y += linkAnchoredText.Y_Offset
	x += linkAnchoredText.X_Offset

	var anchorType string
	switch segmentOfInterest.Orientation {
	case ORIENTATION_HORIZONTAL:

		if segmentOfInterest.EndPoint.X > segmentOfInterest.StartPoint.X {
			anchorType = TEXT_ANCHOR_END.ToString()
		} else {
			anchorType = TEXT_ANCHOR_START.ToString()
		}
	case ORIENTATION_VERTICAL:

		if linkAnchoredText.LinkAnchorType == LINK_LEFT_OR_TOP {
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
	sb.WriteString(">\n")

	for i, line := range lines {

		// if line is empty, it is not displayed by SVG
		if line == "" {
			line = " "
		}

		if i == 0 {
			// Preserve leading/trailing spaces and make them non-breaking
			sb.WriteString(fmt.Sprintf("    <tspan xml:space=\"preserve\">\u00A0%s\u00A0</tspan>\n", line))
		} else {
			// Preserve leading/trailing spaces and make them non-breaking
			sb.WriteString(fmt.Sprintf("    <tspan xml:space=\"preserve\" x=\"%s\" dy=\"1em\">\u00A0%s\u00A0</tspan>\n", formatFloat(x), line))
		}
	}
	sb.WriteString("</text>\n")
}

func (linkAnchoredText *LinkAnchoredText) auto_Y_offset(
	link *Link,
	segment *Segment,
	positionType PositionOnArrowType) (res float64) {

	offset := 0.0
	offsetSign := 1.0
	oneEm := 6.0

	if !linkAnchoredText.AutomaticLayout {
		return offset
	}

	var orientation OrientationType
	if positionType == POSITION_ON_ARROW_END {
		orientation = link.EndOrientation
	} else {
		orientation = link.StartOrientation
	}

	if positionType == POSITION_ON_ARROW_START {
		offsetSign = -offsetSign
	}

	if orientation == ORIENTATION_VERTICAL {
		if segment.EndPoint.Y > segment.StartPoint.Y {
			offsetSign = -offsetSign
		}
	} else {
		if positionType == POSITION_ON_ARROW_END {
			offsetSign = -offsetSign
		}
		if linkAnchoredText.LinkAnchorType == LINK_RIGHT_OR_BOTTOM {
			offsetSign = -offsetSign
		}
	}

	if link.HasEndArrow {
		offset += link.EndArrowSize
	}
	if offsetSign == 1 {
		offset += oneEm
	} else {
		offset += oneEm * 0.4
	}

	lines := strings.Split(linkAnchoredText.Content, "\n")

	res = offset*offsetSign + linkAnchoredText.Y_Offset

	if len(lines) > 1 && linkAnchoredText.LinkAnchorType == LINK_LEFT_OR_TOP {
		res -= float64(len(lines)) * oneEm * 0.4
	}

	return
}
