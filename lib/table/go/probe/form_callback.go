// generated code - do not edit
package probe

import (
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__CellFormCallback(
	cell *models.Cell,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellFormCallback *CellFormCallback) {
	cellFormCallback = new(CellFormCallback)
	cellFormCallback.probe = probe
	cellFormCallback.cell = cell
	cellFormCallback.formGroup = formGroup

	cellFormCallback.CreationMode = (cell == nil)

	return
}

type CellFormCallback struct {
	cell *models.Cell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellFormCallback *CellFormCallback) OnSave() {

	// log.Println("CellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellFormCallback.probe.formStage.Checkout()

	if cellFormCallback.cell == nil {
		cellFormCallback.cell = new(models.Cell).Stage(cellFormCallback.probe.stageOfInterest)
	}
	cell_ := cellFormCallback.cell
	_ = cell_

	for _, formDiv := range cellFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cell_.Name), formDiv)
		case "CellString":
			FormDivSelectFieldToField(&(cell_.CellString), cellFormCallback.probe.stageOfInterest, formDiv)
		case "CellFloat64":
			FormDivSelectFieldToField(&(cell_.CellFloat64), cellFormCallback.probe.stageOfInterest, formDiv)
		case "CellInt":
			FormDivSelectFieldToField(&(cell_.CellInt), cellFormCallback.probe.stageOfInterest, formDiv)
		case "CellBool":
			FormDivSelectFieldToField(&(cell_.CellBool), cellFormCallback.probe.stageOfInterest, formDiv)
		case "CellIcon":
			FormDivSelectFieldToField(&(cell_.CellIcon), cellFormCallback.probe.stageOfInterest, formDiv)
		case "Row:Cells":
			// we need to retrieve the field owner before the change
			var pastRowOwner *models.Row
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Row"
			rf.Fieldname = "Cells"
			reverseFieldOwner := models.GetReverseFieldOwner(
				cellFormCallback.probe.stageOfInterest,
				cell_,
				&rf)

			if reverseFieldOwner != nil {
				pastRowOwner = reverseFieldOwner.(*models.Row)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastRowOwner != nil {
					idx := slices.Index(pastRowOwner.Cells, cell_)
					pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastRowOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _row := range *models.GetGongstructInstancesSet[models.Row](cellFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _row.GetName() == fieldValue.GetName() {
							newRowOwner := _row // we have a match

							// we remove the cell_ instance from the pastRowOwner field
							if pastRowOwner != nil {
								if newRowOwner != pastRowOwner {
									idx := slices.Index(pastRowOwner.Cells, cell_)
									pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
									newRowOwner.Cells = append(newRowOwner.Cells, cell_)
								}
							} else {
								newRowOwner.Cells = append(newRowOwner.Cells, cell_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cell_.Unstage(cellFormCallback.probe.stageOfInterest)
	}

	cellFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Cell](
		cellFormCallback.probe,
	)
	cellFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellFormCallback.CreationMode || cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cellFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellFormCallback(
			nil,
			cellFormCallback.probe,
			newFormGroup,
		)
		cell := new(models.Cell)
		FillUpForm(cell, newFormGroup, cellFormCallback.probe)
		cellFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cellFormCallback.probe)
}
func __gong__New__CellBooleanFormCallback(
	cellboolean *models.CellBoolean,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellbooleanFormCallback *CellBooleanFormCallback) {
	cellbooleanFormCallback = new(CellBooleanFormCallback)
	cellbooleanFormCallback.probe = probe
	cellbooleanFormCallback.cellboolean = cellboolean
	cellbooleanFormCallback.formGroup = formGroup

	cellbooleanFormCallback.CreationMode = (cellboolean == nil)

	return
}

type CellBooleanFormCallback struct {
	cellboolean *models.CellBoolean

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellbooleanFormCallback *CellBooleanFormCallback) OnSave() {

	// log.Println("CellBooleanFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellbooleanFormCallback.probe.formStage.Checkout()

	if cellbooleanFormCallback.cellboolean == nil {
		cellbooleanFormCallback.cellboolean = new(models.CellBoolean).Stage(cellbooleanFormCallback.probe.stageOfInterest)
	}
	cellboolean_ := cellbooleanFormCallback.cellboolean
	_ = cellboolean_

	for _, formDiv := range cellbooleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellboolean_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellboolean_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if cellbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellboolean_.Unstage(cellbooleanFormCallback.probe.stageOfInterest)
	}

	cellbooleanFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CellBoolean](
		cellbooleanFormCallback.probe,
	)
	cellbooleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellbooleanFormCallback.CreationMode || cellbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellbooleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cellbooleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellBooleanFormCallback(
			nil,
			cellbooleanFormCallback.probe,
			newFormGroup,
		)
		cellboolean := new(models.CellBoolean)
		FillUpForm(cellboolean, newFormGroup, cellbooleanFormCallback.probe)
		cellbooleanFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cellbooleanFormCallback.probe)
}
func __gong__New__CellFloat64FormCallback(
	cellfloat64 *models.CellFloat64,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellfloat64FormCallback *CellFloat64FormCallback) {
	cellfloat64FormCallback = new(CellFloat64FormCallback)
	cellfloat64FormCallback.probe = probe
	cellfloat64FormCallback.cellfloat64 = cellfloat64
	cellfloat64FormCallback.formGroup = formGroup

	cellfloat64FormCallback.CreationMode = (cellfloat64 == nil)

	return
}

