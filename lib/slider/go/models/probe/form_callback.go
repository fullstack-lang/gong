// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
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
func __gong__New__CheckboxFormCallback(
	_instance *models.Checkbox,
	probe *Probe,
	formGroup *form.FormGroup,
) (checkboxFormCallback *FormCallback[*models.Checkbox]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCheckboxFields,
	)
}

type CheckboxFormCallback = FormCallback[*models.Checkbox]

func saveCheckboxFields(
	_instance *models.Checkbox,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ValueBool":
			FormDivBasicFieldToField(&(_instance.ValueBool), formDiv)
		case "LabelForTrue":
			FormDivBasicFieldToField(&(_instance.LabelForTrue), formDiv)
		case "LabelForFalse":
			FormDivBasicFieldToField(&(_instance.LabelForFalse), formDiv)
		case "Group:Checkboxes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Checkboxes", func(owner *models.Group) *[]*models.Checkbox { return &owner.Checkboxes })
		}
	}
}

func __gong__New__GroupFormCallback(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupFormCallback *FormCallback[*models.Group]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupFields,
	)
}

type GroupFormCallback = FormCallback[*models.Group]

func saveGroupFields(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Percentage":
			FormDivBasicFieldToField(&(_instance.Percentage), formDiv)
		case "Sliders":
			FormDivSliceOfPointersToField(_instance, "Sliders", &(_instance.Sliders), formDiv, probe)
		case "Checkboxes":
			FormDivSliceOfPointersToField(_instance, "Checkboxes", &(_instance.Checkboxes), formDiv, probe)
		case "Layout:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Layout) *[]*models.Group { return &owner.Groups })
		}
	}
}

func __gong__New__LayoutFormCallback(
	_instance *models.Layout,
	probe *Probe,
	formGroup *form.FormGroup,
) (layoutFormCallback *FormCallback[*models.Layout]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLayoutFields,
	)
}

type LayoutFormCallback = FormCallback[*models.Layout]

func saveLayoutFields(
	_instance *models.Layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "IsWithCustomGutterSize":
			FormDivBasicFieldToField(&(_instance.IsWithCustomGutterSize), formDiv)
		case "GutterSize":
			FormDivBasicFieldToField(&(_instance.GutterSize), formDiv)
		}
	}
}

func __gong__New__SliderFormCallback(
	_instance *models.Slider,
	probe *Probe,
	formGroup *form.FormGroup,
) (sliderFormCallback *FormCallback[*models.Slider]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSliderFields,
	)
}

type SliderFormCallback = FormCallback[*models.Slider]

func saveSliderFields(
	_instance *models.Slider,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsFloat64":
			FormDivBasicFieldToField(&(_instance.IsFloat64), formDiv)
		case "IsInt":
			FormDivBasicFieldToField(&(_instance.IsInt), formDiv)
		case "MinInt":
			FormDivBasicFieldToField(&(_instance.MinInt), formDiv)
		case "MaxInt":
			FormDivBasicFieldToField(&(_instance.MaxInt), formDiv)
		case "StepInt":
			FormDivBasicFieldToField(&(_instance.StepInt), formDiv)
		case "ValueInt":
			FormDivBasicFieldToField(&(_instance.ValueInt), formDiv)
		case "MinFloat64":
			FormDivBasicFieldToField(&(_instance.MinFloat64), formDiv)
		case "MaxFloat64":
			FormDivBasicFieldToField(&(_instance.MaxFloat64), formDiv)
		case "StepFloat64":
			FormDivBasicFieldToField(&(_instance.StepFloat64), formDiv)
		case "ValueFloat64":
			FormDivBasicFieldToField(&(_instance.ValueFloat64), formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(_instance.IsDisabled), formDiv)
		case "Group:Sliders":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sliders", func(owner *models.Group) *[]*models.Slider { return &owner.Sliders })
		}
	}
}

