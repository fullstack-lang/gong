// generated code - do not edit

import { SplitAreaAPI } from './splitarea-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SplitArea {

	static GONGSTRUCT_NAME = "SplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySplitAreaToSplitAreaAPI(splitarea: SplitArea, splitareaAPI: SplitAreaAPI) {

	splitareaAPI.CreatedAt = splitarea.CreatedAt
	splitareaAPI.DeletedAt = splitarea.DeletedAt
	splitareaAPI.ID = splitarea.ID

	// insertion point for basic fields copy operations
	splitareaAPI.Name = splitarea.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySplitAreaAPIToSplitArea update basic, pointers and slice of pointers fields of splitarea
// from respectively the basic fields and encoded fields of pointers and slices of pointers of splitareaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySplitAreaAPIToSplitArea(splitareaAPI: SplitAreaAPI, splitarea: SplitArea, frontRepo: FrontRepo) {

	splitarea.CreatedAt = splitareaAPI.CreatedAt
	splitarea.DeletedAt = splitareaAPI.DeletedAt
	splitarea.ID = splitareaAPI.ID

	// insertion point for basic fields copy operations
	splitarea.Name = splitareaAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
