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

// insertion point
func __gong__New__ButtonFormCallback(
	button *models.Button,
	probe *Probe,
	formGroup *form.FormGroup,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.probe = probe
	buttonFormCallback.button = button
	buttonFormCallback.formGroup = formGroup

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {
	buttonFormCallback.probe.stageOfInterest.Lock()
	defer buttonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.probe.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.probe.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	for _, formDiv := range buttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(button_.Icon), formDiv)
		case "SVGIcon":
			FormDivSelectFieldToField(&(button_.SVGIcon), buttonFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(button_.IsDisabled), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(button_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(button_.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(button_.ToolTipPosition), formDiv)
		case "Table:Buttons":
			// WARNING : this form deals with the N-N association "Table.Buttons []*Button" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Button). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Table
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Table"
				rf.Fieldname = "Buttons"
				formerAssociationSource := button_.GongGetReverseFieldOwner(
					buttonFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Table)
					if !ok {
						log.Fatalln("Source of Table.Buttons []*Button, is not an Table instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Buttons, button_)
					formerSource.Buttons = slices.Delete(formerSource.Buttons, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Table
			for _table := range *models.GetGongstructInstancesSet[models.Table](buttonFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _table.GetName() == newSourceName.GetName() {
					newSource = _table // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Table.Buttons []*Button, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Buttons = append(newSource.Buttons, button_)
		}
	}

	// manage the suppress operation
	if buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		button_.Unstage(buttonFormCallback.probe.stageOfInterest)
	}

	buttonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Button](
		buttonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(buttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			buttonFormCallback.probe,
			newFormGroup,
		)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.probe)
		buttonFormCallback.probe.formStage.Commit()
	}

	buttonFormCallback.probe.ux_tree()
}
func __gong__New__CellFormCallback(
	cell *models.Cell,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (cellFormCallback *CellFormCallback) OnSave() {
	cellFormCallback.probe.stageOfInterest.Lock()
	defer cellFormCallback.probe.stageOfInterest.Unlock()

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
			// WARNING : this form deals with the N-N association "Row.Cells []*Cell" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Cell). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Row
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Row"
				rf.Fieldname = "Cells"
				formerAssociationSource := cell_.GongGetReverseFieldOwner(
					cellFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Row)
					if !ok {
						log.Fatalln("Source of Row.Cells []*Cell, is not an Row instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Cells, cell_)
					formerSource.Cells = slices.Delete(formerSource.Cells, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Row
			for _row := range *models.GetGongstructInstancesSet[models.Row](cellFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _row.GetName() == newSourceName.GetName() {
					newSource = _row // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Row.Cells []*Cell, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Cells = append(newSource.Cells, cell_)
		}
	}

	// manage the suppress operation
	if cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cell_.Unstage(cellFormCallback.probe.stageOfInterest)
	}

	cellFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Cell](
		cellFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cellFormCallback.CreationMode || cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	cellFormCallback.probe.ux_tree()
}
func __gong__New__CellBooleanFormCallback(
	cellboolean *models.CellBoolean,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (cellbooleanFormCallback *CellBooleanFormCallback) OnSave() {
	cellbooleanFormCallback.probe.stageOfInterest.Lock()
	defer cellbooleanFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.CellBoolean](
		cellbooleanFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cellbooleanFormCallback.CreationMode || cellbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellbooleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	cellbooleanFormCallback.probe.ux_tree()
}
func __gong__New__CellFloat64FormCallback(
	cellfloat64 *models.CellFloat64,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (cellfloat64FormCallback *CellFloat64FormCallback) OnSave() {
	cellfloat64FormCallback.probe.stageOfInterest.Lock()
	defer cellfloat64FormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.CellFloat64](
		cellfloat64FormCallback.probe,
	)

	// display a new form by reset the form stage
	if cellfloat64FormCallback.CreationMode || cellfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellfloat64FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	cellfloat64FormCallback.probe.ux_tree()
}
func __gong__New__CellIconFormCallback(
	cellicon *models.CellIcon,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (celliconFormCallback *CellIconFormCallback) OnSave() {
	celliconFormCallback.probe.stageOfInterest.Lock()
	defer celliconFormCallback.probe.stageOfInterest.Unlock()

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
		case "NeedsConfirmation":
			FormDivBasicFieldToField(&(cellicon_.NeedsConfirmation), formDiv)
		case "ConfirmationMessage":
			FormDivBasicFieldToField(&(cellicon_.ConfirmationMessage), formDiv)
		}
	}

	// manage the suppress operation
	if celliconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellicon_.Unstage(celliconFormCallback.probe.stageOfInterest)
	}

	celliconFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.CellIcon](
		celliconFormCallback.probe,
	)

	// display a new form by reset the form stage
	if celliconFormCallback.CreationMode || celliconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		celliconFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	celliconFormCallback.probe.ux_tree()
}
func __gong__New__CellIntFormCallback(
	cellint *models.CellInt,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (cellintFormCallback *CellIntFormCallback) OnSave() {
	cellintFormCallback.probe.stageOfInterest.Lock()
	defer cellintFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.CellInt](
		cellintFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cellintFormCallback.CreationMode || cellintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellintFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	cellintFormCallback.probe.ux_tree()
}
func __gong__New__CellStringFormCallback(
	cellstring *models.CellString,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (cellstringFormCallback *CellStringFormCallback) OnSave() {
	cellstringFormCallback.probe.stageOfInterest.Lock()
	defer cellstringFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.CellString](
		cellstringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cellstringFormCallback.CreationMode || cellstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	cellstringFormCallback.probe.ux_tree()
}
func __gong__New__DisplayedColumnFormCallback(
	displayedcolumn *models.DisplayedColumn,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (displayedcolumnFormCallback *DisplayedColumnFormCallback) OnSave() {
	displayedcolumnFormCallback.probe.stageOfInterest.Lock()
	defer displayedcolumnFormCallback.probe.stageOfInterest.Unlock()

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
			// WARNING : this form deals with the N-N association "Table.DisplayedColumns []*DisplayedColumn" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DisplayedColumn). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Table
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Table"
				rf.Fieldname = "DisplayedColumns"
				formerAssociationSource := displayedcolumn_.GongGetReverseFieldOwner(
					displayedcolumnFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Table)
					if !ok {
						log.Fatalln("Source of Table.DisplayedColumns []*DisplayedColumn, is not an Table instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DisplayedColumns, displayedcolumn_)
					formerSource.DisplayedColumns = slices.Delete(formerSource.DisplayedColumns, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Table
			for _table := range *models.GetGongstructInstancesSet[models.Table](displayedcolumnFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _table.GetName() == newSourceName.GetName() {
					newSource = _table // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Table.DisplayedColumns []*DisplayedColumn, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DisplayedColumns = append(newSource.DisplayedColumns, displayedcolumn_)
		}
	}

	// manage the suppress operation
	if displayedcolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayedcolumn_.Unstage(displayedcolumnFormCallback.probe.stageOfInterest)
	}

	displayedcolumnFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DisplayedColumn](
		displayedcolumnFormCallback.probe,
	)

	// display a new form by reset the form stage
	if displayedcolumnFormCallback.CreationMode || displayedcolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayedcolumnFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	displayedcolumnFormCallback.probe.ux_tree()
}
func __gong__New__RowFormCallback(
	row *models.Row,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (rowFormCallback *RowFormCallback) OnSave() {
	rowFormCallback.probe.stageOfInterest.Lock()
	defer rowFormCallback.probe.stageOfInterest.Unlock()

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
		case "Cells":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Cell](rowFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Cell, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Cell)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rowFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Cell](rowFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			row_.Cells = instanceSlice

		case "IsChecked":
			FormDivBasicFieldToField(&(row_.IsChecked), formDiv)
		case "Table:Rows":
			// WARNING : this form deals with the N-N association "Table.Rows []*Row" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Row). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Table
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Table"
				rf.Fieldname = "Rows"
				formerAssociationSource := row_.GongGetReverseFieldOwner(
					rowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Table)
					if !ok {
						log.Fatalln("Source of Table.Rows []*Row, is not an Table instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Rows, row_)
					formerSource.Rows = slices.Delete(formerSource.Rows, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Table
			for _table := range *models.GetGongstructInstancesSet[models.Table](rowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _table.GetName() == newSourceName.GetName() {
					newSource = _table // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Table.Rows []*Row, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Rows = append(newSource.Rows, row_)
		case "Table:RowsSelectedForBulkDelete":
			// WARNING : this form deals with the N-N association "Table.RowsSelectedForBulkDelete []*Row" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Row). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Table
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Table"
				rf.Fieldname = "RowsSelectedForBulkDelete"
				formerAssociationSource := row_.GongGetReverseFieldOwner(
					rowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Table)
					if !ok {
						log.Fatalln("Source of Table.RowsSelectedForBulkDelete []*Row, is not an Table instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RowsSelectedForBulkDelete, row_)
					formerSource.RowsSelectedForBulkDelete = slices.Delete(formerSource.RowsSelectedForBulkDelete, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Table
			for _table := range *models.GetGongstructInstancesSet[models.Table](rowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _table.GetName() == newSourceName.GetName() {
					newSource = _table // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Table.RowsSelectedForBulkDelete []*Row, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RowsSelectedForBulkDelete = append(newSource.RowsSelectedForBulkDelete, row_)
		}
	}

	// manage the suppress operation
	if rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		row_.Unstage(rowFormCallback.probe.stageOfInterest)
	}

	rowFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Row](
		rowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rowFormCallback.CreationMode || rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rowFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	rowFormCallback.probe.ux_tree()
}
func __gong__New__SVGIconFormCallback(
	svgicon *models.SVGIcon,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgiconFormCallback *SVGIconFormCallback) {
	svgiconFormCallback = new(SVGIconFormCallback)
	svgiconFormCallback.probe = probe
	svgiconFormCallback.svgicon = svgicon
	svgiconFormCallback.formGroup = formGroup

	svgiconFormCallback.CreationMode = (svgicon == nil)

	return
}

type SVGIconFormCallback struct {
	svgicon *models.SVGIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (svgiconFormCallback *SVGIconFormCallback) OnSave() {
	svgiconFormCallback.probe.stageOfInterest.Lock()
	defer svgiconFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SVGIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgiconFormCallback.probe.formStage.Checkout()

	if svgiconFormCallback.svgicon == nil {
		svgiconFormCallback.svgicon = new(models.SVGIcon).Stage(svgiconFormCallback.probe.stageOfInterest)
	}
	svgicon_ := svgiconFormCallback.svgicon
	_ = svgicon_

	for _, formDiv := range svgiconFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(svgicon_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if svgiconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgicon_.Unstage(svgiconFormCallback.probe.stageOfInterest)
	}

	svgiconFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SVGIcon](
		svgiconFormCallback.probe,
	)

	// display a new form by reset the form stage
	if svgiconFormCallback.CreationMode || svgiconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgiconFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(svgiconFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SVGIconFormCallback(
			nil,
			svgiconFormCallback.probe,
			newFormGroup,
		)
		svgicon := new(models.SVGIcon)
		FillUpForm(svgicon, newFormGroup, svgiconFormCallback.probe)
		svgiconFormCallback.probe.formStage.Commit()
	}

	svgiconFormCallback.probe.ux_tree()
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (tableFormCallback *TableFormCallback) OnSave() {
	tableFormCallback.probe.stageOfInterest.Lock()
	defer tableFormCallback.probe.stageOfInterest.Unlock()

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
		case "DisplayedColumns":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DisplayedColumn](tableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DisplayedColumn, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DisplayedColumn)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DisplayedColumn](tableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			table_.DisplayedColumns = instanceSlice

		case "Rows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Row](tableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Row, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Row)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Row](tableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			table_.Rows = instanceSlice

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
		case "SaveButtonLabel":
			FormDivBasicFieldToField(&(table_.SaveButtonLabel), formDiv)
		case "HasBulkDeleteButton":
			FormDivBasicFieldToField(&(table_.HasBulkDeleteButton), formDiv)
		case "BulkDeleteButtonTooltip":
			FormDivBasicFieldToField(&(table_.BulkDeleteButtonTooltip), formDiv)
		case "RowsSelectedForBulkDelete":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Row](tableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Row, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Row)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Row](tableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			table_.RowsSelectedForBulkDelete = instanceSlice

		case "CanDragDropRows":
			FormDivBasicFieldToField(&(table_.CanDragDropRows), formDiv)
		case "HasCloseButton":
			FormDivBasicFieldToField(&(table_.HasCloseButton), formDiv)
		case "SavingInProgress":
			FormDivBasicFieldToField(&(table_.SavingInProgress), formDiv)
		case "NbOfStickyColumns":
			FormDivBasicFieldToField(&(table_.NbOfStickyColumns), formDiv)
		case "Buttons":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Button](tableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Button, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Button)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Button](tableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			table_.Buttons = instanceSlice

		}
	}

	// manage the suppress operation
	if tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		table_.Unstage(tableFormCallback.probe.stageOfInterest)
	}

	tableFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Table](
		tableFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode || tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tableFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	tableFormCallback.probe.ux_tree()
}
