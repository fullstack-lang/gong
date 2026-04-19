// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ProcessFormCallback(
	process *models.Process,
	probe *Probe,
	formGroup *form.FormGroup,
) (processFormCallback *ProcessFormCallback) {
	processFormCallback = new(ProcessFormCallback)
	processFormCallback.probe = probe
	processFormCallback.process = process
	processFormCallback.formGroup = formGroup

	processFormCallback.CreationMode = (process == nil)

	return
}

type ProcessFormCallback struct {
	process *models.Process

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processFormCallback *ProcessFormCallback) OnSave() {
	processFormCallback.probe.stageOfInterest.Lock()
	defer processFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processFormCallback.probe.formStage.Checkout()

	if processFormCallback.process == nil {
		processFormCallback.process = new(models.Process).Stage(processFormCallback.probe.stageOfInterest)
	}
	process_ := processFormCallback.process
	_ = process_

	for _, formDiv := range processFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(process_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		process_.Unstage(processFormCallback.probe.stageOfInterest)
	}

	processFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Process](
		processFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processFormCallback.CreationMode || processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			processFormCallback.probe,
			newFormGroup,
		)
		process := new(models.Process)
		FillUpForm(process, newFormGroup, processFormCallback.probe)
		processFormCallback.probe.formStage.Commit()
	}

	processFormCallback.probe.ux_tree()
}