type CellFloat64FormCallback struct {
	cellfloat64 *models.CellFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellfloat64FormCallback *CellFloat64FormCallback) OnSave() {

	// log.Println("CellFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellfloat64FormCallback.probe.formStage.Checkout()

	if cellfloat64FormCallback.cellfloat64 == nil {
		cellfloat64FormCallback.cellfloat64 = new(models.CellFloat64).Stage(cellfloat64FormCallback.probe.stageOfInterest)
	}
	cellfloat64_ := cellfloat64FormCallback.cellfloat64
	_ = cellfloat64_

	for _, formDiv := range cellfloat64FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellfloat64_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if cellfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellfloat64_.Unstage(cellfloat64FormCallback.probe.stageOfInterest)
	}

	cellfloat64FormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CellFloat64](
		cellfloat64FormCallback.probe,
	)
	cellfloat64FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellfloat64FormCallback.CreationMode || cellfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellfloat64FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cellfloat64FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellFloat64FormCallback(
			nil,
			cellfloat64FormCallback.probe,
			newFormGroup,
		)
		cellfloat64 := new(models.CellFloat64)
		FillUpForm(cellfloat64, newFormGroup, cellfloat64FormCallback.probe)
		cellfloat64FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cellfloat64FormCallback.probe)
}
func __gong__New__CellIconFormCallback(
	cellicon *models.CellIcon,
	probe *Probe,
	formGroup *table.FormGroup,
) (celliconFormCallback *CellIconFormCallback) {
	celliconFormCallback = new(CellIconFormCallback)
	celliconFormCallback.probe = probe
	celliconFormCallback.cellicon = cellicon
	celliconFormCallback.formGroup = formGroup

	celliconFormCallback.CreationMode = (cellicon == nil)

	return
}

type CellIconFormCallback struct {
	cellicon *models.CellIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (celliconFormCallback *CellIconFormCallback) OnSave() {

	// log.Println("CellIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	celliconFormCallback.probe.formStage.Checkout()

	if celliconFormCallback.cellicon == nil {
		celliconFormCallback.cellicon = new(models.CellIcon).Stage(celliconFormCallback.probe.stageOfInterest)
	}
	cellicon_ := celliconFormCallback.cellicon
	_ = cellicon_

	for _, formDiv := range celliconFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellicon_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(cellicon_.Icon), formDiv)
		}
	}

	// manage the suppress operation
	if celliconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellicon_.Unstage(celliconFormCallback.probe.stageOfInterest)
	}

	celliconFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CellIcon](
		celliconFormCallback.probe,
	)
	celliconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if celliconFormCallback.CreationMode || celliconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		celliconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(celliconFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellIconFormCallback(
			nil,
			celliconFormCallback.probe,
			newFormGroup,
		)
		cellicon := new(models.CellIcon)
		FillUpForm(cellicon, newFormGroup, celliconFormCallback.probe)
		celliconFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(celliconFormCallback.probe)
}
func __gong__New__CellIntFormCallback(
	cellint *models.CellInt,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellintFormCallback *CellIntFormCallback) {
	cellintFormCallback = new(CellIntFormCallback)
	cellintFormCallback.probe = probe
	cellintFormCallback.cellint = cellint
	cellintFormCallback.formGroup = formGroup

	cellintFormCallback.CreationMode = (cellint == nil)

	return
}

