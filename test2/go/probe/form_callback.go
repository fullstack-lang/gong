// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AFormCallback(
	a *models.A,
	probe *Probe,
) (aFormCallback *AFormCallback) {
	aFormCallback = new(AFormCallback)
	aFormCallback.probe = probe
	aFormCallback.a = a

	aFormCallback.CreationMode = (a == nil)

	return
}

type AFormCallback struct {
	a *models.A

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (aFormCallback *AFormCallback) OnSave() {

	log.Println("AFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	aFormCallback.probe.formStage.Checkout()

	if aFormCallback.a == nil {
		aFormCallback.a = new(models.A).Stage(aFormCallback.probe.stageOfInterest)
	}
	a_ := aFormCallback.a
	_ = a_

	// get the formGroup
	formGroup := aFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_.Name), formDiv)
		case "NumberField":
			FormDivBasicFieldToField(&(a_.NumberField), formDiv)
		case "B":
			FormDivSelectFieldToField(&(a_.B), aFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	aFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A](
		aFormCallback.probe,
	)
	aFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if aFormCallback.CreationMode {
		aFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AFormCallback(
				nil,
				aFormCallback.probe,
			),
		}).Stage(aFormCallback.probe.formStage)
		a := new(models.A)
		FillUpForm(a, newFormGroup, aFormCallback.probe)
		aFormCallback.probe.formStage.Commit()
	}

	fillUpTree(aFormCallback.probe)
}
func __gong__New__BFormCallback(
	b *models.B,
	probe *Probe,
) (bFormCallback *BFormCallback) {
	bFormCallback = new(BFormCallback)
	bFormCallback.probe = probe
	bFormCallback.b = b

	bFormCallback.CreationMode = (b == nil)

	return
}

type BFormCallback struct {
	b *models.B

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (bFormCallback *BFormCallback) OnSave() {

	log.Println("BFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bFormCallback.probe.formStage.Checkout()

	if bFormCallback.b == nil {
		bFormCallback.b = new(models.B).Stage(bFormCallback.probe.stageOfInterest)
	}
	b_ := bFormCallback.b
	_ = b_

	// get the formGroup
	formGroup := bFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(b_.Name), formDiv)
		case "A:Bs":
			// we need to retrieve the field owner before the change
			var pastAOwner *models.A
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A"
			rf.Fieldname = "Bs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bFormCallback.probe.stageOfInterest,
				bFormCallback.probe.backRepoOfInterest,
				b_,
				&rf)

			if reverseFieldOwner != nil {
				pastAOwner = reverseFieldOwner.(*models.A)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAOwner != nil {
					idx := slices.Index(pastAOwner.Bs, b_)
					pastAOwner.Bs = slices.Delete(pastAOwner.Bs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a := range *models.GetGongstructInstancesSet[models.A](bFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAOwner := _a // we have a match
						if pastAOwner != nil {
							if newAOwner != pastAOwner {
								idx := slices.Index(pastAOwner.Bs, b_)
								pastAOwner.Bs = slices.Delete(pastAOwner.Bs, idx, idx+1)
								newAOwner.Bs = append(newAOwner.Bs, b_)
							}
						} else {
							newAOwner.Bs = append(newAOwner.Bs, b_)
						}
					}
				}
			}
		}
	}

	bFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.B](
		bFormCallback.probe,
	)
	bFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bFormCallback.CreationMode {
		bFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BFormCallback(
				nil,
				bFormCallback.probe,
			),
		}).Stage(bFormCallback.probe.formStage)
		b := new(models.B)
		FillUpForm(b, newFormGroup, bFormCallback.probe)
		bFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bFormCallback.probe)
}
