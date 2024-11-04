// generated code - do not edit

import { RowAPI } from './row-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Cell } from './cell'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Row {

	static GONGSTRUCT_NAME = "Row"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsChecked: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Cells: Array<Cell> = []
}

export function CopyRowToRowAPI(row: Row, rowAPI: RowAPI) {

	rowAPI.CreatedAt = row.CreatedAt
	rowAPI.DeletedAt = row.DeletedAt
	rowAPI.ID = row.ID

	// insertion point for basic fields copy operations
	rowAPI.Name = row.Name
	rowAPI.IsChecked = row.IsChecked

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rowAPI.RowPointersEncoding.Cells = []
	for (let _cell of row.Cells) {
		rowAPI.RowPointersEncoding.Cells.push(_cell.ID)
	}

}

// CopyRowAPIToRow update basic, pointers and slice of pointers fields of row
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rowAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRowAPIToRow(rowAPI: RowAPI, row: Row, frontRepo: FrontRepo) {

	row.CreatedAt = rowAPI.CreatedAt
	row.DeletedAt = rowAPI.DeletedAt
	row.ID = rowAPI.ID

	// insertion point for basic fields copy operations
	row.Name = rowAPI.Name
	row.IsChecked = rowAPI.IsChecked

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(rowAPI.RowPointersEncoding.Cells)) {
		console.error('Rects is not an array:', rowAPI.RowPointersEncoding.Cells);
		return;
	}

	row.Cells = new Array<Cell>()
	for (let _id of rowAPI.RowPointersEncoding.Cells) {
		let _cell = frontRepo.map_ID_Cell.get(_id)
		if (_cell != undefined) {
			row.Cells.push(_cell!)
		}
	}
}
