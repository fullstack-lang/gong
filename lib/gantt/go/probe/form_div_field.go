// generated code - do not edit
package probe

import (
	"log"
	"time"

	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

func FormDivBasicFieldToField[TF models.GongtructBasicField](field *TF, formDiv *form.FormDiv) {

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
	case *time.Time:
		date := formDiv.FormFields[0].FormFieldDate.Value

		// in the angular form div, the time.Time is show twice, once for the Date and once for the Time
		// construing the date back, one needs to truncate the date, otherwise
		// hours, minutes, seconds and nanoseconds would be added twice
		date = date.Truncate(24 * time.Hour)

		time := formDiv.FormFields[1].FormFieldTime.Value
		*fieldWithInterferedType = addTimeComponents(date, time)
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

		if !isNeg {
			*fieldWithInterferedType = -*fieldWithInterferedType
		}

	}
}

func FormDivEnumStringFieldToField[TF models.PointerToGongstructEnumStringField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			log.Println("Unkwnown enum value", value.GetName())
		}
	}
}

func FormDivEnumIntFieldToField[TF models.PointerToGongstructEnumIntField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			log.Println("Unkwnown enum value", value.GetName())
		}
	}
}

func FormDivSelectFieldToField[TF models.PointerToGongstruct](field *TF, stageOfInterest *models.StageStruct, formDiv *form.FormDiv) {

	if formDiv.FormFields[0].FormFieldSelect.Value == nil {
		var zero TF
		if *field != zero {
			*field = zero
		}
	} else {
		for _instance := range *models.GetGongstructInstancesSetFromPointerType[TF](stageOfInterest) {
			if any(_instance).(TF).GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
				*field = any(_instance).(TF)
			}
		}
	}
}

func addTimeComponents(x, y time.Time) time.Time {
	h, m, s := y.Clock()
	x = x.Add(time.Duration(h) * time.Hour)
	x = x.Add(time.Duration(m) * time.Minute)
	x = x.Add(time.Duration(s) * time.Second)
	x = x.Add(time.Duration(y.Nanosecond()) * time.Nanosecond)
	return x
}
