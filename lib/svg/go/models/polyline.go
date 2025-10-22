package models

type Polyline struct {
	Name   string
	Points string
	Presentation
	Animates []*Animate
}
