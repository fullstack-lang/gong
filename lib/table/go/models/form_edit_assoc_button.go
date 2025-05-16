package models

// FormEditAssocButton is a button on the
// front end to edit a N-N association
// the association is stored in the object and
type FormEditAssocButton struct {
	Name string

	Label string

	// AssociationStorage is the encoding into a string of the IDs of the associatied instances
	// the format is json
	AssociationStorage string

	// HasChanged is true when the end user has used this form
	// When true, the instance will be updated to the back
	HasChanged bool

	// swagger:ignore
	OnAssocEditon FormEditAssocButtonInterface
}

// OnAfterUpdate is called when the button is pressed
func (formEditAssocButton *FormEditAssocButton) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *FormEditAssocButton) {

	if stagedInstance.OnAssocEditon != nil {
		stagedInstance.OnAssocEditon.OnButtonPressed()
	}
}

type FormEditAssocButtonInterface interface {
	OnButtonPressed()
}