type CellIntFormCallback struct {
	cellint *models.CellInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellintFormCallback *CellIntFormCallback) OnSave() {

	// log.Println("CellIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellintFormCallback.probe.formStage.Checkout()

	if cellintFormCallback.cellint == nil {
		cellintFormCallback.cellint = new(models.CellInt).Stage(cellintFormCallback.probe.stageOfInterest)
	}
	cellint_ := cellintFormCallback.cellint
	_ = cellint_

	for _, formDiv := range cellintFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellint_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if cellintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellint_.Unstage(cellintFormCallback.probe.stageOfInterest)
	}

	cellintFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CellInt](
		cellintFormCallback.probe,
	)
	cellintFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellintFormCallback.CreationMode || cellintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellintFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cellintFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellIntFormCallback(
			nil,
			cellintFormCallback.probe,
			newFormGroup,
		)
		cellint := new(models.CellInt)
		FillUpForm(cellint, newFormGroup, cellintFormCallback.probe)
		cellintFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cellintFormCallback.probe)
}
func __gong__New__CellStringFormCallback(
	cellstring *models.CellString,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellstringFormCallback *CellStringFormCallback) {
	cellstringFormCallback = new(CellStringFormCallback)
	cellstringFormCallback.probe = probe
	cellstringFormCallback.cellstring = cellstring
	cellstringFormCallback.formGroup = formGroup

	cellstringFormCallback.CreationMode = (cellstring == nil)

	return
}

