package models

// Astruct demonstrate basic metabaron
//
// 		swagger:model Astruct
type Astruct struct {
	Name string

	Anarrayofa []*Astruct

	// Date time.Time

	// Booleanfield bool

	// Floatfield float64

	// Intfield int

	// Anotherbooleanfield bool

	// Duration1 time.Duration

	// Associationtob *Bstruct

	// Anarrayofb []*Bstruct

	// AnarrayofbUse []*AstructBstructUse
}
