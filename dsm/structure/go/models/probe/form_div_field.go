// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func FormDivBasicFieldToField[TF models.GongtructBasicField](field *TF, formDiv *form.FormDiv) {
	gongprobe.FormDivBasicFieldToField(field, formDiv)
}

func FormDivTimeFieldToField(field *time.Time, formDiv *form.FormDiv, isTimeFormOnly bool) {
	gongprobe.FormDivTimeFieldToField(field, formDiv, isTimeFormOnly)
}

func FormDivEnumStringFieldToField[TF models.PointerToGongstructEnumStringField](field TF, formDiv *form.FormDiv) {
	gongprobe.FormDivEnumStringFieldToField(field, formDiv)
}

func FormDivEnumIntFieldToField[TF models.PointerToGongstructEnumIntField](field TF, formDiv *form.FormDiv) {
	gongprobe.FormDivEnumIntFieldToField(field, formDiv)
}

func FormDivSelectFieldToField[TF models.PointerToGongstruct](field *TF, stageOfInterest *models.Stage, formDiv *form.FormDiv) {

	if formDiv.FormEditAssocButton != nil && formDiv.FormEditAssocButton.HasChanged {
		rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
		if err != nil {
			log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
		}
		var zero TF
		if len(rowIDs) == 0 {
			*field = zero
			return
		}
		map_RowID_ID := GetMap_RowID_ID[TF](stageOfInterest)
		lastRowID := rowIDs[len(rowIDs)-1]
		if id, ok := map_RowID_ID[int(lastRowID)]; ok {
			for _instance := range *stageOfInterest.GetInstancesSet[TF]() {
				if stageOfInterest.GetOrder(_instance) == id {
					*field = any(_instance).(TF)
					return
				}
			}
		}
		*field = zero
		return
	}

	if formDiv.FormFields[0].FormFieldSelect.Value == nil {
		var zero TF
		if *field != zero {
			*field = zero
		}
	} else {
		selectedValue := formDiv.FormFields[0].FormFieldSelect.Value.GetName()

		nameCount := make(map[string]int)
		for _instance := range *stageOfInterest.GetInstancesSet[TF]() {
			nameCount[any(_instance).(TF).GetName()]++
		}

		for _instance := range *stageOfInterest.GetInstancesSet[TF]() {
			inst := any(_instance).(TF)
			if GetAssociationOptionName(inst, stageOfInterest, nameCount) == selectedValue {
				*field = inst
				return
			}
		}
	}
}

func addTimeComponents(x, y time.Time) time.Time {
	return gongprobe.AddTimeComponents(x, y)
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
