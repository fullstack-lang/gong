package models

type FormField struct {
	Name          string
	InputTypeEnum InputTypeEnum

	// label for the input field.
	// for instance "First Name" in <mat-label>First Name</mat-label>
	Label string

	// suggestion for the field
	Placeholder string

	FormFieldString  *FormFieldString
	FormFieldFloat64 *FormFieldFloat64
	FormFieldInt     *FormFieldInt

	FormFieldDate     *FormFieldDate
	FormFieldTime     *FormFieldTime
	FormFieldDateTime *FormFieldDateTime

	FormFieldSelect *FormFieldSelect

	// set up specific width
	HasBespokeWidth bool
	BespokeWidthPx  int

	// set up specific height
	HasBespokeHeight bool
	BespokeHeightPx  int
}
