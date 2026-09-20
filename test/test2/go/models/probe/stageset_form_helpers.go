// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

// StageSetBasicFieldtoForm appends a basic field to the form
func StageSetBasicFieldtoForm(
	fieldName string,
	field any,
	formStage *form.Stage,
	formGroup *form.FormGroup,
) {
	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	if boolVal, ok := field.(bool); ok {
		checkBox := (&form.CheckBox{
			Name:  fieldName,
			Value: boolVal,
		}).Stage(formStage)
		formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
		return
	}

	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	switch fieldVal := field.(type) {
	case string:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "string",
			Value: fieldVal,
		}).Stage(formStage)
	case int:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: fieldVal,
		}).Stage(formStage)
	case int8:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int16:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int32:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int64:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint8:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint16:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint32:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint64:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case float32:
		formField.FormFieldFloat64 = (&form.FormFieldFloat64{
			Name:  "float",
			Value: float64(fieldVal),
		}).Stage(formStage)
	case float64:
		formField.FormFieldFloat64 = (&form.FormFieldFloat64{
			Name:  "float",
			Value: fieldVal,
		}).Stage(formStage)
	case time.Time:
		formField.FormFieldDate = (&form.FormFieldDate{
			Name:  fieldName + "Date",
			Value: fieldVal,
		}).Stage(formStage)
	case time.Duration:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "duration",
			Value: fieldVal.String(),
		}).Stage(formStage)
	default:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "string",
			Value: fmt.Sprintf("%v", fieldVal),
		}).Stage(formStage)
	}
}

// StageSetAssociationFieldToForm appends a FormFieldSelect dropdown for a pointer field
func StageSetAssociationFieldToForm[T interface {
	GetName() string
	comparable
}](
	fieldName string,
	field T,
	formGroup *form.FormGroup,
	instancesSet *map[T]struct{},
	formStage *form.Stage,
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

	formFieldSelect.Options = make([]*form.Option, 0)
	if instancesSet != nil {
		instancesSorted := make([]T, 0, len(*instancesSet))
		for inst := range *instancesSet {
			instancesSorted = append(instancesSorted, inst)
		}
		sort.Slice(instancesSorted, func(i, j int) bool {
			return instancesSorted[i].GetName() < instancesSorted[j].GetName()
		})
		for _, instance := range instancesSorted {
			option := (&form.Option{
				Name: instance.GetName(),
			}).Stage(formStage)

			if instance == field {
				formFieldSelect.Value = option
			}
			formFieldSelect.Options = append(formFieldSelect.Options, option)
		}
	}
}

// StageSetFormDivSelectFieldToField updates a pointer field from the selected dropdown option
func StageSetFormDivSelectFieldToField[T interface {
	GetName() string
	comparable
}](
	field *T,
	instancesSet *map[T]struct{},
	formDiv *form.FormDiv,
) {
	if len(formDiv.FormFields) == 0 || formDiv.FormFields[0].FormFieldSelect == nil || formDiv.FormFields[0].FormFieldSelect.Value == nil {
		var zero T
		*field = zero
		return
	}
	selectedName := formDiv.FormFields[0].FormFieldSelect.Value.GetName()
	var zero T
	*field = zero
	if instancesSet != nil {
		for inst := range *instancesSet {
			if inst.GetName() == selectedName {
				*field = inst
				return
			}
		}
	}
}

// StageSetAssociationReverseFieldToForm shows reverse references pointing to this instance
func StageSetAssociationReverseFieldToForm(
	ownerStructName string,
	fieldName string,
	referencingNames []string,
	formGroup *form.FormGroup,
	formStage *form.Stage,
) {
	formDiv := (&form.FormDiv{
		Name: ownerStructName + ":" + fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	val := strings.Join(referencingNames, ", ")
	formField := (&form.FormField{
		Name:  ownerStructName + ":" + fieldName,
		Label: "(" + ownerStructName + ") -> " + fieldName,
		FormFieldString: &form.FormFieldString{
			Value: val,
		},
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)
}
