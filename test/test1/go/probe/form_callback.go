// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test1/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AstructFormCallback(
	astruct *models.Astruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructFormCallback *AstructFormCallback) {
	astructFormCallback = new(AstructFormCallback)
	astructFormCallback.probe = probe
	astructFormCallback.astruct = astruct
	astructFormCallback.formGroup = formGroup

	astructFormCallback.CreationMode = (astruct == nil)

	return
}

type AstructFormCallback struct {
	astruct *models.Astruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (astructFormCallback *AstructFormCallback) OnSave() {
	astructFormCallback.probe.stageOfInterest.Lock()
	defer astructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructFormCallback.probe.formStage.Checkout()

	if astructFormCallback.astruct == nil {
		astructFormCallback.astruct = new(models.Astruct).Stage(astructFormCallback.probe.stageOfInterest)
	}
	astruct_ := astructFormCallback.astruct
	_ = astruct_

	for _, formDiv := range astructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astruct_.Name), formDiv)
		case "Associationtob":
			FormDivSelectFieldToField(&(astruct_.Associationtob), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Anarrayofb":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Bstruct](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Bstruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Bstruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Bstruct](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.Anarrayofb = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "Anarrayofb", &astruct_.Anarrayofb)

		case "Anotherassociationtob_2":
			FormDivSelectFieldToField(&(astruct_.Anotherassociationtob_2), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Date":
			FormDivTimeFieldToField(&(astruct_.Date), formDiv, false)
		case "Date2":
			FormDivTimeFieldToField(&(astruct_.Date2), formDiv, false)
		case "Booleanfield":
			FormDivBasicFieldToField(&(astruct_.Booleanfield), formDiv)
		case "Aenum":
			FormDivEnumStringFieldToField(&(astruct_.Aenum), formDiv)
		case "Aenum_2":
			FormDivEnumStringFieldToField(&(astruct_.Aenum_2), formDiv)
		case "Benum":
			FormDivEnumStringFieldToField(&(astruct_.Benum), formDiv)
		case "CEnum":
			FormDivEnumIntFieldToField(&(astruct_.CEnum), formDiv)
		case "CName":
			FormDivBasicFieldToField(&(astruct_.CName), formDiv)
		case "CFloatfield":
			FormDivBasicFieldToField(&(astruct_.CFloatfield), formDiv)
		case "Bstruct":
			FormDivSelectFieldToField(&(astruct_.Bstruct), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(astruct_.Bstruct2), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Dstruct":
			FormDivSelectFieldToField(&(astruct_.Dstruct), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Dstruct2":
			FormDivSelectFieldToField(&(astruct_.Dstruct2), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Dstruct3":
			FormDivSelectFieldToField(&(astruct_.Dstruct3), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Dstruct4":
			FormDivSelectFieldToField(&(astruct_.Dstruct4), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Dstruct4s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Dstruct](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Dstruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Dstruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Dstruct](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.Dstruct4s = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "Dstruct4s", &astruct_.Dstruct4s)

		case "Floatfield":
			FormDivBasicFieldToField(&(astruct_.Floatfield), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(astruct_.Intfield), formDiv)
		case "Anotherbooleanfield":
			FormDivBasicFieldToField(&(astruct_.Anotherbooleanfield), formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&(astruct_.Duration1), formDiv)
		case "Anarrayofa":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Astruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Astruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.Anarrayofa = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "Anarrayofa", &astruct_.Anarrayofa)

		case "Anotherarrayofb":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Bstruct](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Bstruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Bstruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Bstruct](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.Anotherarrayofb = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "Anotherarrayofb", &astruct_.Anotherarrayofb)

		case "AnarrayofbUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AstructBstructUse](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AstructBstructUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AstructBstructUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AstructBstructUse](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.AnarrayofbUse = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "AnarrayofbUse", &astruct_.AnarrayofbUse)

		case "Anarrayofb2Use":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AstructBstruct2Use](astructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AstructBstruct2Use, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AstructBstruct2Use)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					astructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AstructBstruct2Use](astructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			astruct_.Anarrayofb2Use = instanceSlice
			astructFormCallback.probe.UpdateSliceOfPointersCallback(astruct_, "Anarrayofb2Use", &astruct_.Anarrayofb2Use)

		case "AnAstruct":
			FormDivSelectFieldToField(&(astruct_.AnAstruct), astructFormCallback.probe.stageOfInterest, formDiv)
		case "TextFieldBespokeSize":
			FormDivBasicFieldToField(&(astruct_.TextFieldBespokeSize), formDiv)
		case "TextArea":
			FormDivBasicFieldToField(&(astruct_.TextArea), formDiv)
		case "Astruct:Anarrayofa":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](astructFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their Anarrayofa slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](astructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(astructFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure astruct_ is in _astruct.Anarrayofa
					found := false
					for _, _b := range _astruct.Anarrayofa {
						if _b == astruct_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.Anarrayofa = append(_astruct.Anarrayofa, astruct_)
						astructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofa", &_astruct.Anarrayofa)
					}
				} else {
					// ensure astruct_ is NOT in _astruct.Anarrayofa
					idx := slices.Index(_astruct.Anarrayofa, astruct_)
					if idx != -1 {
						_astruct.Anarrayofa = slices.Delete(_astruct.Anarrayofa, idx, idx+1)
						astructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofa", &_astruct.Anarrayofa)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astruct_.Unstage(astructFormCallback.probe.stageOfInterest)
	}

	astructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Astruct](
		astructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if astructFormCallback.CreationMode || astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(astructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AstructFormCallback(
			nil,
			astructFormCallback.probe,
			newFormGroup,
		)
		astruct := new(models.Astruct)
		FillUpForm(astruct, newFormGroup, astructFormCallback.probe)
		astructFormCallback.probe.formStage.Commit()
	}

	astructFormCallback.probe.ux_tree()
}
func __gong__New__AstructBstruct2UseFormCallback(
	astructbstruct2use *models.AstructBstruct2Use,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) {
	astructbstruct2useFormCallback = new(AstructBstruct2UseFormCallback)
	astructbstruct2useFormCallback.probe = probe
	astructbstruct2useFormCallback.astructbstruct2use = astructbstruct2use
	astructbstruct2useFormCallback.formGroup = formGroup

	astructbstruct2useFormCallback.CreationMode = (astructbstruct2use == nil)

	return
}

type AstructBstruct2UseFormCallback struct {
	astructbstruct2use *models.AstructBstruct2Use

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) OnSave() {
	astructbstruct2useFormCallback.probe.stageOfInterest.Lock()
	defer astructbstruct2useFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AstructBstruct2UseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructbstruct2useFormCallback.probe.formStage.Checkout()

	if astructbstruct2useFormCallback.astructbstruct2use == nil {
		astructbstruct2useFormCallback.astructbstruct2use = new(models.AstructBstruct2Use).Stage(astructbstruct2useFormCallback.probe.stageOfInterest)
	}
	astructbstruct2use_ := astructbstruct2useFormCallback.astructbstruct2use
	_ = astructbstruct2use_

	for _, formDiv := range astructbstruct2useFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astructbstruct2use_.Name), formDiv)
		case "Bstrcut2":
			FormDivSelectFieldToField(&(astructbstruct2use_.Bstrcut2), astructbstruct2useFormCallback.probe.stageOfInterest, formDiv)
		case "Astruct:Anarrayofb2Use":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](astructbstruct2useFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their Anarrayofb2Use slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](astructbstruct2useFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(astructbstruct2useFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure astructbstruct2use_ is in _astruct.Anarrayofb2Use
					found := false
					for _, _b := range _astruct.Anarrayofb2Use {
						if _b == astructbstruct2use_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.Anarrayofb2Use = append(_astruct.Anarrayofb2Use, astructbstruct2use_)
						astructbstruct2useFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofb2Use", &_astruct.Anarrayofb2Use)
					}
				} else {
					// ensure astructbstruct2use_ is NOT in _astruct.Anarrayofb2Use
					idx := slices.Index(_astruct.Anarrayofb2Use, astructbstruct2use_)
					if idx != -1 {
						_astruct.Anarrayofb2Use = slices.Delete(_astruct.Anarrayofb2Use, idx, idx+1)
						astructbstruct2useFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofb2Use", &_astruct.Anarrayofb2Use)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if astructbstruct2useFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstruct2use_.Unstage(astructbstruct2useFormCallback.probe.stageOfInterest)
	}

	astructbstruct2useFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AstructBstruct2Use](
		astructbstruct2useFormCallback.probe,
	)

	// display a new form by reset the form stage
	if astructbstruct2useFormCallback.CreationMode || astructbstruct2useFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstruct2useFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(astructbstruct2useFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AstructBstruct2UseFormCallback(
			nil,
			astructbstruct2useFormCallback.probe,
			newFormGroup,
		)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, newFormGroup, astructbstruct2useFormCallback.probe)
		astructbstruct2useFormCallback.probe.formStage.Commit()
	}

	astructbstruct2useFormCallback.probe.ux_tree()
}
func __gong__New__AstructBstructUseFormCallback(
	astructbstructuse *models.AstructBstructUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructbstructuseFormCallback *AstructBstructUseFormCallback) {
	astructbstructuseFormCallback = new(AstructBstructUseFormCallback)
	astructbstructuseFormCallback.probe = probe
	astructbstructuseFormCallback.astructbstructuse = astructbstructuse
	astructbstructuseFormCallback.formGroup = formGroup

	astructbstructuseFormCallback.CreationMode = (astructbstructuse == nil)

	return
}

