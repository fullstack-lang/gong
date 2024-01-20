// generated code - do not edit

import { CellFloat64DB } from './cellfloat64-db'
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

export function CopyCellFloat64ToCellFloat64DB(cellfloat64: CellFloat64, cellfloat64DB: CellFloat64DB) {

	cellfloat64DB.CreatedAt = cellfloat64.CreatedAt
	cellfloat64DB.DeletedAt = cellfloat64.DeletedAt
	cellfloat64DB.ID = cellfloat64.ID
	
	// insertion point for basic fields copy operations
	cellfloat64DB.Name = cellfloat64.Name
	cellfloat64DB.Value = cellfloat64.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyCellFloat64DBToCellFloat64(cellfloat64DB: CellFloat64DB, cellfloat64: CellFloat64, frontRepo: FrontRepo) {

	cellfloat64.CreatedAt = cellfloat64DB.CreatedAt
	cellfloat64.DeletedAt = cellfloat64DB.DeletedAt
	cellfloat64.ID = cellfloat64DB.ID
	
	// insertion point for basic fields copy operations
	cellfloat64.Name = cellfloat64DB.Name
	cellfloat64.Value = cellfloat64DB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
