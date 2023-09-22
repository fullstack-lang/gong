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

// insertion point
func __gong__New__AstructFormCallback(
	astruct *models.Astruct,
	playground *Playground,
) (astructFormCallback *AstructFormCallback) {
	astructFormCallback = new(AstructFormCallback)
	astructFormCallback.playground = playground
	astructFormCallback.astruct = astruct

	astructFormCallback.CreationMode = (astruct == nil)

	return
}

type AstructFormCallback struct {
	astruct *models.Astruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (astructFormCallback *AstructFormCallback) OnSave() {

	log.Println("AstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructFormCallback.playground.formStage.Checkout()

	if astructFormCallback.astruct == nil {
		astructFormCallback.astruct = new(models.Astruct).Stage(astructFormCallback.playground.stageOfInterest)
	}
	astruct_ := astructFormCallback.astruct
	_ = astruct_

	// get the formGroup
	formGroup := astructFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astruct_.Name), formDiv)
		case "Associationtob":
			FormDivSelectFieldToField(&(astruct_.Associationtob), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Anotherassociationtob_2":
			FormDivSelectFieldToField(&(astruct_.Anotherassociationtob_2), astructFormCallback.playground.stageOfInterest, formDiv)
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
			FormDivSelectFieldToField(&(astruct_.Bstruct), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(astruct_.Bstruct2), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Dstruct":
			FormDivSelectFieldToField(&(astruct_.Dstruct), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Dstruct2":
			FormDivSelectFieldToField(&(astruct_.Dstruct2), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Dstruct3":
			FormDivSelectFieldToField(&(astruct_.Dstruct3), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Dstruct4":
			FormDivSelectFieldToField(&(astruct_.Dstruct4), astructFormCallback.playground.stageOfInterest, formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(astruct_.Floatfield), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(astruct_.Intfield), formDiv)
		case "Anotherbooleanfield":
			FormDivBasicFieldToField(&(astruct_.Anotherbooleanfield), formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&(astruct_.Duration1), formDiv)
		case "AnAstruct":
			FormDivSelectFieldToField(&(astruct_.AnAstruct), astructFormCallback.playground.stageOfInterest, formDiv)
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
				astructFormCallback.playground.stageOfInterest,
				astructFormCallback.playground.backRepoOfInterest,
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
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructFormCallback.playground.stageOfInterest) {

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

	astructFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Astruct](
		astructFormCallback.playground,
	)
	astructFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if astructFormCallback.CreationMode {
		astructFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AstructFormCallback(
				nil,
				astructFormCallback.playground,
			),
		}).Stage(astructFormCallback.playground.formStage)
		astruct := new(models.Astruct)
		FillUpForm(astruct, newFormGroup, astructFormCallback.playground)
		astructFormCallback.playground.formStage.Commit()
	}

}
func __gong__New__AstructBstruct2UseFormCallback(
	astructbstruct2use *models.AstructBstruct2Use,
	playground *Playground,
) (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) {
	astructbstruct2useFormCallback = new(AstructBstruct2UseFormCallback)
	astructbstruct2useFormCallback.playground = playground
	astructbstruct2useFormCallback.astructbstruct2use = astructbstruct2use

	astructbstruct2useFormCallback.CreationMode = (astructbstruct2use == nil)

	return
}

type AstructBstruct2UseFormCallback struct {
	astructbstruct2use *models.AstructBstruct2Use

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) OnSave() {

	log.Println("AstructBstruct2UseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructbstruct2useFormCallback.playground.formStage.Checkout()

	if astructbstruct2useFormCallback.astructbstruct2use == nil {
		astructbstruct2useFormCallback.astructbstruct2use = new(models.AstructBstruct2Use).Stage(astructbstruct2useFormCallback.playground.stageOfInterest)
	}
	astructbstruct2use_ := astructbstruct2useFormCallback.astructbstruct2use
	_ = astructbstruct2use_

	// get the formGroup
	formGroup := astructbstruct2useFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astructbstruct2use_.Name), formDiv)
		case "Bstrcut2":
			FormDivSelectFieldToField(&(astructbstruct2use_.Bstrcut2), astructbstruct2useFormCallback.playground.stageOfInterest, formDiv)
		case "Astruct:Anarrayofb2Use":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofb2Use"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				astructbstruct2useFormCallback.playground.stageOfInterest,
				astructbstruct2useFormCallback.playground.backRepoOfInterest,
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
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstruct2useFormCallback.playground.stageOfInterest) {

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

	astructbstruct2useFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.AstructBstruct2Use](
		astructbstruct2useFormCallback.playground,
	)
	astructbstruct2useFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstruct2useFormCallback.CreationMode {
		astructbstruct2useFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AstructBstruct2UseFormCallback(
				nil,
				astructbstruct2useFormCallback.playground,
			),
		}).Stage(astructbstruct2useFormCallback.playground.formStage)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, newFormGroup, astructbstruct2useFormCallback.playground)
		astructbstruct2useFormCallback.playground.formStage.Commit()
	}

}
func __gong__New__AstructBstructUseFormCallback(
	astructbstructuse *models.AstructBstructUse,
	playground *Playground,
) (astructbstructuseFormCallback *AstructBstructUseFormCallback) {
	astructbstructuseFormCallback = new(AstructBstructUseFormCallback)
	astructbstructuseFormCallback.playground = playground
	astructbstructuseFormCallback.astructbstructuse = astructbstructuse

	astructbstructuseFormCallback.CreationMode = (astructbstructuse == nil)

	return
}

