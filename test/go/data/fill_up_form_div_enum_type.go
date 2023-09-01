package data

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func FillUpFormDivEnumStringType[T models.PointerToGongstruct, TF models.GongstructEnumStringField](
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

		if string(field) == optionValue {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}

func FillUpFormDivEnumIntType[T models.PointerToGongstruct, TF models.GongstructEnumIntField](
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
