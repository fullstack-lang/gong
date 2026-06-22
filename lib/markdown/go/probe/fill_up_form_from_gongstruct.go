// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Content:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Content",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.JpgImage:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "JpgImage",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__JpgImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PngImage:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "PngImage",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PngImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SvgImage:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SvgImage",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgImageFormCallback(
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
