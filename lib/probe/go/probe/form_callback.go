// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/probe/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

// insertion point
func __gong__New__Probe2FormCallback(
	probe2 *models.Probe2,
	probe *Probe,
	formGroup *table.FormGroup,
) (probe2FormCallback *Probe2FormCallback) {
	probe2FormCallback = new(Probe2FormCallback)
	probe2FormCallback.probe = probe
	probe2FormCallback.probe2 = probe2
	probe2FormCallback.formGroup = formGroup

	probe2FormCallback.CreationMode = (probe2 == nil)

	return
}

type Probe2FormCallback struct {
	probe2 *models.Probe2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (probe2FormCallback *Probe2FormCallback) OnSave() {

	// log.Println("Probe2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	probe2FormCallback.probe.formStage.Checkout()

	if probe2FormCallback.probe2 == nil {
		probe2FormCallback.probe2 = new(models.Probe2).Stage(probe2FormCallback.probe.stageOfInterest)
	}
	probe2_ := probe2FormCallback.probe2
	_ = probe2_

	for _, formDiv := range probe2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(probe2_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if probe2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		probe2_.Unstage(probe2FormCallback.probe.stageOfInterest)
	}

	probe2FormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Probe2](
		probe2FormCallback.probe,
	)
	probe2FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if probe2FormCallback.CreationMode || probe2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		probe2FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(probe2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Probe2FormCallback(
			nil,
			probe2FormCallback.probe,
			newFormGroup,
		)
		probe2 := new(models.Probe2)
		FillUpForm(probe2, newFormGroup, probe2FormCallback.probe)
		probe2FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(probe2FormCallback.probe)
}
