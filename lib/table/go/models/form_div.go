package models

type FormDiv struct {
	Name                string
	FormFields          []*FormField
	CheckBoxs           []*CheckBox
	FormEditAssocButton *FormEditAssocButton
	FormSortAssocButton *FormSortAssocButton
}
