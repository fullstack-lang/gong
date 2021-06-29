package models

type AclassBclass2Use struct {
	Name string

	// the field of a MANY-MANY association must be the name of the destination struct
	Bclass2 *Bclass
}
