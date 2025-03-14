// generated code - do not edit

import { XLSheetAPI } from './xlsheet-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLRow } from './xlrow'
import { XLCell } from './xlcell'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLSheet {

	static GONGSTRUCT_NAME = "XLSheet"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	MaxRow: number = 0
	MaxCol: number = 0
	NbRows: number = 0

	// insertion point for pointers and slices of pointers declarations
	Rows: Array<XLRow> = []
	SheetCells: Array<XLCell> = []
}

export function CopyXLSheetToXLSheetAPI(xlsheet: XLSheet, xlsheetAPI: XLSheetAPI) {

	xlsheetAPI.CreatedAt = xlsheet.CreatedAt
	xlsheetAPI.DeletedAt = xlsheet.DeletedAt
	xlsheetAPI.ID = xlsheet.ID

	// insertion point for basic fields copy operations
	xlsheetAPI.Name = xlsheet.Name
	xlsheetAPI.MaxRow = xlsheet.MaxRow
	xlsheetAPI.MaxCol = xlsheet.MaxCol
	xlsheetAPI.NbRows = xlsheet.NbRows

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlsheetAPI.XLSheetPointersEncoding.Rows = []
	for (let _xlrow of xlsheet.Rows) {
		xlsheetAPI.XLSheetPointersEncoding.Rows.push(_xlrow.ID)
	}

	xlsheetAPI.XLSheetPointersEncoding.SheetCells = []
	for (let _xlcell of xlsheet.SheetCells) {
		xlsheetAPI.XLSheetPointersEncoding.SheetCells.push(_xlcell.ID)
	}

}

// CopyXLSheetAPIToXLSheet update basic, pointers and slice of pointers fields of xlsheet
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlsheetAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLSheetAPIToXLSheet(xlsheetAPI: XLSheetAPI, xlsheet: XLSheet, frontRepo: FrontRepo) {

	xlsheet.CreatedAt = xlsheetAPI.CreatedAt
	xlsheet.DeletedAt = xlsheetAPI.DeletedAt
	xlsheet.ID = xlsheetAPI.ID

	// insertion point for basic fields copy operations
	xlsheet.Name = xlsheetAPI.Name
	xlsheet.MaxRow = xlsheetAPI.MaxRow
	xlsheet.MaxCol = xlsheetAPI.MaxCol
	xlsheet.NbRows = xlsheetAPI.NbRows

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(xlsheetAPI.XLSheetPointersEncoding.Rows)) {
		console.error('Rects is not an array:', xlsheetAPI.XLSheetPointersEncoding.Rows);
		return;
	}

	xlsheet.Rows = new Array<XLRow>()
	for (let _id of xlsheetAPI.XLSheetPointersEncoding.Rows) {
		let _xlrow = frontRepo.map_ID_XLRow.get(_id)
		if (_xlrow != undefined) {
			xlsheet.Rows.push(_xlrow!)
		}
	}
	if (!Array.isArray(xlsheetAPI.XLSheetPointersEncoding.SheetCells)) {
		console.error('Rects is not an array:', xlsheetAPI.XLSheetPointersEncoding.SheetCells);
		return;
	}

	xlsheet.SheetCells = new Array<XLCell>()
	for (let _id of xlsheetAPI.XLSheetPointersEncoding.SheetCells) {
		let _xlcell = frontRepo.map_ID_XLCell.get(_id)
		if (_xlcell != undefined) {
			xlsheet.SheetCells.push(_xlcell!)
		}
	}
}
