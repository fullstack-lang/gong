// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Button:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Cell:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Cell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CellBoolean:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CellBoolean Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellBooleanFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CellFloat64:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CellFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFloat64FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CellIcon:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CellIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIconFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CellInt:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CellInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIntFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CellString:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CellString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellStringFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DisplayedColumn:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DisplayedColumn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplayedColumnFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Row:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Row Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SVGIcon:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SVGIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGIconFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Table:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
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
