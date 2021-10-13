package models

import "time"

// Astruct demonstrate basic metabaron
//
// 		swagger:model Astruct
type Astruct struct {
	Name string

	Date time.Time

	Booleanfield bool

	Floatfield float64

	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	AnarrayofbUse []*AstructBstructUse
}
