// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

// insertion point
func __gong__New__DisplaySelectionFormCallback(
	displayselection *models.DisplaySelection,
	probe *Probe,
	formGroup *table.FormGroup,
) (displayselectionFormCallback *DisplaySelectionFormCallback) {
	displayselectionFormCallback = new(DisplaySelectionFormCallback)
	displayselectionFormCallback.probe = probe
	displayselectionFormCallback.displayselection = displayselection
	displayselectionFormCallback.formGroup = formGroup

	displayselectionFormCallback.CreationMode = (displayselection == nil)

	return
}

type DisplaySelectionFormCallback struct {
	displayselection *models.DisplaySelection

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (displayselectionFormCallback *DisplaySelectionFormCallback) OnSave() {

	// log.Println("DisplaySelectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayselectionFormCallback.probe.formStage.Checkout()

	if displayselectionFormCallback.displayselection == nil {
		displayselectionFormCallback.displayselection = new(models.DisplaySelection).Stage(displayselectionFormCallback.probe.stageOfInterest)
	}
	displayselection_ := displayselectionFormCallback.displayselection
	_ = displayselection_

	for _, formDiv := range displayselectionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayselection_.Name), formDiv)
		case "XLFile":
			FormDivSelectFieldToField(&(displayselection_.XLFile), displayselectionFormCallback.probe.stageOfInterest, formDiv)
		case "XLSheet":
			FormDivSelectFieldToField(&(displayselection_.XLSheet), displayselectionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if displayselectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayselection_.Unstage(displayselectionFormCallback.probe.stageOfInterest)
	}

	displayselectionFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.DisplaySelection](
		displayselectionFormCallback.probe,
	)
	displayselectionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if displayselectionFormCallback.CreationMode || displayselectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayselectionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(displayselectionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DisplaySelectionFormCallback(
			nil,
			displayselectionFormCallback.probe,
			newFormGroup,
		)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, newFormGroup, displayselectionFormCallback.probe)
		displayselectionFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(displayselectionFormCallback.probe)
}
func __gong__New__XLCellFormCallback(
	xlcell *models.XLCell,
	probe *Probe,
	formGroup *table.FormGroup,
) (xlcellFormCallback *XLCellFormCallback) {
	xlcellFormCallback = new(XLCellFormCallback)
	xlcellFormCallback.probe = probe
	xlcellFormCallback.xlcell = xlcell
	xlcellFormCallback.formGroup = formGroup

	xlcellFormCallback.CreationMode = (xlcell == nil)

	return
}

