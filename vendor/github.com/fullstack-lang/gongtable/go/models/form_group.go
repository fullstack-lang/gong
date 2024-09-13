package models

type FormGroup struct {
	Name     string
	Label    string
	FormDivs []*FormDiv

	// If yes, a suppress button is present
	HasSuppressButton bool

	// HasSuppressButtonBeenPressed is set to true
	// when the suppress button is pressed. It is the responsability
	// of the backend to reset this value to false
	HasSuppressButtonBeenPressed bool

	// swagger:ignore
	OnSave OnSaveInterface
}

// OnAfterUpdate is called when the form group is updated
func (formGroup *FormGroup) OnAfterUpdate(
	stage *StageStruct,
	stagedInstance, frontInstance *FormGroup) {

	if stagedInstance.OnSave != nil {
		stagedInstance.OnSave.OnSave()
	}
}

type OnSaveInterface interface {
	OnSave()
}
