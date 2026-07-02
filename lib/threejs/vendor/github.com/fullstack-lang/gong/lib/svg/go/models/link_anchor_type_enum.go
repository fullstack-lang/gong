package models

// LinkAnchorType specifies the way a text should be
// attached to a link end.
//
// the link end meets a rect either horizontally or vertically
type LinkAnchorType string

const (
	// If horizontally, the text will be on top of the link end
	// If verticaly, the text will be on the left
	LINK_LEFT_OR_TOP LinkAnchorType = "LINK_LEFT_OR_TOP"
	// bottom / right
	LINK_RIGHT_OR_BOTTOM LinkAnchorType = "LINK_RIGHT_OR_BOTTOM"
)
