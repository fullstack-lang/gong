package data

import (
	"time"

	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func FillUpFormDivBasicField[T models.PointerToGongstruct, TF models.GongtructBasicField](
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
				Name:        "Date",
				Label:       "Date",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formFieldPartDate)

			formFieldDate := (&form.FormFieldDate{
				Name:  "Date",
				Value: fieldWithInterferedType,
			}).Stage(formStage)
			formFieldPartDate.FormFieldDate = formFieldDate
		}
		{
			formFieldPartTime := (&form.FormField{
				Name:        "Time",
				Label:       "Time",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formFieldPartTime)

			formFieldTime := (&form.FormFieldTime{
				Name:  "Time",
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
			Name:  "bool",
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
	case int:
		formDiv := (&form.FormDiv{
			Name: "Intfield",
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:  "Intfield",
			Label: "Intfield",
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldInt := (&form.FormFieldInt{
			Name:  "Intfield",
			Value: fieldWithInterferedType,
		}).Stage(formStage)
		formField.FormFieldInt = formFieldInt
	case float64:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:  "float64",
			Label: "float64",
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldFloat64 := (&form.FormFieldFloat64{
			Name:  "float64",
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
	case models.AEnumType:

		formDiv := (&form.FormDiv{
			Name: "Aenum",
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:        "Aenum",
			Label:       "Aenum",
			Placeholder: "",
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldSelect := (&form.FormFieldSelect{
			Name: "Aenum",
		}).Stage(formStage)
		formField.FormFieldSelect = formFieldSelect

		formField.FormFieldSelect.Options = make([]*form.Option, 0)
		{
			option := (&form.Option{
				Name: models.ENUM_VAL1.ToString(),
			}).Stage(formStage)
			if fieldWithInterferedType == models.ENUM_VAL1 {
				formFieldSelect.Value = option
			}

			formField.FormFieldSelect.Options =
				append(formField.FormFieldSelect.Options, option)
		}
		{
			option := (&form.Option{
				Name: models.ENUM_VAL2.ToString(),
			}).Stage(formStage)
			if fieldWithInterferedType == models.ENUM_VAL2 {
				formFieldSelect.Value = option
			}
			formField.FormFieldSelect.Options =
				append(formField.FormFieldSelect.Options, option)
		}

	}

}
