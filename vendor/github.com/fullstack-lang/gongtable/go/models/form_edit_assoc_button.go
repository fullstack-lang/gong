package models

import "log"

// FormEditAssocButton is a button on the
// front end to edit a 0..1-N association
// when submitted, it will
type FormEditAssocButton struct {
	Name string

	Label string

	OnAssocEditon FormEditAssocButtonInterface
}

// OnAfterUpdate is called when the button is pressed
func (formEditAssocButton *FormEditAssocButton) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormEditAssocButton) {

	log.Println("OnAfterUpdate")

	if stagedInstance.OnAssocEditon != nil {
		stagedInstance.OnAssocEditon.OnButtonPressed()
	}
}

type FormEditAssocButtonInterface interface {
	OnButtonPressed()
}