type CellStringFormCallback struct {
	cellstring *models.CellString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellstringFormCallback *CellStringFormCallback) OnSave() {

	// log.Println("CellStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellstringFormCallback.probe.formStage.Checkout()

	if cellstringFormCallback.cellstring == nil {
		cellstringFormCallback.cellstring = new(models.CellString).Stage(cellstringFormCallback.probe.stageOfInterest)
	}
	cellstring_ := cellstringFormCallback.cellstring
	_ = cellstring_

	for _, formDiv := range cellstringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cellstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(cellstring_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if cellstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellstring_.Unstage(cellstringFormCallback.probe.stageOfInterest)
	}

	cellstringFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CellString](
		cellstringFormCallback.probe,
	)
	cellstringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellstringFormCallback.CreationMode || cellstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cellstringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellStringFormCallback(
			nil,
			cellstringFormCallback.probe,
			newFormGroup,
		)
		cellstring := new(models.CellString)
		FillUpForm(cellstring, newFormGroup, cellstringFormCallback.probe)
		cellstringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cellstringFormCallback.probe)
}
func __gong__New__CheckBoxFormCallback(
	checkbox *models.CheckBox,
	probe *Probe,
	formGroup *table.FormGroup,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.probe = probe
	checkboxFormCallback.checkbox = checkbox
	checkboxFormCallback.formGroup = formGroup

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {

	// log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.probe.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.probe.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	for _, formDiv := range checkboxFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		case "FormDiv:CheckBoxs":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "CheckBoxs"
			reverseFieldOwner := models.GetReverseFieldOwner(
				checkboxFormCallback.probe.stageOfInterest,
				checkbox_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
					pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastFormDivOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](checkboxFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _formdiv.GetName() == fieldValue.GetName() {
							newFormDivOwner := _formdiv // we have a match

							// we remove the checkbox_ instance from the pastFormDivOwner field
							if pastFormDivOwner != nil {
								if newFormDivOwner != pastFormDivOwner {
									idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
									pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
									newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
								}
							} else {
								newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkbox_.Unstage(checkboxFormCallback.probe.stageOfInterest)
	}

	checkboxFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.CheckBox](
		checkboxFormCallback.probe,
	)
	checkboxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode || checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkboxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(checkboxFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CheckBoxFormCallback(
			nil,
			checkboxFormCallback.probe,
			newFormGroup,
		)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.probe)
		checkboxFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(checkboxFormCallback.probe)
}
func __gong__New__DisplayedColumnFormCallback(
	displayedcolumn *models.DisplayedColumn,
	probe *Probe,
	formGroup *table.FormGroup,
) (displayedcolumnFormCallback *DisplayedColumnFormCallback) {
	displayedcolumnFormCallback = new(DisplayedColumnFormCallback)
	displayedcolumnFormCallback.probe = probe
	displayedcolumnFormCallback.displayedcolumn = displayedcolumn
	displayedcolumnFormCallback.formGroup = formGroup

	displayedcolumnFormCallback.CreationMode = (displayedcolumn == nil)

	return
}

type DisplayedColumnFormCallback struct {
	displayedcolumn *models.DisplayedColumn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (displayedcolumnFormCallback *DisplayedColumnFormCallback) OnSave() {

	// log.Println("DisplayedColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayedcolumnFormCallback.probe.formStage.Checkout()

	if displayedcolumnFormCallback.displayedcolumn == nil {
		displayedcolumnFormCallback.displayedcolumn = new(models.DisplayedColumn).Stage(displayedcolumnFormCallback.probe.stageOfInterest)
	}
	displayedcolumn_ := displayedcolumnFormCallback.displayedcolumn
	_ = displayedcolumn_

	for _, formDiv := range displayedcolumnFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayedcolumn_.Name), formDiv)
		case "Table:DisplayedColumns":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "DisplayedColumns"
			reverseFieldOwner := models.GetReverseFieldOwner(
				displayedcolumnFormCallback.probe.stageOfInterest,
				displayedcolumn_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableOwner = reverseFieldOwner.(*models.Table)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastTableOwner != nil {
					idx := slices.Index(pastTableOwner.DisplayedColumns, displayedcolumn_)
					pastTableOwner.DisplayedColumns = slices.Delete(pastTableOwner.DisplayedColumns, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastTableOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _table := range *models.GetGongstructInstancesSet[models.Table](displayedcolumnFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _table.GetName() == fieldValue.GetName() {
							newTableOwner := _table // we have a match

							// we remove the displayedcolumn_ instance from the pastTableOwner field
							if pastTableOwner != nil {
								if newTableOwner != pastTableOwner {
									idx := slices.Index(pastTableOwner.DisplayedColumns, displayedcolumn_)
									pastTableOwner.DisplayedColumns = slices.Delete(pastTableOwner.DisplayedColumns, idx, idx+1)
									newTableOwner.DisplayedColumns = append(newTableOwner.DisplayedColumns, displayedcolumn_)
								}
							} else {
								newTableOwner.DisplayedColumns = append(newTableOwner.DisplayedColumns, displayedcolumn_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if displayedcolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayedcolumn_.Unstage(displayedcolumnFormCallback.probe.stageOfInterest)
	}

	displayedcolumnFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.DisplayedColumn](
		displayedcolumnFormCallback.probe,
	)
	displayedcolumnFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if displayedcolumnFormCallback.CreationMode || displayedcolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayedcolumnFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(displayedcolumnFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DisplayedColumnFormCallback(
			nil,
			displayedcolumnFormCallback.probe,
			newFormGroup,
		)
		displayedcolumn := new(models.DisplayedColumn)
		FillUpForm(displayedcolumn, newFormGroup, displayedcolumnFormCallback.probe)
		displayedcolumnFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(displayedcolumnFormCallback.probe)
}
func __gong__New__FormDivFormCallback(
	formdiv *models.FormDiv,
	probe *Probe,
	formGroup *table.FormGroup,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.probe = probe
	formdivFormCallback.formdiv = formdiv
	formdivFormCallback.formGroup = formGroup

	formdivFormCallback.CreationMode = (formdiv == nil)

	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {

	// log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.probe.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.probe.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	for _, formDiv := range formdivFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "FormGroup:FormDivs":
			// we need to retrieve the field owner before the change
			var pastFormGroupOwner *models.FormGroup
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormGroup"
			rf.Fieldname = "FormDivs"
			reverseFieldOwner := models.GetReverseFieldOwner(
				formdivFormCallback.probe.stageOfInterest,
				formdiv_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormGroupOwner = reverseFieldOwner.(*models.FormGroup)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastFormGroupOwner != nil {
					idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
					pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastFormGroupOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _formgroup := range *models.GetGongstructInstancesSet[models.FormGroup](formdivFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _formgroup.GetName() == fieldValue.GetName() {
							newFormGroupOwner := _formgroup // we have a match

							// we remove the formdiv_ instance from the pastFormGroupOwner field
							if pastFormGroupOwner != nil {
								if newFormGroupOwner != pastFormGroupOwner {
									idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
									pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
									newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
								}
							} else {
								newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if formdivFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formdiv_.Unstage(formdivFormCallback.probe.stageOfInterest)
	}

	formdivFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormDiv](
		formdivFormCallback.probe,
	)
	formdivFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formdivFormCallback.CreationMode || formdivFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formdivFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formdivFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormDivFormCallback(
			nil,
			formdivFormCallback.probe,
			newFormGroup,
		)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, newFormGroup, formdivFormCallback.probe)
		formdivFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formdivFormCallback.probe)
}
func __gong__New__FormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	probe *Probe,
	formGroup *table.FormGroup,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.probe = probe
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton
	formeditassocbuttonFormCallback.formGroup = formGroup

	formeditassocbuttonFormCallback.CreationMode = (formeditassocbutton == nil)

	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {

	// log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.probe.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.probe.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	for _, formDiv := range formeditassocbuttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		case "AssociationStorage":
			FormDivBasicFieldToField(&(formeditassocbutton_.AssociationStorage), formDiv)
		case "HasChanged":
			FormDivBasicFieldToField(&(formeditassocbutton_.HasChanged), formDiv)
		case "IsForSavePurpose":
			FormDivBasicFieldToField(&(formeditassocbutton_.IsForSavePurpose), formDiv)
		}
	}

	// manage the suppress operation
	if formeditassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formeditassocbutton_.Unstage(formeditassocbuttonFormCallback.probe.stageOfInterest)
	}

	formeditassocbuttonFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormEditAssocButton](
		formeditassocbuttonFormCallback.probe,
	)
	formeditassocbuttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formeditassocbuttonFormCallback.CreationMode || formeditassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formeditassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formeditassocbuttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormEditAssocButtonFormCallback(
			nil,
			formeditassocbuttonFormCallback.probe,
			newFormGroup,
		)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, newFormGroup, formeditassocbuttonFormCallback.probe)
		formeditassocbuttonFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formeditassocbuttonFormCallback.probe)
}
func __gong__New__FormFieldFormCallback(
	formfield *models.FormField,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.probe = probe
	formfieldFormCallback.formfield = formfield
	formfieldFormCallback.formGroup = formGroup

	formfieldFormCallback.CreationMode = (formfield == nil)

	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {

	// log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.probe.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.probe.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	for _, formDiv := range formfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfield_.Name), formDiv)
		case "InputTypeEnum":
			FormDivEnumStringFieldToField(&(formfield_.InputTypeEnum), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formfield_.Label), formDiv)
		case "Placeholder":
			FormDivBasicFieldToField(&(formfield_.Placeholder), formDiv)
		case "FormFieldString":
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		case "HasBespokeHeight":
			FormDivBasicFieldToField(&(formfield_.HasBespokeHeight), formDiv)
		case "BespokeHeightPx":
			FormDivBasicFieldToField(&(formfield_.BespokeHeightPx), formDiv)
		case "FormDiv:FormFields":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "FormFields"
			reverseFieldOwner := models.GetReverseFieldOwner(
				formfieldFormCallback.probe.stageOfInterest,
				formfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
					pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastFormDivOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](formfieldFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _formdiv.GetName() == fieldValue.GetName() {
							newFormDivOwner := _formdiv // we have a match

							// we remove the formfield_ instance from the pastFormDivOwner field
							if pastFormDivOwner != nil {
								if newFormDivOwner != pastFormDivOwner {
									idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
									pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
									newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
								}
							} else {
								newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if formfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfield_.Unstage(formfieldFormCallback.probe.stageOfInterest)
	}

	formfieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormField](
		formfieldFormCallback.probe,
	)
	formfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldFormCallback.CreationMode || formfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldFormCallback(
			nil,
			formfieldFormCallback.probe,
			newFormGroup,
		)
		formfield := new(models.FormField)
		FillUpForm(formfield, newFormGroup, formfieldFormCallback.probe)
		formfieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldFormCallback.probe)
}
func __gong__New__FormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.probe = probe
	formfielddateFormCallback.formfielddate = formfielddate
	formfielddateFormCallback.formGroup = formGroup

	formfielddateFormCallback.CreationMode = (formfielddate == nil)

	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {

	// log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.probe.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.probe.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	for _, formDiv := range formfielddateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddate_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if formfielddateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddate_.Unstage(formfielddateFormCallback.probe.stageOfInterest)
	}

	formfielddateFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldDate](
		formfielddateFormCallback.probe,
	)
	formfielddateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddateFormCallback.CreationMode || formfielddateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfielddateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldDateFormCallback(
			nil,
			formfielddateFormCallback.probe,
			newFormGroup,
		)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, newFormGroup, formfielddateFormCallback.probe)
		formfielddateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfielddateFormCallback.probe)
}
func __gong__New__FormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.probe = probe
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime
	formfielddatetimeFormCallback.formGroup = formGroup

	formfielddatetimeFormCallback.CreationMode = (formfielddatetime == nil)

	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {

	// log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.probe.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.probe.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	for _, formDiv := range formfielddatetimeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddatetime_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if formfielddatetimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddatetime_.Unstage(formfielddatetimeFormCallback.probe.stageOfInterest)
	}

	formfielddatetimeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldDateTime](
		formfielddatetimeFormCallback.probe,
	)
	formfielddatetimeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddatetimeFormCallback.CreationMode || formfielddatetimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddatetimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfielddatetimeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldDateTimeFormCallback(
			nil,
			formfielddatetimeFormCallback.probe,
			newFormGroup,
		)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, newFormGroup, formfielddatetimeFormCallback.probe)
		formfielddatetimeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfielddatetimeFormCallback.probe)
}
func __gong__New__FormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.probe = probe
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64
	formfieldfloat64FormCallback.formGroup = formGroup

	formfieldfloat64FormCallback.CreationMode = (formfieldfloat64 == nil)

	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {

	// log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.probe.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.probe.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	for _, formDiv := range formfieldfloat64FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldfloat64_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MaxValue), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldfloat64_.Unstage(formfieldfloat64FormCallback.probe.stageOfInterest)
	}

	formfieldfloat64FormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldFloat64](
		formfieldfloat64FormCallback.probe,
	)
	formfieldfloat64FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldfloat64FormCallback.CreationMode || formfieldfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldfloat64FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldfloat64FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldFloat64FormCallback(
			nil,
			formfieldfloat64FormCallback.probe,
			newFormGroup,
		)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, newFormGroup, formfieldfloat64FormCallback.probe)
		formfieldfloat64FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldfloat64FormCallback.probe)
}
func __gong__New__FormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.probe = probe
	formfieldintFormCallback.formfieldint = formfieldint
	formfieldintFormCallback.formGroup = formGroup

	formfieldintFormCallback.CreationMode = (formfieldint == nil)

	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {

	// log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.probe.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.probe.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	for _, formDiv := range formfieldintFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldint_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldint_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldint_.MaxValue), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldint_.Unstage(formfieldintFormCallback.probe.stageOfInterest)
	}

	formfieldintFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldInt](
		formfieldintFormCallback.probe,
	)
	formfieldintFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldintFormCallback.CreationMode || formfieldintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldintFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldintFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldIntFormCallback(
			nil,
			formfieldintFormCallback.probe,
			newFormGroup,
		)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, newFormGroup, formfieldintFormCallback.probe)
		formfieldintFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldintFormCallback.probe)
}
func __gong__New__FormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.probe = probe
	formfieldselectFormCallback.formfieldselect = formfieldselect
	formfieldselectFormCallback.formGroup = formGroup

	formfieldselectFormCallback.CreationMode = (formfieldselect == nil)

	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {

	// log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.probe.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.probe.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	for _, formDiv := range formfieldselectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.probe.stageOfInterest, formDiv)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldselectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldselect_.Unstage(formfieldselectFormCallback.probe.stageOfInterest)
	}

	formfieldselectFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldSelect](
		formfieldselectFormCallback.probe,
	)
	formfieldselectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldselectFormCallback.CreationMode || formfieldselectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldselectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldselectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldSelectFormCallback(
			nil,
			formfieldselectFormCallback.probe,
			newFormGroup,
		)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, newFormGroup, formfieldselectFormCallback.probe)
		formfieldselectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldselectFormCallback.probe)
}
func __gong__New__FormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.probe = probe
	formfieldstringFormCallback.formfieldstring = formfieldstring
	formfieldstringFormCallback.formGroup = formGroup

	formfieldstringFormCallback.CreationMode = (formfieldstring == nil)

	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {

	// log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.probe.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.probe.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	for _, formDiv := range formfieldstringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(formfieldstring_.IsTextArea), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldstring_.Unstage(formfieldstringFormCallback.probe.stageOfInterest)
	}

	formfieldstringFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldString](
		formfieldstringFormCallback.probe,
	)
	formfieldstringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldstringFormCallback.CreationMode || formfieldstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldstringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldStringFormCallback(
			nil,
			formfieldstringFormCallback.probe,
			newFormGroup,
		)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, newFormGroup, formfieldstringFormCallback.probe)
		formfieldstringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldstringFormCallback.probe)
}
func __gong__New__FormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	probe *Probe,
	formGroup *table.FormGroup,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.probe = probe
	formfieldtimeFormCallback.formfieldtime = formfieldtime
	formfieldtimeFormCallback.formGroup = formGroup

	formfieldtimeFormCallback.CreationMode = (formfieldtime == nil)

	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {

	// log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.probe.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.probe.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	for _, formDiv := range formfieldtimeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldtime_.Value), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldtimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldtime_.Unstage(formfieldtimeFormCallback.probe.stageOfInterest)
	}

	formfieldtimeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormFieldTime](
		formfieldtimeFormCallback.probe,
	)
	formfieldtimeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldtimeFormCallback.CreationMode || formfieldtimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldtimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formfieldtimeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldTimeFormCallback(
			nil,
			formfieldtimeFormCallback.probe,
			newFormGroup,
		)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, newFormGroup, formfieldtimeFormCallback.probe)
		formfieldtimeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formfieldtimeFormCallback.probe)
}
func __gong__New__FormGroupFormCallback(
	formgroup *models.FormGroup,
	probe *Probe,
	formGroup *table.FormGroup,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.probe = probe
	formgroupFormCallback.formgroup = formgroup
	formgroupFormCallback.formGroup = formGroup

	formgroupFormCallback.CreationMode = (formgroup == nil)

	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {

	// log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.probe.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.probe.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	for _, formDiv := range formgroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formgroup_.Label), formDiv)
		case "HasSuppressButton":
			FormDivBasicFieldToField(&(formgroup_.HasSuppressButton), formDiv)
		case "HasSuppressButtonBeenPressed":
			FormDivBasicFieldToField(&(formgroup_.HasSuppressButtonBeenPressed), formDiv)
		}
	}

	// manage the suppress operation
	if formgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formgroup_.Unstage(formgroupFormCallback.probe.stageOfInterest)
	}

	formgroupFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormGroup](
		formgroupFormCallback.probe,
	)
	formgroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formgroupFormCallback.CreationMode || formgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formgroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormGroupFormCallback(
			nil,
			formgroupFormCallback.probe,
			newFormGroup,
		)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, newFormGroup, formgroupFormCallback.probe)
		formgroupFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formgroupFormCallback.probe)
}
func __gong__New__FormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	probe *Probe,
	formGroup *table.FormGroup,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.probe = probe
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton
	formsortassocbuttonFormCallback.formGroup = formGroup

	formsortassocbuttonFormCallback.CreationMode = (formsortassocbutton == nil)

	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {

	// log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.probe.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.probe.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	for _, formDiv := range formsortassocbuttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		}
	}

	// manage the suppress operation
	if formsortassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formsortassocbutton_.Unstage(formsortassocbuttonFormCallback.probe.stageOfInterest)
	}

	formsortassocbuttonFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FormSortAssocButton](
		formsortassocbuttonFormCallback.probe,
	)
	formsortassocbuttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formsortassocbuttonFormCallback.CreationMode || formsortassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formsortassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formsortassocbuttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormSortAssocButtonFormCallback(
			nil,
			formsortassocbuttonFormCallback.probe,
			newFormGroup,
		)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, newFormGroup, formsortassocbuttonFormCallback.probe)
		formsortassocbuttonFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formsortassocbuttonFormCallback.probe)
}
func __gong__New__OptionFormCallback(
	option *models.Option,
	probe *Probe,
	formGroup *table.FormGroup,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.probe = probe
	optionFormCallback.option = option
	optionFormCallback.formGroup = formGroup

	optionFormCallback.CreationMode = (option == nil)

	return
}

