// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
)

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
