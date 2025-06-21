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

type Node_WOP struct {
	// insertion point

	Name string

	FontStyle FontStyleEnum

	BackgroundColor string

	IsExpanded bool

	HasCheckboxButton bool

	IsChecked bool

	IsCheckboxDisabled bool

	CheckboxHasToolTip bool

	CheckboxToolTipText string

	CheckboxToolTipPosition ToolTipPositionEnum

	HasSecondCheckboxButton bool

	IsSecondCheckboxChecked bool

	IsSecondCheckboxDisabled bool

	SecondCheckboxHasToolTip bool

	SecondCheckboxToolTipText string

	SecondCheckboxToolTipPosition ToolTipPositionEnum

	TextAfterSecondCheckbox string

	HasToolTip bool

	ToolTipText string

	ToolTipPosition ToolTipPositionEnum

	IsInEditMode bool

	IsNodeClickable bool

	IsWithPreceedingIcon bool

	PreceedingIcon string
}

func (from *Node) CopyBasicFields(to *Node) {
	// insertion point
	to.Name = from.Name
	to.FontStyle = from.FontStyle
	to.BackgroundColor = from.BackgroundColor
	to.IsExpanded = from.IsExpanded
	to.HasCheckboxButton = from.HasCheckboxButton
	to.IsChecked = from.IsChecked
	to.IsCheckboxDisabled = from.IsCheckboxDisabled
	to.CheckboxHasToolTip = from.CheckboxHasToolTip
	to.CheckboxToolTipText = from.CheckboxToolTipText
	to.CheckboxToolTipPosition = from.CheckboxToolTipPosition
	to.HasSecondCheckboxButton = from.HasSecondCheckboxButton
	to.IsSecondCheckboxChecked = from.IsSecondCheckboxChecked
	to.IsSecondCheckboxDisabled = from.IsSecondCheckboxDisabled
	to.SecondCheckboxHasToolTip = from.SecondCheckboxHasToolTip
	to.SecondCheckboxToolTipText = from.SecondCheckboxToolTipText
	to.SecondCheckboxToolTipPosition = from.SecondCheckboxToolTipPosition
	to.TextAfterSecondCheckbox = from.TextAfterSecondCheckbox
	to.HasToolTip = from.HasToolTip
	to.ToolTipText = from.ToolTipText
	to.ToolTipPosition = from.ToolTipPosition
	to.IsInEditMode = from.IsInEditMode
	to.IsNodeClickable = from.IsNodeClickable
	to.IsWithPreceedingIcon = from.IsWithPreceedingIcon
	to.PreceedingIcon = from.PreceedingIcon
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

type Tree_WOP struct {
	// insertion point

	Name string
}

func (from *Tree) CopyBasicFields(to *Tree) {
	// insertion point
	to.Name = from.Name
}

