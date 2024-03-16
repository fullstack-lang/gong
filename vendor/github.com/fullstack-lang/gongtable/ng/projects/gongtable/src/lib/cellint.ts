// generated code - do not edit

import { CellIntAPI } from './cellint-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellInt {

	static GONGSTRUCT_NAME = "CellInt"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellIntToCellIntAPI(cellint: CellInt, cellintAPI: CellIntAPI) {

	cellintAPI.CreatedAt = cellint.CreatedAt
	cellintAPI.DeletedAt = cellint.DeletedAt
	cellintAPI.ID = cellint.ID

	// insertion point for basic fields copy operations
	cellintAPI.Name = cellint.Name
	cellintAPI.Value = cellint.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellIntAPIToCellInt update basic, pointers and slice of pointers fields of cellint
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellintAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellIntAPIToCellInt(cellintAPI: CellIntAPI, cellint: CellInt, frontRepo: FrontRepo) {

	cellint.CreatedAt = cellintAPI.CreatedAt
	cellint.DeletedAt = cellintAPI.DeletedAt
	cellint.ID = cellintAPI.ID

	// insertion point for basic fields copy operations
	cellint.Name = cellintAPI.Name
	cellint.Value = cellintAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
