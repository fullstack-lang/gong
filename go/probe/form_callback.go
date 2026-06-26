// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__GongBasicFieldFormCallback(
	gongbasicfield *models.GongBasicField,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongbasicfieldFormCallback *GongBasicFieldFormCallback) {
	gongbasicfieldFormCallback = new(GongBasicFieldFormCallback)
	gongbasicfieldFormCallback.probe = probe
	gongbasicfieldFormCallback.gongbasicfield = gongbasicfield
	gongbasicfieldFormCallback.formGroup = formGroup

	gongbasicfieldFormCallback.CreationMode = (gongbasicfield == nil)

	return
}

type GongBasicFieldFormCallback struct {
	gongbasicfield *models.GongBasicField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongbasicfieldFormCallback *GongBasicFieldFormCallback) OnSave() {
	gongbasicfieldFormCallback.probe.stageOfInterest.Lock()
	defer gongbasicfieldFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongBasicFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongbasicfieldFormCallback.probe.formStage.Checkout()

	if gongbasicfieldFormCallback.gongbasicfield == nil {
		gongbasicfieldFormCallback.gongbasicfield = new(models.GongBasicField).Stage(gongbasicfieldFormCallback.probe.stageOfInterest)
	}
	gongbasicfield_ := gongbasicfieldFormCallback.gongbasicfield
	_ = gongbasicfield_

	for _, formDiv := range gongbasicfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongbasicfield_.Name), formDiv)
		case "BasicKindName":
			FormDivBasicFieldToField(&(gongbasicfield_.BasicKindName), formDiv)
		case "GongEnum":
			FormDivSelectFieldToField(&(gongbasicfield_.GongEnum), gongbasicfieldFormCallback.probe.stageOfInterest, formDiv)
		case "DeclaredType":
			FormDivBasicFieldToField(&(gongbasicfield_.DeclaredType), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongbasicfield_.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(gongbasicfield_.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(gongbasicfield_.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(gongbasicfield_.IsAccordionEnd), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongbasicfield_.Index), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(gongbasicfield_.IsTextArea), formDiv)
		case "IsBespokeWidth":
			FormDivBasicFieldToField(&(gongbasicfield_.IsBespokeWidth), formDiv)
		case "BespokeWidth":
			FormDivBasicFieldToField(&(gongbasicfield_.BespokeWidth), formDiv)
		case "IsBespokeHeight":
			FormDivBasicFieldToField(&(gongbasicfield_.IsBespokeHeight), formDiv)
		case "BespokeHeight":
			FormDivBasicFieldToField(&(gongbasicfield_.BespokeHeight), formDiv)
		case "GongStruct:GongBasicFields":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStruct](gongbasicfieldFormCallback.probe.stageOfInterest)
			targetGongStructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStruct instances and update their GongBasicFields slice
			for _gongstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStruct](gongbasicfieldFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongbasicfieldFormCallback.probe.stageOfInterest, _gongstruct)
				
				// if GongStruct is selected
				if targetGongStructIDs[id] {
					// ensure gongbasicfield_ is in _gongstruct.GongBasicFields
					found := false
					for _, _b := range _gongstruct.GongBasicFields {
						if _b == gongbasicfield_ {
							found = true
							break
						}
					}
					if !found {
						_gongstruct.GongBasicFields = append(_gongstruct.GongBasicFields, gongbasicfield_)
						gongbasicfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "GongBasicFields", &_gongstruct.GongBasicFields)
					}
				} else {
					// ensure gongbasicfield_ is NOT in _gongstruct.GongBasicFields
					idx := slices.Index(_gongstruct.GongBasicFields, gongbasicfield_)
					if idx != -1 {
						_gongstruct.GongBasicFields = slices.Delete(_gongstruct.GongBasicFields, idx, idx+1)
						gongbasicfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "GongBasicFields", &_gongstruct.GongBasicFields)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongbasicfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongbasicfield_.Unstage(gongbasicfieldFormCallback.probe.stageOfInterest)
	}

	gongbasicfieldFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongBasicField](
		gongbasicfieldFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongbasicfieldFormCallback.CreationMode || gongbasicfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongbasicfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongbasicfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongBasicFieldFormCallback(
			nil,
			gongbasicfieldFormCallback.probe,
			newFormGroup,
		)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, newFormGroup, gongbasicfieldFormCallback.probe)
		gongbasicfieldFormCallback.probe.formStage.Commit()
	}

	gongbasicfieldFormCallback.probe.ux_tree()
}
func __gong__New__GongEnumFormCallback(
	gongenum *models.GongEnum,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumFormCallback *GongEnumFormCallback) {
	gongenumFormCallback = new(GongEnumFormCallback)
	gongenumFormCallback.probe = probe
	gongenumFormCallback.gongenum = gongenum
	gongenumFormCallback.formGroup = formGroup

	gongenumFormCallback.CreationMode = (gongenum == nil)

	return
}

