// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Astruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AstructBstruct2Use:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AstructBstruct2Use Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructBstruct2UseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AstructBstructUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AstructBstructUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructBstructUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bstruct:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BstructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dstruct:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Dstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DstructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.F0123456789012345678901234567890:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "F0123456789012345678901234567890 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__F0123456789012345678901234567890FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Gstruct:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Gstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GstructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
