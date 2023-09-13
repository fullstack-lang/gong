package models

import "log"

type FormGroup struct {
	Name     string
	Label    string
	FormDivs []*FormDiv

	// swagger:ignore
	OnSave OnSaveInterface
}

// OnAfterUpdate is called when the form group is updated
func (formGroup *FormGroup) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormGroup) {

	log.Println("FormGroup: OnAfterUpdate")

	if stagedInstance.OnSave != nil {
		stagedInstance.OnSave.OnSave()
	}
}

type OnSaveInterface interface {
	OnSave()
}