type GongEnumFormCallback struct {
	gongenum *models.GongEnum

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongenumFormCallback *GongEnumFormCallback) OnSave() {
	gongenumFormCallback.probe.stageOfInterest.Lock()
	defer gongenumFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongEnumFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumFormCallback.probe.formStage.Checkout()

	if gongenumFormCallback.gongenum == nil {
		gongenumFormCallback.gongenum = new(models.GongEnum).Stage(gongenumFormCallback.probe.stageOfInterest)
	}
	gongenum_ := gongenumFormCallback.gongenum
	_ = gongenum_

	for _, formDiv := range gongenumFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenum_.Name), formDiv)
		case "Type":
			FormDivEnumIntFieldToField(&(gongenum_.Type), formDiv)
		case "GongEnumValues":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumValue](gongenumFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongEnumValue, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongEnumValue)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongenumFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongEnumValue](gongenumFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongenum_.GongEnumValues = instanceSlice
			gongenumFormCallback.probe.UpdateSliceOfPointersCallback(gongenum_, "GongEnumValues", &gongenum_.GongEnumValues)

		}
	}

	// manage the suppress operation
	if gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenum_.Unstage(gongenumFormCallback.probe.stageOfInterest)
	}

	gongenumFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongEnum](
		gongenumFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongenumFormCallback.CreationMode || gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongenumFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumFormCallback(
			nil,
			gongenumFormCallback.probe,
			newFormGroup,
		)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, newFormGroup, gongenumFormCallback.probe)
		gongenumFormCallback.probe.formStage.Commit()
	}

	gongenumFormCallback.probe.ux_tree()
}
func __gong__New__GongEnumValueFormCallback(
	gongenumvalue *models.GongEnumValue,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumvalueFormCallback *GongEnumValueFormCallback) {
	gongenumvalueFormCallback = new(GongEnumValueFormCallback)
	gongenumvalueFormCallback.probe = probe
	gongenumvalueFormCallback.gongenumvalue = gongenumvalue
	gongenumvalueFormCallback.formGroup = formGroup

	gongenumvalueFormCallback.CreationMode = (gongenumvalue == nil)

	return
}

