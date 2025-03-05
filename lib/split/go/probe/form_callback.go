// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__SplitAreaFormCallback(
	splitarea *models.SplitArea,
	probe *Probe,
	formGroup *table.FormGroup,
) (splitareaFormCallback *SplitAreaFormCallback) {
	splitareaFormCallback = new(SplitAreaFormCallback)
	splitareaFormCallback.probe = probe
	splitareaFormCallback.splitarea = splitarea
	splitareaFormCallback.formGroup = formGroup

	splitareaFormCallback.CreationMode = (splitarea == nil)

	return
}

type SplitAreaFormCallback struct {
	splitarea *models.SplitArea

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (splitareaFormCallback *SplitAreaFormCallback) OnSave() {

	log.Println("SplitAreaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	splitareaFormCallback.probe.formStage.Checkout()

	if splitareaFormCallback.splitarea == nil {
		splitareaFormCallback.splitarea = new(models.SplitArea).Stage(splitareaFormCallback.probe.stageOfInterest)
	}
	splitarea_ := splitareaFormCallback.splitarea
	_ = splitarea_

	for _, formDiv := range splitareaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(splitarea_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if splitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		splitarea_.Unstage(splitareaFormCallback.probe.stageOfInterest)
	}

	splitareaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SplitArea](
		splitareaFormCallback.probe,
	)
	splitareaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if splitareaFormCallback.CreationMode || splitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		splitareaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(splitareaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SplitAreaFormCallback(
			nil,
			splitareaFormCallback.probe,
			newFormGroup,
		)
		splitarea := new(models.SplitArea)
		FillUpForm(splitarea, newFormGroup, splitareaFormCallback.probe)
		splitareaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(splitareaFormCallback.probe)
}
