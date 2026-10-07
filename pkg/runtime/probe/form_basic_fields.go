package probe

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type GongstructEnumIntField interface {
	~int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(string) error
}

type PointerToGongstructEnumIntField interface {
	FromCodeString(string) error
}

func BasicFieldtoForm[TF GongstructBasicField](
	fieldName string, field TF, formStage *form.Stage, formGroup *form.FormGroup,
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

func EnumTypeStringToForm[TF GongstructEnumStringField](
	fieldName string, field TF, formStage *form.Stage, formGroup *form.FormGroup,
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

func EnumTypeIntToForm[TF GongstructEnumIntField](
	fieldName string, field TF, formStage *form.Stage, formGroup *form.FormGroup,
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

func FormDivBasicFieldToField[TF GongstructBasicField](field *TF, formDiv *form.FormDiv) {
	switch fieldWithInterferedType := any(field).(type) {
	case *string:
		newValue := formDiv.FormFields[0].FormFieldString.Value
		*fieldWithInterferedType = newValue
	case *bool:
		value := formDiv.CheckBoxs[0].Value
		*fieldWithInterferedType = value
	case *int:
		value := formDiv.FormFields[0].FormFieldInt.Value
		*fieldWithInterferedType = value
	case *float64:
		value := formDiv.FormFields[0].FormFieldFloat64.Value
		*fieldWithInterferedType = value

	case *time.Duration:
		isNeg := formDiv.CheckBoxs[0].Value

		days := formDiv.FormFields[0].FormFieldInt.Value
		hours := formDiv.FormFields[1].FormFieldInt.Value
		minutes := formDiv.FormFields[2].FormFieldInt.Value
		seconds := formDiv.FormFields[3].FormFieldInt.Value

		*fieldWithInterferedType =
			time.Duration(days)*time.Hour*24 +
				time.Duration(hours)*time.Hour +
				time.Duration(minutes)*time.Minute +
				time.Duration(seconds)*time.Second

		if isNeg {
			*fieldWithInterferedType = -*fieldWithInterferedType
		}
	}
}

func FormDivTimeFieldToField(field *time.Time, formDiv *form.FormDiv, isTimeFormOnly bool) {
	date := formDiv.FormFields[0].FormFieldDate.Value

	// in the angular form div, the time.Time is show twice, once for the Date and once for the Time
	// construing the date back, one needs to truncate the date, otherwise
	// hours, minutes, seconds and nanoseconds would be added twice
	date = date.Truncate(24 * time.Hour)

	if !isTimeFormOnly {
		time := formDiv.FormFields[1].FormFieldTime.Value
		*field = AddTimeComponents(date, time)
	} else {
		*field = date
	}
}

func FormDivEnumStringFieldToField[TF PointerToGongstructEnumStringField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		_ = field.FromCodeString(value.GetName())
	}
}

func FormDivEnumIntFieldToField[TF PointerToGongstructEnumIntField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		_ = field.FromCodeString(value.GetName())
	}
}

func AddTimeComponents(x, y time.Time) time.Time {
	h, m, s := y.Clock()
	x = x.Add(time.Duration(h) * time.Hour)
	x = x.Add(time.Duration(m) * time.Minute)
	x = x.Add(time.Duration(s) * time.Second)
	x = x.Add(time.Duration(y.Nanosecond()) * time.Nanosecond)
	return x
}

func EncodeIntSliceToString(data []uint) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal uint slice to JSON: %w", err)
	}
	return string(jsonData), nil
}

func DecodeStringToIntSlice(str string) ([]uint, error) {
	var decodedData []uint
	err := json.Unmarshal([]byte(str), &decodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON string to uint slice: %w", err)
	}
	return decodedData, nil
}
