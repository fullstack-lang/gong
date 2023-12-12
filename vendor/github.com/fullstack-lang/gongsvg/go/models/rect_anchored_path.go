package models

type RectAnchoredPath struct {
	Name string

	Definition string

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType

	// if true, rect has the same Dimension that the rect it is anchored to
	WidthFollowRect     bool
	HeightFollowRect    bool
	ScalePropotionnally bool

	Presentation
}
