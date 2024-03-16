// generated code - do not edit

import { DisplayedColumnAPI } from './displayedcolumn-api'
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

export function CopyDisplayedColumnToDisplayedColumnAPI(displayedcolumn: DisplayedColumn, displayedcolumnAPI: DisplayedColumnAPI) {

	displayedcolumnAPI.CreatedAt = displayedcolumn.CreatedAt
	displayedcolumnAPI.DeletedAt = displayedcolumn.DeletedAt
	displayedcolumnAPI.ID = displayedcolumn.ID

	// insertion point for basic fields copy operations
	displayedcolumnAPI.Name = displayedcolumn.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDisplayedColumnAPIToDisplayedColumn update basic, pointers and slice of pointers fields of displayedcolumn
// from respectively the basic fields and encoded fields of pointers and slices of pointers of displayedcolumnAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDisplayedColumnAPIToDisplayedColumn(displayedcolumnAPI: DisplayedColumnAPI, displayedcolumn: DisplayedColumn, frontRepo: FrontRepo) {

	displayedcolumn.CreatedAt = displayedcolumnAPI.CreatedAt
	displayedcolumn.DeletedAt = displayedcolumnAPI.DeletedAt
	displayedcolumn.ID = displayedcolumnAPI.ID

	// insertion point for basic fields copy operations
	displayedcolumn.Name = displayedcolumnAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
