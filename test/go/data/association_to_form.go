package data

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func AssociationFieldToForm[T models.PointerToGongstruct, TF models.PointerToGongstruct, TFinfered models.Gongstruct](
	fieldName string, field TF, stageOfInterest *models.StageStruct, instance T, formStage *form.StageStruct, formGroup *form.FormGroup,
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
		Name:       "association",
		CanBeEmpty: true,
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for _instance := range *models.GetGongstructInstancesSet[TFinfered](stageOfInterest) {
		option := (&form.Option{
			Name: any(_instance).(TF).GetName(),
		}).Stage(formStage)

		if any(_instance).(TF) == field {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
