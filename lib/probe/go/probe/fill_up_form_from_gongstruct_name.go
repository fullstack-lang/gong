// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/probe/go/models"
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
	case "Probe2":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Probe2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Probe2FormCallback(
			nil,
			probe,
			formGroup,
		)
		probe2 := new(models.Probe2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(probe2, formGroup, probe)
	}
	formStage.Commit()
}