type AstructBstructUseFormCallback struct {
	astructbstructuse *models.AstructBstructUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (astructbstructuseFormCallback *AstructBstructUseFormCallback) OnSave() {
	astructbstructuseFormCallback.probe.stageOfInterest.Lock()
	defer astructbstructuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AstructBstructUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructbstructuseFormCallback.probe.formStage.Checkout()

	if astructbstructuseFormCallback.astructbstructuse == nil {
		astructbstructuseFormCallback.astructbstructuse = new(models.AstructBstructUse).Stage(astructbstructuseFormCallback.probe.stageOfInterest)
	}
	astructbstructuse_ := astructbstructuseFormCallback.astructbstructuse
	_ = astructbstructuse_

	for _, formDiv := range astructbstructuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astructbstructuse_.Name), formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(astructbstructuse_.Bstruct2), astructbstructuseFormCallback.probe.stageOfInterest, formDiv)
		case "Astruct:AnarrayofbUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](astructbstructuseFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their AnarrayofbUse slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](astructbstructuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(astructbstructuseFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure astructbstructuse_ is in _astruct.AnarrayofbUse
					found := false
					for _, _b := range _astruct.AnarrayofbUse {
						if _b == astructbstructuse_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.AnarrayofbUse = append(_astruct.AnarrayofbUse, astructbstructuse_)
						astructbstructuseFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "AnarrayofbUse", &_astruct.AnarrayofbUse)
					}
				} else {
					// ensure astructbstructuse_ is NOT in _astruct.AnarrayofbUse
					idx := slices.Index(_astruct.AnarrayofbUse, astructbstructuse_)
					if idx != -1 {
						_astruct.AnarrayofbUse = slices.Delete(_astruct.AnarrayofbUse, idx, idx+1)
						astructbstructuseFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "AnarrayofbUse", &_astruct.AnarrayofbUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if astructbstructuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstructuse_.Unstage(astructbstructuseFormCallback.probe.stageOfInterest)
	}

	astructbstructuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AstructBstructUse](
		astructbstructuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if astructbstructuseFormCallback.CreationMode || astructbstructuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstructuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(astructbstructuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AstructBstructUseFormCallback(
			nil,
			astructbstructuseFormCallback.probe,
			newFormGroup,
		)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, newFormGroup, astructbstructuseFormCallback.probe)
		astructbstructuseFormCallback.probe.formStage.Commit()
	}

	astructbstructuseFormCallback.probe.ux_tree()
}
func __gong__New__BstructFormCallback(
	bstruct *models.Bstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (bstructFormCallback *BstructFormCallback) {
	bstructFormCallback = new(BstructFormCallback)
	bstructFormCallback.probe = probe
	bstructFormCallback.bstruct = bstruct
	bstructFormCallback.formGroup = formGroup

	bstructFormCallback.CreationMode = (bstruct == nil)

	return
}

type BstructFormCallback struct {
	bstruct *models.Bstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bstructFormCallback *BstructFormCallback) OnSave() {
	bstructFormCallback.probe.stageOfInterest.Lock()
	defer bstructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bstructFormCallback.probe.formStage.Checkout()

	if bstructFormCallback.bstruct == nil {
		bstructFormCallback.bstruct = new(models.Bstruct).Stage(bstructFormCallback.probe.stageOfInterest)
	}
	bstruct_ := bstructFormCallback.bstruct
	_ = bstruct_

	for _, formDiv := range bstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bstruct_.Name), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(bstruct_.Floatfield), formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&(bstruct_.Floatfield2), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(bstruct_.Intfield), formDiv)
		case "Astruct:Anarrayofb":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](bstructFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their Anarrayofb slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](bstructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bstructFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure bstruct_ is in _astruct.Anarrayofb
					found := false
					for _, _b := range _astruct.Anarrayofb {
						if _b == bstruct_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.Anarrayofb = append(_astruct.Anarrayofb, bstruct_)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofb", &_astruct.Anarrayofb)
					}
				} else {
					// ensure bstruct_ is NOT in _astruct.Anarrayofb
					idx := slices.Index(_astruct.Anarrayofb, bstruct_)
					if idx != -1 {
						_astruct.Anarrayofb = slices.Delete(_astruct.Anarrayofb, idx, idx+1)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anarrayofb", &_astruct.Anarrayofb)
					}
				}
			}
		case "Astruct:Anotherarrayofb":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](bstructFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their Anotherarrayofb slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](bstructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bstructFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure bstruct_ is in _astruct.Anotherarrayofb
					found := false
					for _, _b := range _astruct.Anotherarrayofb {
						if _b == bstruct_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.Anotherarrayofb = append(_astruct.Anotherarrayofb, bstruct_)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anotherarrayofb", &_astruct.Anotherarrayofb)
					}
				} else {
					// ensure bstruct_ is NOT in _astruct.Anotherarrayofb
					idx := slices.Index(_astruct.Anotherarrayofb, bstruct_)
					if idx != -1 {
						_astruct.Anotherarrayofb = slices.Delete(_astruct.Anotherarrayofb, idx, idx+1)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Anotherarrayofb", &_astruct.Anotherarrayofb)
					}
				}
			}
		case "Dstruct:Anarrayofb":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Dstruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Dstruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Dstruct](bstructFormCallback.probe.stageOfInterest)
			targetDstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Dstruct instances and update their Anarrayofb slice
			for _dstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Dstruct](bstructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bstructFormCallback.probe.stageOfInterest, _dstruct)
				
				// if Dstruct is selected
				if targetDstructIDs[id] {
					// ensure bstruct_ is in _dstruct.Anarrayofb
					found := false
					for _, _b := range _dstruct.Anarrayofb {
						if _b == bstruct_ {
							found = true
							break
						}
					}
					if !found {
						_dstruct.Anarrayofb = append(_dstruct.Anarrayofb, bstruct_)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_dstruct, "Anarrayofb", &_dstruct.Anarrayofb)
					}
				} else {
					// ensure bstruct_ is NOT in _dstruct.Anarrayofb
					idx := slices.Index(_dstruct.Anarrayofb, bstruct_)
					if idx != -1 {
						_dstruct.Anarrayofb = slices.Delete(_dstruct.Anarrayofb, idx, idx+1)
						bstructFormCallback.probe.UpdateSliceOfPointersCallback(_dstruct, "Anarrayofb", &_dstruct.Anarrayofb)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bstruct_.Unstage(bstructFormCallback.probe.stageOfInterest)
	}

	bstructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Bstruct](
		bstructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bstructFormCallback.CreationMode || bstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BstructFormCallback(
			nil,
			bstructFormCallback.probe,
			newFormGroup,
		)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, newFormGroup, bstructFormCallback.probe)
		bstructFormCallback.probe.formStage.Commit()
	}

	bstructFormCallback.probe.ux_tree()
}
func __gong__New__DstructFormCallback(
	dstruct *models.Dstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (dstructFormCallback *DstructFormCallback) {
	dstructFormCallback = new(DstructFormCallback)
	dstructFormCallback.probe = probe
	dstructFormCallback.dstruct = dstruct
	dstructFormCallback.formGroup = formGroup

	dstructFormCallback.CreationMode = (dstruct == nil)

	return
}

type DstructFormCallback struct {
	dstruct *models.Dstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (dstructFormCallback *DstructFormCallback) OnSave() {
	dstructFormCallback.probe.stageOfInterest.Lock()
	defer dstructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dstructFormCallback.probe.formStage.Checkout()

	if dstructFormCallback.dstruct == nil {
		dstructFormCallback.dstruct = new(models.Dstruct).Stage(dstructFormCallback.probe.stageOfInterest)
	}
	dstruct_ := dstructFormCallback.dstruct
	_ = dstruct_

	for _, formDiv := range dstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dstruct_.Name), formDiv)
		case "Anarrayofb":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Bstruct](dstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Bstruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Bstruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					dstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Bstruct](dstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			dstruct_.Anarrayofb = instanceSlice
			dstructFormCallback.probe.UpdateSliceOfPointersCallback(dstruct_, "Anarrayofb", &dstruct_.Anarrayofb)

		case "Gstruct":
			FormDivSelectFieldToField(&(dstruct_.Gstruct), dstructFormCallback.probe.stageOfInterest, formDiv)
		case "Gstructs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Gstruct](dstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Gstruct, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Gstruct)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					dstructFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Gstruct](dstructFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			dstruct_.Gstructs = instanceSlice
			dstructFormCallback.probe.UpdateSliceOfPointersCallback(dstruct_, "Gstructs", &dstruct_.Gstructs)

		case "Astruct:Dstruct4s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Astruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Astruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Astruct](dstructFormCallback.probe.stageOfInterest)
			targetAstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Astruct instances and update their Dstruct4s slice
			for _astruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Astruct](dstructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dstructFormCallback.probe.stageOfInterest, _astruct)
				
				// if Astruct is selected
				if targetAstructIDs[id] {
					// ensure dstruct_ is in _astruct.Dstruct4s
					found := false
					for _, _b := range _astruct.Dstruct4s {
						if _b == dstruct_ {
							found = true
							break
						}
					}
					if !found {
						_astruct.Dstruct4s = append(_astruct.Dstruct4s, dstruct_)
						dstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Dstruct4s", &_astruct.Dstruct4s)
					}
				} else {
					// ensure dstruct_ is NOT in _astruct.Dstruct4s
					idx := slices.Index(_astruct.Dstruct4s, dstruct_)
					if idx != -1 {
						_astruct.Dstruct4s = slices.Delete(_astruct.Dstruct4s, idx, idx+1)
						dstructFormCallback.probe.UpdateSliceOfPointersCallback(_astruct, "Dstruct4s", &_astruct.Dstruct4s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if dstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dstruct_.Unstage(dstructFormCallback.probe.stageOfInterest)
	}

	dstructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Dstruct](
		dstructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if dstructFormCallback.CreationMode || dstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(dstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DstructFormCallback(
			nil,
			dstructFormCallback.probe,
			newFormGroup,
		)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, newFormGroup, dstructFormCallback.probe)
		dstructFormCallback.probe.formStage.Commit()
	}

	dstructFormCallback.probe.ux_tree()
}
func __gong__New__F0123456789012345678901234567890FormCallback(
	f0123456789012345678901234567890 *models.F0123456789012345678901234567890,
	probe *Probe,
	formGroup *form.FormGroup,
) (f0123456789012345678901234567890FormCallback *F0123456789012345678901234567890FormCallback) {
	f0123456789012345678901234567890FormCallback = new(F0123456789012345678901234567890FormCallback)
	f0123456789012345678901234567890FormCallback.probe = probe
	f0123456789012345678901234567890FormCallback.f0123456789012345678901234567890 = f0123456789012345678901234567890
	f0123456789012345678901234567890FormCallback.formGroup = formGroup

	f0123456789012345678901234567890FormCallback.CreationMode = (f0123456789012345678901234567890 == nil)

	return
}

