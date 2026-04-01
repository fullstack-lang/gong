package models

// Identifiable is a constraint for types that can provide a string identifier.
// This is for type that are referenced by others
// for instance
//
// DATATYPE_DEFINITION...
// ATTRIBUTE_DEFINITION...
type Identifiable interface {
	PointerToGongstruct
	GetIdentifier() string
}

type AttributeDefinition interface {
	Identifiable
	GetDatatypeDefinitionRef() string
	GetIsEditable() bool
	GetLongName() string
}

type AttributeDefinitionRendering interface {
	GongstructIF
	GetName() string
	GetShowInTablePtr() *bool
	GetShowInTitlePtr() *bool
	GetShowInSubjectPtr() *bool
	GetRankPtr() *int
}

// Things like A_ATTRIBUTE_DEFINITION_XHTML_REF
type AttributeDefinitionRef interface {
	PointerToGongstruct

	GetRef() string
}

type Attribute interface {
	GetValue() string
	GetAttributeDefinitionRef() string
	// GetAttributeDefinition() AttributeDefinition
}

type DatatypeDefinition interface {
	PointerToGongstruct

	GetLongName() string
}

// for SPEC OBJECTS and SPEC RELATIONS
type ObjectWithValues interface {
	GetValues() *A_ATTRIBUTE_VALUE_XHTML_1
}
