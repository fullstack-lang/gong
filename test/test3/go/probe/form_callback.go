// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
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
		case "As":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.A](aFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.A, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.A)

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
			a_.As = instanceSlice

		case "A:As":
			// WARNING : this form deals with the N-N association "A.As []*A" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of A). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A"
				rf.Fieldname = "As"
				formerAssociationSource := models.GetReverseFieldOwner(
					aFormCallback.probe.stageOfInterest,
					a_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A)
					if !ok {
						log.Fatalln("Source of A.As []*A, is not an A instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.As, a_)
					formerSource.As = slices.Delete(formerSource.As, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.A
			for _a := range *models.GetGongstructInstancesSet[models.A](aFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _a.GetName() == newSourceName.GetName() {
					newSource = _a // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A.As []*A, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.As = append(newSource.As, a_)
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
