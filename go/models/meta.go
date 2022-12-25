package models

// Meta is a fallback solution explained to the meta keyword proposal
type Meta struct {
	Name string
	Text string

	// references to symbols in the code
	MetaReferences []*MetaReference
}
