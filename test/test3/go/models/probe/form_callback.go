// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
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
		case "Date":
			FormDivTimeFieldToField(&(_instance.Date), formDiv, true)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "FloatValue":
			FormDivBasicFieldToField(&(_instance.FloatValue), formDiv)
		case "IntValue":
			FormDivBasicFieldToField(&(_instance.IntValue), formDiv)
		case "EnumString":
			FormDivEnumStringFieldToField(&(_instance.EnumString), formDiv)
		case "EnumInt":
			FormDivEnumIntFieldToField(&(_instance.EnumInt), formDiv)
		case "B":
			FormDivSelectFieldToField(&(_instance.B), probe.stageOfInterest, formDiv)
		case "Bs":
			FormDivSliceOfPointersToField(_instance, "Bs", &(_instance.Bs), formDiv, probe)
		case "C":
			FormDivSelectFieldToField(&(_instance.C), probe.stageOfInterest, formDiv)
		case "Cs":
			FormDivSliceOfPointersToField(_instance, "Cs", &(_instance.Cs), formDiv, probe)
		case "UUID":
			FormDivBasicFieldToField(&(_instance.UUID), formDiv)
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

func __gong__New__CFormCallback(
	_instance *models.C,
	probe *Probe,
	formGroup *form.FormGroup,
) (cFormCallback *FormCallback[*models.C]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCFields,
	)
}

type CFormCallback = FormCallback[*models.C]

func saveCFields(
	_instance *models.C,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "A:Cs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Cs", func(owner *models.A) *[]*models.C { return &owner.Cs })
		}
	}
}

