package models

// Link represent the UML Link in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model Link
type Link struct {
	Name string

	Structname string

	// Identifier is the identifier of the struct field referenced by the
	// UML field of the classshape in the modeled package
	//gong:ident
	Identifier string

	Fieldtypename      string
	TargetMultiplicity MultiplicityType
	SourceMultiplicity MultiplicityType

	// Vertices at the middle
	Middlevertice *Vertice
}
