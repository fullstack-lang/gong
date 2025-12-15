package models

type TextAnchorType string

// values for https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor
const (
	TEXT_ANCHOR_START  TextAnchorType = "start"
	TEXT_ANCHOR_CENTER TextAnchorType = "middle"
	TEXT_ANCHOR_END    TextAnchorType = "end"
)
