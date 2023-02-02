package models

// Umlsc diagram struct store a class diagram
// temporary here
// swagger:model Umlsc
type Umlsc struct {
	Name string

	// this is the memory model (and not the "memory motel" of the Rolling Stones)
	// it is not ignored by swagger because it is used by the angular model
	States []*UmlState

	// in this version, only one state is active mong the states
	// (with the embedded states version, that might change)
	Activestate string

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsInDrawMode bool
}
