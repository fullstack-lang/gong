// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
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
		case "Icon":
			FormDivBasicFieldToField(&(_instance.Icon), formDiv)
		case "SVGIcon":
			FormDivSelectFieldToField(&(_instance.SVGIcon), probe.stageOfInterest, formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(_instance.IsDisabled), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.ToolTipPosition), formDiv)
		case "ClientOnX":
			FormDivBasicFieldToField(&(_instance.ClientOnX), formDiv)
		case "ClientOnY":
			FormDivBasicFieldToField(&(_instance.ClientOnY), formDiv)
		case "Menu:Buttons":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Buttons", func(owner *models.Menu) *[]*models.Button { return &owner.Buttons })
		case "Node:Buttons":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Buttons", func(owner *models.Node) *[]*models.Button { return &owner.Buttons })
		}
	}
}

func __gong__New__MenuFormCallback(
	_instance *models.Menu,
	probe *Probe,
	formGroup *form.FormGroup,
) (menuFormCallback *FormCallback[*models.Menu]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMenuFields,
	)
}

type MenuFormCallback = FormCallback[*models.Menu]

func saveMenuFields(
	_instance *models.Menu,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Buttons":
			FormDivSliceOfPointersToField(_instance, "Buttons", &(_instance.Buttons), formDiv, probe)
		}
	}
}

func __gong__New__NodeFormCallback(
	_instance *models.Node,
	probe *Probe,
	formGroup *form.FormGroup,
) (nodeFormCallback *FormCallback[*models.Node]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNodeFields,
	)
}

type NodeFormCallback = FormCallback[*models.Node]

func saveNodeFields(
	_instance *models.Node,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsWithPrefix":
			FormDivBasicFieldToField(&(_instance.IsWithPrefix), formDiv)
		case "Prefix":
			FormDivBasicFieldToField(&(_instance.Prefix), formDiv)
		case "FontStyle":
			FormDivEnumStringFieldToField(&(_instance.FontStyle), formDiv)
		case "BackgroundColor":
			FormDivBasicFieldToField(&(_instance.BackgroundColor), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "HasCheckboxButton":
			FormDivBasicFieldToField(&(_instance.HasCheckboxButton), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsCheckboxDisabled":
			FormDivBasicFieldToField(&(_instance.IsCheckboxDisabled), formDiv)
		case "CheckboxHasToolTip":
			FormDivBasicFieldToField(&(_instance.CheckboxHasToolTip), formDiv)
		case "CheckboxToolTipText":
			FormDivBasicFieldToField(&(_instance.CheckboxToolTipText), formDiv)
		case "CheckboxToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.CheckboxToolTipPosition), formDiv)
		case "HasSecondCheckboxButton":
			FormDivBasicFieldToField(&(_instance.HasSecondCheckboxButton), formDiv)
		case "IsSecondCheckboxChecked":
			FormDivBasicFieldToField(&(_instance.IsSecondCheckboxChecked), formDiv)
		case "IsSecondCheckboxDisabled":
			FormDivBasicFieldToField(&(_instance.IsSecondCheckboxDisabled), formDiv)
		case "SecondCheckboxHasToolTip":
			FormDivBasicFieldToField(&(_instance.SecondCheckboxHasToolTip), formDiv)
		case "SecondCheckboxToolTipText":
			FormDivBasicFieldToField(&(_instance.SecondCheckboxToolTipText), formDiv)
		case "SecondCheckboxToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.SecondCheckboxToolTipPosition), formDiv)
		case "TextAfterSecondCheckbox":
			FormDivBasicFieldToField(&(_instance.TextAfterSecondCheckbox), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.ToolTipPosition), formDiv)
		case "ClientOnY":
			FormDivBasicFieldToField(&(_instance.ClientOnY), formDiv)
		case "IsInEditMode":
			FormDivBasicFieldToField(&(_instance.IsInEditMode), formDiv)
		case "IsNodeClickable":
			FormDivBasicFieldToField(&(_instance.IsNodeClickable), formDiv)
		case "IsWithPreceedingIcon":
			FormDivBasicFieldToField(&(_instance.IsWithPreceedingIcon), formDiv)
		case "PreceedingIcon":
			FormDivBasicFieldToField(&(_instance.PreceedingIcon), formDiv)
		case "PreceedingSVGIcon":
			FormDivSelectFieldToField(&(_instance.PreceedingSVGIcon), probe.stageOfInterest, formDiv)
		case "Children":
			FormDivSliceOfPointersToField(_instance, "Children", &(_instance.Children), formDiv, probe)
		case "Buttons":
			FormDivSliceOfPointersToField(_instance, "Buttons", &(_instance.Buttons), formDiv, probe)
		case "Menu":
			FormDivSelectFieldToField(&(_instance.Menu), probe.stageOfInterest, formDiv)
		case "Node:Children":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Children", func(owner *models.Node) *[]*models.Node { return &owner.Children })
		case "Tree:RootNodes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootNodes", func(owner *models.Tree) *[]*models.Node { return &owner.RootNodes })
		}
	}
}

func __gong__New__SVGIconFormCallback(
	_instance *models.SVGIcon,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgiconFormCallback *FormCallback[*models.SVGIcon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSVGIconFields,
	)
}

type SVGIconFormCallback = FormCallback[*models.SVGIcon]

func saveSVGIconFields(
	_instance *models.SVGIcon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(_instance.SVG), formDiv)
		}
	}
}

func __gong__New__TreeFormCallback(
	_instance *models.Tree,
	probe *Probe,
	formGroup *form.FormGroup,
) (treeFormCallback *FormCallback[*models.Tree]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTreeFields,
	)
}

type TreeFormCallback = FormCallback[*models.Tree]

func saveTreeFields(
	_instance *models.Tree,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RootNodes":
			FormDivSliceOfPointersToField(_instance, "RootNodes", &(_instance.RootNodes), formDiv, probe)
		case "HaveSearch":
			FormDivBasicFieldToField(&(_instance.HaveSearch), formDiv)
		}
	}
}

