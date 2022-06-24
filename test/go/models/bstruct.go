package models

// Bstruct demonstrate basic metabaron
//
// 	swagger:model Bstruct
type Bstruct struct {
	Name string

	Floatfield float64

	Intfield int

	// swagger:ignore this field is ignored by gongc
	Struct interface{} `gorm:"-"` // instruction for gorm ORM

	// swagger:ignore
	ToBeIgnored int
}
