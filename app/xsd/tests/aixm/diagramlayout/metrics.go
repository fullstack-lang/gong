package diagramlayout

import (
	"math"
	"regexp"
	"strings"

	"github.com/fullstack-lang/gong/lib/strutils"
)

var reDocModelLink = regexp.MustCompile(`\[models\.([^\]]+)\]`)

// BoldTextWidth estimates the rendered SVG pixel width of bold ~16px sans-serif text
// used for GongStructShape, GongEnumShape, and GongNoteShape titles in lib/doc.
func BoldTextWidth(s string) float64 {
	w := 0.0
	for _, r := range s {
		switch r {
		case 'i', 'l', 'I', '.', ':', ' ', '(', ')', '!', '|', '\'', ',':
			w += 4.8
		case 'j', 'f', 't', 'r', '-', '/':
			w += 6.4
		case 'm', 'w':
			w += 12.4
		case 'M', 'W':
			w += 14.0
		default:
			if (r >= 'A' && r <= 'Z') || r == '_' {
				w += 10.2
			} else {
				w += 8.1
			}
		}
	}
	return w
}

// NormalTextWidth estimates the rendered SVG pixel width of regular ~15-16px sans-serif text
// used for AttributeShapes, GongEnumValueShapes, and GongNoteShape body lines.
func NormalTextWidth(s string) float64 {
	return BoldTextWidth(s) * 0.89
}

// LinkLabelWidth estimates the rendered SVG pixel width of LinkAnchoredText in lib/doc,
// which uses FontSize="15" and LetterSpacing="0.1em".
func LinkLabelWidth(s string) float64 {
	if s == "" {
		return 0
	}
	return NormalTextWidth(s)*0.95 + float64(len(s))*1.55
}

// StructBoxMinWidth computes the minimum collision-free width for a GongStructShape
// given its displayed header title (with centered title + top-right "(nb)" instance badge)
// and visible "Attr : Type" lines.
func StructBoxMinWidth(displayTitle string, attrLines []string, showNbInstances bool) float64 {
	badgeSymmetricPad := 48.0
	if showNbInstances {
		// Title is centered at W/2 while "(nb)" is right-aligned at W - 9.
		// Reserving 62px on each side guarantees >= 14px gap even for 4-5 digit counts.
		badgeSymmetricPad = 124.0
	}
	width := math.Max(205.0, BoldTextWidth(displayTitle)+badgeSymmetricPad)
	for _, line := range attrLines {
		lw := NormalTextWidth(line) + 28.0
		if lw > width {
			width = lw
		}
	}
	return math.Ceil(width)
}

// EnumBoxSize computes the collision-free (width, height) for a GongEnumShape.
func EnumBoxSize(displayTitle string, valueLines []string) (float64, float64) {
	width := math.Max(220.0, BoldTextWidth(displayTitle)+52.0)
	for _, val := range valueLines {
		vw := NormalTextWidth(val) + 30.0
		if vw > width {
			width = vw
		}
	}
	height := math.Max(66.0, 48.0+float64(len(valueLines))*20.0)
	return math.Ceil(width), math.Ceil(height)
}

// NoteBoxSize computes clean (width, height) for a GongNoteShape matching stage_ux_svg.go wrapping.
func NoteBoxSize(body string, preferredWidth float64) (float64, float64) {
	bodyLines := strings.SplitN(body, "\n", 2)
	noteTitle := ""
	noteContent := ""
	if len(bodyLines) > 0 {
		noteTitle = strings.TrimSpace(bodyLines[0])
	}
	if len(bodyLines) > 1 {
		noteContent = strings.TrimSpace(bodyLines[1])
		noteContent = reDocModelLink.ReplaceAllString(noteContent, "$1")
	}

	width := math.Max(preferredWidth, BoldTextWidth(noteTitle)+56.0)
	if width < 440.0 {
		width = 440.0
	}

	for range 4 {
		wrapCols := int(width / 8.0)
		if wrapCols < 30 {
			wrapCols = 30
		}
		wrapped := strutils.WrapString(noteContent, wrapCols)
		lines := strings.Split(wrapped, "\n")
		maxLineW := 0.0
		for _, line := range lines {
			lw := NormalTextWidth(line) + 36.0
			if lw > maxLineW {
				maxLineW = lw
			}
		}
		if maxLineW <= width || width >= 640.0 {
			height := math.Max(104.0, 50.0+float64(len(lines))*19.0)
			return math.Ceil(width), math.Ceil(height)
		}
		width = math.Min(640.0, math.Ceil(maxLineW))
	}

	wrapped := strutils.WrapString(noteContent, int(width/8.0))
	lines := strings.Split(wrapped, "\n")
	height := math.Max(104.0, 50.0+float64(len(lines))*19.0)
	return math.Ceil(width), math.Ceil(height)
}
