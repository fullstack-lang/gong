// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type DisplaySelection_WOP struct {
	// insertion point

	Name string
}

func (from *DisplaySelection) GongCopyBasicFields(to *DisplaySelection) {
	// insertion point
	to.Name = from.Name
}

type XLCell_WOP struct {
	// insertion point

	Name string

	X int

	Y int
}

func (from *XLCell) GongCopyBasicFields(to *XLCell) {
	// insertion point
	*to = *from
}

type XLFile_WOP struct {
	// insertion point

	Name string

	NbSheets int
}

func (from *XLFile) GongCopyBasicFields(to *XLFile) {
	// insertion point
	to.Name = from.Name
	to.NbSheets = from.NbSheets
}

type XLRow_WOP struct {
	// insertion point

	Name string

	RowIndex int
}

func (from *XLRow) GongCopyBasicFields(to *XLRow) {
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

func (from *XLSheet) GongCopyBasicFields(to *XLSheet) {
	// insertion point
	to.Name = from.Name
	to.MaxRow = from.MaxRow
	to.MaxCol = from.MaxCol
	to.NbRows = from.NbRows
}

// end of insertion point
