// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
func __gong__New__AsSplitFormCallback(
	_instance *models.AsSplit,
	probe *Probe,
	formGroup *form.FormGroup,
) (assplitFormCallback *FormCallback[*models.AsSplit]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAsSplitFields,
	)
}

type AsSplitFormCallback = FormCallback[*models.AsSplit]

func saveAsSplitFields(
	_instance *models.AsSplit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(_instance.Direction), formDiv)
		case "AsSplitAreas":
			FormDivSliceOfPointersToField(_instance, "AsSplitAreas", &(_instance.AsSplitAreas), formDiv, probe)
		case "IsSizeInPixel":
			FormDivBasicFieldToField(&(_instance.IsSizeInPixel), formDiv)
		case "IsWithCustomGutterSize":
			FormDivBasicFieldToField(&(_instance.IsWithCustomGutterSize), formDiv)
		case "GutterSize":
			FormDivBasicFieldToField(&(_instance.GutterSize), formDiv)
		}
	}
}

func __gong__New__AsSplitAreaFormCallback(
	_instance *models.AsSplitArea,
	probe *Probe,
	formGroup *form.FormGroup,
) (assplitareaFormCallback *FormCallback[*models.AsSplitArea]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAsSplitAreaFields,
	)
}

type AsSplitAreaFormCallback = FormCallback[*models.AsSplitArea]

func saveAsSplitAreaFields(
	_instance *models.AsSplitArea,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowNameInHeader":
			FormDivBasicFieldToField(&(_instance.ShowNameInHeader), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(_instance.Size), formDiv)
		case "IsAny":
			FormDivBasicFieldToField(&(_instance.IsAny), formDiv)
		case "AsSplit":
			FormDivSelectFieldToField(&(_instance.AsSplit), probe.stageOfInterest, formDiv)
		case "Button":
			FormDivSelectFieldToField(&(_instance.Button), probe.stageOfInterest, formDiv)
		case "Cursor":
			FormDivSelectFieldToField(&(_instance.Cursor), probe.stageOfInterest, formDiv)
		case "Form":
			FormDivSelectFieldToField(&(_instance.Form), probe.stageOfInterest, formDiv)
		case "Load":
			FormDivSelectFieldToField(&(_instance.Load), probe.stageOfInterest, formDiv)
		case "Markdown":
			FormDivSelectFieldToField(&(_instance.Markdown), probe.stageOfInterest, formDiv)
		case "Slider":
			FormDivSelectFieldToField(&(_instance.Slider), probe.stageOfInterest, formDiv)
		case "Split":
			FormDivSelectFieldToField(&(_instance.Split), probe.stageOfInterest, formDiv)
		case "Svg":
			FormDivSelectFieldToField(&(_instance.Svg), probe.stageOfInterest, formDiv)
		case "Table":
			FormDivSelectFieldToField(&(_instance.Table), probe.stageOfInterest, formDiv)
		case "Tone":
			FormDivSelectFieldToField(&(_instance.Tone), probe.stageOfInterest, formDiv)
		case "Tree":
			FormDivSelectFieldToField(&(_instance.Tree), probe.stageOfInterest, formDiv)
		case "Threejs":
			FormDivSelectFieldToField(&(_instance.Threejs), probe.stageOfInterest, formDiv)
		case "Xlsx":
			FormDivSelectFieldToField(&(_instance.Xlsx), probe.stageOfInterest, formDiv)
		case "HasDiv":
			FormDivBasicFieldToField(&(_instance.HasDiv), formDiv)
		case "DivStyle":
			FormDivBasicFieldToField(&(_instance.DivStyle), formDiv)
		case "AsSplit:AsSplitAreas":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AsSplitAreas", func(owner *models.AsSplit) *[]*models.AsSplitArea { return &owner.AsSplitAreas })
		case "View:RootAsSplitAreas":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootAsSplitAreas", func(owner *models.View) *[]*models.AsSplitArea { return &owner.RootAsSplitAreas })
		}
	}
}

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
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

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
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		case "Style":
			FormDivBasicFieldToField(&(_instance.Style), formDiv)
		}
	}
}

func __gong__New__FavIconFormCallback(
	_instance *models.FavIcon,
	probe *Probe,
	formGroup *form.FormGroup,
) (faviconFormCallback *FormCallback[*models.FavIcon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFavIconFields,
	)
}

type FavIconFormCallback = FormCallback[*models.FavIcon]

func saveFavIconFields(
	_instance *models.FavIcon,
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

func __gong__New__FormFormCallback(
	_instance *models.Form,
	probe *Probe,
	formGroup *form.FormGroup,
) (formFormCallback *FormCallback[*models.Form]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormFields,
	)
}

type FormFormCallback = FormCallback[*models.Form]

func saveFormFields(
	_instance *models.Form,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__LoadFormCallback(
	_instance *models.Load,
	probe *Probe,
	formGroup *form.FormGroup,
) (loadFormCallback *FormCallback[*models.Load]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLoadFields,
	)
}

type LoadFormCallback = FormCallback[*models.Load]

