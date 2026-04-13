// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *Form2FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Form2", true)
			} else {
				FillUpFormFromGongstruct(onSave.form2, probe)
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
	case "Form2":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Form2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Form2FormCallback(
			nil,
			probe,
			formGroup,
		)
		form2 := new(models.Form2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(form2, formGroup, probe)
	}
	formStage.Commit()
}
