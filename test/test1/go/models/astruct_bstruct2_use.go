package models

type AstructBstruct2Use struct {
	Name string

	// the field of a MANY-MANY association must be the name of the destination struct
	Bstrcut2 *Bstruct
}
