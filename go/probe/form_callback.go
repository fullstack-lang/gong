// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__GongBasicFieldFormCallback(
	gongbasicfield *models.GongBasicField,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongbasicfieldFormCallback *GongBasicFieldFormCallback) OnSave() {

	log.Println("GongBasicFieldFormCallback, OnSave")

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
		case "Index":
			FormDivBasicFieldToField(&(gongbasicfield_.Index), formDiv)
		case "IsDocLink":
			FormDivBasicFieldToField(&(gongbasicfield_.IsDocLink), formDiv)
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
			// we need to retrieve the field owner before the change
			var pastGongStructOwner *models.GongStruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongBasicFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongbasicfieldFormCallback.probe.stageOfInterest,
				gongbasicfieldFormCallback.probe.backRepoOfInterest,
				gongbasicfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructOwner = reverseFieldOwner.(*models.GongStruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructOwner != nil {
					idx := slices.Index(pastGongStructOwner.GongBasicFields, gongbasicfield_)
					pastGongStructOwner.GongBasicFields = slices.Delete(pastGongStructOwner.GongBasicFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](gongbasicfieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructOwner := _gongstruct // we have a match
						if pastGongStructOwner != nil {
							if newGongStructOwner != pastGongStructOwner {
								idx := slices.Index(pastGongStructOwner.GongBasicFields, gongbasicfield_)
								pastGongStructOwner.GongBasicFields = slices.Delete(pastGongStructOwner.GongBasicFields, idx, idx+1)
								newGongStructOwner.GongBasicFields = append(newGongStructOwner.GongBasicFields, gongbasicfield_)
							}
						} else {
							newGongStructOwner.GongBasicFields = append(newGongStructOwner.GongBasicFields, gongbasicfield_)
						}
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
	fillUpTable[models.GongBasicField](
		gongbasicfieldFormCallback.probe,
	)
	gongbasicfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongbasicfieldFormCallback.CreationMode || gongbasicfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongbasicfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongbasicfieldFormCallback.probe)
}
func __gong__New__GongEnumFormCallback(
	gongenum *models.GongEnum,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongenumFormCallback *GongEnumFormCallback) OnSave() {

	log.Println("GongEnumFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenum_.Unstage(gongenumFormCallback.probe.stageOfInterest)
	}

	gongenumFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongEnum](
		gongenumFormCallback.probe,
	)
	gongenumFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumFormCallback.CreationMode || gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongenumFormCallback.probe)
}
func __gong__New__GongEnumValueFormCallback(
	gongenumvalue *models.GongEnumValue,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongenumvalueFormCallback *GongEnumValueFormCallback) OnSave() {

	log.Println("GongEnumValueFormCallback, OnSave")

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
			// we need to retrieve the field owner before the change
			var pastGongEnumOwner *models.GongEnum
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnum"
			rf.Fieldname = "GongEnumValues"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongenumvalueFormCallback.probe.stageOfInterest,
				gongenumvalueFormCallback.probe.backRepoOfInterest,
				gongenumvalue_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongEnumOwner = reverseFieldOwner.(*models.GongEnum)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongEnumOwner != nil {
					idx := slices.Index(pastGongEnumOwner.GongEnumValues, gongenumvalue_)
					pastGongEnumOwner.GongEnumValues = slices.Delete(pastGongEnumOwner.GongEnumValues, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongenum := range *models.GetGongstructInstancesSet[models.GongEnum](gongenumvalueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongenum.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongEnumOwner := _gongenum // we have a match
						if pastGongEnumOwner != nil {
							if newGongEnumOwner != pastGongEnumOwner {
								idx := slices.Index(pastGongEnumOwner.GongEnumValues, gongenumvalue_)
								pastGongEnumOwner.GongEnumValues = slices.Delete(pastGongEnumOwner.GongEnumValues, idx, idx+1)
								newGongEnumOwner.GongEnumValues = append(newGongEnumOwner.GongEnumValues, gongenumvalue_)
							}
						} else {
							newGongEnumOwner.GongEnumValues = append(newGongEnumOwner.GongEnumValues, gongenumvalue_)
						}
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
	fillUpTable[models.GongEnumValue](
		gongenumvalueFormCallback.probe,
	)
	gongenumvalueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueFormCallback.CreationMode || gongenumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongenumvalueFormCallback.probe)
}
func __gong__New__GongLinkFormCallback(
	gonglink *models.GongLink,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gonglinkFormCallback *GongLinkFormCallback) OnSave() {

	log.Println("GongLinkFormCallback, OnSave")

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
			// we need to retrieve the field owner before the change
			var pastGongNoteOwner *models.GongNote
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongNote"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gonglinkFormCallback.probe.stageOfInterest,
				gonglinkFormCallback.probe.backRepoOfInterest,
				gonglink_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongNoteOwner = reverseFieldOwner.(*models.GongNote)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongNoteOwner != nil {
					idx := slices.Index(pastGongNoteOwner.Links, gonglink_)
					pastGongNoteOwner.Links = slices.Delete(pastGongNoteOwner.Links, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongnote := range *models.GetGongstructInstancesSet[models.GongNote](gonglinkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongnote.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongNoteOwner := _gongnote // we have a match
						if pastGongNoteOwner != nil {
							if newGongNoteOwner != pastGongNoteOwner {
								idx := slices.Index(pastGongNoteOwner.Links, gonglink_)
								pastGongNoteOwner.Links = slices.Delete(pastGongNoteOwner.Links, idx, idx+1)
								newGongNoteOwner.Links = append(newGongNoteOwner.Links, gonglink_)
							}
						} else {
							newGongNoteOwner.Links = append(newGongNoteOwner.Links, gonglink_)
						}
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
	fillUpTable[models.GongLink](
		gonglinkFormCallback.probe,
	)
	gonglinkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gonglinkFormCallback.CreationMode || gonglinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gonglinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gonglinkFormCallback.probe)
}
func __gong__New__GongNoteFormCallback(
	gongnote *models.GongNote,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongnoteFormCallback *GongNoteFormCallback) OnSave() {

	log.Println("GongNoteFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnote_.Unstage(gongnoteFormCallback.probe.stageOfInterest)
	}

	gongnoteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongNote](
		gongnoteFormCallback.probe,
	)
	gongnoteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongnoteFormCallback.CreationMode || gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongnoteFormCallback.probe)
}
func __gong__New__GongStructFormCallback(
	gongstruct *models.GongStruct,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongstructFormCallback *GongStructFormCallback) OnSave() {

	log.Println("GongStructFormCallback, OnSave")

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
	fillUpTable[models.GongStruct](
		gongstructFormCallback.probe,
	)
	gongstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongstructFormCallback.CreationMode || gongstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongstructFormCallback.probe)
}
func __gong__New__GongTimeFieldFormCallback(
	gongtimefield *models.GongTimeField,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (gongtimefieldFormCallback *GongTimeFieldFormCallback) OnSave() {

	log.Println("GongTimeFieldFormCallback, OnSave")

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
		case "BespokeTimeFormat":
			FormDivBasicFieldToField(&(gongtimefield_.BespokeTimeFormat), formDiv)
		case "GongStruct:GongTimeFields":
			// we need to retrieve the field owner before the change
			var pastGongStructOwner *models.GongStruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongTimeFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongtimefieldFormCallback.probe.stageOfInterest,
				gongtimefieldFormCallback.probe.backRepoOfInterest,
				gongtimefield_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructOwner = reverseFieldOwner.(*models.GongStruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructOwner != nil {
					idx := slices.Index(pastGongStructOwner.GongTimeFields, gongtimefield_)
					pastGongStructOwner.GongTimeFields = slices.Delete(pastGongStructOwner.GongTimeFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](gongtimefieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructOwner := _gongstruct // we have a match
						if pastGongStructOwner != nil {
							if newGongStructOwner != pastGongStructOwner {
								idx := slices.Index(pastGongStructOwner.GongTimeFields, gongtimefield_)
								pastGongStructOwner.GongTimeFields = slices.Delete(pastGongStructOwner.GongTimeFields, idx, idx+1)
								newGongStructOwner.GongTimeFields = append(newGongStructOwner.GongTimeFields, gongtimefield_)
							}
						} else {
							newGongStructOwner.GongTimeFields = append(newGongStructOwner.GongTimeFields, gongtimefield_)
						}
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
	fillUpTable[models.GongTimeField](
		gongtimefieldFormCallback.probe,
	)
	gongtimefieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongtimefieldFormCallback.CreationMode || gongtimefieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongtimefieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(gongtimefieldFormCallback.probe)
}
func __gong__New__MetaFormCallback(
	meta *models.Meta,
	probe *Probe,
	formGroup *table.FormGroup,
) (metaFormCallback *MetaFormCallback) {
	metaFormCallback = new(MetaFormCallback)
	metaFormCallback.probe = probe
	metaFormCallback.meta = meta
	metaFormCallback.formGroup = formGroup

	metaFormCallback.CreationMode = (meta == nil)

	return
}

type MetaFormCallback struct {
	meta *models.Meta

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metaFormCallback *MetaFormCallback) OnSave() {

	log.Println("MetaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metaFormCallback.probe.formStage.Checkout()

	if metaFormCallback.meta == nil {
		metaFormCallback.meta = new(models.Meta).Stage(metaFormCallback.probe.stageOfInterest)
	}
	meta_ := metaFormCallback.meta
	_ = meta_

	for _, formDiv := range metaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(meta_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(meta_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if metaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meta_.Unstage(metaFormCallback.probe.stageOfInterest)
	}

	metaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Meta](
		metaFormCallback.probe,
	)
	metaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metaFormCallback.CreationMode || metaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MetaFormCallback(
			nil,
			metaFormCallback.probe,
			newFormGroup,
		)
		meta := new(models.Meta)
		FillUpForm(meta, newFormGroup, metaFormCallback.probe)
		metaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metaFormCallback.probe)
}
func __gong__New__MetaReferenceFormCallback(
	metareference *models.MetaReference,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (metareferenceFormCallback *MetaReferenceFormCallback) OnSave() {

	log.Println("MetaReferenceFormCallback, OnSave")

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
		case "Meta:MetaReferences":
			// we need to retrieve the field owner before the change
			var pastMetaOwner *models.Meta
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Meta"
			rf.Fieldname = "MetaReferences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				metareferenceFormCallback.probe.stageOfInterest,
				metareferenceFormCallback.probe.backRepoOfInterest,
				metareference_,
				&rf)

			if reverseFieldOwner != nil {
				pastMetaOwner = reverseFieldOwner.(*models.Meta)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMetaOwner != nil {
					idx := slices.Index(pastMetaOwner.MetaReferences, metareference_)
					pastMetaOwner.MetaReferences = slices.Delete(pastMetaOwner.MetaReferences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _meta := range *models.GetGongstructInstancesSet[models.Meta](metareferenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _meta.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMetaOwner := _meta // we have a match
						if pastMetaOwner != nil {
							if newMetaOwner != pastMetaOwner {
								idx := slices.Index(pastMetaOwner.MetaReferences, metareference_)
								pastMetaOwner.MetaReferences = slices.Delete(pastMetaOwner.MetaReferences, idx, idx+1)
								newMetaOwner.MetaReferences = append(newMetaOwner.MetaReferences, metareference_)
							}
						} else {
							newMetaOwner.MetaReferences = append(newMetaOwner.MetaReferences, metareference_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareference_.Unstage(metareferenceFormCallback.probe.stageOfInterest)
	}

	metareferenceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MetaReference](
		metareferenceFormCallback.probe,
	)
	metareferenceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metareferenceFormCallback.CreationMode || metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareferenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(metareferenceFormCallback.probe)
}
func __gong__New__ModelPkgFormCallback(
	modelpkg *models.ModelPkg,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (modelpkgFormCallback *ModelPkgFormCallback) OnSave() {

	log.Println("ModelPkgFormCallback, OnSave")

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
	fillUpTable[models.ModelPkg](
		modelpkgFormCallback.probe,
	)
	modelpkgFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if modelpkgFormCallback.CreationMode || modelpkgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		modelpkgFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(modelpkgFormCallback.probe)
}
func __gong__New__PointerToGongStructFieldFormCallback(
	pointertogongstructfield *models.PointerToGongStructField,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) OnSave() {

	log.Println("PointerToGongStructFieldFormCallback, OnSave")

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
		case "IsType":
			FormDivBasicFieldToField(&(pointertogongstructfield_.IsType), formDiv)
		case "GongStruct:PointerToGongStructFields":
			// we need to retrieve the field owner before the change
			var pastGongStructOwner *models.GongStruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "PointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				pointertogongstructfieldFormCallback.probe.stageOfInterest,
				pointertogongstructfieldFormCallback.probe.backRepoOfInterest,
				pointertogongstructfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructOwner = reverseFieldOwner.(*models.GongStruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructOwner != nil {
					idx := slices.Index(pastGongStructOwner.PointerToGongStructFields, pointertogongstructfield_)
					pastGongStructOwner.PointerToGongStructFields = slices.Delete(pastGongStructOwner.PointerToGongStructFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](pointertogongstructfieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructOwner := _gongstruct // we have a match
						if pastGongStructOwner != nil {
							if newGongStructOwner != pastGongStructOwner {
								idx := slices.Index(pastGongStructOwner.PointerToGongStructFields, pointertogongstructfield_)
								pastGongStructOwner.PointerToGongStructFields = slices.Delete(pastGongStructOwner.PointerToGongStructFields, idx, idx+1)
								newGongStructOwner.PointerToGongStructFields = append(newGongStructOwner.PointerToGongStructFields, pointertogongstructfield_)
							}
						} else {
							newGongStructOwner.PointerToGongStructFields = append(newGongStructOwner.PointerToGongStructFields, pointertogongstructfield_)
						}
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
	fillUpTable[models.PointerToGongStructField](
		pointertogongstructfieldFormCallback.probe,
	)
	pointertogongstructfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pointertogongstructfieldFormCallback.CreationMode || pointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(pointertogongstructfieldFormCallback.probe)
}
func __gong__New__SliceOfPointerToGongStructFieldFormCallback(
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) OnSave() {

	log.Println("SliceOfPointerToGongStructFieldFormCallback, OnSave")

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
		case "GongStruct:SliceOfPointerToGongStructFields":
			// we need to retrieve the field owner before the change
			var pastGongStructOwner *models.GongStruct
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "SliceOfPointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest,
				sliceofpointertogongstructfieldFormCallback.probe.backRepoOfInterest,
				sliceofpointertogongstructfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructOwner = reverseFieldOwner.(*models.GongStruct)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructOwner != nil {
					idx := slices.Index(pastGongStructOwner.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
					pastGongStructOwner.SliceOfPointerToGongStructFields = slices.Delete(pastGongStructOwner.SliceOfPointerToGongStructFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstruct.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructOwner := _gongstruct // we have a match
						if pastGongStructOwner != nil {
							if newGongStructOwner != pastGongStructOwner {
								idx := slices.Index(pastGongStructOwner.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
								pastGongStructOwner.SliceOfPointerToGongStructFields = slices.Delete(pastGongStructOwner.SliceOfPointerToGongStructFields, idx, idx+1)
								newGongStructOwner.SliceOfPointerToGongStructFields = append(newGongStructOwner.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
							}
						} else {
							newGongStructOwner.SliceOfPointerToGongStructFields = append(newGongStructOwner.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
						}
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
	fillUpTable[models.SliceOfPointerToGongStructField](
		sliceofpointertogongstructfieldFormCallback.probe,
	)
	sliceofpointertogongstructfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sliceofpointertogongstructfieldFormCallback.CreationMode || sliceofpointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliceofpointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
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

	fillUpTree(sliceofpointertogongstructfieldFormCallback.probe)
}
