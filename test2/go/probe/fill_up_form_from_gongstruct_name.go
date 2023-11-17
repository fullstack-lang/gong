// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "A":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " A Form",
			OnSave: __gong__New__AFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		a := new(models.A)
		FillUpForm(a, formGroup, probe)
	case "B":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " B Form",
			OnSave: __gong__New__BFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		b := new(models.B)
		FillUpForm(b, formGroup, probe)
	}
	formStage.Commit()
}
