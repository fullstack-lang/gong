// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
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
func __gong__New__AFormCallback(
	_instance *models.A,
	probe *Probe,
	formGroup *form.FormGroup,
) (aFormCallback *FormCallback[*models.A]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAFields,
	)
}

type AFormCallback = FormCallback[*models.A]

func saveAFields(
	_instance *models.A,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NumberField":
			FormDivBasicFieldToField(&(_instance.NumberField), formDiv)
		case "B":
			FormDivSelectFieldToField(&(_instance.B), probe.stageOfInterest, formDiv)
		case "Bs":
			FormDivSliceOfPointersToField(_instance, "Bs", &(_instance.Bs), formDiv, probe)
		case "Foo":
			FormDivBasicFieldToField(&(_instance.Foo), formDiv)
		case "Bar":
			FormDivBasicFieldToField(&(_instance.Bar), formDiv)
		case "Zorgh":
			FormDivBasicFieldToField(&(_instance.Zorgh), formDiv)
		}
	}
}

func __gong__New__BFormCallback(
	_instance *models.B,
	probe *Probe,
	formGroup *form.FormGroup,
) (bFormCallback *FormCallback[*models.B]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBFields,
	)
}

type BFormCallback = FormCallback[*models.B]

func saveBFields(
	_instance *models.B,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "A:Bs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bs", func(owner *models.A) *[]*models.B { return &owner.Bs })
		}
	}
}

