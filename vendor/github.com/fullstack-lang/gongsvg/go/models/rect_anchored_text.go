package models

type RectAnchoredText struct {
	Name    string
	Content string

	// "bold", "normal", ...
	FontWeight string
	FontSize   int

	// Specifies the style of the font (e.g., normal, italic, oblique).
	FontStyle string

	X_Offset float64
	Y_Offset float64

	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType

	Presentation
	Animates []*Animate
}
