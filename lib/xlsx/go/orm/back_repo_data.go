// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	DisplaySelectionAPIs []*DisplaySelectionAPI

	XLCellAPIs []*XLCellAPI

	XLFileAPIs []*XLFileAPI

	XLRowAPIs []*XLRowAPI

	XLSheetAPIs []*XLSheetAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, displayselectionDB := range backRepo.BackRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB {

		var displayselectionAPI DisplaySelectionAPI
		displayselectionAPI.ID = displayselectionDB.ID
		displayselectionAPI.DisplaySelectionPointersEncoding = displayselectionDB.DisplaySelectionPointersEncoding
		displayselectionDB.CopyBasicFieldsToDisplaySelection_WOP(&displayselectionAPI.DisplaySelection_WOP)

		backRepoData.DisplaySelectionAPIs = append(backRepoData.DisplaySelectionAPIs, &displayselectionAPI)
	}

	for _, xlcellDB := range backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellDB {

		var xlcellAPI XLCellAPI
		xlcellAPI.ID = xlcellDB.ID
		xlcellAPI.XLCellPointersEncoding = xlcellDB.XLCellPointersEncoding
		xlcellDB.CopyBasicFieldsToXLCell_WOP(&xlcellAPI.XLCell_WOP)

		backRepoData.XLCellAPIs = append(backRepoData.XLCellAPIs, &xlcellAPI)
	}

	for _, xlfileDB := range backRepo.BackRepoXLFile.Map_XLFileDBID_XLFileDB {

		var xlfileAPI XLFileAPI
		xlfileAPI.ID = xlfileDB.ID
		xlfileAPI.XLFilePointersEncoding = xlfileDB.XLFilePointersEncoding
		xlfileDB.CopyBasicFieldsToXLFile_WOP(&xlfileAPI.XLFile_WOP)

		backRepoData.XLFileAPIs = append(backRepoData.XLFileAPIs, &xlfileAPI)
	}

	for _, xlrowDB := range backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowDB {

		var xlrowAPI XLRowAPI
		xlrowAPI.ID = xlrowDB.ID
		xlrowAPI.XLRowPointersEncoding = xlrowDB.XLRowPointersEncoding
		xlrowDB.CopyBasicFieldsToXLRow_WOP(&xlrowAPI.XLRow_WOP)

		backRepoData.XLRowAPIs = append(backRepoData.XLRowAPIs, &xlrowAPI)
	}

	for _, xlsheetDB := range backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetDB {

		var xlsheetAPI XLSheetAPI
		xlsheetAPI.ID = xlsheetDB.ID
		xlsheetAPI.XLSheetPointersEncoding = xlsheetDB.XLSheetPointersEncoding
		xlsheetDB.CopyBasicFieldsToXLSheet_WOP(&xlsheetAPI.XLSheet_WOP)

		backRepoData.XLSheetAPIs = append(backRepoData.XLSheetAPIs, &xlsheetAPI)
	}

}
