package models

// Link represent the UML Link in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model Link
type Link struct {
	Name string

	// Identifier is the identifier of the struct field referenced by the
	// UML field of the classshape in the modeled package
	//gong:ident
	Identifier string

	//gong:ident
	Fieldtypename string

	// position of the field text relative to the
	// arrow end
	FieldOffsetX float64
	FieldOffsetY float64

	TargetMultiplicity MultiplicityType

	// position of the TargetMultiplicity text relative to the
	// arrow end
	TargetMultiplicityOffsetX float64
	TargetMultiplicityOffsetY float64

	SourceMultiplicity MultiplicityType

	// position of the SourceMultiplicity text relative to the
	// arrow end
	SourceMultiplicityOffsetX float64
	SourceMultiplicityOffsetY float64

	// For link with control points
	// Vertices at the middle
	Middlevertice *Vertice

	// for links with floating anchors
	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64
}
