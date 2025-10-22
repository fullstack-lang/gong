package models

// RectLinkLink define a link between a rectangle  and a target link
type RectLinkLink struct {
	Name  string
	Start *Rect
	End   *Link

	// TargetAnchorPosition defines the anchor point on the target link between
	// the start and and the end rectangle
	// 0 zero means the anchor is on the anchor at the start rectangle and 1 means the anchor is on the anchor on the end rectangle
	TargetAnchorPosition float64

	Presentation
}
