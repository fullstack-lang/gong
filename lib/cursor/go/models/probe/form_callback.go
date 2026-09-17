// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/cursor/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: any(instance) == nil,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	if any(cb.Instance) == nil {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__CursorFormCallback(
	_instance *models.Cursor,
	probe *Probe,
	formGroup *form.FormGroup,
) (cursorFormCallback *FormCallback[*models.Cursor]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCursorFields,
	)
}

type CursorFormCallback = FormCallback[*models.Cursor]

func saveCursorFields(
	_instance *models.Cursor,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(_instance.Y1), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(_instance.Y2), formDiv)
		case "DurationSeconds":
			FormDivBasicFieldToField(&(_instance.DurationSeconds), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "IsPlaying":
			FormDivBasicFieldToField(&(_instance.IsPlaying), formDiv)
		}
	}
}

