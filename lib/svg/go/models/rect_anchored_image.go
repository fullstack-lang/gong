package models

type RectAnchoredPngImage struct {
	Name string

	X, Y, Width, Height, RX float64

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType

	// gong:text gong:width 600 gong:height 400
	Base64Content string
}