type GongEnumValueFormCallback struct {
	gongenumvalue *models.GongEnumValue

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongenumvalueFormCallback *GongEnumValueFormCallback) OnSave() {
	gongenumvalueFormCallback.probe.stageOfInterest.Lock()
	defer gongenumvalueFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongEnumValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueFormCallback.probe.formStage.Checkout()

	if gongenumvalueFormCallback.gongenumvalue == nil {
		gongenumvalueFormCallback.gongenumvalue = new(models.GongEnumValue).Stage(gongenumvalueFormCallback.probe.stageOfInterest)
	}
	gongenumvalue_ := gongenumvalueFormCallback.gongenumvalue
	_ = gongenumvalue_

	for _, formDiv := range gongenumvalueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalue_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(gongenumvalue_.Value), formDiv)
		case "GongEnum:GongEnumValues":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongEnum instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongEnum instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongEnum](gongenumvalueFormCallback.probe.stageOfInterest)
			targetGongEnumIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongEnumIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongEnum instances and update their GongEnumValues slice
			for _gongenum := range *models.GetGongstructInstancesSetFromPointerType[*models.GongEnum](gongenumvalueFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongenumvalueFormCallback.probe.stageOfInterest, _gongenum)
				
				// if GongEnum is selected
				if targetGongEnumIDs[id] {
					// ensure gongenumvalue_ is in _gongenum.GongEnumValues
					found := false
					for _, _b := range _gongenum.GongEnumValues {
						if _b == gongenumvalue_ {
							found = true
							break
						}
					}
					if !found {
						_gongenum.GongEnumValues = append(_gongenum.GongEnumValues, gongenumvalue_)
						gongenumvalueFormCallback.probe.UpdateSliceOfPointersCallback(_gongenum, "GongEnumValues", &_gongenum.GongEnumValues)
					}
				} else {
					// ensure gongenumvalue_ is NOT in _gongenum.GongEnumValues
					idx := slices.Index(_gongenum.GongEnumValues, gongenumvalue_)
					if idx != -1 {
						_gongenum.GongEnumValues = slices.Delete(_gongenum.GongEnumValues, idx, idx+1)
						gongenumvalueFormCallback.probe.UpdateSliceOfPointersCallback(_gongenum, "GongEnumValues", &_gongenum.GongEnumValues)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalue_.Unstage(gongenumvalueFormCallback.probe.stageOfInterest)
	}

	gongenumvalueFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongEnumValue](
		gongenumvalueFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongenumvalueFormCallback.CreationMode || gongenumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongenumvalueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumValueFormCallback(
			nil,
			gongenumvalueFormCallback.probe,
			newFormGroup,
		)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, newFormGroup, gongenumvalueFormCallback.probe)
		gongenumvalueFormCallback.probe.formStage.Commit()
	}

	gongenumvalueFormCallback.probe.ux_tree()
}
func __gong__New__GongLinkFormCallback(
	gonglink *models.GongLink,
	probe *Probe,
	formGroup *form.FormGroup,
) (gonglinkFormCallback *GongLinkFormCallback) {
	gonglinkFormCallback = new(GongLinkFormCallback)
	gonglinkFormCallback.probe = probe
	gonglinkFormCallback.gonglink = gonglink
	gonglinkFormCallback.formGroup = formGroup

	gonglinkFormCallback.CreationMode = (gonglink == nil)

	return
}

type GongLinkFormCallback struct {
	gonglink *models.GongLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gonglinkFormCallback *GongLinkFormCallback) OnSave() {
	gonglinkFormCallback.probe.stageOfInterest.Lock()
	defer gonglinkFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gonglinkFormCallback.probe.formStage.Checkout()

	if gonglinkFormCallback.gonglink == nil {
		gonglinkFormCallback.gonglink = new(models.GongLink).Stage(gonglinkFormCallback.probe.stageOfInterest)
	}
	gonglink_ := gonglinkFormCallback.gonglink
	_ = gonglink_

	for _, formDiv := range gonglinkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gonglink_.Name), formDiv)
		case "Recv":
			FormDivBasicFieldToField(&(gonglink_.Recv), formDiv)
		case "ImportPath":
			FormDivBasicFieldToField(&(gonglink_.ImportPath), formDiv)
		case "GongNote:Links":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongNote instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongNote instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongNote](gonglinkFormCallback.probe.stageOfInterest)
			targetGongNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongNote instances and update their Links slice
			for _gongnote := range *models.GetGongstructInstancesSetFromPointerType[*models.GongNote](gonglinkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gonglinkFormCallback.probe.stageOfInterest, _gongnote)
				
				// if GongNote is selected
				if targetGongNoteIDs[id] {
					// ensure gonglink_ is in _gongnote.Links
					found := false
					for _, _b := range _gongnote.Links {
						if _b == gonglink_ {
							found = true
							break
						}
					}
					if !found {
						_gongnote.Links = append(_gongnote.Links, gonglink_)
						gonglinkFormCallback.probe.UpdateSliceOfPointersCallback(_gongnote, "Links", &_gongnote.Links)
					}
				} else {
					// ensure gonglink_ is NOT in _gongnote.Links
					idx := slices.Index(_gongnote.Links, gonglink_)
					if idx != -1 {
						_gongnote.Links = slices.Delete(_gongnote.Links, idx, idx+1)
						gonglinkFormCallback.probe.UpdateSliceOfPointersCallback(_gongnote, "Links", &_gongnote.Links)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gonglinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gonglink_.Unstage(gonglinkFormCallback.probe.stageOfInterest)
	}

	gonglinkFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongLink](
		gonglinkFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gonglinkFormCallback.CreationMode || gonglinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gonglinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gonglinkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongLinkFormCallback(
			nil,
			gonglinkFormCallback.probe,
			newFormGroup,
		)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, newFormGroup, gonglinkFormCallback.probe)
		gonglinkFormCallback.probe.formStage.Commit()
	}

	gonglinkFormCallback.probe.ux_tree()
}
func __gong__New__GongNoteFormCallback(
	gongnote *models.GongNote,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnoteFormCallback *GongNoteFormCallback) {
	gongnoteFormCallback = new(GongNoteFormCallback)
	gongnoteFormCallback.probe = probe
	gongnoteFormCallback.gongnote = gongnote
	gongnoteFormCallback.formGroup = formGroup

	gongnoteFormCallback.CreationMode = (gongnote == nil)

	return
}

