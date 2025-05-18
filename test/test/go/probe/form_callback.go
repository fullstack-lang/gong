// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AstructFormCallback(
	astruct *models.Astruct,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (astructFormCallback *AstructFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.Anarrayofb = instanceSlice

		case "Anotherassociationtob_2":
			FormDivSelectFieldToField(&(astruct_.Anotherassociationtob_2), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Date":
			FormDivBasicFieldToField(&(astruct_.Date), formDiv)
		case "Date2":
			FormDivBasicFieldToField(&(astruct_.Date2), formDiv)
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.Dstruct4s = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.Anarrayofa = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.Anotherarrayofb = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.AnarrayofbUse = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			astruct_.Anarrayofb2Use = instanceSlice

		case "AnAstruct":
			FormDivSelectFieldToField(&(astruct_.AnAstruct), astructFormCallback.probe.stageOfInterest, formDiv)
		case "StructRef":
			FormDivBasicFieldToField(&(astruct_.StructRef), formDiv)
		case "FieldRef":
			FormDivBasicFieldToField(&(astruct_.FieldRef), formDiv)
		case "EnumIntRef":
			FormDivBasicFieldToField(&(astruct_.EnumIntRef), formDiv)
		case "EnumStringRef":
			FormDivBasicFieldToField(&(astruct_.EnumStringRef), formDiv)
		case "EnumValue":
			FormDivBasicFieldToField(&(astruct_.EnumValue), formDiv)
		case "ConstIdentifierValue":
			FormDivBasicFieldToField(&(astruct_.ConstIdentifierValue), formDiv)
		case "TextFieldBespokeSize":
			FormDivBasicFieldToField(&(astruct_.TextFieldBespokeSize), formDiv)
		case "TextArea":
			FormDivBasicFieldToField(&(astruct_.TextArea), formDiv)
		case "Astruct:Anarrayofa":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofa"
			reverseFieldOwner := models.GetReverseFieldOwner(
				astructFormCallback.probe.stageOfInterest,
				astruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofa, astruct_)
					pastAstructOwner.Anarrayofa = slices.Delete(pastAstructOwner.Anarrayofa, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the astruct_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.Anarrayofa, astruct_)
									pastAstructOwner.Anarrayofa = slices.Delete(pastAstructOwner.Anarrayofa, idx, idx+1)
									newAstructOwner.Anarrayofa = append(newAstructOwner.Anarrayofa, astruct_)
								}
							} else {
								newAstructOwner.Anarrayofa = append(newAstructOwner.Anarrayofa, astruct_)
							}
						}
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
	updateAndCommitTable[models.Astruct](
		astructFormCallback.probe,
	)
	astructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructFormCallback.CreationMode || astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(astructFormCallback.probe)
}
func __gong__New__AstructBstruct2UseFormCallback(
	astructbstruct2use *models.AstructBstruct2Use,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) OnSave() {

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
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofb2Use"
			reverseFieldOwner := models.GetReverseFieldOwner(
				astructbstruct2useFormCallback.probe.stageOfInterest,
				astructbstruct2use_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofb2Use, astructbstruct2use_)
					pastAstructOwner.Anarrayofb2Use = slices.Delete(pastAstructOwner.Anarrayofb2Use, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstruct2useFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the astructbstruct2use_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.Anarrayofb2Use, astructbstruct2use_)
									pastAstructOwner.Anarrayofb2Use = slices.Delete(pastAstructOwner.Anarrayofb2Use, idx, idx+1)
									newAstructOwner.Anarrayofb2Use = append(newAstructOwner.Anarrayofb2Use, astructbstruct2use_)
								}
							} else {
								newAstructOwner.Anarrayofb2Use = append(newAstructOwner.Anarrayofb2Use, astructbstruct2use_)
							}
						}
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
	updateAndCommitTable[models.AstructBstruct2Use](
		astructbstruct2useFormCallback.probe,
	)
	astructbstruct2useFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstruct2useFormCallback.CreationMode || astructbstruct2useFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstruct2useFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(astructbstruct2useFormCallback.probe)
}
func __gong__New__AstructBstructUseFormCallback(
	astructbstructuse *models.AstructBstructUse,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (astructbstructuseFormCallback *AstructBstructUseFormCallback) OnSave() {

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
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "AnarrayofbUse"
			reverseFieldOwner := models.GetReverseFieldOwner(
				astructbstructuseFormCallback.probe.stageOfInterest,
				astructbstructuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.AnarrayofbUse, astructbstructuse_)
					pastAstructOwner.AnarrayofbUse = slices.Delete(pastAstructOwner.AnarrayofbUse, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstructuseFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the astructbstructuse_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.AnarrayofbUse, astructbstructuse_)
									pastAstructOwner.AnarrayofbUse = slices.Delete(pastAstructOwner.AnarrayofbUse, idx, idx+1)
									newAstructOwner.AnarrayofbUse = append(newAstructOwner.AnarrayofbUse, astructbstructuse_)
								}
							} else {
								newAstructOwner.AnarrayofbUse = append(newAstructOwner.AnarrayofbUse, astructbstructuse_)
							}
						}
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
	updateAndCommitTable[models.AstructBstructUse](
		astructbstructuseFormCallback.probe,
	)
	astructbstructuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstructuseFormCallback.CreationMode || astructbstructuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstructuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(astructbstructuseFormCallback.probe)
}
func __gong__New__BstructFormCallback(
	bstruct *models.Bstruct,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (bstructFormCallback *BstructFormCallback) OnSave() {

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
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofb"
			reverseFieldOwner := models.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofb, bstruct_)
					pastAstructOwner.Anarrayofb = slices.Delete(pastAstructOwner.Anarrayofb, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the bstruct_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.Anarrayofb, bstruct_)
									pastAstructOwner.Anarrayofb = slices.Delete(pastAstructOwner.Anarrayofb, idx, idx+1)
									newAstructOwner.Anarrayofb = append(newAstructOwner.Anarrayofb, bstruct_)
								}
							} else {
								newAstructOwner.Anarrayofb = append(newAstructOwner.Anarrayofb, bstruct_)
							}
						}
					}
				}
			}
		case "Astruct:Anotherarrayofb":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anotherarrayofb"
			reverseFieldOwner := models.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anotherarrayofb, bstruct_)
					pastAstructOwner.Anotherarrayofb = slices.Delete(pastAstructOwner.Anotherarrayofb, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the bstruct_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.Anotherarrayofb, bstruct_)
									pastAstructOwner.Anotherarrayofb = slices.Delete(pastAstructOwner.Anotherarrayofb, idx, idx+1)
									newAstructOwner.Anotherarrayofb = append(newAstructOwner.Anotherarrayofb, bstruct_)
								}
							} else {
								newAstructOwner.Anotherarrayofb = append(newAstructOwner.Anotherarrayofb, bstruct_)
							}
						}
					}
				}
			}
		case "Dstruct:Anarrayofb":
			// we need to retrieve the field owner before the change
			var pastDstructOwner *models.Dstruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Dstruct"
			rf.Fieldname = "Anarrayofb"
			reverseFieldOwner := models.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastDstructOwner = reverseFieldOwner.(*models.Dstruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastDstructOwner != nil {
					idx := slices.Index(pastDstructOwner.Anarrayofb, bstruct_)
					pastDstructOwner.Anarrayofb = slices.Delete(pastDstructOwner.Anarrayofb, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastDstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _dstruct := range *models.GetGongstructInstancesSet[models.Dstruct](bstructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _dstruct.GetName() == fieldValue.GetName() {
							newDstructOwner := _dstruct // we have a match

							// we remove the bstruct_ instance from the pastDstructOwner field
							if pastDstructOwner != nil {
								if newDstructOwner != pastDstructOwner {
									idx := slices.Index(pastDstructOwner.Anarrayofb, bstruct_)
									pastDstructOwner.Anarrayofb = slices.Delete(pastDstructOwner.Anarrayofb, idx, idx+1)
									newDstructOwner.Anarrayofb = append(newDstructOwner.Anarrayofb, bstruct_)
								}
							} else {
								newDstructOwner.Anarrayofb = append(newDstructOwner.Anarrayofb, bstruct_)
							}
						}
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
	updateAndCommitTable[models.Bstruct](
		bstructFormCallback.probe,
	)
	bstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bstructFormCallback.CreationMode || bstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(bstructFormCallback.probe)
}
func __gong__New__DstructFormCallback(
	dstruct *models.Dstruct,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (dstructFormCallback *DstructFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			dstruct_.Anarrayofb = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			dstruct_.Gstructs = instanceSlice

		case "Astruct:Dstruct4s":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Dstruct4s"
			reverseFieldOwner := models.GetReverseFieldOwner(
				dstructFormCallback.probe.stageOfInterest,
				dstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Dstruct4s, dstruct_)
					pastAstructOwner.Dstruct4s = slices.Delete(pastAstructOwner.Dstruct4s, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](dstructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _astruct.GetName() == fieldValue.GetName() {
							newAstructOwner := _astruct // we have a match

							// we remove the dstruct_ instance from the pastAstructOwner field
							if pastAstructOwner != nil {
								if newAstructOwner != pastAstructOwner {
									idx := slices.Index(pastAstructOwner.Dstruct4s, dstruct_)
									pastAstructOwner.Dstruct4s = slices.Delete(pastAstructOwner.Dstruct4s, idx, idx+1)
									newAstructOwner.Dstruct4s = append(newAstructOwner.Dstruct4s, dstruct_)
								}
							} else {
								newAstructOwner.Dstruct4s = append(newAstructOwner.Dstruct4s, dstruct_)
							}
						}
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
	updateAndCommitTable[models.Dstruct](
		dstructFormCallback.probe,
	)
	dstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dstructFormCallback.CreationMode || dstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(dstructFormCallback.probe)
}
func __gong__New__FstructFormCallback(
	fstruct *models.Fstruct,
	probe *Probe,
	formGroup *table.FormGroup,
) (fstructFormCallback *FstructFormCallback) {
	fstructFormCallback = new(FstructFormCallback)
	fstructFormCallback.probe = probe
	fstructFormCallback.fstruct = fstruct
	fstructFormCallback.formGroup = formGroup

	fstructFormCallback.CreationMode = (fstruct == nil)

	return
}

type FstructFormCallback struct {
	fstruct *models.Fstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fstructFormCallback *FstructFormCallback) OnSave() {

	// log.Println("FstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fstructFormCallback.probe.formStage.Checkout()

	if fstructFormCallback.fstruct == nil {
		fstructFormCallback.fstruct = new(models.Fstruct).Stage(fstructFormCallback.probe.stageOfInterest)
	}
	fstruct_ := fstructFormCallback.fstruct
	_ = fstruct_

	for _, formDiv := range fstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(fstruct_.Name), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(fstruct_.Date), formDiv)
		}
	}

	// manage the suppress operation
	if fstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fstruct_.Unstage(fstructFormCallback.probe.stageOfInterest)
	}

	fstructFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Fstruct](
		fstructFormCallback.probe,
	)
	fstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fstructFormCallback.CreationMode || fstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(fstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FstructFormCallback(
			nil,
			fstructFormCallback.probe,
			newFormGroup,
		)
		fstruct := new(models.Fstruct)
		FillUpForm(fstruct, newFormGroup, fstructFormCallback.probe)
		fstructFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(fstructFormCallback.probe)
}
func __gong__New__GstructFormCallback(
	gstruct *models.Gstruct,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gstructFormCallback *GstructFormCallback) OnSave() {

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
			// we need to retrieve the field owner before the change
			var pastDstructOwner *models.Dstruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Dstruct"
			rf.Fieldname = "Gstructs"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gstructFormCallback.probe.stageOfInterest,
				gstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastDstructOwner = reverseFieldOwner.(*models.Dstruct)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastDstructOwner != nil {
					idx := slices.Index(pastDstructOwner.Gstructs, gstruct_)
					pastDstructOwner.Gstructs = slices.Delete(pastDstructOwner.Gstructs, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastDstructOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _dstruct := range *models.GetGongstructInstancesSet[models.Dstruct](gstructFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _dstruct.GetName() == fieldValue.GetName() {
							newDstructOwner := _dstruct // we have a match

							// we remove the gstruct_ instance from the pastDstructOwner field
							if pastDstructOwner != nil {
								if newDstructOwner != pastDstructOwner {
									idx := slices.Index(pastDstructOwner.Gstructs, gstruct_)
									pastDstructOwner.Gstructs = slices.Delete(pastDstructOwner.Gstructs, idx, idx+1)
									newDstructOwner.Gstructs = append(newDstructOwner.Gstructs, gstruct_)
								}
							} else {
								newDstructOwner.Gstructs = append(newDstructOwner.Gstructs, gstruct_)
							}
						}
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
	updateAndCommitTable[models.Gstruct](
		gstructFormCallback.probe,
	)
	gstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gstructFormCallback.CreationMode || gstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(gstructFormCallback.probe)
}
