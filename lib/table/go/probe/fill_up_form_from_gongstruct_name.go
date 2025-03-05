// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

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
	case "Cell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Cell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFormCallback(
			nil,
			probe,
			formGroup,
		)
		cell := new(models.Cell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cell, formGroup, probe)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CellBoolean Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellBooleanFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellboolean := new(models.CellBoolean)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellboolean, formGroup, probe)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CellFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFloat64FormCallback(
			nil,
			probe,
			formGroup,
		)
		cellfloat64 := new(models.CellFloat64)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellfloat64, formGroup, probe)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CellIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellicon := new(models.CellIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellicon, formGroup, probe)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CellInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIntFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellint := new(models.CellInt)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellint, formGroup, probe)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CellString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellStringFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellstring := new(models.CellString)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellstring, formGroup, probe)
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
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
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DisplayedColumn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplayedColumnFormCallback(
			nil,
			probe,
			formGroup,
		)
		displayedcolumn := new(models.DisplayedColumn)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(displayedcolumn, formGroup, probe)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
			Name:  form.FormGroupDefaultName.ToString(),
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
	case "Row":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Row Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RowFormCallback(
			nil,
			probe,
			formGroup,
		)
		row := new(models.Row)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(row, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			probe,
			formGroup,
		)
		table := new(models.Table)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(table, formGroup, probe)
	}
	formStage.Commit()
}
