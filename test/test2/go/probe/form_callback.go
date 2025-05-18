// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AFormCallback(
	a *models.A,
	probe *Probe,
	formGroup *table.FormGroup,
) (aFormCallback *AFormCallback) {
	aFormCallback = new(AFormCallback)
	aFormCallback.probe = probe
	aFormCallback.a = a
	aFormCallback.formGroup = formGroup

	aFormCallback.CreationMode = (a == nil)

	return
}

type AFormCallback struct {
	a *models.A

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (aFormCallback *AFormCallback) OnSave() {

	// log.Println("AFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	aFormCallback.probe.formStage.Checkout()

	if aFormCallback.a == nil {
		aFormCallback.a = new(models.A).Stage(aFormCallback.probe.stageOfInterest)
	}
	a_ := aFormCallback.a
	_ = a_

	for _, formDiv := range aFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_.Name), formDiv)
		case "NumberField":
			FormDivBasicFieldToField(&(a_.NumberField), formDiv)
		case "B":
			FormDivSelectFieldToField(&(a_.B), aFormCallback.probe.stageOfInterest, formDiv)
		case "Bs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.B](aFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.B, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.B)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					aFormCallback.probe.stageOfInterest,
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
			a_.Bs = instanceSlice

		}
	}

	// manage the suppress operation
	if aFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_.Unstage(aFormCallback.probe.stageOfInterest)
	}

	aFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.A](
		aFormCallback.probe,
	)
	aFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if aFormCallback.CreationMode || aFormCallback.formGroup.HasSuppressButtonBeenPressed {
		aFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(aFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AFormCallback(
			nil,
			aFormCallback.probe,
			newFormGroup,
		)
		a := new(models.A)
		FillUpForm(a, newFormGroup, aFormCallback.probe)
		aFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(aFormCallback.probe)
}
func __gong__New__BFormCallback(
	b *models.B,
	probe *Probe,
	formGroup *table.FormGroup,
) (bFormCallback *BFormCallback) {
	bFormCallback = new(BFormCallback)
	bFormCallback.probe = probe
	bFormCallback.b = b
	bFormCallback.formGroup = formGroup

	bFormCallback.CreationMode = (b == nil)

	return
}

type BFormCallback struct {
	b *models.B

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bFormCallback *BFormCallback) OnSave() {

	// log.Println("BFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bFormCallback.probe.formStage.Checkout()

	if bFormCallback.b == nil {
		bFormCallback.b = new(models.B).Stage(bFormCallback.probe.stageOfInterest)
	}
	b_ := bFormCallback.b
	_ = b_

	for _, formDiv := range bFormCallback.formGroup.FormDivs {
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
			reverseFieldOwner := models.GetReverseFieldOwner(
				bFormCallback.probe.stageOfInterest,
				b_,
				&rf)

			if reverseFieldOwner != nil {
				pastAOwner = reverseFieldOwner.(*models.A)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastAOwner != nil {
					idx := slices.Index(pastAOwner.Bs, b_)
					pastAOwner.Bs = slices.Delete(pastAOwner.Bs, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastAOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _a := range *models.GetGongstructInstancesSet[models.A](bFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _a.GetName() == fieldValue.GetName() {
							newAOwner := _a // we have a match

							// we remove the b_ instance from the pastAOwner field
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
	}

	// manage the suppress operation
	if bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		b_.Unstage(bFormCallback.probe.stageOfInterest)
	}

	bFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.B](
		bFormCallback.probe,
	)
	bFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bFormCallback.CreationMode || bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(bFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BFormCallback(
			nil,
			bFormCallback.probe,
			newFormGroup,
		)
		b := new(models.B)
		FillUpForm(b, newFormGroup, bFormCallback.probe)
		bFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(bFormCallback.probe)
}
