package models

// AstructBstructUse is the gong way of defining a MANY-MANY association
//
// The go generated code is the same as for any struct.
// The "Use" at the end of the struct name will generate something different in the front
type AstructBstructUse struct {
	Name string

	// the field of a MANY-MANY association must be the name of the destination struct
	Bstruct2 *Bstruct
}
