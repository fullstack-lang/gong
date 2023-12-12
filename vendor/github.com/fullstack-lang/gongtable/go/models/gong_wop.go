// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Cell_WOP struct {
	// insertion point
	Name string
}

type CellBoolean_WOP struct {
	// insertion point
	Name string
	Value bool
}

type CellFloat64_WOP struct {
	// insertion point
	Name string
	Value float64
}

type CellIcon_WOP struct {
	// insertion point
	Name string
	Icon string
}

type CellInt_WOP struct {
	// insertion point
	Name string
	Value int
}

type CellString_WOP struct {
	// insertion point
	Name string
	Value string
}

type CheckBox_WOP struct {
	// insertion point
	Name string
	Value bool
}

type DisplayedColumn_WOP struct {
	// insertion point
	Name string
}

type FormDiv_WOP struct {
	// insertion point
	Name string
}

type FormEditAssocButton_WOP struct {
	// insertion point
	Name string
	Label string
}

type FormField_WOP struct {
	// insertion point
	Name string
	InputTypeEnum InputTypeEnum
	Label string
	Placeholder string
	HasBespokeWidth bool
	BespokeWidthPx int
	HasBespokeHeight bool
	BespokeHeightPx int
}

type FormFieldDate_WOP struct {
	// insertion point
	Name string
	Value time.Time
}

type FormFieldDateTime_WOP struct {
	// insertion point
	Name string
	Value time.Time
}

type FormFieldFloat64_WOP struct {
	// insertion point
	Name string
	Value float64
	HasMinValidator bool
	MinValue float64
	HasMaxValidator bool
	MaxValue float64
}

type FormFieldInt_WOP struct {
	// insertion point
	Name string
	Value int
	HasMinValidator bool
	MinValue int
	HasMaxValidator bool
	MaxValue int
}

type FormFieldSelect_WOP struct {
	// insertion point
	Name string
	CanBeEmpty bool
}

type FormFieldString_WOP struct {
	// insertion point
	Name string
	Value string
	IsTextArea bool
}

type FormFieldTime_WOP struct {
	// insertion point
	Name string
	Value time.Time
	Step float64
}

type FormGroup_WOP struct {
	// insertion point
	Name string
	Label string
	HasSuppressButton bool
	HasSuppressButtonBeenPressed bool
}

type FormSortAssocButton_WOP struct {
	// insertion point
	Name string
	Label string
}

type Option_WOP struct {
	// insertion point
	Name string
}

type Row_WOP struct {
	// insertion point
	Name string
	IsChecked bool
}

type Table_WOP struct {
	// insertion point
	Name string
	HasFiltering bool
	HasColumnSorting bool
	HasPaginator bool
	HasCheckableRows bool
	HasSaveButton bool
	CanDragDropRows bool
	HasCloseButton bool
	SavingInProgress bool
	NbOfStickyColumns int
}

