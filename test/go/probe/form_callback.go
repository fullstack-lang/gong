// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

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

	log.Println("AstructFormCallback, OnSave")

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
		case "Anotherassociationtob_2":
			FormDivSelectFieldToField(&(astruct_.Anotherassociationtob_2), astructFormCallback.probe.stageOfInterest, formDiv)
		case "Date":
			FormDivBasicFieldToField(&(astruct_.Date), formDiv)
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
		case "Floatfield":
			FormDivBasicFieldToField(&(astruct_.Floatfield), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(astruct_.Intfield), formDiv)
		case "Anotherbooleanfield":
			FormDivBasicFieldToField(&(astruct_.Anotherbooleanfield), formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&(astruct_.Duration1), formDiv)
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
			reverseFieldOwner := orm.GetReverseFieldOwner(
				astructFormCallback.probe.stageOfInterest,
				astructFormCallback.probe.backRepoOfInterest,
				astruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofa, astruct_)
					pastAstructOwner.Anarrayofa = slices.Delete(pastAstructOwner.Anarrayofa, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _astruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAstructOwner := _astruct // we have a match
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

	// manage the suppress operation
	if astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astruct_.Unstage(astructFormCallback.probe.stageOfInterest)
	}

	astructFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Astruct](
		astructFormCallback.probe,
	)
	astructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructFormCallback.CreationMode || astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(astructFormCallback.probe)
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

	log.Println("AstructBstruct2UseFormCallback, OnSave")

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
			reverseFieldOwner := orm.GetReverseFieldOwner(
				astructbstruct2useFormCallback.probe.stageOfInterest,
				astructbstruct2useFormCallback.probe.backRepoOfInterest,
				astructbstruct2use_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofb2Use, astructbstruct2use_)
					pastAstructOwner.Anarrayofb2Use = slices.Delete(pastAstructOwner.Anarrayofb2Use, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstruct2useFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _astruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAstructOwner := _astruct // we have a match
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

	// manage the suppress operation
	if astructbstruct2useFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstruct2use_.Unstage(astructbstruct2useFormCallback.probe.stageOfInterest)
	}

	astructbstruct2useFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AstructBstruct2Use](
		astructbstruct2useFormCallback.probe,
	)
	astructbstruct2useFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstruct2useFormCallback.CreationMode || astructbstruct2useFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstruct2useFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(astructbstruct2useFormCallback.probe)
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

	log.Println("AstructBstructUseFormCallback, OnSave")

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
			reverseFieldOwner := orm.GetReverseFieldOwner(
				astructbstructuseFormCallback.probe.stageOfInterest,
				astructbstructuseFormCallback.probe.backRepoOfInterest,
				astructbstructuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.AnarrayofbUse, astructbstructuse_)
					pastAstructOwner.AnarrayofbUse = slices.Delete(pastAstructOwner.AnarrayofbUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstructuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _astruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAstructOwner := _astruct // we have a match
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

	// manage the suppress operation
	if astructbstructuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstructuse_.Unstage(astructbstructuseFormCallback.probe.stageOfInterest)
	}

	astructbstructuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AstructBstructUse](
		astructbstructuseFormCallback.probe,
	)
	astructbstructuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstructuseFormCallback.CreationMode || astructbstructuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructbstructuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(astructbstructuseFormCallback.probe)
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

	log.Println("BstructFormCallback, OnSave")

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
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstructFormCallback.probe.backRepoOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anarrayofb, bstruct_)
					pastAstructOwner.Anarrayofb = slices.Delete(pastAstructOwner.Anarrayofb, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _astruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAstructOwner := _astruct // we have a match
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
		case "Astruct:Anotherarrayofb":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anotherarrayofb"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstructFormCallback.probe.backRepoOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastAstructOwner = reverseFieldOwner.(*models.Astruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAstructOwner != nil {
					idx := slices.Index(pastAstructOwner.Anotherarrayofb, bstruct_)
					pastAstructOwner.Anotherarrayofb = slices.Delete(pastAstructOwner.Anotherarrayofb, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _astruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAstructOwner := _astruct // we have a match
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
		case "Dstruct:Anarrayofb":
			// we need to retrieve the field owner before the change
			var pastDstructOwner *models.Dstruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Dstruct"
			rf.Fieldname = "Anarrayofb"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bstructFormCallback.probe.stageOfInterest,
				bstructFormCallback.probe.backRepoOfInterest,
				bstruct_,
				&rf)

			if reverseFieldOwner != nil {
				pastDstructOwner = reverseFieldOwner.(*models.Dstruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDstructOwner != nil {
					idx := slices.Index(pastDstructOwner.Anarrayofb, bstruct_)
					pastDstructOwner.Anarrayofb = slices.Delete(pastDstructOwner.Anarrayofb, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _dstruct := range *models.GetGongstructInstancesSet[models.Dstruct](bstructFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _dstruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDstructOwner := _dstruct // we have a match
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

	// manage the suppress operation
	if bstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bstruct_.Unstage(bstructFormCallback.probe.stageOfInterest)
	}

	bstructFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bstruct](
		bstructFormCallback.probe,
	)
	bstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bstructFormCallback.CreationMode || bstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(bstructFormCallback.probe)
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

	log.Println("DstructFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if dstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dstruct_.Unstage(dstructFormCallback.probe.stageOfInterest)
	}

	dstructFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Dstruct](
		dstructFormCallback.probe,
	)
	dstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dstructFormCallback.CreationMode || dstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(dstructFormCallback.probe)
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

	log.Println("FstructFormCallback, OnSave")

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
	fillUpTable[models.Fstruct](
		fstructFormCallback.probe,
	)
	fstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fstructFormCallback.CreationMode || fstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(fstructFormCallback.probe)
}
