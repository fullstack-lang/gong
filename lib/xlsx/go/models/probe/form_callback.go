// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
func __gong__New__DisplaySelectionFormCallback(
	_instance *models.DisplaySelection,
	probe *Probe,
	formGroup *form.FormGroup,
) (displayselectionFormCallback *FormCallback[*models.DisplaySelection]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDisplaySelectionFields,
	)
}

type DisplaySelectionFormCallback = FormCallback[*models.DisplaySelection]

func saveDisplaySelectionFields(
	_instance *models.DisplaySelection,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "XLFile":
			FormDivSelectFieldToField(&(_instance.XLFile), probe.stageOfInterest, formDiv)
		case "XLSheet":
			FormDivSelectFieldToField(&(_instance.XLSheet), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__XLCellFormCallback(
	_instance *models.XLCell,
	probe *Probe,
	formGroup *form.FormGroup,
) (xlcellFormCallback *FormCallback[*models.XLCell]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXLCellFields,
	)
}

type XLCellFormCallback = FormCallback[*models.XLCell]

func saveXLCellFields(
	_instance *models.XLCell,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "XLRow:Cells":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Cells", func(owner *models.XLRow) *[]*models.XLCell { return &owner.Cells })
		case "XLSheet:SheetCells":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SheetCells", func(owner *models.XLSheet) *[]*models.XLCell { return &owner.SheetCells })
		}
	}
}

func __gong__New__XLFileFormCallback(
	_instance *models.XLFile,
	probe *Probe,
	formGroup *form.FormGroup,
) (xlfileFormCallback *FormCallback[*models.XLFile]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXLFileFields,
	)
}

type XLFileFormCallback = FormCallback[*models.XLFile]

func saveXLFileFields(
	_instance *models.XLFile,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NbSheets":
			FormDivBasicFieldToField(&(_instance.NbSheets), formDiv)
		case "Sheets":
			FormDivSliceOfPointersToField(_instance, "Sheets", &(_instance.Sheets), formDiv, probe)
		}
	}
}

func __gong__New__XLRowFormCallback(
	_instance *models.XLRow,
	probe *Probe,
	formGroup *form.FormGroup,
) (xlrowFormCallback *FormCallback[*models.XLRow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXLRowFields,
	)
}

type XLRowFormCallback = FormCallback[*models.XLRow]

func saveXLRowFields(
	_instance *models.XLRow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RowIndex":
			FormDivBasicFieldToField(&(_instance.RowIndex), formDiv)
		case "Cells":
			FormDivSliceOfPointersToField(_instance, "Cells", &(_instance.Cells), formDiv, probe)
		case "XLSheet:Rows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Rows", func(owner *models.XLSheet) *[]*models.XLRow { return &owner.Rows })
		}
	}
}

func __gong__New__XLSheetFormCallback(
	_instance *models.XLSheet,
	probe *Probe,
	formGroup *form.FormGroup,
) (xlsheetFormCallback *FormCallback[*models.XLSheet]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXLSheetFields,
	)
}

type XLSheetFormCallback = FormCallback[*models.XLSheet]

func saveXLSheetFields(
	_instance *models.XLSheet,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MaxRow":
			FormDivBasicFieldToField(&(_instance.MaxRow), formDiv)
		case "MaxCol":
			FormDivBasicFieldToField(&(_instance.MaxCol), formDiv)
		case "NbRows":
			FormDivBasicFieldToField(&(_instance.NbRows), formDiv)
		case "Rows":
			FormDivSliceOfPointersToField(_instance, "Rows", &(_instance.Rows), formDiv, probe)
		case "SheetCells":
			FormDivSliceOfPointersToField(_instance, "SheetCells", &(_instance.SheetCells), formDiv, probe)
		case "XLFile:Sheets":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sheets", func(owner *models.XLFile) *[]*models.XLSheet { return &owner.Sheets })
		}
	}
}

