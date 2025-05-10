package models

type RectAnchoredText struct {
	Name string

	//gong:text width:300 height:300
	Content string

	TextAttributes

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType

	Presentation
	Animates []*Animate
}
