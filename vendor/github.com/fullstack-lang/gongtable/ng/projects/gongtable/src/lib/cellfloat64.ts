// generated code - do not edit

import { CellFloat64API } from './cellfloat64-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellFloat64 {

	static GONGSTRUCT_NAME = "CellFloat64"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellFloat64ToCellFloat64API(cellfloat64: CellFloat64, cellfloat64API: CellFloat64API) {

	cellfloat64API.CreatedAt = cellfloat64.CreatedAt
	cellfloat64API.DeletedAt = cellfloat64.DeletedAt
	cellfloat64API.ID = cellfloat64.ID

	// insertion point for basic fields copy operations
	cellfloat64API.Name = cellfloat64.Name
	cellfloat64API.Value = cellfloat64.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellFloat64APIToCellFloat64 update basic, pointers and slice of pointers fields of cellfloat64
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellfloat64API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellFloat64APIToCellFloat64(cellfloat64API: CellFloat64API, cellfloat64: CellFloat64, frontRepo: FrontRepo) {

	cellfloat64.CreatedAt = cellfloat64API.CreatedAt
	cellfloat64.DeletedAt = cellfloat64API.DeletedAt
	cellfloat64.ID = cellfloat64API.ID

	// insertion point for basic fields copy operations
	cellfloat64.Name = cellfloat64API.Name
	cellfloat64.Value = cellfloat64API.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
