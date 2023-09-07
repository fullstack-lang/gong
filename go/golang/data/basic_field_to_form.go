package data

const BasicFieldtoFormTemplate = `// generated code - do not edit
package data

import (
	"time"

	form "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
)

func BasicFieldtoForm[T models.PointerToGongstruct, TF models.GongtructBasicField](
	fieldName string, field TF, instance T, formStage *form.StageStruct, formGroup *form.FormGroup,
) {

	switch fieldWithInterferedType := any(field).(type) {
	case string:
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

		formFieldString := (&form.FormFieldString{
			Name:  "string",
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formField.FormFieldString = formFieldString
	case time.Time:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		{
			formFieldPartDate := (&form.FormField{
				Name:        fieldName + "Date",
				Label:       fieldName + "Date",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formFieldPartDate)

			formFieldDate := (&form.FormFieldDate{
				Name:  fieldName + "Date",
				Value: fieldWithInterferedType,
			}).Stage(formStage)
			formFieldPartDate.FormFieldDate = formFieldDate
		}
		{
			formFieldPartTime := (&form.FormField{
				Name:        fieldName + "Time",
				Label:       fieldName + "Time",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formFieldPartTime)

			formFieldTime := (&form.FormFieldTime{
				Name:  fieldName + "Time",
				Value: fieldWithInterferedType,
			}).Stage(formStage)
			formFieldPartTime.FormFieldTime = formFieldTime
		}
	case bool:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

		checkBox := (&form.CheckBox{
			Name:  fieldName,
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
	case int:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:  fieldName,
			Label: fieldName,
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldInt := (&form.FormFieldInt{
			Name:  fieldName,
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formField.FormFieldInt = formFieldInt
	case float64:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:  fieldName,
			Label: fieldName,
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldFloat64 := (&form.FormFieldFloat64{
			Name:  fieldName,
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formField.FormFieldFloat64 = formFieldFloat64
	case time.Duration:

		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

		{
			formFieldHours := (&form.FormField{
				Name:  "Hours",
				Label: "Hours",
			}).Stage(formStage)
			formFieldHours.HasBespokeWidth = true
			formFieldHours.BespokeWidthPx = 90

			formDiv.FormFields = append(formDiv.FormFields, formFieldHours)

			formFieldIntHours := (&form.FormFieldInt{
				Name:  "Hours",
				Value: int(fieldWithInterferedType.Hours()) % 24,
			}).Stage(formStage)
			formFieldIntHours.HasMaxValidator = true
			formFieldIntHours.MaxValue = 23
			formFieldIntHours.HasMinValidator = true
			formFieldIntHours.MinValue = 0
			formFieldHours.FormFieldInt = formFieldIntHours
		}

		{
			formFieldMinutes := (&form.FormField{
				Name:  "Minutes",
				Label: "Minutes",
			}).Stage(formStage)
			formFieldMinutes.HasBespokeWidth = true
			formFieldMinutes.BespokeWidthPx = 90
			formDiv.FormFields = append(formDiv.FormFields, formFieldMinutes)

			formFieldIntMinutes := (&form.FormFieldInt{
				Name:  "Minutes",
				Value: int(fieldWithInterferedType.Minutes()) % 60,
			}).Stage(formStage)
			formFieldIntMinutes.HasMaxValidator = true
			formFieldIntMinutes.MaxValue = 59
			formFieldIntMinutes.HasMinValidator = true
			formFieldIntMinutes.MinValue = 0
			formFieldMinutes.FormFieldInt = formFieldIntMinutes
		}
		{
			formFieldSeconds := (&form.FormField{
				Name:  "Seconds",
				Label: "Seconds",
			}).Stage(formStage)
			formFieldSeconds.HasBespokeWidth = true
			formFieldSeconds.BespokeWidthPx = 90
			formDiv.FormFields = append(formDiv.FormFields, formFieldSeconds)

			formFieldIntSeconds := (&form.FormFieldInt{
				Name:  "Seconds",
				Value: int(fieldWithInterferedType.Seconds()) % 60,
			}).Stage(formStage)
			formFieldIntSeconds.HasMaxValidator = true
			formFieldIntSeconds.MaxValue = 59
			formFieldIntSeconds.HasMinValidator = true
			formFieldIntSeconds.MinValue = 0
			formFieldSeconds.FormFieldInt = formFieldIntSeconds
		}
	}

}
`
