// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
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
func __gong__New__CheckBoxFormCallback(
	_instance *models.CheckBox,
	probe *Probe,
	formGroup *form.FormGroup,
) (checkboxFormCallback *FormCallback[*models.CheckBox]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCheckBoxFields,
	)
}

type CheckBoxFormCallback = FormCallback[*models.CheckBox]

func saveCheckBoxFields(
	_instance *models.CheckBox,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "FormDiv:CheckBoxs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "CheckBoxs", func(owner *models.FormDiv) *[]*models.CheckBox { return &owner.CheckBoxs })
		}
	}
}

func __gong__New__FormDivFormCallback(
	_instance *models.FormDiv,
	probe *Probe,
	formGroup *form.FormGroup,
) (formdivFormCallback *FormCallback[*models.FormDiv]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormDivFields,
	)
}

type FormDivFormCallback = FormCallback[*models.FormDiv]

func saveFormDivFields(
	_instance *models.FormDiv,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "FormFields":
			FormDivSliceOfPointersToField(_instance, "FormFields", &(_instance.FormFields), formDiv, probe)
		case "CheckBoxs":
			FormDivSliceOfPointersToField(_instance, "CheckBoxs", &(_instance.CheckBoxs), formDiv, probe)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(_instance.FormEditAssocButton), probe.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(_instance.FormSortAssocButton), probe.stageOfInterest, formDiv)
		case "IsADivider":
			FormDivBasicFieldToField(&(_instance.IsADivider), formDiv)
		case "IsAStartAccordionGroup":
			FormDivBasicFieldToField(&(_instance.IsAStartAccordionGroup), formDiv)
		case "AccordionGroupName":
			FormDivBasicFieldToField(&(_instance.AccordionGroupName), formDiv)
		case "IsAEndAccordionGroup":
			FormDivBasicFieldToField(&(_instance.IsAEndAccordionGroup), formDiv)
		case "FormGroup:FormDivs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "FormDivs", func(owner *models.FormGroup) *[]*models.FormDiv { return &owner.FormDivs })
		}
	}
}

func __gong__New__FormEditAssocButtonFormCallback(
	_instance *models.FormEditAssocButton,
	probe *Probe,
	formGroup *form.FormGroup,
) (formeditassocbuttonFormCallback *FormCallback[*models.FormEditAssocButton]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormEditAssocButtonFields,
	)
}

type FormEditAssocButtonFormCallback = FormCallback[*models.FormEditAssocButton]

func saveFormEditAssocButtonFields(
	_instance *models.FormEditAssocButton,
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
		case "AssociationStorage":
			FormDivBasicFieldToField(&(_instance.AssociationStorage), formDiv)
		case "HasChanged":
			FormDivBasicFieldToField(&(_instance.HasChanged), formDiv)
		case "IsForSavePurpose":
			FormDivBasicFieldToField(&(_instance.IsForSavePurpose), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "MatTooltipShowDelay":
			FormDivBasicFieldToField(&(_instance.MatTooltipShowDelay), formDiv)
		}
	}
}

func __gong__New__FormFieldFormCallback(
	_instance *models.FormField,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldFormCallback *FormCallback[*models.FormField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldFields,
	)
}

type FormFieldFormCallback = FormCallback[*models.FormField]

func saveFormFieldFields(
	_instance *models.FormField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "InputTypeEnum":
			FormDivEnumStringFieldToField(&(_instance.InputTypeEnum), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "Placeholder":
			FormDivBasicFieldToField(&(_instance.Placeholder), formDiv)
		case "FormFieldString":
			FormDivSelectFieldToField(&(_instance.FormFieldString), probe.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(_instance.FormFieldFloat64), probe.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(_instance.FormFieldInt), probe.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(_instance.FormFieldDate), probe.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(_instance.FormFieldTime), probe.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(_instance.FormFieldDateTime), probe.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(_instance.FormFieldSelect), probe.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(_instance.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(_instance.BespokeWidthPx), formDiv)
		case "HasBespokeHeight":
			FormDivBasicFieldToField(&(_instance.HasBespokeHeight), formDiv)
		case "BespokeHeightPx":
			FormDivBasicFieldToField(&(_instance.BespokeHeightPx), formDiv)
		case "FormDiv:FormFields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "FormFields", func(owner *models.FormDiv) *[]*models.FormField { return &owner.FormFields })
		}
	}
}

func __gong__New__FormFieldDateFormCallback(
	_instance *models.FormFieldDate,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfielddateFormCallback *FormCallback[*models.FormFieldDate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldDateFields,
	)
}

type FormFieldDateFormCallback = FormCallback[*models.FormFieldDate]

func saveFormFieldDateFields(
	_instance *models.FormFieldDate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivTimeFieldToField(&(_instance.Value), formDiv, false)
		}
	}
}

func __gong__New__FormFieldDateTimeFormCallback(
	_instance *models.FormFieldDateTime,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfielddatetimeFormCallback *FormCallback[*models.FormFieldDateTime]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldDateTimeFields,
	)
}

