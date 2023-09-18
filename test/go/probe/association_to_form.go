// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

// AssociationFieldToForm will append a div to the form
// with the values of options equal to the name of each possible instances in the association type
func AssociationFieldToForm[FieldType models.PointerToGongstruct](
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
	for instance := range *models.GetGongstructInstancesSetFromPointerType[FieldType](playground.stageOfInterest) {
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

func AssociationReverseFieldToForm[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	owner OwnerType,
	fieldName string,
	instance FieldType,
	sliceField []FieldType,
	formGroup *form.FormGroup,
	playground *Playground,
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
	for _instance := range *models.GetGongstructInstancesSetFromPointerType[FieldType](playground.stageOfInterest) {
		option := (&form.Option{
			Name: _instance.GetName(),
		}).Stage(playground.formStage)

		// set up select value if field matches the instance
		if _instance == instance {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
