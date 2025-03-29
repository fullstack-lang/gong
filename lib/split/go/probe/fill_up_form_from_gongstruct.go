// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AsSplit:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AsSplit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AsSplitArea:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AsSplitArea Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitAreaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Button:
		formGroup := (&gongtable.FormGroup{
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
	case *models.Cursor:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Cursor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CursorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Doc:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Doc Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Form:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Form Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Load:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Load Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LoadFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slider:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Slider Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliderFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Split:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Split Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SplitFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Svg:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Svg Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Table:
		formGroup := (&gongtable.FormGroup{
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
	case *models.Tone:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tree:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TreeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.View:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "View Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewFormCallback(
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
