// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Content:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.JpgImage:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "JpgImage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__JpgImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PngImage:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "PngImage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PngImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SvgImage:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SvgImage Form",
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
	formStage.Commit()
}
