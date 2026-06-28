package models

type WithGoIdentifier struct {
	// HasNameConflict
	//
	// In XML Schema Definition (XSD), it is technically possible to have a
	// named complex type and an element with the same name or elements with
	// the same (see musicxml.xsd for instance),
	HasNameConflict bool
	GoIdentifier    string
}
