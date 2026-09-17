// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
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
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CheckBox Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CheckBoxFormCallback(
			nil,
			probe,
			formGroup,
		)
		checkbox := new(models.CheckBox)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(checkbox, formGroup, probe)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormDiv Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormDivFormCallback(
			nil,
			probe,
			formGroup,
		)
		formdiv := new(models.FormDiv)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formdiv, formGroup, probe)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormEditAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormEditAssocButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		formeditassocbutton := new(models.FormEditAssocButton)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formeditassocbutton, formGroup, probe)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfield := new(models.FormField)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfield, formGroup, probe)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldDate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfielddate := new(models.FormFieldDate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfielddate, formGroup, probe)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldDateTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateTimeFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfielddatetime := new(models.FormFieldDateTime)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfielddatetime, formGroup, probe)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFloat64FormCallback(
			nil,
			probe,
			formGroup,
		)
		formfieldfloat64 := new(models.FormFieldFloat64)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfieldfloat64, formGroup, probe)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldIntFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfieldint := new(models.FormFieldInt)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfieldint, formGroup, probe)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldSelect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldSelectFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfieldselect := new(models.FormFieldSelect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfieldselect, formGroup, probe)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldStringFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfieldstring := new(models.FormFieldString)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfieldstring, formGroup, probe)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormFieldTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldTimeFormCallback(
			nil,
			probe,
			formGroup,
		)
		formfieldtime := new(models.FormFieldTime)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formfieldtime, formGroup, probe)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		formgroup := new(models.FormGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formgroup, formGroup, probe)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FormSortAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormSortAssocButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		formsortassocbutton := new(models.FormSortAssocButton)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formsortassocbutton, formGroup, probe)
	case "Option":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Option Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OptionFormCallback(
			nil,
			probe,
			formGroup,
		)
		option := new(models.Option)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(option, formGroup, probe)
	}
	formStage.Commit()
}
