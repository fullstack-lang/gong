package models

type RectAnchoredPath struct {
	Name string

	Definition string

	X_Offset float64
	Y_Offset float64

	// gong:width 400
	RectAnchorType RectAnchorType

	// if true, rect has the same Dimension that the rect it is anchored to
	// rect must scale proportionnaly
	ScalePropotionnally bool

	// AppliedScaling is the scale that is applied is ScalePropotionnally is set
	// The value is initialized to 1 then scales with the scaling action
	// of the end user
	AppliedScaling float64

	Presentation
}
