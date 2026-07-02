package models

type LinkType string

// values for EnumType
const (
	LINK_TYPE_LINE_WITH_CONTROL_POINTS LinkType = "LINK_TYPE_LINE_WITH_CONTROL_POINTS"

	//  These connectors allow you to anchor the link at any point along the edge of the
	// shape, giving you more flexibility in connecting shapes while still maintaining
	// the orthogonal (vertical/horizontal) orientation of the connection lines.
	LINK_TYPE_FLOATING_ORTHOGONAL LinkType = "LINK_TYPE_FLOATING_ORTHOGONAL"
)
