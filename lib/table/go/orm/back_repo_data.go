// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ButtonAPIs []*ButtonAPI

	CellAPIs []*CellAPI

	CellBooleanAPIs []*CellBooleanAPI

	CellFloat64APIs []*CellFloat64API

	CellIconAPIs []*CellIconAPI

	CellIntAPIs []*CellIntAPI

	CellStringAPIs []*CellStringAPI

	DisplayedColumnAPIs []*DisplayedColumnAPI

	RowAPIs []*RowAPI

	SVGIconAPIs []*SVGIconAPI

	TableAPIs []*TableAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, buttonDB := range backRepo.BackRepoButton.Map_ButtonDBID_ButtonDB {

		var buttonAPI ButtonAPI
		buttonAPI.ID = buttonDB.ID
		buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
		buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)

		backRepoData.ButtonAPIs = append(backRepoData.ButtonAPIs, &buttonAPI)
	}

	for _, cellDB := range backRepo.BackRepoCell.Map_CellDBID_CellDB {

		var cellAPI CellAPI
		cellAPI.ID = cellDB.ID
		cellAPI.CellPointersEncoding = cellDB.CellPointersEncoding
		cellDB.CopyBasicFieldsToCell_WOP(&cellAPI.Cell_WOP)

		backRepoData.CellAPIs = append(backRepoData.CellAPIs, &cellAPI)
	}

	for _, cellbooleanDB := range backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {

		var cellbooleanAPI CellBooleanAPI
		cellbooleanAPI.ID = cellbooleanDB.ID
		cellbooleanAPI.CellBooleanPointersEncoding = cellbooleanDB.CellBooleanPointersEncoding
		cellbooleanDB.CopyBasicFieldsToCellBoolean_WOP(&cellbooleanAPI.CellBoolean_WOP)

		backRepoData.CellBooleanAPIs = append(backRepoData.CellBooleanAPIs, &cellbooleanAPI)
	}

	for _, cellfloat64DB := range backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64DB {

		var cellfloat64API CellFloat64API
		cellfloat64API.ID = cellfloat64DB.ID
		cellfloat64API.CellFloat64PointersEncoding = cellfloat64DB.CellFloat64PointersEncoding
		cellfloat64DB.CopyBasicFieldsToCellFloat64_WOP(&cellfloat64API.CellFloat64_WOP)

		backRepoData.CellFloat64APIs = append(backRepoData.CellFloat64APIs, &cellfloat64API)
	}

	for _, celliconDB := range backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconDB {

		var celliconAPI CellIconAPI
		celliconAPI.ID = celliconDB.ID
		celliconAPI.CellIconPointersEncoding = celliconDB.CellIconPointersEncoding
		celliconDB.CopyBasicFieldsToCellIcon_WOP(&celliconAPI.CellIcon_WOP)

		backRepoData.CellIconAPIs = append(backRepoData.CellIconAPIs, &celliconAPI)
	}

	for _, cellintDB := range backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntDB {

		var cellintAPI CellIntAPI
		cellintAPI.ID = cellintDB.ID
		cellintAPI.CellIntPointersEncoding = cellintDB.CellIntPointersEncoding
		cellintDB.CopyBasicFieldsToCellInt_WOP(&cellintAPI.CellInt_WOP)

		backRepoData.CellIntAPIs = append(backRepoData.CellIntAPIs, &cellintAPI)
	}

	for _, cellstringDB := range backRepo.BackRepoCellString.Map_CellStringDBID_CellStringDB {

		var cellstringAPI CellStringAPI
		cellstringAPI.ID = cellstringDB.ID
		cellstringAPI.CellStringPointersEncoding = cellstringDB.CellStringPointersEncoding
		cellstringDB.CopyBasicFieldsToCellString_WOP(&cellstringAPI.CellString_WOP)

		backRepoData.CellStringAPIs = append(backRepoData.CellStringAPIs, &cellstringAPI)
	}

	for _, displayedcolumnDB := range backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {

		var displayedcolumnAPI DisplayedColumnAPI
		displayedcolumnAPI.ID = displayedcolumnDB.ID
		displayedcolumnAPI.DisplayedColumnPointersEncoding = displayedcolumnDB.DisplayedColumnPointersEncoding
		displayedcolumnDB.CopyBasicFieldsToDisplayedColumn_WOP(&displayedcolumnAPI.DisplayedColumn_WOP)

		backRepoData.DisplayedColumnAPIs = append(backRepoData.DisplayedColumnAPIs, &displayedcolumnAPI)
	}

	for _, rowDB := range backRepo.BackRepoRow.Map_RowDBID_RowDB {

		var rowAPI RowAPI
		rowAPI.ID = rowDB.ID
		rowAPI.RowPointersEncoding = rowDB.RowPointersEncoding
		rowDB.CopyBasicFieldsToRow_WOP(&rowAPI.Row_WOP)

		backRepoData.RowAPIs = append(backRepoData.RowAPIs, &rowAPI)
	}

	for _, svgiconDB := range backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {

		var svgiconAPI SVGIconAPI
		svgiconAPI.ID = svgiconDB.ID
		svgiconAPI.SVGIconPointersEncoding = svgiconDB.SVGIconPointersEncoding
		svgiconDB.CopyBasicFieldsToSVGIcon_WOP(&svgiconAPI.SVGIcon_WOP)

		backRepoData.SVGIconAPIs = append(backRepoData.SVGIconAPIs, &svgiconAPI)
	}

	for _, tableDB := range backRepo.BackRepoTable.Map_TableDBID_TableDB {

		var tableAPI TableAPI
		tableAPI.ID = tableDB.ID
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

		backRepoData.TableAPIs = append(backRepoData.TableAPIs, &tableAPI)
	}

}