type GongNoteFormCallback struct {
	gongnote *models.GongNote

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongnoteFormCallback *GongNoteFormCallback) OnSave() {
	gongnoteFormCallback.probe.stageOfInterest.Lock()
	defer gongnoteFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongNoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteFormCallback.probe.formStage.Checkout()

	if gongnoteFormCallback.gongnote == nil {
		gongnoteFormCallback.gongnote = new(models.GongNote).Stage(gongnoteFormCallback.probe.stageOfInterest)
	}
	gongnote_ := gongnoteFormCallback.gongnote
	_ = gongnote_

	for _, formDiv := range gongnoteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnote_.Name), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnote_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnote_.BodyHTML), formDiv)
		case "Links":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongLink](gongnoteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongLink, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongLink)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongnoteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongLink](gongnoteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongnote_.Links = instanceSlice
			gongnoteFormCallback.probe.UpdateSliceOfPointersCallback(gongnote_, "Links", &gongnote_.Links)

		}
	}

	// manage the suppress operation
	if gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnote_.Unstage(gongnoteFormCallback.probe.stageOfInterest)
	}

	gongnoteFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongNote](
		gongnoteFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongnoteFormCallback.CreationMode || gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongnoteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteFormCallback(
			nil,
			gongnoteFormCallback.probe,
			newFormGroup,
		)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, newFormGroup, gongnoteFormCallback.probe)
		gongnoteFormCallback.probe.formStage.Commit()
	}

	gongnoteFormCallback.probe.ux_tree()
}
func __gong__New__GongStructFormCallback(
	gongstruct *models.GongStruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongstructFormCallback *GongStructFormCallback) {
	gongstructFormCallback = new(GongStructFormCallback)
	gongstructFormCallback.probe = probe
	gongstructFormCallback.gongstruct = gongstruct
	gongstructFormCallback.formGroup = formGroup

	gongstructFormCallback.CreationMode = (gongstruct == nil)

	return
}

