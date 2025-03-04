package models

type GongEnumValueEntry struct {
	// Name of the enum value
	Name string

	// Identifier is the identifier of the enum value referenced by the shape in the modeled package
	//gong:ident
	Identifier string
}
