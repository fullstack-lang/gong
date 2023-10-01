// generated code - do not edit
package x

import (
	"github.com/fullstack-lang/gong/test2/go/models/x"
	form "github.com/fullstack-lang/gongtable/go/models"
)

// AssociationFieldToForm will append a div to the form
// with the values of options equal to the name of each possible instances in the association type
func AssociationFieldToForm[FieldType x.PointerToGongstruct](
	fieldName string, field FieldType, formGroup *form.FormGroup, playground *Playground,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(playground.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(playground.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       "association",
		CanBeEmpty: true,
	}).Stage(playground.formStage)
	formField.FormFieldSelect = formFieldSelect

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for instance := range *x.GetGongstructInstancesSetFromPointerType[FieldType](playground.stageOfInterest) {
		option := (&form.Option{
			Name: instance.GetName(),
		}).Stage(playground.formStage)

		// set up select value if field matches the instance
		if instance == field {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}

func AssociationReverseFieldToForm[OwnerType x.PointerToGongstruct, FieldType x.PointerToGongstruct](
	owner OwnerType,
	fieldName string,
	instance FieldType,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	formDiv := (&form.FormDiv{
		Name: x.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
	}).Stage(playground.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        x.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Label:       x.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Placeholder: "",
	}).Stage(playground.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       x.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		CanBeEmpty: true,
	}).Stage(playground.formStage)
	formField.FormFieldSelect = formFieldSelect

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for _instance := range *x.GetGongstructInstancesSetFromPointerType[OwnerType](playground.stageOfInterest) {
		option := (&form.Option{
			Name: _instance.GetName(),
		}).Stage(playground.formStage)

		// set up select value if field matches the instance
		if owner != nil && _instance == owner {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
