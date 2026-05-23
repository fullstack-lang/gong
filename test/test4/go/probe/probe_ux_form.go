// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test4/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AstructFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Astruct", true)
			} else {
				FillUpFormFromGongstruct(onSave.astruct, probe)
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
	case "Astruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Astruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructFormCallback(
			nil,
			probe,
			formGroup,
		)
		astruct := new(models.Astruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(astruct, formGroup, probe)
	}
	formStage.Commit()
}
