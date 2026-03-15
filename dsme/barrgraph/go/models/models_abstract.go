package models

import "time"

type ArtElement interface {
	IsArtElement()
	GetName() string
	GetIsInRenameMode() bool
	SetIsInRenameMode(bool)
}

type AbstractTypeFields struct {
	// nodes can be edited
	IsInRenameMode bool
}

func (r *AbstractTypeFields) GetIsInRenameMode() bool {
	return r.IsInRenameMode
}

func (r *AbstractTypeFields) SetIsInRenameMode(isInRenameMode bool) {
	r.IsInRenameMode = isInRenameMode
}

type Place struct {
	Name string
}

type Desk struct {
	Name            string
	SelectedDiagram *Diagram
}

type Movement struct {
	Name string

	AbstractTypeFields

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

type Artist struct {
	Name string

	AbstractTypeFields

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

type ArtefactType struct {
	Name string

	AbstractTypeFields
}

func (*ArtefactType) IsArtElement() {
}

// An Influence is between one Artist/Movement/ArtefactType and another
//
// since gong does not yet support interface, one have to mutliply source/target types
type Influence struct {
	Name string

	SourceMovement     *Movement
	SourceArtefactType *ArtefactType
	SourceArtist       *Artist

	source ArtElement

	TargetMovement     *Movement
	TargetArtefactType *ArtefactType
	TargetArtist       *Artist

	target ArtElement

	// hypothetical, some influences are with ashed lines
	// For instance, Marchine Art to Brancusi (indeed)
	IsHypothtical bool
}
