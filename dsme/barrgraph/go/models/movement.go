package models

import (
	"time"
)

type Movement struct {
	Name string

	Date time.Time
	// NON GEOMETRICAL and GEOMTRICAL ABSRTACT ART
	// have no dates shown
	HideDate bool

	Places []*Place

	// Barr’s primary intention with this chart was to map the genealogy specifically
	// of non-representational (abstract) art. He wanted to show how European avant-garde
	// movements evolved and influenced each other to ultimately arrive at pure abstraction.
	//
	// However, he ran into a historical problem: many of the most influential movements
	// of the early 20th century were not strictly or entirely abstract. By adding the word
	// "(ABSTRACT)" in parentheses, Barr was deliberately isolating a specific thread within those movements
	HasTaxonomicFilter bool // some mouvements are abstracts
	TaxonomicFilter    string

	// the "Modern" in "Modern Architecture" is rendered above.
	// We surmise this is a Alfred Barr"s intentional caterogy choice
	IsFeatured    bool
	FeaturePrefix string // the "Modern" in "Modern Architecture"

	IsMajor bool // cubism
	IsMinor bool // orphism

	// For DE STILJ and NEOPLATICISM, the
	// rendering is different the "and" is in lower cases
	// and NEOPLATICISM is on the next line
	// So, for this movement, the
	// Name is "DE STILJ" and the AdditionnalName
	// AdditionnalName is NEOPLATICISM
	AdditionnalName string
}

func (*Movement) IsArtElement() {
}

func (shape *MovementShape) GetArtElement() *Movement {
	return shape.Movement
}

type MovementShape struct {
	Name     string
	Movement *Movement

	X, Y float64

	Width, Height float64

	IsHidden bool
}
