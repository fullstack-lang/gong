// generated code - do not edit

import { CellStringAPI } from './cellstring-api'
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

export function CopyCellStringToCellStringAPI(cellstring: CellString, cellstringAPI: CellStringAPI) {

	cellstringAPI.CreatedAt = cellstring.CreatedAt
	cellstringAPI.DeletedAt = cellstring.DeletedAt
	cellstringAPI.ID = cellstring.ID

	// insertion point for basic fields copy operations
	cellstringAPI.Name = cellstring.Name
	cellstringAPI.Value = cellstring.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellStringAPIToCellString update basic, pointers and slice of pointers fields of cellstring
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellstringAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellStringAPIToCellString(cellstringAPI: CellStringAPI, cellstring: CellString, frontRepo: FrontRepo) {

	cellstring.CreatedAt = cellstringAPI.CreatedAt
	cellstring.DeletedAt = cellstringAPI.DeletedAt
	cellstring.ID = cellstringAPI.ID

	// insertion point for basic fields copy operations
	cellstring.Name = cellstringAPI.Name
	cellstring.Value = cellstringAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