type GongStructFormCallback struct {
	gongstruct *models.GongStruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongstructFormCallback *GongStructFormCallback) OnSave() {
	gongstructFormCallback.probe.stageOfInterest.Lock()
	defer gongstructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongStructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructFormCallback.probe.formStage.Checkout()

	if gongstructFormCallback.gongstruct == nil {
		gongstructFormCallback.gongstruct = new(models.GongStruct).Stage(gongstructFormCallback.probe.stageOfInterest)
	}
	gongstruct_ := gongstructFormCallback.gongstruct
	_ = gongstruct_

	for _, formDiv := range gongstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstruct_.Name), formDiv)
		case "GongBasicFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongBasicField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongBasicField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongBasicField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongBasicField](gongstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstruct_.GongBasicFields = instanceSlice
			gongstructFormCallback.probe.UpdateSliceOfPointersCallback(gongstruct_, "GongBasicFields", &gongstruct_.GongBasicFields)

		case "GongTimeFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongTimeField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongTimeField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongTimeField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongTimeField](gongstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstruct_.GongTimeFields = instanceSlice
			gongstructFormCallback.probe.UpdateSliceOfPointersCallback(gongstruct_, "GongTimeFields", &gongstruct_.GongTimeFields)

		case "PointerToGongStructFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PointerToGongStructField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PointerToGongStructField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstruct_.PointerToGongStructFields = instanceSlice
			gongstructFormCallback.probe.UpdateSliceOfPointersCallback(gongstruct_, "PointerToGongStructFields", &gongstruct_.PointerToGongStructFields)

		case "SliceOfPointerToGongStructFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SliceOfPointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SliceOfPointerToGongStructField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SliceOfPointerToGongStructField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SliceOfPointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstruct_.SliceOfPointerToGongStructFields = instanceSlice
			gongstructFormCallback.probe.UpdateSliceOfPointersCallback(gongstruct_, "SliceOfPointerToGongStructFields", &gongstruct_.SliceOfPointerToGongStructFields)

		case "HasOnAfterUpdateSignature":
			FormDivBasicFieldToField(&(gongstruct_.HasOnAfterUpdateSignature), formDiv)
		case "IsIgnoredForFront":
			FormDivBasicFieldToField(&(gongstruct_.IsIgnoredForFront), formDiv)
		}
	}

	// manage the suppress operation
	if gongstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstruct_.Unstage(gongstructFormCallback.probe.stageOfInterest)
	}

	gongstructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongStruct](
		gongstructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongstructFormCallback.CreationMode || gongstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongStructFormCallback(
			nil,
			gongstructFormCallback.probe,
			newFormGroup,
		)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, newFormGroup, gongstructFormCallback.probe)
		gongstructFormCallback.probe.formStage.Commit()
	}

	gongstructFormCallback.probe.ux_tree()
}
func __gong__New__GongTimeFieldFormCallback(
	gongtimefield *models.GongTimeField,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongtimefieldFormCallback *GongTimeFieldFormCallback) {
	gongtimefieldFormCallback = new(GongTimeFieldFormCallback)
	gongtimefieldFormCallback.probe = probe
	gongtimefieldFormCallback.gongtimefield = gongtimefield
	gongtimefieldFormCallback.formGroup = formGroup

	gongtimefieldFormCallback.CreationMode = (gongtimefield == nil)

	return
}

type GongTimeFieldFormCallback struct {
	gongtimefield *models.GongTimeField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongtimefieldFormCallback *GongTimeFieldFormCallback) OnSave() {
	gongtimefieldFormCallback.probe.stageOfInterest.Lock()
	defer gongtimefieldFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongTimeFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongtimefieldFormCallback.probe.formStage.Checkout()

	if gongtimefieldFormCallback.gongtimefield == nil {
		gongtimefieldFormCallback.gongtimefield = new(models.GongTimeField).Stage(gongtimefieldFormCallback.probe.stageOfInterest)
	}
	gongtimefield_ := gongtimefieldFormCallback.gongtimefield
	_ = gongtimefield_

	for _, formDiv := range gongtimefieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongtimefield_.Name), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongtimefield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongtimefield_.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(gongtimefield_.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(gongtimefield_.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(gongtimefield_.IsAccordionEnd), formDiv)
		case "BespokeTimeFormat":
			FormDivBasicFieldToField(&(gongtimefield_.BespokeTimeFormat), formDiv)
		case "GongStruct:GongTimeFields":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStruct](gongtimefieldFormCallback.probe.stageOfInterest)
			targetGongStructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStruct instances and update their GongTimeFields slice
			for _gongstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStruct](gongtimefieldFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongtimefieldFormCallback.probe.stageOfInterest, _gongstruct)
				
				// if GongStruct is selected
				if targetGongStructIDs[id] {
					// ensure gongtimefield_ is in _gongstruct.GongTimeFields
					found := false
					for _, _b := range _gongstruct.GongTimeFields {
						if _b == gongtimefield_ {
							found = true
							break
						}
					}
					if !found {
						_gongstruct.GongTimeFields = append(_gongstruct.GongTimeFields, gongtimefield_)
						gongtimefieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "GongTimeFields", &_gongstruct.GongTimeFields)
					}
				} else {
					// ensure gongtimefield_ is NOT in _gongstruct.GongTimeFields
					idx := slices.Index(_gongstruct.GongTimeFields, gongtimefield_)
					if idx != -1 {
						_gongstruct.GongTimeFields = slices.Delete(_gongstruct.GongTimeFields, idx, idx+1)
						gongtimefieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "GongTimeFields", &_gongstruct.GongTimeFields)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongtimefieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongtimefield_.Unstage(gongtimefieldFormCallback.probe.stageOfInterest)
	}

	gongtimefieldFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongTimeField](
		gongtimefieldFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongtimefieldFormCallback.CreationMode || gongtimefieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongtimefieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongtimefieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongTimeFieldFormCallback(
			nil,
			gongtimefieldFormCallback.probe,
			newFormGroup,
		)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, newFormGroup, gongtimefieldFormCallback.probe)
		gongtimefieldFormCallback.probe.formStage.Commit()
	}

	gongtimefieldFormCallback.probe.ux_tree()
}
func __gong__New__MetaReferenceFormCallback(
	metareference *models.MetaReference,
	probe *Probe,
	formGroup *form.FormGroup,
) (metareferenceFormCallback *MetaReferenceFormCallback) {
	metareferenceFormCallback = new(MetaReferenceFormCallback)
	metareferenceFormCallback.probe = probe
	metareferenceFormCallback.metareference = metareference
	metareferenceFormCallback.formGroup = formGroup

	metareferenceFormCallback.CreationMode = (metareference == nil)

	return
}

