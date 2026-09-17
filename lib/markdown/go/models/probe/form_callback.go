// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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
func __gong__New__ContentFormCallback(
	_instance *models.Content,
	probe *Probe,
	formGroup *form.FormGroup,
) (contentFormCallback *FormCallback[*models.Content]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveContentFields,
	)
}

type ContentFormCallback = FormCallback[*models.Content]

func saveContentFields(
	_instance *models.Content,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		}
	}
}

func __gong__New__JpgImageFormCallback(
	_instance *models.JpgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (jpgimageFormCallback *FormCallback[*models.JpgImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveJpgImageFields,
	)
}

type JpgImageFormCallback = FormCallback[*models.JpgImage]

func saveJpgImageFields(
	_instance *models.JpgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(_instance.Base64Content), formDiv)
		}
	}
}

func __gong__New__PngImageFormCallback(
	_instance *models.PngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (pngimageFormCallback *FormCallback[*models.PngImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePngImageFields,
	)
}

type PngImageFormCallback = FormCallback[*models.PngImage]

func savePngImageFields(
	_instance *models.PngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(_instance.Base64Content), formDiv)
		}
	}
}

func __gong__New__SvgImageFormCallback(
	_instance *models.SvgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgimageFormCallback *FormCallback[*models.SvgImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSvgImageFields,
	)
}

type SvgImageFormCallback = FormCallback[*models.SvgImage]

func saveSvgImageFields(
	_instance *models.SvgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		}
	}
}

