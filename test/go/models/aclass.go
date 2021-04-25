package models

import (
	"time"
)

// Aclass demonstrate basic metabaron
//
// 		swagger:model Aclass
type Aclass struct {
	Name string

	Date time.Time

	Booleanfield bool

	Aenum   AEnumType
	Aenum_2 AEnumType
	Benum   BEnumType

	Cclass

	Floatfield float64

	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	Associationtob          *Bclass
	Anotherassociationtob_2 *Bclass

	Anarrayofb      []*Bclass
	Anotherarrayofb []*Bclass

	Anarrayofa []*Aclass
}
