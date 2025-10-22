package models

type Ellipse struct {
	Name           string
	CX, CY, RX, RY float64

	Presentation

	Animates []*Animate
}
