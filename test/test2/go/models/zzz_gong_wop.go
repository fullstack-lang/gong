// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type A_WOP struct {
	// insertion point

	Name string

	NumberField int

	Foo int

	Bar float64

	Zorgh string
}

func (from *A) CopyBasicFields(to *A) {
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

func (from *B) CopyBasicFields(to *B) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
