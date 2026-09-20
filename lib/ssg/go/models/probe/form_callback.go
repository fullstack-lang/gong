// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
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
func __gong__New__ChapterFormCallback(
	_instance *models.Chapter,
	probe *Probe,
	formGroup *form.FormGroup,
) (chapterFormCallback *FormCallback[*models.Chapter]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveChapterFields,
	)
}

type ChapterFormCallback = FormCallback[*models.Chapter]

func saveChapterFields(
	_instance *models.Chapter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(_instance.MardownContent), formDiv)
		case "Sections":
			FormDivSliceOfPointersToField(_instance, "Sections", &(_instance.Sections), formDiv, probe)
		case "Pages":
			FormDivSliceOfPointersToField(_instance, "Pages", &(_instance.Pages), formDiv, probe)
		case "SubChapters":
			FormDivSliceOfPointersToField(_instance, "SubChapters", &(_instance.SubChapters), formDiv, probe)
		case "Chapter:SubChapters":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubChapters", func(owner *models.Chapter) *[]*models.Chapter { return &owner.SubChapters })
		case "Content:Chapters":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Chapters", func(owner *models.Content) *[]*models.Chapter { return &owner.Chapters })
		}
	}
}

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
		case "MardownContent":
			FormDivBasicFieldToField(&(_instance.MardownContent), formDiv)
		case "ContentPath":
			FormDivBasicFieldToField(&(_instance.ContentPath), formDiv)
		case "OutputPath":
			FormDivBasicFieldToField(&(_instance.OutputPath), formDiv)
		case "StaticPath":
			FormDivBasicFieldToField(&(_instance.StaticPath), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "IsBespokeLogoFileName":
			FormDivBasicFieldToField(&(_instance.IsBespokeLogoFileName), formDiv)
		case "BespokeLogoFileName":
			FormDivBasicFieldToField(&(_instance.BespokeLogoFileName), formDiv)
		case "IsBespokePageTileLogoFileName":
			FormDivBasicFieldToField(&(_instance.IsBespokePageTileLogoFileName), formDiv)
		case "BespokePageTileLogoFileName":
			FormDivBasicFieldToField(&(_instance.BespokePageTileLogoFileName), formDiv)
		case "Target":
			FormDivEnumStringFieldToField(&(_instance.Target), formDiv)
		case "Chapters":
			FormDivSliceOfPointersToField(_instance, "Chapters", &(_instance.Chapters), formDiv, probe)
		case "VersionInfo":
			FormDivBasicFieldToField(&(_instance.VersionInfo), formDiv)
		}
	}
}

func __gong__New__DownloadableFileFormCallback(
	_instance *models.DownloadableFile,
	probe *Probe,
	formGroup *form.FormGroup,
) (downloadablefileFormCallback *FormCallback[*models.DownloadableFile]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDownloadableFileFields,
	)
}

type DownloadableFileFormCallback = FormCallback[*models.DownloadableFile]

func saveDownloadableFileFields(
	_instance *models.DownloadableFile,
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

func __gong__New__PageFormCallback(
	_instance *models.Page,
	probe *Probe,
	formGroup *form.FormGroup,
) (pageFormCallback *FormCallback[*models.Page]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePageFields,
	)
}

type PageFormCallback = FormCallback[*models.Page]

func savePageFields(
	_instance *models.Page,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(_instance.MardownContent), formDiv)
		case "Sections":
			FormDivSliceOfPointersToField(_instance, "Sections", &(_instance.Sections), formDiv, probe)
		case "Chapter:Pages":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Pages", func(owner *models.Chapter) *[]*models.Page { return &owner.Pages })
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

func __gong__New__SectionFormCallback(
	_instance *models.Section,
	probe *Probe,
	formGroup *form.FormGroup,
) (sectionFormCallback *FormCallback[*models.Section]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSectionFields,
	)
}

type SectionFormCallback = FormCallback[*models.Section]

func saveSectionFields(
	_instance *models.Section,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(_instance.MardownContent), formDiv)
		case "IsImage":
			FormDivBasicFieldToField(&(_instance.IsImage), formDiv)
		case "SvgImage":
			FormDivSelectFieldToField(&(_instance.SvgImage), probe.stageOfInterest, formDiv)
		case "PngImage":
			FormDivSelectFieldToField(&(_instance.PngImage), probe.stageOfInterest, formDiv)
		case "JpgImage":
			FormDivSelectFieldToField(&(_instance.JpgImage), probe.stageOfInterest, formDiv)
		case "IsDownloadableFile":
			FormDivBasicFieldToField(&(_instance.IsDownloadableFile), formDiv)
		case "DownloadableFile":
			FormDivSelectFieldToField(&(_instance.DownloadableFile), probe.stageOfInterest, formDiv)
		case "Chapter:Sections":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sections", func(owner *models.Chapter) *[]*models.Section { return &owner.Sections })
		case "Page:Sections":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sections", func(owner *models.Page) *[]*models.Section { return &owner.Sections })
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

