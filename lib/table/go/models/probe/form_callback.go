// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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
		case "Table:Buttons":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Buttons", func(owner *models.Table) *[]*models.Button { return &owner.Buttons })
		}
	}
}

func __gong__New__CellFormCallback(
	_instance *models.Cell,
	probe *Probe,
	formGroup *form.FormGroup,
) (cellFormCallback *FormCallback[*models.Cell]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellFields,
	)
}

type CellFormCallback = FormCallback[*models.Cell]

func saveCellFields(
	_instance *models.Cell,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "CellString":
			FormDivSelectFieldToField(&(_instance.CellString), probe.stageOfInterest, formDiv)
		case "CellFloat64":
			FormDivSelectFieldToField(&(_instance.CellFloat64), probe.stageOfInterest, formDiv)
		case "CellInt":
			FormDivSelectFieldToField(&(_instance.CellInt), probe.stageOfInterest, formDiv)
		case "CellBool":
			FormDivSelectFieldToField(&(_instance.CellBool), probe.stageOfInterest, formDiv)
		case "CellIcon":
			FormDivSelectFieldToField(&(_instance.CellIcon), probe.stageOfInterest, formDiv)
		case "Row:Cells":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Cells", func(owner *models.Row) *[]*models.Cell { return &owner.Cells })
		}
	}
}

func __gong__New__CellBooleanFormCallback(
	_instance *models.CellBoolean,
	probe *Probe,
	formGroup *form.FormGroup,
) (cellbooleanFormCallback *FormCallback[*models.CellBoolean]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellBooleanFields,
	)
}

type CellBooleanFormCallback = FormCallback[*models.CellBoolean]

func saveCellBooleanFields(
	_instance *models.CellBoolean,
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
		}
	}
}

func __gong__New__CellFloat64FormCallback(
	_instance *models.CellFloat64,
	probe *Probe,
	formGroup *form.FormGroup,
) (cellfloat64FormCallback *FormCallback[*models.CellFloat64]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellFloat64Fields,
	)
}

type CellFloat64FormCallback = FormCallback[*models.CellFloat64]

func saveCellFloat64Fields(
	_instance *models.CellFloat64,
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
		}
	}
}

func __gong__New__CellIconFormCallback(
	_instance *models.CellIcon,
	probe *Probe,
	formGroup *form.FormGroup,
) (celliconFormCallback *FormCallback[*models.CellIcon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellIconFields,
	)
}

type CellIconFormCallback = FormCallback[*models.CellIcon]

func saveCellIconFields(
	_instance *models.CellIcon,
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
		case "NeedsConfirmation":
			FormDivBasicFieldToField(&(_instance.NeedsConfirmation), formDiv)
		case "ConfirmationMessage":
			FormDivBasicFieldToField(&(_instance.ConfirmationMessage), formDiv)
		}
	}
}

func __gong__New__CellIntFormCallback(
	_instance *models.CellInt,
	probe *Probe,
	formGroup *form.FormGroup,
) (cellintFormCallback *FormCallback[*models.CellInt]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellIntFields,
	)
}

type CellIntFormCallback = FormCallback[*models.CellInt]

func saveCellIntFields(
	_instance *models.CellInt,
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
		}
	}
}

func __gong__New__CellStringFormCallback(
	_instance *models.CellString,
	probe *Probe,
	formGroup *form.FormGroup,
) (cellstringFormCallback *FormCallback[*models.CellString]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCellStringFields,
	)
}

type CellStringFormCallback = FormCallback[*models.CellString]

func saveCellStringFields(
	_instance *models.CellString,
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
		}
	}
}

func __gong__New__DisplayedColumnFormCallback(
	_instance *models.DisplayedColumn,
	probe *Probe,
	formGroup *form.FormGroup,
) (displayedcolumnFormCallback *FormCallback[*models.DisplayedColumn]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDisplayedColumnFields,
	)
}

type DisplayedColumnFormCallback = FormCallback[*models.DisplayedColumn]

func saveDisplayedColumnFields(
	_instance *models.DisplayedColumn,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Table:DisplayedColumns":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DisplayedColumns", func(owner *models.Table) *[]*models.DisplayedColumn { return &owner.DisplayedColumns })
		}
	}
}

func __gong__New__RowFormCallback(
	_instance *models.Row,
	probe *Probe,
	formGroup *form.FormGroup,
) (rowFormCallback *FormCallback[*models.Row]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRowFields,
	)
}

type RowFormCallback = FormCallback[*models.Row]

func saveRowFields(
	_instance *models.Row,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Cells":
			FormDivSliceOfPointersToField(_instance, "Cells", &(_instance.Cells), formDiv, probe)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "Table:Rows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Rows", func(owner *models.Table) *[]*models.Row { return &owner.Rows })
		case "Table:RowsSelectedForBulkDelete":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RowsSelectedForBulkDelete", func(owner *models.Table) *[]*models.Row { return &owner.RowsSelectedForBulkDelete })
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
		case "DisplayedColumns":
			FormDivSliceOfPointersToField(_instance, "DisplayedColumns", &(_instance.DisplayedColumns), formDiv, probe)
		case "Rows":
			FormDivSliceOfPointersToField(_instance, "Rows", &(_instance.Rows), formDiv, probe)
		case "HasFiltering":
			FormDivBasicFieldToField(&(_instance.HasFiltering), formDiv)
		case "HasColumnSorting":
			FormDivBasicFieldToField(&(_instance.HasColumnSorting), formDiv)
		case "HasPaginator":
			FormDivBasicFieldToField(&(_instance.HasPaginator), formDiv)
		case "HasCheckableRows":
			FormDivBasicFieldToField(&(_instance.HasCheckableRows), formDiv)
		case "HasSaveButton":
			FormDivBasicFieldToField(&(_instance.HasSaveButton), formDiv)
		case "SaveButtonLabel":
			FormDivBasicFieldToField(&(_instance.SaveButtonLabel), formDiv)
		case "HasBulkDeleteButton":
			FormDivBasicFieldToField(&(_instance.HasBulkDeleteButton), formDiv)
		case "BulkDeleteButtonTooltip":
			FormDivBasicFieldToField(&(_instance.BulkDeleteButtonTooltip), formDiv)
		case "RowsSelectedForBulkDelete":
			FormDivSliceOfPointersToField(_instance, "RowsSelectedForBulkDelete", &(_instance.RowsSelectedForBulkDelete), formDiv, probe)
		case "CanDragDropRows":
			FormDivBasicFieldToField(&(_instance.CanDragDropRows), formDiv)
		case "HasCloseButton":
			FormDivBasicFieldToField(&(_instance.HasCloseButton), formDiv)
		case "SavingInProgress":
			FormDivBasicFieldToField(&(_instance.SavingInProgress), formDiv)
		case "NbOfStickyColumns":
			FormDivBasicFieldToField(&(_instance.NbOfStickyColumns), formDiv)
		case "Buttons":
			FormDivSliceOfPointersToField(_instance, "Buttons", &(_instance.Buttons), formDiv, probe)
		}
	}
}

