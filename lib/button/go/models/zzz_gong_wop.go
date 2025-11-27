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

	Label string

	Icon string

	IsDisabled bool

	Color MatButtonPaletteType

	MatButtonType MatButtonType

	MatButtonAppearance MatButtonAppearance
}

func (from *Button) CopyBasicFields(to *Button) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.Icon = from.Icon
	to.IsDisabled = from.IsDisabled
	to.Color = from.Color
	to.MatButtonType = from.MatButtonType
	to.MatButtonAppearance = from.MatButtonAppearance
}

type Group_WOP struct {
	// insertion point

	Name string

	Percentage float64

	NbColumns int
}

func (from *Group) CopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.Percentage = from.Percentage
	to.NbColumns = from.NbColumns
}

type Layout_WOP struct {
	// insertion point

	Name string
}

func (from *Layout) CopyBasicFields(to *Layout) {
	// insertion point
	to.Name = from.Name
}

