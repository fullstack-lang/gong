package models

type Text struct {
	Name    string
	X, Y    float64
	Content string
	Presentation

	TextAttributes

	Animates []*Animate
}
