package models

type RectAnchorType string

// values for EnumType
const (
	RECT_TOP                RectAnchorType = "RECT_TOP"
	RECT_TOP_LEFT           RectAnchorType = "RECT_TOP_LEFT"
	RECT_TOP_RIGHT          RectAnchorType = "RECT_TOP_RIGHT"
	RECT_BOTTOM             RectAnchorType = "RECT_BOTTOM"
	RECT_BOTTOM_LEFT        RectAnchorType = "RECT_BOTTOM_LEFT"
	RECT_BOTTOM_LEFT_LEFT   RectAnchorType = "RECT_BOTTOM_LEFT_LEFT"
	RECT_BOTTOM_BOTTOM_LEFT RectAnchorType = "RECT_BOTTOM_BOTTOM_LEFT"
	RECT_BOTTOM_RIGHT       RectAnchorType = "RECT_BOTTOM_RIGHT"

	// the path is inside the shape, but on the right
	RECT_BOTTOM_INSIDE_RIGHT RectAnchorType = "RECT_BOTTOM_INSIDE_RIGHT"

	RECT_LEFT   RectAnchorType = "RECT_LEFT"
	RECT_RIGHT  RectAnchorType = "RECT_RIGHT"
	RECT_CENTER RectAnchorType = "RECT_CENTER"
)
