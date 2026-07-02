package models

type FormDiv struct {
	Name                string
	FormFields          []*FormField
	CheckBoxs           []*CheckBox
	FormEditAssocButton *FormEditAssocButton
	FormSortAssocButton *FormSortAssocButton

	// IsADivider indicates if the FormDiv is a divider
	IsADivider bool

	// indicates if the FormDiv is a start tab group
	IsAStartAccordionGroup bool

	// Name of the Accordion group (relevant if IsAStartAccordionGroup is true)
	AccordionGroupName string

	// indicates if the FormDiv is an end Accordion group
	IsAEndAccordionGroup bool
}
