package models

import "time"

// Aclass demonstrate basic metabaron
//
// 		swagger:model Aclass
type Aclass struct {
	Name string

	Date time.Time

	Booleanfield bool

	Floatfield float64

	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	// Anarrayofa []*Aclass
}
