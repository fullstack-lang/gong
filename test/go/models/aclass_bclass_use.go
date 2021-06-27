package models

// AclassBclassUse is the gong way of defining a MANY-MANY association
//
// The go generated code is the same as for any struct.
// The "Use" at the end of the struct name will generate something different in the front
type AclassBclassUse struct {
	Name   string
	Bclass *Bclass
}
