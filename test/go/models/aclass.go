package models

import (
	"time"
)

// Aclass demonstrate basic metabaron
//
// 		swagger:model Aclass
type Aclass struct {

	// a "Name" field is necessary to generate a GongStruct
	Name string

	// time.Time is a supported type of gong
	Date time.Time

	// bool is a supported type of gong
	Booleanfield bool

	// enums is a supported type of gong (if they are string)
	Aenum   AEnumType
	Aenum_2 AEnumType
	Benum   BEnumType

	// an embedded struct is supported (without a field name)
	Cclass

	// float64 is a supported type of gong
	Floatfield float64

	// int is a supported type of gong
	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	// ONE-ZERO/ONE association is a supported type of gong
	Associationtob          *Bclass
	Anotherassociationtob_2 *Bclass

	// ONE-MANY association is a supported type of gong
	// not that that in go, this is a MANY-MANY association. In gong, the association
	// is implemented in the destination of the association, therefore, a Bclass instance
	// is related to ZERO/ONE Aclass instance by Arrayofb field
	Anarrayofb []*Bclass

	// gong supports different fields with the same destination struct, even itself
	Anotherarrayofb []*Bclass
	Anarrayofa      []*Aclass

	// MANY-MANY association, (because AclassBclassUse ends with "Use")
	AnarrayofbUse []*AclassBclassUse
}