type AstructBstructUseFormCallback struct {
	astructbstructuse *models.AstructBstructUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (astructbstructuseFormCallback *AstructBstructUseFormCallback) OnSave() {

	log.Println("AstructBstructUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructbstructuseFormCallback.playground.formStage.Checkout()

	if astructbstructuseFormCallback.astructbstructuse == nil {
		astructbstructuseFormCallback.astructbstructuse = new(models.AstructBstructUse).Stage(astructbstructuseFormCallback.playground.stageOfInterest)
	}
	astructbstructuse_ := astructbstructuseFormCallback.astructbstructuse
	_ = astructbstructuse_

	// get the formGroup
	formGroup := astructbstructuseFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(astructbstructuse_.Name), formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(astructbstructuse_.Bstruct2), astructbstructuseFormCallback.playground.stageOfInterest, formDiv)
		case "Astruct:AnarrayofbUse":
			// we need to retrieve the field owner before the change
			var pastAstructOwner *models.Astruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "AnarrayofbUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				astructbstructuseFormCallback.playground.stageOfInterest,
				astructbstructuseFormCallback.playground.backRepoOfInterest,
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
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](astructbstructuseFormCallback.playground.stageOfInterest) {

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

	astructbstructuseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.AstructBstructUse](
		astructbstructuseFormCallback.playground,
	)
	astructbstructuseFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if astructbstructuseFormCallback.CreationMode {
		astructbstructuseFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AstructBstructUseFormCallback(
				nil,
				astructbstructuseFormCallback.playground,
			),
		}).Stage(astructbstructuseFormCallback.playground.formStage)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, newFormGroup, astructbstructuseFormCallback.playground)
		astructbstructuseFormCallback.playground.formStage.Commit()
	}

}
func __gong__New__BstructFormCallback(
	bstruct *models.Bstruct,
	playground *Playground,
) (bstructFormCallback *BstructFormCallback) {
	bstructFormCallback = new(BstructFormCallback)
	bstructFormCallback.playground = playground
	bstructFormCallback.bstruct = bstruct

	bstructFormCallback.CreationMode = (bstruct == nil)

	return
}

type BstructFormCallback struct {
	bstruct *models.Bstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (bstructFormCallback *BstructFormCallback) OnSave() {

	log.Println("BstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bstructFormCallback.playground.formStage.Checkout()

	if bstructFormCallback.bstruct == nil {
		bstructFormCallback.bstruct = new(models.Bstruct).Stage(bstructFormCallback.playground.stageOfInterest)
	}
	bstruct_ := bstructFormCallback.bstruct
	_ = bstruct_

	// get the formGroup
	formGroup := bstructFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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
				bstructFormCallback.playground.stageOfInterest,
				bstructFormCallback.playground.backRepoOfInterest,
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
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.playground.stageOfInterest) {

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
				bstructFormCallback.playground.stageOfInterest,
				bstructFormCallback.playground.backRepoOfInterest,
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
				for _astruct := range *models.GetGongstructInstancesSet[models.Astruct](bstructFormCallback.playground.stageOfInterest) {

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
				bstructFormCallback.playground.stageOfInterest,
				bstructFormCallback.playground.backRepoOfInterest,
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
				for _dstruct := range *models.GetGongstructInstancesSet[models.Dstruct](bstructFormCallback.playground.stageOfInterest) {

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

	bstructFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Bstruct](
		bstructFormCallback.playground,
	)
	bstructFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if bstructFormCallback.CreationMode {
		bstructFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BstructFormCallback(
				nil,
				bstructFormCallback.playground,
			),
		}).Stage(bstructFormCallback.playground.formStage)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, newFormGroup, bstructFormCallback.playground)
		bstructFormCallback.playground.formStage.Commit()
	}

}
func __gong__New__DstructFormCallback(
	dstruct *models.Dstruct,
	playground *Playground,
) (dstructFormCallback *DstructFormCallback) {
	dstructFormCallback = new(DstructFormCallback)
	dstructFormCallback.playground = playground
	dstructFormCallback.dstruct = dstruct

	dstructFormCallback.CreationMode = (dstruct == nil)

	return
}

type DstructFormCallback struct {
	dstruct *models.Dstruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (dstructFormCallback *DstructFormCallback) OnSave() {

	log.Println("DstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dstructFormCallback.playground.formStage.Checkout()

	if dstructFormCallback.dstruct == nil {
		dstructFormCallback.dstruct = new(models.Dstruct).Stage(dstructFormCallback.playground.stageOfInterest)
	}
	dstruct_ := dstructFormCallback.dstruct
	_ = dstruct_

	// get the formGroup
	formGroup := dstructFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dstruct_.Name), formDiv)
		}
	}

	dstructFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Dstruct](
		dstructFormCallback.playground,
	)
	dstructFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if dstructFormCallback.CreationMode {
		dstructFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DstructFormCallback(
				nil,
				dstructFormCallback.playground,
			),
		}).Stage(dstructFormCallback.playground.formStage)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, newFormGroup, dstructFormCallback.playground)
		dstructFormCallback.playground.formStage.Commit()
	}

}
