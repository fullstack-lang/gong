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

	Date time.Time

	FloatValue float64

	IntValue int

	Duration time.Duration

	EnumString EnumTypeString

	EnumInt EnumTypeInt
}

func (from *A) CopyBasicFields(to *A) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
	to.FloatValue = from.FloatValue
	to.IntValue = from.IntValue
	to.Duration = from.Duration
	to.EnumString = from.EnumString
	to.EnumInt = from.EnumInt
}

type B_WOP struct {
	// insertion point

	Name string
}

func (from *B) CopyBasicFields(to *B) {
	// insertion point
	to.Name = from.Name
}

