// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DisplaySelection",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplaySelectionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLCell:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "XLCell",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLCellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLFile:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "XLFile",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLFileFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLRow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "XLRow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLRowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLSheet:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "XLSheet",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLSheetFormCallback(
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