type MetaReferenceFormCallback struct {
	metareference *models.MetaReference

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (metareferenceFormCallback *MetaReferenceFormCallback) OnSave() {
	metareferenceFormCallback.probe.stageOfInterest.Lock()
	defer metareferenceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MetaReferenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metareferenceFormCallback.probe.formStage.Checkout()

	if metareferenceFormCallback.metareference == nil {
		metareferenceFormCallback.metareference = new(models.MetaReference).Stage(metareferenceFormCallback.probe.stageOfInterest)
	}
	metareference_ := metareferenceFormCallback.metareference
	_ = metareference_

	for _, formDiv := range metareferenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metareference_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareference_.Unstage(metareferenceFormCallback.probe.stageOfInterest)
	}

	metareferenceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MetaReference](
		metareferenceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if metareferenceFormCallback.CreationMode || metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareferenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(metareferenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MetaReferenceFormCallback(
			nil,
			metareferenceFormCallback.probe,
			newFormGroup,
		)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, newFormGroup, metareferenceFormCallback.probe)
		metareferenceFormCallback.probe.formStage.Commit()
	}

	metareferenceFormCallback.probe.ux_tree()
}
func __gong__New__ModelPkgFormCallback(
	modelpkg *models.ModelPkg,
	probe *Probe,
	formGroup *form.FormGroup,
) (modelpkgFormCallback *ModelPkgFormCallback) {
	modelpkgFormCallback = new(ModelPkgFormCallback)
	modelpkgFormCallback.probe = probe
	modelpkgFormCallback.modelpkg = modelpkg
	modelpkgFormCallback.formGroup = formGroup

	modelpkgFormCallback.CreationMode = (modelpkg == nil)

	return
}

