// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
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
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
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

	var zero T
	if cb.Instance == zero {
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
func __gong__New__FreqencyFormCallback(
	_instance *models.Freqency,
	probe *Probe,
	formGroup *form.FormGroup,
) (freqencyFormCallback *FormCallback[*models.Freqency]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFreqencyFields,
	)
}

type FreqencyFormCallback = FormCallback[*models.Freqency]

func saveFreqencyFields(
	_instance *models.Freqency,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note:Frequencies":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Frequencies", func(owner *models.Note) *[]*models.Freqency { return &owner.Frequencies })
		}
	}
}

func __gong__New__NoteFormCallback(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *FormCallback[*models.Note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteFields,
	)
}

type NoteFormCallback = FormCallback[*models.Note]

func saveNoteFields(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Frequencies":
			FormDivSliceOfPointersToField(_instance, "Frequencies", &(_instance.Frequencies), formDiv, probe)
		case "Start":
			FormDivBasicFieldToField(&(_instance.Start), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Velocity":
			FormDivBasicFieldToField(&(_instance.Velocity), formDiv)
		case "Info":
			FormDivBasicFieldToField(&(_instance.Info), formDiv)
		}
	}
}

func __gong__New__PlayerFormCallback(
	_instance *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) (playerFormCallback *FormCallback[*models.Player]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlayerFields,
	)
}

type PlayerFormCallback = FormCallback[*models.Player]

func savePlayerFields(
	_instance *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Status":
			FormDivEnumStringFieldToField(&(_instance.Status), formDiv)
		}
	}
}

