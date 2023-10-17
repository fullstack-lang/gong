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
	playground *Playground,
) (aFormCallback *AFormCallback) {
	aFormCallback = new(AFormCallback)
	aFormCallback.playground = playground
	aFormCallback.a = a

	aFormCallback.CreationMode = (a == nil)

	return
}

type AFormCallback struct {
	a *models.A

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (aFormCallback *AFormCallback) OnSave() {

	log.Println("AFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	aFormCallback.playground.formStage.Checkout()

	if aFormCallback.a == nil {
		aFormCallback.a = new(models.A).Stage(aFormCallback.playground.stageOfInterest)
	}
	a_ := aFormCallback.a
	_ = a_

	// get the formGroup
	formGroup := aFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_.Name), formDiv)
		case "B":
			FormDivSelectFieldToField(&(a_.B), aFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	aFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.A](
		aFormCallback.playground,
	)
	aFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if aFormCallback.CreationMode {
		aFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AFormCallback(
				nil,
				aFormCallback.playground,
			),
		}).Stage(aFormCallback.playground.formStage)
		a := new(models.A)
		FillUpForm(a, newFormGroup, aFormCallback.playground)
		aFormCallback.playground.formStage.Commit()
	}

	fillUpTree(aFormCallback.playground)
}
func __gong__New__BFormCallback(
	b *models.B,
	playground *Playground,
) (bFormCallback *BFormCallback) {
	bFormCallback = new(BFormCallback)
	bFormCallback.playground = playground
	bFormCallback.b = b

	bFormCallback.CreationMode = (b == nil)

	return
}

type BFormCallback struct {
	b *models.B

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (bFormCallback *BFormCallback) OnSave() {

	log.Println("BFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bFormCallback.playground.formStage.Checkout()

	if bFormCallback.b == nil {
		bFormCallback.b = new(models.B).Stage(bFormCallback.playground.stageOfInterest)
	}
	b_ := bFormCallback.b
	_ = b_

	// get the formGroup
	formGroup := bFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				bFormCallback.playground.stageOfInterest,
				bFormCallback.playground.backRepoOfInterest,
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
				for _a := range *models.GetGongstructInstancesSet[models.A](bFormCallback.playground.stageOfInterest) {

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

	bFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.B](
		bFormCallback.playground,
	)
	bFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if bFormCallback.CreationMode {
		bFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BFormCallback(
				nil,
				bFormCallback.playground,
			),
		}).Stage(bFormCallback.playground.formStage)
		b := new(models.B)
		FillUpForm(b, newFormGroup, bFormCallback.playground)
		bFormCallback.playground.formStage.Commit()
	}

	fillUpTree(bFormCallback.playground)
}
