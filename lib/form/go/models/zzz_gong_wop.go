// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type CheckBox_WOP struct {
	// insertion point

	Name string

	Value bool
}

func (from *CheckBox) GongCopyBasicFields(to *CheckBox) {
	// insertion point
	*to = *from
}

type FormDiv_WOP struct {
	// insertion point

	Name string

	IsADivider bool

	IsAStartAccordionGroup bool

	AccordionGroupName string

	IsAEndAccordionGroup bool
}

func (from *FormDiv) GongCopyBasicFields(to *FormDiv) {
	// insertion point
	to.Name = from.Name
	to.IsADivider = from.IsADivider
	to.IsAStartAccordionGroup = from.IsAStartAccordionGroup
	to.AccordionGroupName = from.AccordionGroupName
	to.IsAEndAccordionGroup = from.IsAEndAccordionGroup
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

func (from *FormEditAssocButton) GongCopyBasicFields(to *FormEditAssocButton) {
	// insertion point
	*to = *from
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

func (from *FormField) GongCopyBasicFields(to *FormField) {
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

func (from *FormFieldDate) GongCopyBasicFields(to *FormFieldDate) {
	// insertion point
	*to = *from
}

type FormFieldDateTime_WOP struct {
	// insertion point

	Name string

	Value time.Time
}

func (from *FormFieldDateTime) GongCopyBasicFields(to *FormFieldDateTime) {
	// insertion point
	*to = *from
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

func (from *FormFieldFloat64) GongCopyBasicFields(to *FormFieldFloat64) {
	// insertion point
	*to = *from
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

func (from *FormFieldInt) GongCopyBasicFields(to *FormFieldInt) {
	// insertion point
	*to = *from
}

type FormFieldSelect_WOP struct {
	// insertion point

	Name string

	CanBeEmpty bool

	PreserveInitialOrder bool
}

func (from *FormFieldSelect) GongCopyBasicFields(to *FormFieldSelect) {
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

func (from *FormFieldString) GongCopyBasicFields(to *FormFieldString) {
	// insertion point
	*to = *from
}

type FormFieldTime_WOP struct {
	// insertion point

	Name string

	Value time.Time

	Step float64
}

func (from *FormFieldTime) GongCopyBasicFields(to *FormFieldTime) {
	// insertion point
	*to = *from
}

type FormGroup_WOP struct {
	// insertion point

	Name string

	Label string

	TypeLabel string

	HasSuppressButton bool

	HasSuppressButtonBeenPressed bool
}

func (from *FormGroup) GongCopyBasicFields(to *FormGroup) {
	// insertion point
	to.Name = from.Name
	to.Label = from.Label
	to.TypeLabel = from.TypeLabel
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

func (from *FormSortAssocButton) GongCopyBasicFields(to *FormSortAssocButton) {
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

func (from *Option) GongCopyBasicFields(to *Option) {
	// insertion point
	*to = *from
}

// end of insertion point
