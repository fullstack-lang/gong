package models

// Cstruct demonstrate basic metabaron
//
// 		swagger:model Cstruct
type Cstruct struct {
	CName string

	CFloatfield float64

	Bstruct *Bstruct
}
