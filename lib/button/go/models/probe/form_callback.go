// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
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
func __gong__New__ButtonFormCallback(
	_instance *models.Button,
	probe *Probe,
	formGroup *form.FormGroup,
) (buttonFormCallback *FormCallback[*models.Button]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveButtonFields,
	)
}

type ButtonFormCallback = FormCallback[*models.Button]

func saveButtonFields(
	_instance *models.Button,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(_instance.Icon), formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(_instance.IsDisabled), formDiv)
		case "Color":
			FormDivEnumStringFieldToField(&(_instance.Color), formDiv)
		case "MatButtonType":
			FormDivEnumStringFieldToField(&(_instance.MatButtonType), formDiv)
		case "MatButtonAppearance":
			FormDivEnumStringFieldToField(&(_instance.MatButtonAppearance), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.ToolTipPosition), formDiv)
		case "Group:Buttons":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Buttons", func(owner *models.Group) *[]*models.Button { return &owner.Buttons })
		}
	}
}

func __gong__New__ButtonToggleFormCallback(
	_instance *models.ButtonToggle,
	probe *Probe,
	formGroup *form.FormGroup,
) (buttontoggleFormCallback *FormCallback[*models.ButtonToggle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveButtonToggleFields,
	)
}

type ButtonToggleFormCallback = FormCallback[*models.ButtonToggle]

func saveButtonToggleFields(
	_instance *models.ButtonToggle,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(_instance.Icon), formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(_instance.IsDisabled), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "GroupToogle:ButtonToggles":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ButtonToggles", func(owner *models.GroupToogle) *[]*models.ButtonToggle { return &owner.ButtonToggles })
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
		case "Buttons":
			FormDivSliceOfPointersToField(_instance, "Buttons", &(_instance.Buttons), formDiv, probe)
		case "NbColumns":
			FormDivBasicFieldToField(&(_instance.NbColumns), formDiv)
		case "Layout:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Layout) *[]*models.Group { return &owner.Groups })
		}
	}
}

func __gong__New__GroupToogleFormCallback(
	_instance *models.GroupToogle,
	probe *Probe,
	formGroup *form.FormGroup,
) (grouptoogleFormCallback *FormCallback[*models.GroupToogle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupToogleFields,
	)
}

type GroupToogleFormCallback = FormCallback[*models.GroupToogle]

func saveGroupToogleFields(
	_instance *models.GroupToogle,
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
		case "ButtonToggles":
			FormDivSliceOfPointersToField(_instance, "ButtonToggles", &(_instance.ButtonToggles), formDiv, probe)
		case "IsSingleSelector":
			FormDivBasicFieldToField(&(_instance.IsSingleSelector), formDiv)
		case "Layout:GroupToogles":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GroupToogles", func(owner *models.Layout) *[]*models.GroupToogle { return &owner.GroupToogles })
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
		case "GroupToogles":
			FormDivSliceOfPointersToField(_instance, "GroupToogles", &(_instance.GroupToogles), formDiv, probe)
		}
	}
}

