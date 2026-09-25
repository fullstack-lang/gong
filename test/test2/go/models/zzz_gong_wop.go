// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type A_WOP struct {
	// insertion point

	Name string

	NumberField int

	Foo int

	Bar float64

	Zorgh string
}

func (from *A) GongCopyBasicFields(to *A) {
	// insertion point
	to.Name = from.Name
	to.NumberField = from.NumberField
	to.Foo = from.Foo
	to.Bar = from.Bar
	to.Zorgh = from.Zorgh
}

type B_WOP struct {
	// insertion point

	Name string
}

func (from *B) GongCopyBasicFields(to *B) {
	// insertion point
	*to = *from
}

// end of insertion point
