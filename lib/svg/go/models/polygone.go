package models

type Polygone struct {
	Name   string
	Points string
	Presentation
	ShapeConditions
	Animates []*Animate
}
