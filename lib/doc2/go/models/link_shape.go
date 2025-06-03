package models

// LinkShape represent the UML LinkShape in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model LinkShape
type LinkShape struct {
	Name string

	// Identifier is the identifier of the struct field referenced by the
	// UML field of the classshape in the modeled package
	//gong:ident
	Identifier string

	// for storing the reference as a renaming target for gopls
	// for instance 'ref_models.Astruct.Anarrayofb'
	//gong:meta
	IdentifierMeta any

	//gong:ident
	Fieldtypename string

	//gong:meta
	FieldTypeIdentifierMeta any

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

	// for links with floating anchors
	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	// Vertice in the middle
	X                float64
	Y                float64
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64
}
