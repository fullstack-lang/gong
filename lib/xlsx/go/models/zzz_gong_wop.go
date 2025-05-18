// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour
var _ = __GONG_time_The_fool_doth_think_he_is_wise__

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type DisplaySelection_WOP struct {
	// insertion point

	Name string
}

func (from *DisplaySelection) CopyBasicFields(to *DisplaySelection) {
	// insertion point
	to.Name = from.Name
}

type XLCell_WOP struct {
	// insertion point

	Name string

	X int

	Y int
}

func (from *XLCell) CopyBasicFields(to *XLCell) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type XLFile_WOP struct {
	// insertion point

	Name string

	NbSheets int
}

func (from *XLFile) CopyBasicFields(to *XLFile) {
	// insertion point
	to.Name = from.Name
	to.NbSheets = from.NbSheets
}

type XLRow_WOP struct {
	// insertion point

	Name string

	RowIndex int
}

func (from *XLRow) CopyBasicFields(to *XLRow) {
	// insertion point
	to.Name = from.Name
	to.RowIndex = from.RowIndex
}

type XLSheet_WOP struct {
	// insertion point

	Name string

	MaxRow int

	MaxCol int

	NbRows int
}

func (from *XLSheet) CopyBasicFields(to *XLSheet) {
	// insertion point
	to.Name = from.Name
	to.MaxRow = from.MaxRow
	to.MaxCol = from.MaxCol
	to.NbRows = from.NbRows
}

