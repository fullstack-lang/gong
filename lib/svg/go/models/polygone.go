package models

type Polygone struct {
	Name   string
	Points string
	Presentation
	Animates []*Animate
}
