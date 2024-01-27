// generated code - do not edit

import { CellBooleanDB } from './cellboolean-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellBoolean {

	static GONGSTRUCT_NAME = "CellBoolean"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellBooleanToCellBooleanDB(cellboolean: CellBoolean, cellbooleanDB: CellBooleanDB) {

	cellbooleanDB.CreatedAt = cellboolean.CreatedAt
	cellbooleanDB.DeletedAt = cellboolean.DeletedAt
	cellbooleanDB.ID = cellboolean.ID

	// insertion point for basic fields copy operations
	cellbooleanDB.Name = cellboolean.Name
	cellbooleanDB.Value = cellboolean.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellBooleanDBToCellBoolean update basic, pointers and slice of pointers fields of cellboolean
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellbooleanDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellBooleanDBToCellBoolean(cellbooleanDB: CellBooleanDB, cellboolean: CellBoolean, frontRepo: FrontRepo) {

	cellboolean.CreatedAt = cellbooleanDB.CreatedAt
	cellboolean.DeletedAt = cellbooleanDB.DeletedAt
	cellboolean.ID = cellbooleanDB.ID

	// insertion point for basic fields copy operations
	cellboolean.Name = cellbooleanDB.Name
	cellboolean.Value = cellbooleanDB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
