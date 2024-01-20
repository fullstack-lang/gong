// generated code - do not edit

import { CellIconDB } from './cellicon-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellIcon {

	static GONGSTRUCT_NAME = "CellIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellIconToCellIconDB(cellicon: CellIcon, celliconDB: CellIconDB) {

	celliconDB.CreatedAt = cellicon.CreatedAt
	celliconDB.DeletedAt = cellicon.DeletedAt
	celliconDB.ID = cellicon.ID
	
	// insertion point for basic fields copy operations
	celliconDB.Name = cellicon.Name
	celliconDB.Icon = cellicon.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyCellIconDBToCellIcon(celliconDB: CellIconDB, cellicon: CellIcon, frontRepo: FrontRepo) {

	cellicon.CreatedAt = celliconDB.CreatedAt
	cellicon.DeletedAt = celliconDB.DeletedAt
	cellicon.ID = celliconDB.ID
	
	// insertion point for basic fields copy operations
	cellicon.Name = celliconDB.Name
	cellicon.Icon = celliconDB.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
