// generated code - do not edit
package probe

import (
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/cursor/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__CursorFormCallback(
	cursor *models.Cursor,
	probe *Probe,
	formGroup *table.FormGroup,
) (cursorFormCallback *CursorFormCallback) {
	cursorFormCallback = new(CursorFormCallback)
	cursorFormCallback.probe = probe
	cursorFormCallback.cursor = cursor
	cursorFormCallback.formGroup = formGroup

	cursorFormCallback.CreationMode = (cursor == nil)

	return
}

type CursorFormCallback struct {
	cursor *models.Cursor

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cursorFormCallback *CursorFormCallback) OnSave() {

	// log.Println("CursorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cursorFormCallback.probe.formStage.Checkout()

	if cursorFormCallback.cursor == nil {
		cursorFormCallback.cursor = new(models.Cursor).Stage(cursorFormCallback.probe.stageOfInterest)
	}
	cursor_ := cursorFormCallback.cursor
	_ = cursor_

	for _, formDiv := range cursorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cursor_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(cursor_.StartX), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(cursor_.EndX), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(cursor_.Y1), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(cursor_.Y2), formDiv)
		case "DurationSeconds":
			FormDivBasicFieldToField(&(cursor_.DurationSeconds), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(cursor_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(cursor_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(cursor_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(cursor_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(cursor_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(cursor_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(cursor_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(cursor_.Transform), formDiv)
		case "IsPlaying":
			FormDivBasicFieldToField(&(cursor_.IsPlaying), formDiv)
		}
	}

	// manage the suppress operation
	if cursorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cursor_.Unstage(cursorFormCallback.probe.stageOfInterest)
	}

	cursorFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Cursor](
		cursorFormCallback.probe,
	)
	cursorFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cursorFormCallback.CreationMode || cursorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cursorFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cursorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CursorFormCallback(
			nil,
			cursorFormCallback.probe,
			newFormGroup,
		)
		cursor := new(models.Cursor)
		FillUpForm(cursor, newFormGroup, cursorFormCallback.probe)
		cursorFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cursorFormCallback.probe)
}
