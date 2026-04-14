// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Button_WOP struct {
	// insertion point

	Name string

	Icon string

	IsDisabled bool

	HasToolTip bool

	ToolTipText string

	ToolTipPosition ToolTipPositionEnum
}

func (from *Button) CopyBasicFields(to *Button) {
	// insertion point
	to.Name = from.Name
	to.Icon = from.Icon
	to.IsDisabled = from.IsDisabled
	to.HasToolTip = from.HasToolTip
	to.ToolTipText = from.ToolTipText
	to.ToolTipPosition = from.ToolTipPosition
}

type Cell_WOP struct {
	// insertion point

	Name string
}

func (from *Cell) CopyBasicFields(to *Cell) {
	// insertion point
	to.Name = from.Name
}

type CellBoolean_WOP struct {
	// insertion point

	Name string

	Value bool
}

func (from *CellBoolean) CopyBasicFields(to *CellBoolean) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type CellFloat64_WOP struct {
	// insertion point

	Name string

	Value float64
}

func (from *CellFloat64) CopyBasicFields(to *CellFloat64) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type CellIcon_WOP struct {
	// insertion point

	Name string

	Icon string

	NeedsConfirmation bool

	ConfirmationMessage string
}

func (from *CellIcon) CopyBasicFields(to *CellIcon) {
	// insertion point
	to.Name = from.Name
	to.Icon = from.Icon
	to.NeedsConfirmation = from.NeedsConfirmation
	to.ConfirmationMessage = from.ConfirmationMessage
}

type CellInt_WOP struct {
	// insertion point

	Name string

	Value int
}

func (from *CellInt) CopyBasicFields(to *CellInt) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type CellString_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *CellString) CopyBasicFields(to *CellString) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type DisplayedColumn_WOP struct {
	// insertion point

	Name string
}

func (from *DisplayedColumn) CopyBasicFields(to *DisplayedColumn) {
	// insertion point
	to.Name = from.Name
}

type Row_WOP struct {
	// insertion point

	Name string

	IsChecked bool
}

func (from *Row) CopyBasicFields(to *Row) {
	// insertion point
	to.Name = from.Name
	to.IsChecked = from.IsChecked
}

type SVGIcon_WOP struct {
	// insertion point

	Name string

	SVG string
}

func (from *SVGIcon) CopyBasicFields(to *SVGIcon) {
	// insertion point
	to.Name = from.Name
	to.SVG = from.SVG
}

type Table_WOP struct {
	// insertion point

	Name string

	HasFiltering bool

	HasColumnSorting bool

	HasPaginator bool

	HasCheckableRows bool

	HasSaveButton bool

	SaveButtonLabel string

	CanDragDropRows bool

	HasCloseButton bool

	SavingInProgress bool

	NbOfStickyColumns int
}

func (from *Table) CopyBasicFields(to *Table) {
	// insertion point
	to.Name = from.Name
	to.HasFiltering = from.HasFiltering
	to.HasColumnSorting = from.HasColumnSorting
	to.HasPaginator = from.HasPaginator
	to.HasCheckableRows = from.HasCheckableRows
	to.HasSaveButton = from.HasSaveButton
	to.SaveButtonLabel = from.SaveButtonLabel
	to.CanDragDropRows = from.CanDragDropRows
	to.HasCloseButton = from.HasCloseButton
	to.SavingInProgress = from.SavingInProgress
	to.NbOfStickyColumns = from.NbOfStickyColumns
}

