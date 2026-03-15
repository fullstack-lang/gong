package models

import (
	"time"
)

type Artist struct {
	Name string

	IsDead      bool
	DateOfDeath time.Time
	Place       *Place
}

func (*Artist) IsArtElement() {
}

type ArtistShape struct {
	Name   string
	Artist *Artist

	X, Y float64

	Width, Height float64

	IsHidden bool
}

func (shape *ArtistShape) GetArtElement() *Artist {
	return shape.Artist
}
