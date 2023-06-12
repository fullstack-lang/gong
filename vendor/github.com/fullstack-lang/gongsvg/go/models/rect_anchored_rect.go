package models

type RectAnchoredRect struct {
	Name string

	X, Y, Width, Height, RX float64

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType

	// if true, rect has the same Dimension that the rect it is anchored to
	WidthFollowRect  bool
	HeightFollowRect bool

	Presentation
}
