// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "CheckBox Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CheckBoxFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Form2:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Form2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Form2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormDiv:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormDiv Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormDivFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormEditAssocButton:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormEditAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormEditAssocButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormField:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldDate:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldDate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldDateTime:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldDateTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateTimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldFloat64:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFloat64FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldInt:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldIntFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldSelect:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldSelect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldSelectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldString:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldStringFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldTime:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormFieldTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldTimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormGroup:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormGroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormSortAssocButton:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FormSortAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormSortAssocButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Option:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Option Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OptionFormCallback(
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
