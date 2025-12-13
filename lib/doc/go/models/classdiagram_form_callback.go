package models

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"
)

type ClassDiagramFormCallback struct {
	classdiagram *Classdiagram
	formGroup    *form.FormGroup

	stager *Stager
}

func (classdiagramFormCallback *ClassDiagramFormCallback) OnSave() {

	// log.Println("ClassDiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.stager.formStage.Checkout()

	newValue := classdiagramFormCallback.formGroup.FormDivs[0].FormFields[0].FormFieldString.Value

	classdiagramFormCallback.classdiagram.Description = newValue
}
