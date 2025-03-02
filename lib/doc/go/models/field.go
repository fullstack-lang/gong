package models

// Field represent the UML Field of a Class (a "struct" in go)
// swagger:model Field
type Field struct {
	Name string

	// Identifier is the identifier of the struct field referenced
	// by the UML classshape in the modeled package
	//gong:ident
	Identifier string

	FieldTypeAsString string
	Structname        string
	Fieldtypename     string
}