type OptionFormCallback struct {
	option *models.Option

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (optionFormCallback *OptionFormCallback) OnSave() {

	// log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.probe.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.probe.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	for _, formDiv := range optionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		case "FormFieldSelect:Options":
			// we need to retrieve the field owner before the change
			var pastFormFieldSelectOwner *models.FormFieldSelect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormFieldSelect"
			rf.Fieldname = "Options"
			reverseFieldOwner := models.GetReverseFieldOwner(
				optionFormCallback.probe.stageOfInterest,
				option_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormFieldSelectOwner = reverseFieldOwner.(*models.FormFieldSelect)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastFormFieldSelectOwner != nil {
					idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
					pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastFormFieldSelectOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _formfieldselect := range *models.GetGongstructInstancesSet[models.FormFieldSelect](optionFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _formfieldselect.GetName() == fieldValue.GetName() {
							newFormFieldSelectOwner := _formfieldselect // we have a match

							// we remove the option_ instance from the pastFormFieldSelectOwner field
							if pastFormFieldSelectOwner != nil {
								if newFormFieldSelectOwner != pastFormFieldSelectOwner {
									idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
									pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
									newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
								}
							} else {
								newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if optionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		option_.Unstage(optionFormCallback.probe.stageOfInterest)
	}

	optionFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Option](
		optionFormCallback.probe,
	)
	optionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if optionFormCallback.CreationMode || optionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		optionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(optionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OptionFormCallback(
			nil,
			optionFormCallback.probe,
			newFormGroup,
		)
		option := new(models.Option)
		FillUpForm(option, newFormGroup, optionFormCallback.probe)
		optionFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(optionFormCallback.probe)
}
func __gong__New__RowFormCallback(
	row *models.Row,
	probe *Probe,
	formGroup *table.FormGroup,
) (rowFormCallback *RowFormCallback) {
	rowFormCallback = new(RowFormCallback)
	rowFormCallback.probe = probe
	rowFormCallback.row = row
	rowFormCallback.formGroup = formGroup

	rowFormCallback.CreationMode = (row == nil)

	return
}

type RowFormCallback struct {
	row *models.Row

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rowFormCallback *RowFormCallback) OnSave() {

	// log.Println("RowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rowFormCallback.probe.formStage.Checkout()

	if rowFormCallback.row == nil {
		rowFormCallback.row = new(models.Row).Stage(rowFormCallback.probe.stageOfInterest)
	}
	row_ := rowFormCallback.row
	_ = row_

	for _, formDiv := range rowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(row_.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(row_.IsChecked), formDiv)
		case "Table:Rows":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "Rows"
			reverseFieldOwner := models.GetReverseFieldOwner(
				rowFormCallback.probe.stageOfInterest,
				row_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableOwner = reverseFieldOwner.(*models.Table)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastTableOwner != nil {
					idx := slices.Index(pastTableOwner.Rows, row_)
					pastTableOwner.Rows = slices.Delete(pastTableOwner.Rows, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastTableOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _table := range *models.GetGongstructInstancesSet[models.Table](rowFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _table.GetName() == fieldValue.GetName() {
							newTableOwner := _table // we have a match

							// we remove the row_ instance from the pastTableOwner field
							if pastTableOwner != nil {
								if newTableOwner != pastTableOwner {
									idx := slices.Index(pastTableOwner.Rows, row_)
									pastTableOwner.Rows = slices.Delete(pastTableOwner.Rows, idx, idx+1)
									newTableOwner.Rows = append(newTableOwner.Rows, row_)
								}
							} else {
								newTableOwner.Rows = append(newTableOwner.Rows, row_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		row_.Unstage(rowFormCallback.probe.stageOfInterest)
	}

	rowFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Row](
		rowFormCallback.probe,
	)
	rowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rowFormCallback.CreationMode || rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RowFormCallback(
			nil,
			rowFormCallback.probe,
			newFormGroup,
		)
		row := new(models.Row)
		FillUpForm(row, newFormGroup, rowFormCallback.probe)
		rowFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rowFormCallback.probe)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
	formGroup *table.FormGroup,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.probe = probe
	tableFormCallback.table = table
	tableFormCallback.formGroup = formGroup

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tableFormCallback *TableFormCallback) OnSave() {

	// log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.probe.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.probe.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	for _, formDiv := range tableFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "HasFiltering":
			FormDivBasicFieldToField(&(table_.HasFiltering), formDiv)
		case "HasColumnSorting":
			FormDivBasicFieldToField(&(table_.HasColumnSorting), formDiv)
		case "HasPaginator":
			FormDivBasicFieldToField(&(table_.HasPaginator), formDiv)
		case "HasCheckableRows":
			FormDivBasicFieldToField(&(table_.HasCheckableRows), formDiv)
		case "HasSaveButton":
			FormDivBasicFieldToField(&(table_.HasSaveButton), formDiv)
		case "CanDragDropRows":
			FormDivBasicFieldToField(&(table_.CanDragDropRows), formDiv)
		case "HasCloseButton":
			FormDivBasicFieldToField(&(table_.HasCloseButton), formDiv)
		case "SavingInProgress":
			FormDivBasicFieldToField(&(table_.SavingInProgress), formDiv)
		case "NbOfStickyColumns":
			FormDivBasicFieldToField(&(table_.NbOfStickyColumns), formDiv)
		}
	}

	// manage the suppress operation
	if tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		table_.Unstage(tableFormCallback.probe.stageOfInterest)
	}

	tableFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Table](
		tableFormCallback.probe,
	)
	tableFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode || tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tableFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tableFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			tableFormCallback.probe,
			newFormGroup,
		)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.probe)
		tableFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tableFormCallback.probe)
}
