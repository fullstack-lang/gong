// generated code - do not edit

import { CellBooleanAPI } from './cellboolean-api'
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

export function CopyCellBooleanToCellBooleanAPI(cellboolean: CellBoolean, cellbooleanAPI: CellBooleanAPI) {

	cellbooleanAPI.CreatedAt = cellboolean.CreatedAt
	cellbooleanAPI.DeletedAt = cellboolean.DeletedAt
	cellbooleanAPI.ID = cellboolean.ID

	// insertion point for basic fields copy operations
	cellbooleanAPI.Name = cellboolean.Name
	cellbooleanAPI.Value = cellboolean.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellBooleanAPIToCellBoolean update basic, pointers and slice of pointers fields of cellboolean
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellbooleanAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellBooleanAPIToCellBoolean(cellbooleanAPI: CellBooleanAPI, cellboolean: CellBoolean, frontRepo: FrontRepo) {

	cellboolean.CreatedAt = cellbooleanAPI.CreatedAt
	cellboolean.DeletedAt = cellbooleanAPI.DeletedAt
	cellboolean.ID = cellbooleanAPI.ID

	// insertion point for basic fields copy operations
	cellboolean.Name = cellbooleanAPI.Name
	cellboolean.Value = cellbooleanAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
