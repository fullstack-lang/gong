// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type CheckBox_WOP struct {
	// insertion point

	Name string

	Value bool
}

func (from *CheckBox) CopyBasicFields(to *CheckBox) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Form2_WOP struct {
	// insertion point

	Name string
}

func (from *Form2) CopyBasicFields(to *Form2) {
	// insertion point
	to.Name = from.Name
}

type FormDiv_WOP struct {
	// insertion point

	Name string
}

func (from *FormDiv) CopyBasicFields(to *FormDiv) {
	// insertion point
	to.Name = from.Name
}

type FormEditAssocButton_WOP struct {
	// insertion point

	Name string

	Label string

	AssociationStorage string

	HasChanged bool

	IsForSavePurpose bool

	HasToolTip bool

	ToolTipText string

	MatTooltipShowDelay string
}

func (from *FormEditAssocButton) CopyBasicFields(to *FormEditAssocButton) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.AssociationStorage = from.AssociationStorage
	to.HasChanged = from.HasChanged
	to.IsForSavePurpose = from.IsForSavePurpose
	to.HasToolTip = from.HasToolTip
	to.ToolTipText = from.ToolTipText
	to.MatTooltipShowDelay = from.MatTooltipShowDelay
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

func (from *FormField) CopyBasicFields(to *FormField) {
	// insertion point
	to.Name = from.Name
	to.InputTypeEnum = from.InputTypeEnum
	to.Label = from.Label
	to.Placeholder = from.Placeholder
	to.HasBespokeWidth = from.HasBespokeWidth
	to.BespokeWidthPx = from.BespokeWidthPx
	to.HasBespokeHeight = from.HasBespokeHeight
	to.BespokeHeightPx = from.BespokeHeightPx
}

type FormFieldDate_WOP struct {
	// insertion point

	Name string

	Value time.Time
}

func (from *FormFieldDate) CopyBasicFields(to *FormFieldDate) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type FormFieldDateTime_WOP struct {
	// insertion point

	Name string

	Value time.Time
}

func (from *FormFieldDateTime) CopyBasicFields(to *FormFieldDateTime) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
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

func (from *FormFieldFloat64) CopyBasicFields(to *FormFieldFloat64) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
	to.HasMinValidator = from.HasMinValidator
	to.MinValue = from.MinValue
	to.HasMaxValidator = from.HasMaxValidator
	to.MaxValue = from.MaxValue
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

func (from *FormFieldInt) CopyBasicFields(to *FormFieldInt) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
	to.HasMinValidator = from.HasMinValidator
	to.MinValue = from.MinValue
	to.HasMaxValidator = from.HasMaxValidator
	to.MaxValue = from.MaxValue
}

type FormFieldSelect_WOP struct {
	// insertion point

	Name string

	CanBeEmpty bool

	PreserveInitialOrder bool
}

func (from *FormFieldSelect) CopyBasicFields(to *FormFieldSelect) {
	// insertion point
	to.Name = from.Name
	to.CanBeEmpty = from.CanBeEmpty
	to.PreserveInitialOrder = from.PreserveInitialOrder
}

type FormFieldString_WOP struct {
	// insertion point

	Name string

	Value string

	IsTextArea bool
}

func (from *FormFieldString) CopyBasicFields(to *FormFieldString) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
	to.IsTextArea = from.IsTextArea
}

type FormFieldTime_WOP struct {
	// insertion point

	Name string

	Value time.Time

	Step float64
}

func (from *FormFieldTime) CopyBasicFields(to *FormFieldTime) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
	to.Step = from.Step
}

type FormGroup_WOP struct {
	// insertion point

	Name string

	Label string

	HasSuppressButton bool

	HasSuppressButtonBeenPressed bool
}

func (from *FormGroup) CopyBasicFields(to *FormGroup) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.HasSuppressButton = from.HasSuppressButton
	to.HasSuppressButtonBeenPressed = from.HasSuppressButtonBeenPressed
}

type FormSortAssocButton_WOP struct {
	// insertion point

	Name string

	Label string

	HasToolTip bool

	ToolTipText string

	MatTooltipShowDelay string
}

func (from *FormSortAssocButton) CopyBasicFields(to *FormSortAssocButton) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.HasToolTip = from.HasToolTip
	to.ToolTipText = from.ToolTipText
	to.MatTooltipShowDelay = from.MatTooltipShowDelay
}

type Option_WOP struct {
	// insertion point

	Name string
}

func (from *Option) CopyBasicFields(to *Option) {
	// insertion point
	to.Name = from.Name
}

