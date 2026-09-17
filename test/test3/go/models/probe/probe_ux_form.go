// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "A":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AFormCallback(
			nil,
			probe,
			formGroup,
		)
		a := new(models.A)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a, formGroup, probe)
	case "B":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "B Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BFormCallback(
			nil,
			probe,
			formGroup,
		)
		b := new(models.B)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(b, formGroup, probe)
	case "C":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "C Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CFormCallback(
			nil,
			probe,
			formGroup,
		)
		c := new(models.C)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(c, formGroup, probe)
	}
	formStage.Commit()
}
