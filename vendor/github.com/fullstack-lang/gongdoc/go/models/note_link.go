package models

// A NoteLink is a link from a MetaShape to
// a Link or to a Classshape (guessed from the Type)
type NoteLink struct {
	Name string

	Type ReferenceType

	Classshape *Classshape
	Link       *Link

	// Vertices at the middle
	Middlevertice *Vertice
}
