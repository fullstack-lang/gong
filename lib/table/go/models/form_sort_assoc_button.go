package models

// FormSortAssocButton is a button on the
// front end to sort a 0..1-N association
// when submitted, it will update the data
type FormSortAssocButton struct {
	Name string

	Label string

	// swagger:ignore
	OnSortEdition FormSortAssocButtonInterface

	HasToolTip  bool
	ToolTipText string
}

// OnAfterUpdate is called when the button is pressed
func (formSortAssocButton *FormSortAssocButton) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *FormSortAssocButton) {

	if stagedInstance.OnSortEdition != nil {
		stagedInstance.OnSortEdition.OnButtonPressed()
	}

}

type FormSortAssocButtonInterface interface {
	OnButtonPressed()
}
