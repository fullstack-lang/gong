// generated code - do not edit

import { CellIconAPI } from './cellicon-api'
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

export function CopyCellIconToCellIconAPI(cellicon: CellIcon, celliconAPI: CellIconAPI) {

	celliconAPI.CreatedAt = cellicon.CreatedAt
	celliconAPI.DeletedAt = cellicon.DeletedAt
	celliconAPI.ID = cellicon.ID

	// insertion point for basic fields copy operations
	celliconAPI.Name = cellicon.Name
	celliconAPI.Icon = cellicon.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellIconAPIToCellIcon update basic, pointers and slice of pointers fields of cellicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of celliconAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellIconAPIToCellIcon(celliconAPI: CellIconAPI, cellicon: CellIcon, frontRepo: FrontRepo) {

	cellicon.CreatedAt = celliconAPI.CreatedAt
	cellicon.DeletedAt = celliconAPI.DeletedAt
	cellicon.ID = celliconAPI.ID

	// insertion point for basic fields copy operations
	cellicon.Name = celliconAPI.Name
	cellicon.Icon = celliconAPI.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
