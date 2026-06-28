// generated code - do not edit
package probe

import (
	"math"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/books/go/models"
)

func BasicFieldtoForm[TF models.GongtructBasicField](
	fieldName string, field TF, instance models.GongstructIF, formStage *form.Stage, formGroup *form.FormGroup,
	isTextArea bool, isBespokeWidth bool, bespokeWidth int, isBespokeHeight bool, bespokeHeight int, isTimeFormOnly bool,
) {

	switch fieldWithInterferedType := any(field).(type) {
	case string:
		formDiv := (&form.FormDiv{
			Name: fieldName,
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&form.FormField{
			Name:             fieldName,
			Label:            fieldName,
			Placeholder:      "",
			HasBespokeWidth:  isBespokeWidth,
			BespokeWidthPx:   bespokeWidth,
			HasBespokeHeight: isBespokeHeight,
			BespokeHeightPx:  bespokeHeight,
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldString := (&form.FormFieldString{
			Name:       "string",
			Value:      fieldWithInterferedType,
			IsTextArea: isTextArea,
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
		if !isTimeFormOnly {
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
			checkBox := (&form.CheckBox{
				Name:  fieldName + " - negative",
				Value: fieldWithInterferedType < 0,
			}).Stage(formStage)
			formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
		}

		{
			formFieldDays := (&form.FormField{
				Name:  fieldName + " - Days",
				Label: "Days",
			}).Stage(formStage)
			formFieldDays.HasBespokeWidth = true
			formFieldDays.BespokeWidthPx = 90

			formDiv.FormFields = append(formDiv.FormFields, formFieldDays)

			value := int(math.Abs(fieldWithInterferedType.Hours() / 24))
			formFieldIntDays := (&form.FormFieldInt{
				Name:  fieldName + " - Days",
				Value: value,
			}).Stage(formStage)
			formFieldIntDays.HasMinValidator = true
			formFieldIntDays.MinValue = 0
			formFieldDays.FormFieldInt = formFieldIntDays
		}

		{
			formFieldHours := (&form.FormField{
				Name:  fieldName + " - Hours",
				Label: "Hours",
			}).Stage(formStage)
			formFieldHours.HasBespokeWidth = true
			formFieldHours.BespokeWidthPx = 90

			formDiv.FormFields = append(formDiv.FormFields, formFieldHours)

			formFieldIntHours := (&form.FormFieldInt{
				Name:  fieldName + " - Hours",
				Value: int(math.Abs(fieldWithInterferedType.Hours())) % 24,
			}).Stage(formStage)
			formFieldIntHours.HasMaxValidator = true
			formFieldIntHours.MaxValue = 23
			formFieldIntHours.HasMinValidator = true
			formFieldIntHours.MinValue = 0
			formFieldHours.FormFieldInt = formFieldIntHours
		}

		{
			formFieldMinutes := (&form.FormField{
				Name:  fieldName + " - Minutes",
				Label: "Minutes",
			}).Stage(formStage)
			formFieldMinutes.HasBespokeWidth = true
			formFieldMinutes.BespokeWidthPx = 90
			formDiv.FormFields = append(formDiv.FormFields, formFieldMinutes)

			formFieldIntMinutes := (&form.FormFieldInt{
				Name:  fieldName + " - Minutes",
				Value: int(math.Abs(fieldWithInterferedType.Minutes())) % 60,
			}).Stage(formStage)
			formFieldIntMinutes.HasMaxValidator = true
			formFieldIntMinutes.MaxValue = 59
			formFieldIntMinutes.HasMinValidator = true
			formFieldIntMinutes.MinValue = 0
			formFieldMinutes.FormFieldInt = formFieldIntMinutes
		}
		{
			formFieldSeconds := (&form.FormField{
				Name:  fieldName + " - Seconds",
				Label: "Seconds",
			}).Stage(formStage)
			formFieldSeconds.HasBespokeWidth = true
			formFieldSeconds.BespokeWidthPx = 90
			formDiv.FormFields = append(formDiv.FormFields, formFieldSeconds)

			formFieldIntSeconds := (&form.FormFieldInt{
				Name:  fieldName + " - Seconds",
				Value: int(math.Abs(fieldWithInterferedType.Seconds())) % 60,
			}).Stage(formStage)
			formFieldIntSeconds.HasMaxValidator = true
			formFieldIntSeconds.MaxValue = 59
			formFieldIntSeconds.HasMinValidator = true
			formFieldIntSeconds.MinValue = 0
			formFieldSeconds.FormFieldInt = formFieldIntSeconds
		}
	}

}
