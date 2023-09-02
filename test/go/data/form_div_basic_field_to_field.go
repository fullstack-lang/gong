package data

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
)

func FormDivBasicFieldToField[TF models.GongtructBasicField](field *TF, formDiv *form.FormDiv,
) {

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
		hours := formDiv.FormFields[0].FormFieldInt.Value
		minutes := formDiv.FormFields[1].FormFieldInt.Value
		seconds := formDiv.FormFields[2].FormFieldInt.Value

		*fieldWithInterferedType = time.Duration(hours)*time.Hour +
			time.Duration(minutes)*time.Minute +
			time.Duration(seconds)*time.Second

	}
}
