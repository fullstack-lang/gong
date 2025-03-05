// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

func EnumTypeStringToForm[T models.PointerToGongstruct, TF models.GongstructEnumStringField](
	fieldName string, field TF, instance T, formStage *form.StageStruct, formGroup *form.FormGroup,
) {
	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name: "enum",
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for idx, optionCode := range field.Codes() {
		optionValue := field.CodeValues()[idx]

		option := (&form.Option{
			Name: optionCode,
		}).Stage(formStage)

		if field.ToString() == optionValue {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}

func EnumTypeIntToForm[T models.PointerToGongstruct, TF models.GongstructEnumIntField](
	fieldName string, field TF, instance T, formStage *form.StageStruct, formGroup *form.FormGroup,
) {
	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name: "enum",
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for idx, optionCode := range field.Codes() {
		optionValue := field.CodeValues()[idx]

		option := (&form.Option{
			Name: optionCode,
		}).Stage(formStage)

		if field == TF(optionValue) {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
