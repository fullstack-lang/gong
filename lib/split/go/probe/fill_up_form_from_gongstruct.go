// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AsSplit:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
	case *models.Cursor:
		formGroup := (&form.FormGroup{
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
	case *models.FavIcon:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FavIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FavIconFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Form:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
	case *models.LogoOnTheLeft:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LogoOnTheLeft Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LogoOnTheLeftFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LogoOnTheRight:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LogoOnTheRight Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LogoOnTheRightFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Markdown:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Markdown Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkdownFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slider:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
	case *models.Title:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Title Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TitleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tone:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
	case *models.Xlsx:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Xlsx Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XlsxFormCallback(
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
