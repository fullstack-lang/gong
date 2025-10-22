package models

type Polyline struct {
	Name   string
	Points string
	Presentation
	ShapeConditions
	Animates []*Animate
}
