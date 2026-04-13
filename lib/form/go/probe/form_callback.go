// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__Form2FormCallback(
	form2 *models.Form2,
	probe *Probe,
	formGroup *form.FormGroup,
) (form2FormCallback *Form2FormCallback) {
	form2FormCallback = new(Form2FormCallback)
	form2FormCallback.probe = probe
	form2FormCallback.form2 = form2
	form2FormCallback.formGroup = formGroup

	form2FormCallback.CreationMode = (form2 == nil)

	return
}

type Form2FormCallback struct {
	form2 *models.Form2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (form2FormCallback *Form2FormCallback) OnSave() {
	form2FormCallback.probe.stageOfInterest.Lock()
	defer form2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("Form2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	form2FormCallback.probe.formStage.Checkout()

	if form2FormCallback.form2 == nil {
		form2FormCallback.form2 = new(models.Form2).Stage(form2FormCallback.probe.stageOfInterest)
	}
	form2_ := form2FormCallback.form2
	_ = form2_

	for _, formDiv := range form2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(form2_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if form2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		form2_.Unstage(form2FormCallback.probe.stageOfInterest)
	}

	form2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Form2](
		form2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if form2FormCallback.CreationMode || form2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		form2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(form2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Form2FormCallback(
			nil,
			form2FormCallback.probe,
			newFormGroup,
		)
		form2 := new(models.Form2)
		FillUpForm(form2, newFormGroup, form2FormCallback.probe)
		form2FormCallback.probe.formStage.Commit()
	}

	form2FormCallback.probe.ux_tree()
}
