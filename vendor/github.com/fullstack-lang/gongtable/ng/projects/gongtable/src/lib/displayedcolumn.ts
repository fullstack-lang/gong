// generated code - do not edit

import { DisplayedColumnDB } from './displayedcolumn-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplayedColumn {

	static GONGSTRUCT_NAME = "DisplayedColumn"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDisplayedColumnToDisplayedColumnDB(displayedcolumn: DisplayedColumn, displayedcolumnDB: DisplayedColumnDB) {

	displayedcolumnDB.CreatedAt = displayedcolumn.CreatedAt
	displayedcolumnDB.DeletedAt = displayedcolumn.DeletedAt
	displayedcolumnDB.ID = displayedcolumn.ID

	// insertion point for basic fields copy operations
	displayedcolumnDB.Name = displayedcolumn.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDisplayedColumnDBToDisplayedColumn update basic, pointers and slice of pointers fields of displayedcolumn
// from respectively the basic fields and encoded fields of pointers and slices of pointers of displayedcolumnDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDisplayedColumnDBToDisplayedColumn(displayedcolumnDB: DisplayedColumnDB, displayedcolumn: DisplayedColumn, frontRepo: FrontRepo) {

	displayedcolumn.CreatedAt = displayedcolumnDB.CreatedAt
	displayedcolumn.DeletedAt = displayedcolumnDB.DeletedAt
	displayedcolumn.ID = displayedcolumnDB.ID

	// insertion point for basic fields copy operations
	displayedcolumn.Name = displayedcolumnDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
