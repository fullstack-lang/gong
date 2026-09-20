// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
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
func __gong__New__CommandFormCallback(
	_instance *models.Command,
	probe *Probe,
	formGroup *form.FormGroup,
) (commandFormCallback *FormCallback[*models.Command]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCommandFields,
	)
}

type CommandFormCallback = FormCallback[*models.Command]

func saveCommandFields(
	_instance *models.Command,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Command":
			FormDivEnumStringFieldToField(&(_instance.Command), formDiv)
		case "CommandDate":
			FormDivBasicFieldToField(&(_instance.CommandDate), formDiv)
		case "Engine":
			FormDivSelectFieldToField(&(_instance.Engine), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__DummyAgentFormCallback(
	_instance *models.DummyAgent,
	probe *Probe,
	formGroup *form.FormGroup,
) (dummyagentFormCallback *FormCallback[*models.DummyAgent]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDummyAgentFields,
	)
}

type DummyAgentFormCallback = FormCallback[*models.DummyAgent]

func saveDummyAgentFields(
	_instance *models.DummyAgent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "TechName":
			FormDivBasicFieldToField(&(_instance.TechName), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EngineFormCallback(
	_instance *models.Engine,
	probe *Probe,
	formGroup *form.FormGroup,
) (engineFormCallback *FormCallback[*models.Engine]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEngineFields,
	)
}

type EngineFormCallback = FormCallback[*models.Engine]

func saveEngineFields(
	_instance *models.Engine,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EndTime":
			FormDivBasicFieldToField(&(_instance.EndTime), formDiv)
		case "CurrentTime":
			FormDivBasicFieldToField(&(_instance.CurrentTime), formDiv)
		case "DisplayFormat":
			FormDivBasicFieldToField(&(_instance.DisplayFormat), formDiv)
		case "SecondsSinceStart":
			FormDivBasicFieldToField(&(_instance.SecondsSinceStart), formDiv)
		case "Fired":
			FormDivBasicFieldToField(&(_instance.Fired), formDiv)
		case "ControlMode":
			FormDivEnumStringFieldToField(&(_instance.ControlMode), formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(_instance.State), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(_instance.Speed), formDiv)
		}
	}
}

func __gong__New__EventFormCallback(
	_instance *models.Event,
	probe *Probe,
	formGroup *form.FormGroup,
) (eventFormCallback *FormCallback[*models.Event]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEventFields,
	)
}

type EventFormCallback = FormCallback[*models.Event]

func saveEventFields(
	_instance *models.Event,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		}
	}
}

func __gong__New__StatusFormCallback(
	_instance *models.Status,
	probe *Probe,
	formGroup *form.FormGroup,
) (statusFormCallback *FormCallback[*models.Status]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStatusFields,
	)
}

type StatusFormCallback = FormCallback[*models.Status]

func saveStatusFields(
	_instance *models.Status,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "CurrentCommand":
			FormDivEnumStringFieldToField(&(_instance.CurrentCommand), formDiv)
		case "CompletionDate":
			FormDivBasicFieldToField(&(_instance.CompletionDate), formDiv)
		case "CurrentSpeedCommand":
			FormDivEnumStringFieldToField(&(_instance.CurrentSpeedCommand), formDiv)
		case "SpeedCommandCompletionDate":
			FormDivBasicFieldToField(&(_instance.SpeedCommandCompletionDate), formDiv)
		}
	}
}

func __gong__New__UpdateStateFormCallback(
	_instance *models.UpdateState,
	probe *Probe,
	formGroup *form.FormGroup,
) (updatestateFormCallback *FormCallback[*models.UpdateState]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUpdateStateFields,
	)
}

type UpdateStateFormCallback = FormCallback[*models.UpdateState]

func saveUpdateStateFields(
	_instance *models.UpdateState,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Period":
			FormDivBasicFieldToField(&(_instance.Period), formDiv)
		}
	}
}

