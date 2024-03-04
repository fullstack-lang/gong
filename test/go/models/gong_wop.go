// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Astruct_WOP struct {
	// insertion point
	Name string
	Date time.Time
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
	StructRef string
	FieldRef string
	EnumIntRef string
	EnumStringRef string
	EnumValue string
	ConstIdentifierValue string
	TextFieldBespokeSize string
	TextArea string
}

func (from *Astruct) CopyBasicFields(to *Astruct) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
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
	to.StructRef = from.StructRef
	to.FieldRef = from.FieldRef
	to.EnumIntRef = from.EnumIntRef
	to.EnumStringRef = from.EnumStringRef
	to.EnumValue = from.EnumValue
	to.ConstIdentifierValue = from.ConstIdentifierValue
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

type Fstruct_WOP struct {
	// insertion point
	Name string
}

func (from *Fstruct) CopyBasicFields(to *Fstruct) {
	// insertion point
	to.Name = from.Name
}