type ModelPkgFormCallback struct {
	modelpkg *models.ModelPkg

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (modelpkgFormCallback *ModelPkgFormCallback) OnSave() {
	modelpkgFormCallback.probe.stageOfInterest.Lock()
	defer modelpkgFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ModelPkgFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	modelpkgFormCallback.probe.formStage.Checkout()

	if modelpkgFormCallback.modelpkg == nil {
		modelpkgFormCallback.modelpkg = new(models.ModelPkg).Stage(modelpkgFormCallback.probe.stageOfInterest)
	}
	modelpkg_ := modelpkgFormCallback.modelpkg
	_ = modelpkg_

	for _, formDiv := range modelpkgFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(modelpkg_.Name), formDiv)
		case "PkgPath":
			FormDivBasicFieldToField(&(modelpkg_.PkgPath), formDiv)
		case "PathToGoSubDirectory":
			FormDivBasicFieldToField(&(modelpkg_.PathToGoSubDirectory), formDiv)
		case "OrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.OrmPkgGenPath), formDiv)
		case "DbOrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbOrmPkgGenPath), formDiv)
		case "DbLiteOrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbLiteOrmPkgGenPath), formDiv)
		case "DbPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbPkgGenPath), formDiv)
		case "ControllersPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.ControllersPkgGenPath), formDiv)
		case "FullstackPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.FullstackPkgGenPath), formDiv)
		case "StackPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.StackPkgGenPath), formDiv)
		case "Level1StackPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.Level1StackPkgGenPath), formDiv)
		case "StaticPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.StaticPkgGenPath), formDiv)
		case "ProbePkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.ProbePkgGenPath), formDiv)
		case "NgWorkspacePath":
			FormDivBasicFieldToField(&(modelpkg_.NgWorkspacePath), formDiv)
		case "NgWorkspaceName":
			FormDivBasicFieldToField(&(modelpkg_.NgWorkspaceName), formDiv)
		case "NgDataLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(modelpkg_.NgDataLibrarySourceCodeDirectory), formDiv)
		case "NgSpecificLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(modelpkg_.NgSpecificLibrarySourceCodeDirectory), formDiv)
		case "MaterialLibDatamodelTargetPath":
			FormDivBasicFieldToField(&(modelpkg_.MaterialLibDatamodelTargetPath), formDiv)
		}
	}

	// manage the suppress operation
	if modelpkgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		modelpkg_.Unstage(modelpkgFormCallback.probe.stageOfInterest)
	}

	modelpkgFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ModelPkg](
		modelpkgFormCallback.probe,
	)

	// display a new form by reset the form stage
	if modelpkgFormCallback.CreationMode || modelpkgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		modelpkgFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(modelpkgFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ModelPkgFormCallback(
			nil,
			modelpkgFormCallback.probe,
			newFormGroup,
		)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, newFormGroup, modelpkgFormCallback.probe)
		modelpkgFormCallback.probe.formStage.Commit()
	}

	modelpkgFormCallback.probe.ux_tree()
}
func __gong__New__PointerToGongStructFieldFormCallback(
	pointertogongstructfield *models.PointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) {
	pointertogongstructfieldFormCallback = new(PointerToGongStructFieldFormCallback)
	pointertogongstructfieldFormCallback.probe = probe
	pointertogongstructfieldFormCallback.pointertogongstructfield = pointertogongstructfield
	pointertogongstructfieldFormCallback.formGroup = formGroup

	pointertogongstructfieldFormCallback.CreationMode = (pointertogongstructfield == nil)

	return
}

type PointerToGongStructFieldFormCallback struct {
	pointertogongstructfield *models.PointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) OnSave() {
	pointertogongstructfieldFormCallback.probe.stageOfInterest.Lock()
	defer pointertogongstructfieldFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointertogongstructfieldFormCallback.probe.formStage.Checkout()

	if pointertogongstructfieldFormCallback.pointertogongstructfield == nil {
		pointertogongstructfieldFormCallback.pointertogongstructfield = new(models.PointerToGongStructField).Stage(pointertogongstructfieldFormCallback.probe.stageOfInterest)
	}
	pointertogongstructfield_ := pointertogongstructfieldFormCallback.pointertogongstructfield
	_ = pointertogongstructfield_

	for _, formDiv := range pointertogongstructfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(pointertogongstructfield_.GongStruct), pointertogongstructfieldFormCallback.probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(pointertogongstructfield_.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(pointertogongstructfield_.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(pointertogongstructfield_.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(pointertogongstructfield_.IsAccordionEnd), formDiv)
		case "IsType":
			FormDivBasicFieldToField(&(pointertogongstructfield_.IsType), formDiv)
		case "GongStruct:PointerToGongStructFields":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStruct](pointertogongstructfieldFormCallback.probe.stageOfInterest)
			targetGongStructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStruct instances and update their PointerToGongStructFields slice
			for _gongstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStruct](pointertogongstructfieldFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(pointertogongstructfieldFormCallback.probe.stageOfInterest, _gongstruct)
				
				// if GongStruct is selected
				if targetGongStructIDs[id] {
					// ensure pointertogongstructfield_ is in _gongstruct.PointerToGongStructFields
					found := false
					for _, _b := range _gongstruct.PointerToGongStructFields {
						if _b == pointertogongstructfield_ {
							found = true
							break
						}
					}
					if !found {
						_gongstruct.PointerToGongStructFields = append(_gongstruct.PointerToGongStructFields, pointertogongstructfield_)
						pointertogongstructfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "PointerToGongStructFields", &_gongstruct.PointerToGongStructFields)
					}
				} else {
					// ensure pointertogongstructfield_ is NOT in _gongstruct.PointerToGongStructFields
					idx := slices.Index(_gongstruct.PointerToGongStructFields, pointertogongstructfield_)
					if idx != -1 {
						_gongstruct.PointerToGongStructFields = slices.Delete(_gongstruct.PointerToGongStructFields, idx, idx+1)
						pointertogongstructfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "PointerToGongStructFields", &_gongstruct.PointerToGongStructFields)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if pointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointertogongstructfield_.Unstage(pointertogongstructfieldFormCallback.probe.stageOfInterest)
	}

	pointertogongstructfieldFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PointerToGongStructField](
		pointertogongstructfieldFormCallback.probe,
	)

	// display a new form by reset the form stage
	if pointertogongstructfieldFormCallback.CreationMode || pointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(pointertogongstructfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PointerToGongStructFieldFormCallback(
			nil,
			pointertogongstructfieldFormCallback.probe,
			newFormGroup,
		)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, newFormGroup, pointertogongstructfieldFormCallback.probe)
		pointertogongstructfieldFormCallback.probe.formStage.Commit()
	}

	pointertogongstructfieldFormCallback.probe.ux_tree()
}
func __gong__New__SliceOfPointerToGongStructFieldFormCallback(
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) {
	sliceofpointertogongstructfieldFormCallback = new(SliceOfPointerToGongStructFieldFormCallback)
	sliceofpointertogongstructfieldFormCallback.probe = probe
	sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = sliceofpointertogongstructfield
	sliceofpointertogongstructfieldFormCallback.formGroup = formGroup

	sliceofpointertogongstructfieldFormCallback.CreationMode = (sliceofpointertogongstructfield == nil)

	return
}

