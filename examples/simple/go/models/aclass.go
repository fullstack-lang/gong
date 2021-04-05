package models

// Aclass demonstrate basic metabaron
//
// 		swagger:model Aclass
type Aclass struct {
	Name string

	Booleanfield bool

	Aenum   AEnumType
	Aenum_2 AEnumType
	Benum   BEnumType

	Cclass

	Floatfield float64

	Intfield int

	Anotherbooleanfield bool

	Associationtob          *Bclass
	Anotherassociationtob_2 *Bclass

	Anarrayofb      []*Bclass
	Anotherarrayofb []*Bclass

	Anarrayofa []*Aclass
}