type XLCellFormCallback struct {
	xlcell *models.XLCell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xlcellFormCallback *XLCellFormCallback) OnSave() {

	// log.Println("XLCellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlcellFormCallback.probe.formStage.Checkout()

	if xlcellFormCallback.xlcell == nil {
		xlcellFormCallback.xlcell = new(models.XLCell).Stage(xlcellFormCallback.probe.stageOfInterest)
	}
	xlcell_ := xlcellFormCallback.xlcell
	_ = xlcell_

	for _, formDiv := range xlcellFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlcell_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(xlcell_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(xlcell_.Y), formDiv)
		case "XLRow:Cells":
			// WARNING : this form deals with the N-N association "XLRow.Cells []*XLCell" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of XLCell). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.XLRow
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "XLRow"
				rf.Fieldname = "Cells"
				formerAssociationSource := models.GetReverseFieldOwner(
					xlcellFormCallback.probe.stageOfInterest,
					xlcell_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.XLRow)
					if !ok {
						log.Fatalln("Source of XLRow.Cells []*XLCell, is not an XLRow instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Cells, xlcell_)
					formerSource.Cells = slices.Delete(formerSource.Cells, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Cells, xlcell_)
				formerSource.Cells = slices.Delete(formerSource.Cells, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.XLRow
			for _xlrow := range *models.GetGongstructInstancesSet[models.XLRow](xlcellFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _xlrow.GetName() == newSourceName.GetName() {
					newSource = _xlrow // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of XLRow.Cells []*XLCell, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Cells = append(newSource.Cells, xlcell_)
		case "XLSheet:SheetCells":
			// WARNING : this form deals with the N-N association "XLSheet.SheetCells []*XLCell" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of XLCell). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.XLSheet
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "XLSheet"
				rf.Fieldname = "SheetCells"
				formerAssociationSource := models.GetReverseFieldOwner(
					xlcellFormCallback.probe.stageOfInterest,
					xlcell_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.XLSheet)
					if !ok {
						log.Fatalln("Source of XLSheet.SheetCells []*XLCell, is not an XLSheet instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.SheetCells, xlcell_)
					formerSource.SheetCells = slices.Delete(formerSource.SheetCells, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.SheetCells, xlcell_)
				formerSource.SheetCells = slices.Delete(formerSource.SheetCells, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.XLSheet
			for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlcellFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _xlsheet.GetName() == newSourceName.GetName() {
					newSource = _xlsheet // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of XLSheet.SheetCells []*XLCell, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SheetCells = append(newSource.SheetCells, xlcell_)
		}
	}

	// manage the suppress operation
	if xlcellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlcell_.Unstage(xlcellFormCallback.probe.stageOfInterest)
	}

	xlcellFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.XLCell](
		xlcellFormCallback.probe,
	)
	xlcellFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlcellFormCallback.CreationMode || xlcellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlcellFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xlcellFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XLCellFormCallback(
			nil,
			xlcellFormCallback.probe,
			newFormGroup,
		)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, newFormGroup, xlcellFormCallback.probe)
		xlcellFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xlcellFormCallback.probe)
}
func __gong__New__XLFileFormCallback(
	xlfile *models.XLFile,
	probe *Probe,
	formGroup *table.FormGroup,
) (xlfileFormCallback *XLFileFormCallback) {
	xlfileFormCallback = new(XLFileFormCallback)
	xlfileFormCallback.probe = probe
	xlfileFormCallback.xlfile = xlfile
	xlfileFormCallback.formGroup = formGroup

	xlfileFormCallback.CreationMode = (xlfile == nil)

	return
}

type XLFileFormCallback struct {
	xlfile *models.XLFile

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xlfileFormCallback *XLFileFormCallback) OnSave() {

	// log.Println("XLFileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlfileFormCallback.probe.formStage.Checkout()

	if xlfileFormCallback.xlfile == nil {
		xlfileFormCallback.xlfile = new(models.XLFile).Stage(xlfileFormCallback.probe.stageOfInterest)
	}
	xlfile_ := xlfileFormCallback.xlfile
	_ = xlfile_

	for _, formDiv := range xlfileFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlfile_.Name), formDiv)
		case "NbSheets":
			FormDivBasicFieldToField(&(xlfile_.NbSheets), formDiv)
		}
	}

	// manage the suppress operation
	if xlfileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlfile_.Unstage(xlfileFormCallback.probe.stageOfInterest)
	}

	xlfileFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.XLFile](
		xlfileFormCallback.probe,
	)
	xlfileFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlfileFormCallback.CreationMode || xlfileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlfileFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xlfileFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XLFileFormCallback(
			nil,
			xlfileFormCallback.probe,
			newFormGroup,
		)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, newFormGroup, xlfileFormCallback.probe)
		xlfileFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xlfileFormCallback.probe)
}
func __gong__New__XLRowFormCallback(
	xlrow *models.XLRow,
	probe *Probe,
	formGroup *table.FormGroup,
) (xlrowFormCallback *XLRowFormCallback) {
	xlrowFormCallback = new(XLRowFormCallback)
	xlrowFormCallback.probe = probe
	xlrowFormCallback.xlrow = xlrow
	xlrowFormCallback.formGroup = formGroup

	xlrowFormCallback.CreationMode = (xlrow == nil)

	return
}