func saveLoadFields(
	_instance *models.Load,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__LogoOnTheLeftFormCallback(
	_instance *models.LogoOnTheLeft,
	probe *Probe,
	formGroup *form.FormGroup,
) (logoontheleftFormCallback *FormCallback[*models.LogoOnTheLeft]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLogoOnTheLeftFields,
	)
}

type LogoOnTheLeftFormCallback = FormCallback[*models.LogoOnTheLeft]

func saveLogoOnTheLeftFields(
	_instance *models.LogoOnTheLeft,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(_instance.SVG), formDiv)
		}
	}
}

func __gong__New__LogoOnTheRightFormCallback(
	_instance *models.LogoOnTheRight,
	probe *Probe,
	formGroup *form.FormGroup,
) (logoontherightFormCallback *FormCallback[*models.LogoOnTheRight]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLogoOnTheRightFields,
	)
}

type LogoOnTheRightFormCallback = FormCallback[*models.LogoOnTheRight]

func saveLogoOnTheRightFields(
	_instance *models.LogoOnTheRight,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(_instance.SVG), formDiv)
		}
	}
}

func __gong__New__MarkdownFormCallback(
	_instance *models.Markdown,
	probe *Probe,
	formGroup *form.FormGroup,
) (markdownFormCallback *FormCallback[*models.Markdown]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMarkdownFields,
	)
}

type MarkdownFormCallback = FormCallback[*models.Markdown]

func saveMarkdownFields(
	_instance *models.Markdown,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
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
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__SplitFormCallback(
	_instance *models.Split,
	probe *Probe,
	formGroup *form.FormGroup,
) (splitFormCallback *FormCallback[*models.Split]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSplitFields,
	)
}

type SplitFormCallback = FormCallback[*models.Split]

func saveSplitFields(
	_instance *models.Split,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__SvgFormCallback(
	_instance *models.Svg,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgFormCallback *FormCallback[*models.Svg]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSvgFields,
	)
}

type SvgFormCallback = FormCallback[*models.Svg]

func saveSvgFields(
	_instance *models.Svg,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		case "Style":
			FormDivBasicFieldToField(&(_instance.Style), formDiv)
		}
	}
}

func __gong__New__TableFormCallback(
	_instance *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
) (tableFormCallback *FormCallback[*models.Table]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTableFields,
	)
}

type TableFormCallback = FormCallback[*models.Table]

func saveTableFields(
	_instance *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__ThreejsFormCallback(
	_instance *models.Threejs,
	probe *Probe,
	formGroup *form.FormGroup,
) (threejsFormCallback *FormCallback[*models.Threejs]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveThreejsFields,
	)
}

type ThreejsFormCallback = FormCallback[*models.Threejs]

func saveThreejsFields(
	_instance *models.Threejs,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__TitleFormCallback(
	_instance *models.Title,
	probe *Probe,
	formGroup *form.FormGroup,
) (titleFormCallback *FormCallback[*models.Title]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTitleFields,
	)
}

type TitleFormCallback = FormCallback[*models.Title]

func saveTitleFields(
	_instance *models.Title,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__ToneFormCallback(
	_instance *models.Tone,
	probe *Probe,
	formGroup *form.FormGroup,
) (toneFormCallback *FormCallback[*models.Tone]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveToneFields,
	)
}

type ToneFormCallback = FormCallback[*models.Tone]

func saveToneFields(
	_instance *models.Tone,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
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
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

func __gong__New__ViewFormCallback(
	_instance *models.View,
	probe *Probe,
	formGroup *form.FormGroup,
) (viewFormCallback *FormCallback[*models.View]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveViewFields,
	)
}

type ViewFormCallback = FormCallback[*models.View]

func saveViewFields(
	_instance *models.View,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowViewName":
			FormDivBasicFieldToField(&(_instance.ShowViewName), formDiv)
		case "RootAsSplitAreas":
			FormDivSliceOfPointersToField(_instance, "RootAsSplitAreas", &(_instance.RootAsSplitAreas), formDiv, probe)
		case "IsSelectedView":
			FormDivBasicFieldToField(&(_instance.IsSelectedView), formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(_instance.Direction), formDiv)
		case "IsSecondaryView":
			FormDivBasicFieldToField(&(_instance.IsSecondaryView), formDiv)
		case "IsSizeInPixel":
			FormDivBasicFieldToField(&(_instance.IsSizeInPixel), formDiv)
		case "IsWithCustomGutterSize":
			FormDivBasicFieldToField(&(_instance.IsWithCustomGutterSize), formDiv)
		case "GutterSize":
			FormDivBasicFieldToField(&(_instance.GutterSize), formDiv)
		}
	}
}

func __gong__New__XlsxFormCallback(
	_instance *models.Xlsx,
	probe *Probe,
	formGroup *form.FormGroup,
) (xlsxFormCallback *FormCallback[*models.Xlsx]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXlsxFields,
	)
}

type XlsxFormCallback = FormCallback[*models.Xlsx]

func saveXlsxFields(
	_instance *models.Xlsx,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(_instance.StackName), formDiv)
		}
	}
}

