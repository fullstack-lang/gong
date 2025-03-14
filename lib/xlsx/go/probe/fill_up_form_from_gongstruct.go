// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DisplaySelection Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplaySelectionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLCell:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XLCell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLCellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLFile:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XLFile Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLFileFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLRow:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XLRow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLRowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XLSheet:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XLSheet Form",
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
	formStage.Commit()
}
