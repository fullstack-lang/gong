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

// insertion point
func __gong__New__DisplaySelectionFormCallback(
	displayselection *models.DisplaySelection,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (displayselectionFormCallback *DisplaySelectionFormCallback) OnSave() {
	displayselectionFormCallback.probe.stageOfInterest.Lock()
	defer displayselectionFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.DisplaySelection](
		displayselectionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if displayselectionFormCallback.CreationMode || displayselectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		displayselectionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	displayselectionFormCallback.probe.ux_tree()
}
func __gong__New__XLCellFormCallback(
	xlcell *models.XLCell,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (xlcellFormCallback *XLCellFormCallback) OnSave() {
	xlcellFormCallback.probe.stageOfInterest.Lock()
	defer xlcellFormCallback.probe.stageOfInterest.Unlock()

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the XLRow instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target XLRow instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.XLRow](xlcellFormCallback.probe.stageOfInterest)
			targetXLRowIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetXLRowIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all XLRow instances and update their Cells slice
			for _xlrow := range *models.GetGongstructInstancesSetFromPointerType[*models.XLRow](xlcellFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(xlcellFormCallback.probe.stageOfInterest, _xlrow)
				
				// if XLRow is selected
				if targetXLRowIDs[id] {
					// ensure xlcell_ is in _xlrow.Cells
					found := false
					for _, _b := range _xlrow.Cells {
						if _b == xlcell_ {
							found = true
							break
						}
					}
					if !found {
						_xlrow.Cells = append(_xlrow.Cells, xlcell_)
						xlcellFormCallback.probe.UpdateSliceOfPointersCallback(_xlrow, "Cells", &_xlrow.Cells)
					}
				} else {
					// ensure xlcell_ is NOT in _xlrow.Cells
					idx := slices.Index(_xlrow.Cells, xlcell_)
					if idx != -1 {
						_xlrow.Cells = slices.Delete(_xlrow.Cells, idx, idx+1)
						xlcellFormCallback.probe.UpdateSliceOfPointersCallback(_xlrow, "Cells", &_xlrow.Cells)
					}
				}
			}
		case "XLSheet:SheetCells":
			// 1. Decode the AssociationStorage which contains the rowIDs of the XLSheet instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target XLSheet instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.XLSheet](xlcellFormCallback.probe.stageOfInterest)
			targetXLSheetIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetXLSheetIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all XLSheet instances and update their SheetCells slice
			for _xlsheet := range *models.GetGongstructInstancesSetFromPointerType[*models.XLSheet](xlcellFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(xlcellFormCallback.probe.stageOfInterest, _xlsheet)
				
				// if XLSheet is selected
				if targetXLSheetIDs[id] {
					// ensure xlcell_ is in _xlsheet.SheetCells
					found := false
					for _, _b := range _xlsheet.SheetCells {
						if _b == xlcell_ {
							found = true
							break
						}
					}
					if !found {
						_xlsheet.SheetCells = append(_xlsheet.SheetCells, xlcell_)
						xlcellFormCallback.probe.UpdateSliceOfPointersCallback(_xlsheet, "SheetCells", &_xlsheet.SheetCells)
					}
				} else {
					// ensure xlcell_ is NOT in _xlsheet.SheetCells
					idx := slices.Index(_xlsheet.SheetCells, xlcell_)
					if idx != -1 {
						_xlsheet.SheetCells = slices.Delete(_xlsheet.SheetCells, idx, idx+1)
						xlcellFormCallback.probe.UpdateSliceOfPointersCallback(_xlsheet, "SheetCells", &_xlsheet.SheetCells)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if xlcellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlcell_.Unstage(xlcellFormCallback.probe.stageOfInterest)
	}

	xlcellFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.XLCell](
		xlcellFormCallback.probe,
	)

	// display a new form by reset the form stage
	if xlcellFormCallback.CreationMode || xlcellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlcellFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	xlcellFormCallback.probe.ux_tree()
}
func __gong__New__XLFileFormCallback(
	xlfile *models.XLFile,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (xlfileFormCallback *XLFileFormCallback) OnSave() {
	xlfileFormCallback.probe.stageOfInterest.Lock()
	defer xlfileFormCallback.probe.stageOfInterest.Unlock()

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
		case "Sheets":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.XLSheet](xlfileFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.XLSheet, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.XLSheet)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					xlfileFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.XLSheet](xlfileFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			xlfile_.Sheets = instanceSlice
			xlfileFormCallback.probe.UpdateSliceOfPointersCallback(xlfile_, "Sheets", &xlfile_.Sheets)

		}
	}

	// manage the suppress operation
	if xlfileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlfile_.Unstage(xlfileFormCallback.probe.stageOfInterest)
	}

	xlfileFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.XLFile](
		xlfileFormCallback.probe,
	)

	// display a new form by reset the form stage
	if xlfileFormCallback.CreationMode || xlfileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlfileFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	xlfileFormCallback.probe.ux_tree()
}
func __gong__New__XLRowFormCallback(
	xlrow *models.XLRow,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (xlrowFormCallback *XLRowFormCallback) OnSave() {
	xlrowFormCallback.probe.stageOfInterest.Lock()
	defer xlrowFormCallback.probe.stageOfInterest.Unlock()

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
		case "Cells":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.XLCell](xlrowFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.XLCell, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.XLCell)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					xlrowFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.XLCell](xlrowFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			xlrow_.Cells = instanceSlice
			xlrowFormCallback.probe.UpdateSliceOfPointersCallback(xlrow_, "Cells", &xlrow_.Cells)

		case "XLSheet:Rows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the XLSheet instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target XLSheet instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.XLSheet](xlrowFormCallback.probe.stageOfInterest)
			targetXLSheetIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetXLSheetIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all XLSheet instances and update their Rows slice
			for _xlsheet := range *models.GetGongstructInstancesSetFromPointerType[*models.XLSheet](xlrowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(xlrowFormCallback.probe.stageOfInterest, _xlsheet)
				
				// if XLSheet is selected
				if targetXLSheetIDs[id] {
					// ensure xlrow_ is in _xlsheet.Rows
					found := false
					for _, _b := range _xlsheet.Rows {
						if _b == xlrow_ {
							found = true
							break
						}
					}
					if !found {
						_xlsheet.Rows = append(_xlsheet.Rows, xlrow_)
						xlrowFormCallback.probe.UpdateSliceOfPointersCallback(_xlsheet, "Rows", &_xlsheet.Rows)
					}
				} else {
					// ensure xlrow_ is NOT in _xlsheet.Rows
					idx := slices.Index(_xlsheet.Rows, xlrow_)
					if idx != -1 {
						_xlsheet.Rows = slices.Delete(_xlsheet.Rows, idx, idx+1)
						xlrowFormCallback.probe.UpdateSliceOfPointersCallback(_xlsheet, "Rows", &_xlsheet.Rows)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if xlrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlrow_.Unstage(xlrowFormCallback.probe.stageOfInterest)
	}

	xlrowFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.XLRow](
		xlrowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if xlrowFormCallback.CreationMode || xlrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	xlrowFormCallback.probe.ux_tree()
}
func __gong__New__XLSheetFormCallback(
	xlsheet *models.XLSheet,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (xlsheetFormCallback *XLSheetFormCallback) OnSave() {
	xlsheetFormCallback.probe.stageOfInterest.Lock()
	defer xlsheetFormCallback.probe.stageOfInterest.Unlock()

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
		case "Rows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.XLRow](xlsheetFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.XLRow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.XLRow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					xlsheetFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.XLRow](xlsheetFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			xlsheet_.Rows = instanceSlice
			xlsheetFormCallback.probe.UpdateSliceOfPointersCallback(xlsheet_, "Rows", &xlsheet_.Rows)

		case "SheetCells":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.XLCell](xlsheetFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.XLCell, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.XLCell)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					xlsheetFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.XLCell](xlsheetFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			xlsheet_.SheetCells = instanceSlice
			xlsheetFormCallback.probe.UpdateSliceOfPointersCallback(xlsheet_, "SheetCells", &xlsheet_.SheetCells)

		case "XLFile:Sheets":
			// 1. Decode the AssociationStorage which contains the rowIDs of the XLFile instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target XLFile instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.XLFile](xlsheetFormCallback.probe.stageOfInterest)
			targetXLFileIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetXLFileIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all XLFile instances and update their Sheets slice
			for _xlfile := range *models.GetGongstructInstancesSetFromPointerType[*models.XLFile](xlsheetFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(xlsheetFormCallback.probe.stageOfInterest, _xlfile)
				
				// if XLFile is selected
				if targetXLFileIDs[id] {
					// ensure xlsheet_ is in _xlfile.Sheets
					found := false
					for _, _b := range _xlfile.Sheets {
						if _b == xlsheet_ {
							found = true
							break
						}
					}
					if !found {
						_xlfile.Sheets = append(_xlfile.Sheets, xlsheet_)
						xlsheetFormCallback.probe.UpdateSliceOfPointersCallback(_xlfile, "Sheets", &_xlfile.Sheets)
					}
				} else {
					// ensure xlsheet_ is NOT in _xlfile.Sheets
					idx := slices.Index(_xlfile.Sheets, xlsheet_)
					if idx != -1 {
						_xlfile.Sheets = slices.Delete(_xlfile.Sheets, idx, idx+1)
						xlsheetFormCallback.probe.UpdateSliceOfPointersCallback(_xlfile, "Sheets", &_xlfile.Sheets)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if xlsheetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsheet_.Unstage(xlsheetFormCallback.probe.stageOfInterest)
	}

	xlsheetFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.XLSheet](
		xlsheetFormCallback.probe,
	)

	// display a new form by reset the form stage
	if xlsheetFormCallback.CreationMode || xlsheetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsheetFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	xlsheetFormCallback.probe.ux_tree()
}
