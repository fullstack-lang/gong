// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type A_WOP struct {
	// insertion point

	Name string

	Date time.Time

	Duration time.Duration

	FloatValue float64

	IntValue int

	EnumString EnumTypeString

	EnumInt EnumTypeInt

	UUID string
}

func (from *A) GongCopyBasicFields(to *A) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
	to.Duration = from.Duration
	to.FloatValue = from.FloatValue
	to.IntValue = from.IntValue
	to.EnumString = from.EnumString
	to.EnumInt = from.EnumInt
	to.UUID = from.UUID
}

type B_WOP struct {
	// insertion point

	Name string
}

func (from *B) GongCopyBasicFields(to *B) {
	// insertion point
	to.Name = from.Name
}

type C_WOP struct {
	// insertion point

	Name string
}

func (from *C) GongCopyBasicFields(to *C) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
