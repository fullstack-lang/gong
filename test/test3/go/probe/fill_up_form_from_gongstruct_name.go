// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A", true)
			} else {
				FillUpFormFromGongstruct(onSave.a, probe)
			}
		case *BFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "B", true)
			} else {
				FillUpFormFromGongstruct(onSave.b, probe)
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
	}
	formStage.Commit()
}
