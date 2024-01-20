// generated code - do not edit

import { CellIntDB } from './cellint-db'
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

export function CopyCellIntToCellIntDB(cellint: CellInt, cellintDB: CellIntDB) {

	cellintDB.CreatedAt = cellint.CreatedAt
	cellintDB.DeletedAt = cellint.DeletedAt
	cellintDB.ID = cellint.ID
	
	// insertion point for basic fields copy operations
	cellintDB.Name = cellint.Name
	cellintDB.Value = cellint.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyCellIntDBToCellInt(cellintDB: CellIntDB, cellint: CellInt, frontRepo: FrontRepo) {

	cellint.CreatedAt = cellintDB.CreatedAt
	cellint.DeletedAt = cellintDB.DeletedAt
	cellint.ID = cellintDB.ID
	
	// insertion point for basic fields copy operations
	cellint.Name = cellintDB.Name
	cellint.Value = cellintDB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
