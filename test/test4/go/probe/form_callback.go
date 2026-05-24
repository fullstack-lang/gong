// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test4/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AstructFormCallback(
	astruct *models.Astruct,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (astructFormCallback *AstructFormCallback) OnSave() {
	astructFormCallback.probe.stageOfInterest.Lock()
	defer astructFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AstructFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astruct_.Unstage(astructFormCallback.probe.stageOfInterest)
	}

	astructFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Astruct](
		astructFormCallback.probe,
	)

	// display a new form by reset the form stage
	if astructFormCallback.CreationMode || astructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		astructFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
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

	astructFormCallback.probe.ux_tree()
}
