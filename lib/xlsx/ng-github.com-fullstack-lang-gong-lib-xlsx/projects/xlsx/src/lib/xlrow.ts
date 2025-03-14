// generated code - do not edit

import { XLRowAPI } from './xlrow-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLCell } from './xlcell'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLRow {

	static GONGSTRUCT_NAME = "XLRow"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RowIndex: number = 0

	// insertion point for pointers and slices of pointers declarations
	Cells: Array<XLCell> = []
}

export function CopyXLRowToXLRowAPI(xlrow: XLRow, xlrowAPI: XLRowAPI) {

	xlrowAPI.CreatedAt = xlrow.CreatedAt
	xlrowAPI.DeletedAt = xlrow.DeletedAt
	xlrowAPI.ID = xlrow.ID

	// insertion point for basic fields copy operations
	xlrowAPI.Name = xlrow.Name
	xlrowAPI.RowIndex = xlrow.RowIndex

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlrowAPI.XLRowPointersEncoding.Cells = []
	for (let _xlcell of xlrow.Cells) {
		xlrowAPI.XLRowPointersEncoding.Cells.push(_xlcell.ID)
	}

}

// CopyXLRowAPIToXLRow update basic, pointers and slice of pointers fields of xlrow
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlrowAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLRowAPIToXLRow(xlrowAPI: XLRowAPI, xlrow: XLRow, frontRepo: FrontRepo) {

	xlrow.CreatedAt = xlrowAPI.CreatedAt
	xlrow.DeletedAt = xlrowAPI.DeletedAt
	xlrow.ID = xlrowAPI.ID

	// insertion point for basic fields copy operations
	xlrow.Name = xlrowAPI.Name
	xlrow.RowIndex = xlrowAPI.RowIndex

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(xlrowAPI.XLRowPointersEncoding.Cells)) {
		console.error('Rects is not an array:', xlrowAPI.XLRowPointersEncoding.Cells);
		return;
	}

	xlrow.Cells = new Array<XLCell>()
	for (let _id of xlrowAPI.XLRowPointersEncoding.Cells) {
		let _xlcell = frontRepo.map_ID_XLCell.get(_id)
		if (_xlcell != undefined) {
			xlrow.Cells.push(_xlcell!)
		}
	}
}
