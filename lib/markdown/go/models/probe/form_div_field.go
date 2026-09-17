// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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
		*field = addTimeComponents(date, time)
	} else {
		*field = date
	}
}

func FormDivEnumStringFieldToField[TF models.PointerToGongstructEnumStringField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			// log.Println("Unkwnown enum value", value.GetName())
		}
	}
}

func FormDivEnumIntFieldToField[TF models.PointerToGongstructEnumIntField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			// log.Println("Unkwnown enum value", value.GetName())
		}
	}
}

func FormDivSelectFieldToField[TF models.PointerToGongstruct](field *TF, stageOfInterest *models.Stage, formDiv *form.FormDiv) {

	if formDiv.FormFields[0].FormFieldSelect.Value == nil {
		var zero TF
		if *field != zero {
			*field = zero
		}
	} else {
		for _instance := range *stageOfInterest.GetInstancesSet[TF]() {
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

func FormDivSliceOfPointersToField[AssocType models.PointerToGongstruct](
	owner any,
	fieldName string,
	sliceField *[]AssocType,
	formDiv *form.FormDiv,
	probe *Probe,
) {
	if formDiv.FormEditAssocButton == nil {
		return
	}
	instanceSet := *probe.stageOfInterest.GetInstancesSet[AssocType]()
	instanceSlice := make([]AssocType, 0)

	// make a map of all instances by their ID
	map_id_instances := make(map[uint]AssocType)

	for instance := range instanceSet {
		id := probe.stageOfInterest.GetOrder(instance)
		map_id_instances[id] = instance
	}

	rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
	if err != nil {
		log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
	}
	map_RowID_ID := GetMap_RowID_ID[AssocType](probe.stageOfInterest)

	for _, rowID := range rowIDs {
		if id, ok := map_RowID_ID[int(rowID)]; ok {
			instanceSlice = append(instanceSlice, map_id_instances[id])
		} else {
			log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
		}
	}
	*sliceField = instanceSlice
	probe.UpdateSliceOfPointersCallback(owner, fieldName, sliceField)
}

func FormDivReverseSliceOfPointersToField[OwnerType models.PointerToGongstruct, TargetType models.PointerToGongstruct](
	target TargetType,
	formDiv *form.FormDiv,
	probe *Probe,
	fieldName string,
	getSlice func(owner OwnerType) *[]TargetType,
) {
	if formDiv.FormEditAssocButton == nil {
		return
	}
	rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
	if err != nil {
		log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
	}

	map_RowID_ID := GetMap_RowID_ID[OwnerType](probe.stageOfInterest)
	targetOwnerIDs := make(map[uint]bool)
	for _, rowID := range rowIDs {
		if id, ok := map_RowID_ID[int(rowID)]; ok {
			targetOwnerIDs[id] = true
		} else {
			log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
		}
	}

	for owner := range *probe.stageOfInterest.GetInstancesSet[OwnerType]() {
		id := probe.stageOfInterest.GetOrder(owner)
		slicePtr := getSlice(owner)
		if targetOwnerIDs[id] {
			if !slices.Contains(*slicePtr, target) {
				*slicePtr = append(*slicePtr, target)
				probe.UpdateSliceOfPointersCallback(owner, fieldName, slicePtr)
			}
		} else {
			if idx := slices.Index(*slicePtr, target); idx != -1 {
				*slicePtr = slices.Delete(*slicePtr, idx, idx+1)
				probe.UpdateSliceOfPointersCallback(owner, fieldName, slicePtr)
			}
		}
	}
}
