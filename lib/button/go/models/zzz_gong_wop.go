// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

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

	HasToolTip bool

	ToolTipText string

	ToolTipPosition ToolTipPositionEnum
}

func (from *Button) GongCopyBasicFields(to *Button) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.Icon = from.Icon
	to.IsDisabled = from.IsDisabled
	to.Color = from.Color
	to.MatButtonType = from.MatButtonType
	to.MatButtonAppearance = from.MatButtonAppearance
	to.HasToolTip = from.HasToolTip
	to.ToolTipText = from.ToolTipText
	to.ToolTipPosition = from.ToolTipPosition
}

type ButtonToggle_WOP struct {
	// insertion point

	Name string

	Label string

	Icon string

	IsDisabled bool

	IsChecked bool
}

func (from *ButtonToggle) GongCopyBasicFields(to *ButtonToggle) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.Icon = from.Icon
	to.IsDisabled = from.IsDisabled
	to.IsChecked = from.IsChecked
}

type Group_WOP struct {
	// insertion point

	Name string

	Percentage float64

	NbColumns int
}

func (from *Group) GongCopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.Percentage = from.Percentage
	to.NbColumns = from.NbColumns
}

type GroupToogle_WOP struct {
	// insertion point

	Name string

	Percentage float64

	IsSingleSelector bool
}

func (from *GroupToogle) GongCopyBasicFields(to *GroupToogle) {
	// insertion point
	to.Name = from.Name
	to.Percentage = from.Percentage
	to.IsSingleSelector = from.IsSingleSelector
}

type Layout_WOP struct {
	// insertion point

	Name string
}

func (from *Layout) GongCopyBasicFields(to *Layout) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
