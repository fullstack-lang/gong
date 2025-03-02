// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Button_WOP struct {
	// insertion point
	Name string
	Icon string
}

func (from *Button) CopyBasicFields(to *Button) {
	// insertion point
	to.Name = from.Name
	to.Icon = from.Icon
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
	HasSecondCheckboxButton bool
	IsSecondCheckboxChecked bool
	IsSecondCheckboxDisabled bool
	TextAfterSecondCheckbox string
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
	to.HasSecondCheckboxButton = from.HasSecondCheckboxButton
	to.IsSecondCheckboxChecked = from.IsSecondCheckboxChecked
	to.IsSecondCheckboxDisabled = from.IsSecondCheckboxDisabled
	to.TextAfterSecondCheckbox = from.TextAfterSecondCheckbox
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

