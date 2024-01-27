// generated code - do not edit

import { CellStringDB } from './cellstring-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellString {

	static GONGSTRUCT_NAME = "CellString"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellStringToCellStringDB(cellstring: CellString, cellstringDB: CellStringDB) {

	cellstringDB.CreatedAt = cellstring.CreatedAt
	cellstringDB.DeletedAt = cellstring.DeletedAt
	cellstringDB.ID = cellstring.ID

	// insertion point for basic fields copy operations
	cellstringDB.Name = cellstring.Name
	cellstringDB.Value = cellstring.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellStringDBToCellString update basic, pointers and slice of pointers fields of cellstring
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellstringDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellStringDBToCellString(cellstringDB: CellStringDB, cellstring: CellString, frontRepo: FrontRepo) {

	cellstring.CreatedAt = cellstringDB.CreatedAt
	cellstring.DeletedAt = cellstringDB.DeletedAt
	cellstring.ID = cellstringDB.ID

	// insertion point for basic fields copy operations
	cellstring.Name = cellstringDB.Name
	cellstring.Value = cellstringDB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
