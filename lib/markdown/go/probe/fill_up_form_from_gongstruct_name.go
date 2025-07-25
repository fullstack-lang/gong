// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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
	case "Content":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		content := new(models.Content)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(content, formGroup, probe)
	}
	formStage.Commit()
}
