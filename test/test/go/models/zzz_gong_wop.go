// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Astruct_WOP struct {
	// insertion point

	Name string

	Field any

	Date time.Time

	Date2 time.Time

	Booleanfield bool

	Aenum AEnumType

	Aenum_2 AEnumType

	Benum BEnumType

	CEnum CEnumTypeInt

	CName string

	CFloatfield float64

	Floatfield float64

	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	TextFieldBespokeSize string

	TextArea string
}

func (from *Astruct) CopyBasicFields(to *Astruct) {
	// insertion point
	to.Name = from.Name
	to.Field = from.Field
	to.Date = from.Date
	to.Date2 = from.Date2
	to.Booleanfield = from.Booleanfield
	to.Aenum = from.Aenum
	to.Aenum_2 = from.Aenum_2
	to.Benum = from.Benum
	to.CEnum = from.CEnum
	to.CName = from.CName
	to.CFloatfield = from.CFloatfield
	to.Floatfield = from.Floatfield
	to.Intfield = from.Intfield
	to.Anotherbooleanfield = from.Anotherbooleanfield
	to.Duration1 = from.Duration1
	to.TextFieldBespokeSize = from.TextFieldBespokeSize
	to.TextArea = from.TextArea
}

type AstructBstruct2Use_WOP struct {
	// insertion point

	Name string
}

func (from *AstructBstruct2Use) CopyBasicFields(to *AstructBstruct2Use) {
	// insertion point
	to.Name = from.Name
}

type AstructBstructUse_WOP struct {
	// insertion point

	Name string
}

func (from *AstructBstructUse) CopyBasicFields(to *AstructBstructUse) {
	// insertion point
	to.Name = from.Name
}

type Bstruct_WOP struct {
	// insertion point

	Name string

	Floatfield float64

	Floatfield2 float64

	Intfield int
}

func (from *Bstruct) CopyBasicFields(to *Bstruct) {
	// insertion point
	to.Name = from.Name
	to.Floatfield = from.Floatfield
	to.Floatfield2 = from.Floatfield2
	to.Intfield = from.Intfield
}

type Dstruct_WOP struct {
	// insertion point

	Name string
}

func (from *Dstruct) CopyBasicFields(to *Dstruct) {
	// insertion point
	to.Name = from.Name
}

type F0123456789012345678901234567890_WOP struct {
	// insertion point

	Name string

	Date time.Time
}

func (from *F0123456789012345678901234567890) CopyBasicFields(to *F0123456789012345678901234567890) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
}

type Gstruct_WOP struct {
	// insertion point

	Name string

	Floatfield float64

	Floatfield2 float64

	Intfield int
}

func (from *Gstruct) CopyBasicFields(to *Gstruct) {
	// insertion point
	to.Name = from.Name
	to.Floatfield = from.Floatfield
	to.Floatfield2 = from.Floatfield2
	to.Intfield = from.Intfield
}