type FormFieldDateTimeFormCallback = FormCallback[*models.FormFieldDateTime]

func saveFormFieldDateTimeFields(
	_instance *models.FormFieldDateTime,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivTimeFieldToField(&(_instance.Value), formDiv, false)
		}
	}
}

func __gong__New__FormFieldFloat64FormCallback(
	_instance *models.FormFieldFloat64,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldfloat64FormCallback *FormCallback[*models.FormFieldFloat64]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldFloat64Fields,
	)
}

type FormFieldFloat64FormCallback = FormCallback[*models.FormFieldFloat64]

func saveFormFieldFloat64Fields(
	_instance *models.FormFieldFloat64,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(_instance.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(_instance.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(_instance.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(_instance.MaxValue), formDiv)
		}
	}
}

func __gong__New__FormFieldIntFormCallback(
	_instance *models.FormFieldInt,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldintFormCallback *FormCallback[*models.FormFieldInt]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldIntFields,
	)
}

type FormFieldIntFormCallback = FormCallback[*models.FormFieldInt]

func saveFormFieldIntFields(
	_instance *models.FormFieldInt,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(_instance.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(_instance.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(_instance.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(_instance.MaxValue), formDiv)
		}
	}
}

func __gong__New__FormFieldSelectFormCallback(
	_instance *models.FormFieldSelect,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldselectFormCallback *FormCallback[*models.FormFieldSelect]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldSelectFields,
	)
}

type FormFieldSelectFormCallback = FormCallback[*models.FormFieldSelect]

func saveFormFieldSelectFields(
	_instance *models.FormFieldSelect,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(_instance.Value), probe.stageOfInterest, formDiv)
		case "Options":
			FormDivSliceOfPointersToField(_instance, "Options", &(_instance.Options), formDiv, probe)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(_instance.CanBeEmpty), formDiv)
		case "PreserveInitialOrder":
			FormDivBasicFieldToField(&(_instance.PreserveInitialOrder), formDiv)
		}
	}
}

func __gong__New__FormFieldStringFormCallback(
	_instance *models.FormFieldString,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldstringFormCallback *FormCallback[*models.FormFieldString]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldStringFields,
	)
}

type FormFieldStringFormCallback = FormCallback[*models.FormFieldString]

func saveFormFieldStringFields(
	_instance *models.FormFieldString,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(_instance.IsTextArea), formDiv)
		}
	}
}

func __gong__New__FormFieldTimeFormCallback(
	_instance *models.FormFieldTime,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldtimeFormCallback *FormCallback[*models.FormFieldTime]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFieldTimeFields,
	)
}

type FormFieldTimeFormCallback = FormCallback[*models.FormFieldTime]

func saveFormFieldTimeFields(
	_instance *models.FormFieldTime,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivTimeFieldToField(&(_instance.Value), formDiv, false)
		case "Step":
			FormDivBasicFieldToField(&(_instance.Step), formDiv)
		}
	}
}

func __gong__New__FormGroupFormCallback(
	_instance *models.FormGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) (formgroupFormCallback *FormCallback[*models.FormGroup]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormGroupFields,
	)
}

type FormGroupFormCallback = FormCallback[*models.FormGroup]

func saveFormGroupFields(
	_instance *models.FormGroup,
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
		case "TypeLabel":
			FormDivBasicFieldToField(&(_instance.TypeLabel), formDiv)
		case "FormDivs":
			FormDivSliceOfPointersToField(_instance, "FormDivs", &(_instance.FormDivs), formDiv, probe)
		case "HasSuppressButton":
			FormDivBasicFieldToField(&(_instance.HasSuppressButton), formDiv)
		case "HasSuppressButtonBeenPressed":
			FormDivBasicFieldToField(&(_instance.HasSuppressButtonBeenPressed), formDiv)
		}
	}
}

func __gong__New__FormSortAssocButtonFormCallback(
	_instance *models.FormSortAssocButton,
	probe *Probe,
	formGroup *form.FormGroup,
) (formsortassocbuttonFormCallback *FormCallback[*models.FormSortAssocButton]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormSortAssocButtonFields,
	)
}

type FormSortAssocButtonFormCallback = FormCallback[*models.FormSortAssocButton]

func saveFormSortAssocButtonFields(
	_instance *models.FormSortAssocButton,
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
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "MatTooltipShowDelay":
			FormDivBasicFieldToField(&(_instance.MatTooltipShowDelay), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(_instance.FormEditAssocButton), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__OptionFormCallback(
	_instance *models.Option,
	probe *Probe,
	formGroup *form.FormGroup,
) (optionFormCallback *FormCallback[*models.Option]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOptionFields,
	)
}

type OptionFormCallback = FormCallback[*models.Option]

func saveOptionFields(
	_instance *models.Option,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "FormFieldSelect:Options":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Options", func(owner *models.FormFieldSelect) *[]*models.Option { return &owner.Options })
		}
	}
}

