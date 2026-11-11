package models

type Protocol struct {
	Name string

	// Acronym is an optional short name for the protocol, used in a data flow diagram
	// as a replacement of the full name when present.
	Acronym string

	//gong:text width:300 height:300
	Description string

	ProtocolLayers []*ProtocolLayer

	IsProtocolLayersNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Protocol)(nil)

// for instance IP, ethernet, REST
type ProtocolLayer struct {
	Name string

	Acronym string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*ProtocolLayer)(nil)