type SliceOfPointerToGongStructFieldFormCallback struct {
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) OnSave() {
	sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest.Lock()
	defer sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SliceOfPointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliceofpointertogongstructfieldFormCallback.probe.formStage.Checkout()

	if sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield == nil {
		sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = new(models.SliceOfPointerToGongStructField).Stage(sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest)
	}
	sliceofpointertogongstructfield_ := sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield
	_ = sliceofpointertogongstructfield_

	for _, formDiv := range sliceofpointertogongstructfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(sliceofpointertogongstructfield_.GongStruct), sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.IsAccordionEnd), formDiv)
		case "GongStruct:SliceOfPointerToGongStructFields":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStruct](sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest)
			targetGongStructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStruct instances and update their SliceOfPointerToGongStructFields slice
			for _gongstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStruct](sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest, _gongstruct)
				
				// if GongStruct is selected
				if targetGongStructIDs[id] {
					// ensure sliceofpointertogongstructfield_ is in _gongstruct.SliceOfPointerToGongStructFields
					found := false
					for _, _b := range _gongstruct.SliceOfPointerToGongStructFields {
						if _b == sliceofpointertogongstructfield_ {
							found = true
							break
						}
					}
					if !found {
						_gongstruct.SliceOfPointerToGongStructFields = append(_gongstruct.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
						sliceofpointertogongstructfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "SliceOfPointerToGongStructFields", &_gongstruct.SliceOfPointerToGongStructFields)
					}
				} else {
					// ensure sliceofpointertogongstructfield_ is NOT in _gongstruct.SliceOfPointerToGongStructFields
					idx := slices.Index(_gongstruct.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
					if idx != -1 {
						_gongstruct.SliceOfPointerToGongStructFields = slices.Delete(_gongstruct.SliceOfPointerToGongStructFields, idx, idx+1)
						sliceofpointertogongstructfieldFormCallback.probe.UpdateSliceOfPointersCallback(_gongstruct, "SliceOfPointerToGongStructFields", &_gongstruct.SliceOfPointerToGongStructFields)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if sliceofpointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliceofpointertogongstructfield_.Unstage(sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest)
	}

	sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SliceOfPointerToGongStructField](
		sliceofpointertogongstructfieldFormCallback.probe,
	)

	// display a new form by reset the form stage
	if sliceofpointertogongstructfieldFormCallback.CreationMode || sliceofpointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliceofpointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(sliceofpointertogongstructfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SliceOfPointerToGongStructFieldFormCallback(
			nil,
			sliceofpointertogongstructfieldFormCallback.probe,
			newFormGroup,
		)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, newFormGroup, sliceofpointertogongstructfieldFormCallback.probe)
		sliceofpointertogongstructfieldFormCallback.probe.formStage.Commit()
	}

	sliceofpointertogongstructfieldFormCallback.probe.ux_tree()
}
