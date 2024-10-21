// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test3/go/models"
)

// AssociationFieldToForm will append a div to the form
// with the values of options equal to the name of each possible instances in the association type
func AssociationFieldToForm[FieldType models.PointerToGongstruct](
	fieldName string, field FieldType, formGroup *form.FormGroup, probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(probe.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       "association",
		CanBeEmpty: true,
	}).Stage(probe.formStage)
	formField.FormFieldSelect = formFieldSelect

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for instance := range *models.GetGongstructInstancesSetFromPointerType[FieldType](probe.stageOfInterest) {
		option := (&form.Option{
			Name: instance.GetName(),
		}).Stage(probe.formStage)

		// set up select value if field matches the instance
		if instance == field {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}

func AssociationReverseFieldToForm[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	owner OwnerType,
	fieldName string,
	instance FieldType,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Label:       models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Placeholder: "",
	}).Stage(probe.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		CanBeEmpty: true,
	}).Stage(probe.formStage)
	formField.FormFieldSelect = formFieldSelect

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for _instance := range *models.GetGongstructInstancesSetFromPointerType[OwnerType](probe.stageOfInterest) {
		option := (&form.Option{
			Name: _instance.GetName(),
		}).Stage(probe.formStage)

		// set up select value if field matches the instance
		if any(owner) != nil && _instance == owner {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
