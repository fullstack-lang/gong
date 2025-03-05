// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	case "SplitArea":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SplitArea Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SplitAreaFormCallback(
			nil,
			probe,
			formGroup,
		)
		splitarea := new(models.SplitArea)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(splitarea, formGroup, probe)
	}
	formStage.Commit()
}
