// generated code - do not edit

import { RowDB } from './row-db'
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

export function CopyRowToRowDB(row: Row, rowDB: RowDB) {

	rowDB.CreatedAt = row.CreatedAt
	rowDB.DeletedAt = row.DeletedAt
	rowDB.ID = row.ID

	// insertion point for basic fields copy operations
	rowDB.Name = row.Name
	rowDB.IsChecked = row.IsChecked

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rowDB.RowPointersEncoding.Cells = []
	for (let _cell of row.Cells) {
		rowDB.RowPointersEncoding.Cells.push(_cell.ID)
	}

}

// CopyRowDBToRow update basic, pointers and slice of pointers fields of row
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rowDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRowDBToRow(rowDB: RowDB, row: Row, frontRepo: FrontRepo) {

	row.CreatedAt = rowDB.CreatedAt
	row.DeletedAt = rowDB.DeletedAt
	row.ID = rowDB.ID

	// insertion point for basic fields copy operations
	row.Name = rowDB.Name
	row.IsChecked = rowDB.IsChecked

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	row.Cells = new Array<Cell>()
	for (let _id of rowDB.RowPointersEncoding.Cells) {
		let _cell = frontRepo.map_ID_Cell.get(_id)
		if (_cell != undefined) {
			row.Cells.push(_cell!)
		}
	}
}
