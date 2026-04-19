// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ProcessFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Process", true)
			} else {
				FillUpFormFromGongstruct(onSave.process, probe)
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
	case "Process":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Process Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			probe,
			formGroup,
		)
		process := new(models.Process)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(process, formGroup, probe)
	}
	formStage.Commit()
}
