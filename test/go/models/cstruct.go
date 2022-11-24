package models

// Cstruct demonstrates embedded
// structs
type Cstruct struct {
	CName string

	CFloatfield float64

	Bstruct  *Bstruct
	Bstruct2 *Bstruct

	Dstruct  *Dstruct
	Dstruct2 *Dstruct
}
