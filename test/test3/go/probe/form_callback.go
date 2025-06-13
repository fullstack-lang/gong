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
