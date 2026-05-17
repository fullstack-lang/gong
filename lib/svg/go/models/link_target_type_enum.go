package models

type LinkTargetType string

// values for LinkTargetType
const (
	// Opens the link in a new tab or window.
	LINK_TARGET_BLANK LinkTargetType = "_blank"
	// Opens the link in the same tab/frame (this is the default behavior if you don't include a target attribute).
	LINK_TARGET_SELF LinkTargetType = "_self"
	// Opens the link in the parent frame (used when dealing with nested <iframe> elements).
	LINK_TARGET_PARENT LinkTargetType = "_parent"
	// Breaks out of all frames and opens the link in the full body of the current window.
	LINK_TARGET_TOP LinkTargetType = "_top"
)