type F0123456789012345678901234567890FormCallback struct {
	f0123456789012345678901234567890 *models.F0123456789012345678901234567890

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (f0123456789012345678901234567890FormCallback *F0123456789012345678901234567890FormCallback) OnSave() {
	f0123456789012345678901234567890FormCallback.probe.stageOfInterest.Lock()
	defer f0123456789012345678901234567890FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("F0123456789012345678901234567890FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	f0123456789012345678901234567890FormCallback.probe.formStage.Checkout()

	if f0123456789012345678901234567890FormCallback.f0123456789012345678901234567890 == nil {
		f0123456789012345678901234567890FormCallback.f0123456789012345678901234567890 = new(models.F0123456789012345678901234567890).Stage(f0123456789012345678901234567890FormCallback.probe.stageOfInterest)
	}
	f0123456789012345678901234567890_ := f0123456789012345678901234567890FormCallback.f0123456789012345678901234567890
	_ = f0123456789012345678901234567890_

	for _, formDiv := range f0123456789012345678901234567890FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(f0123456789012345678901234567890_.Name), formDiv)
		case "Date":
			FormDivTimeFieldToField(&(f0123456789012345678901234567890_.Date), formDiv, false)
		}
	}

	// manage the suppress operation
	if f0123456789012345678901234567890FormCallback.formGroup.HasSuppressButtonBeenPressed {
		f0123456789012345678901234567890_.Unstage(f0123456789012345678901234567890FormCallback.probe.stageOfInterest)
	}

	f0123456789012345678901234567890FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.F0123456789012345678901234567890](
		f0123456789012345678901234567890FormCallback.probe,
	)

	// display a new form by reset the form stage
	if f0123456789012345678901234567890FormCallback.CreationMode || f0123456789012345678901234567890FormCallback.formGroup.HasSuppressButtonBeenPressed {
		f0123456789012345678901234567890FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(f0123456789012345678901234567890FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__F0123456789012345678901234567890FormCallback(
			nil,
			f0123456789012345678901234567890FormCallback.probe,
			newFormGroup,
		)
		f0123456789012345678901234567890 := new(models.F0123456789012345678901234567890)
		FillUpForm(f0123456789012345678901234567890, newFormGroup, f0123456789012345678901234567890FormCallback.probe)
		f0123456789012345678901234567890FormCallback.probe.formStage.Commit()
	}

	f0123456789012345678901234567890FormCallback.probe.ux_tree()
}
func __gong__New__GstructFormCallback(
	gstruct *models.Gstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (gstructFormCallback *GstructFormCallback) {
	gstructFormCallback = new(GstructFormCallback)
	gstructFormCallback.probe = probe
	gstructFormCallback.gstruct = gstruct
	gstructFormCallback.formGroup = formGroup

	gstructFormCallback.CreationMode = (gstruct == nil)

	return
}

type GstructFormCallback struct {
	gstruct *models.Gstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gstructFormCallback *GstructFormCallback) OnSave() {
	gstructFormCallback.probe.stageOfInterest.Lock()
	defer gstructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gstructFormCallback.probe.formStage.Checkout()

	if gstructFormCallback.gstruct == nil {
		gstructFormCallback.gstruct = new(models.Gstruct).Stage(gstructFormCallback.probe.stageOfInterest)
	}
	gstruct_ := gstructFormCallback.gstruct
	_ = gstruct_

	for _, formDiv := range gstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gstruct_.Name), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(gstruct_.Floatfield), formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&(gstruct_.Floatfield2), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(gstruct_.Intfield), formDiv)
		case "Dstruct:Gstructs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Dstruct instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Dstruct instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Dstruct](gstructFormCallback.probe.stageOfInterest)
			targetDstructIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDstructIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Dstruct instances and update their Gstructs slice
			for _dstruct := range *models.GetGongstructInstancesSetFromPointerType[*models.Dstruct](gstructFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gstructFormCallback.probe.stageOfInterest, _dstruct)
				
				// if Dstruct is selected
				if targetDstructIDs[id] {
					// ensure gstruct_ is in _dstruct.Gstructs
					found := false
					for _, _b := range _dstruct.Gstructs {
						if _b == gstruct_ {
							found = true
							break
						}
					}
					if !found {
						_dstruct.Gstructs = append(_dstruct.Gstructs, gstruct_)
						gstructFormCallback.probe.UpdateSliceOfPointersCallback(_dstruct, "Gstructs", &_dstruct.Gstructs)
					}
				} else {
					// ensure gstruct_ is NOT in _dstruct.Gstructs
					idx := slices.Index(_dstruct.Gstructs, gstruct_)
					if idx != -1 {
						_dstruct.Gstructs = slices.Delete(_dstruct.Gstructs, idx, idx+1)
						gstructFormCallback.probe.UpdateSliceOfPointersCallback(_dstruct, "Gstructs", &_dstruct.Gstructs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gstruct_.Unstage(gstructFormCallback.probe.stageOfInterest)
	}

	gstructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Gstruct](
		gstructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gstructFormCallback.CreationMode || gstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GstructFormCallback(
			nil,
			gstructFormCallback.probe,
			newFormGroup,
		)
		gstruct := new(models.Gstruct)
		FillUpForm(gstruct, newFormGroup, gstructFormCallback.probe)
		gstructFormCallback.probe.formStage.Commit()
	}

	gstructFormCallback.probe.ux_tree()
}