type XLRowFormCallback struct {
	xlrow *models.XLRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xlrowFormCallback *XLRowFormCallback) OnSave() {

	// log.Println("XLRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlrowFormCallback.probe.formStage.Checkout()

	if xlrowFormCallback.xlrow == nil {
		xlrowFormCallback.xlrow = new(models.XLRow).Stage(xlrowFormCallback.probe.stageOfInterest)
	}
	xlrow_ := xlrowFormCallback.xlrow
	_ = xlrow_

	for _, formDiv := range xlrowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlrow_.Name), formDiv)
		case "RowIndex":
			FormDivBasicFieldToField(&(xlrow_.RowIndex), formDiv)
		case "XLSheet:Rows":
			// WARNING : this form deals with the N-N association "XLSheet.Rows []*XLRow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of XLRow). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.XLSheet
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "XLSheet"
				rf.Fieldname = "Rows"
				formerAssociationSource := models.GetReverseFieldOwner(
					xlrowFormCallback.probe.stageOfInterest,
					xlrow_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.XLSheet)
					if !ok {
						log.Fatalln("Source of XLSheet.Rows []*XLRow, is not an XLSheet instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Rows, xlrow_)
					formerSource.Rows = slices.Delete(formerSource.Rows, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Rows, xlrow_)
				formerSource.Rows = slices.Delete(formerSource.Rows, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.XLSheet
			for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlrowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _xlsheet.GetName() == newSourceName.GetName() {
					newSource = _xlsheet // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of XLSheet.Rows []*XLRow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Rows = append(newSource.Rows, xlrow_)
		}
	}

	// manage the suppress operation
	if xlrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlrow_.Unstage(xlrowFormCallback.probe.stageOfInterest)
	}

	xlrowFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.XLRow](
		xlrowFormCallback.probe,
	)
	xlrowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlrowFormCallback.CreationMode || xlrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xlrowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XLRowFormCallback(
			nil,
			xlrowFormCallback.probe,
			newFormGroup,
		)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, newFormGroup, xlrowFormCallback.probe)
		xlrowFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xlrowFormCallback.probe)
}
func __gong__New__XLSheetFormCallback(
	xlsheet *models.XLSheet,
	probe *Probe,
	formGroup *table.FormGroup,
) (xlsheetFormCallback *XLSheetFormCallback) {
	xlsheetFormCallback = new(XLSheetFormCallback)
	xlsheetFormCallback.probe = probe
	xlsheetFormCallback.xlsheet = xlsheet
	xlsheetFormCallback.formGroup = formGroup

	xlsheetFormCallback.CreationMode = (xlsheet == nil)

	return
}

type XLSheetFormCallback struct {
	xlsheet *models.XLSheet

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xlsheetFormCallback *XLSheetFormCallback) OnSave() {

	// log.Println("XLSheetFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlsheetFormCallback.probe.formStage.Checkout()

	if xlsheetFormCallback.xlsheet == nil {
		xlsheetFormCallback.xlsheet = new(models.XLSheet).Stage(xlsheetFormCallback.probe.stageOfInterest)
	}
	xlsheet_ := xlsheetFormCallback.xlsheet
	_ = xlsheet_

	for _, formDiv := range xlsheetFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlsheet_.Name), formDiv)
		case "MaxRow":
			FormDivBasicFieldToField(&(xlsheet_.MaxRow), formDiv)
		case "MaxCol":
			FormDivBasicFieldToField(&(xlsheet_.MaxCol), formDiv)
		case "NbRows":
			FormDivBasicFieldToField(&(xlsheet_.NbRows), formDiv)
		case "XLFile:Sheets":
			// WARNING : this form deals with the N-N association "XLFile.Sheets []*XLSheet" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of XLSheet). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.XLFile
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "XLFile"
				rf.Fieldname = "Sheets"
				formerAssociationSource := models.GetReverseFieldOwner(
					xlsheetFormCallback.probe.stageOfInterest,
					xlsheet_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.XLFile)
					if !ok {
						log.Fatalln("Source of XLFile.Sheets []*XLSheet, is not an XLFile instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Sheets, xlsheet_)
					formerSource.Sheets = slices.Delete(formerSource.Sheets, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Sheets, xlsheet_)
				formerSource.Sheets = slices.Delete(formerSource.Sheets, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.XLFile
			for _xlfile := range *models.GetGongstructInstancesSet[models.XLFile](xlsheetFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _xlfile.GetName() == newSourceName.GetName() {
					newSource = _xlfile // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of XLFile.Sheets []*XLSheet, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sheets = append(newSource.Sheets, xlsheet_)
		}
	}

	// manage the suppress operation
	if xlsheetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsheet_.Unstage(xlsheetFormCallback.probe.stageOfInterest)
	}

	xlsheetFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.XLSheet](
		xlsheetFormCallback.probe,
	)
	xlsheetFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlsheetFormCallback.CreationMode || xlsheetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsheetFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xlsheetFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XLSheetFormCallback(
			nil,
			xlsheetFormCallback.probe,
			newFormGroup,
		)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, newFormGroup, xlsheetFormCallback.probe)
		xlsheetFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xlsheetFormCallback.probe)
}
