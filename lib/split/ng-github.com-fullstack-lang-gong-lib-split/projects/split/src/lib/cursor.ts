// generated code - do not edit

import { CursorAPI } from './cursor-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Cursor {

	static GONGSTRUCT_NAME = "Cursor"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCursorToCursorAPI(cursor: Cursor, cursorAPI: CursorAPI) {

	cursorAPI.CreatedAt = cursor.CreatedAt
	cursorAPI.DeletedAt = cursor.DeletedAt
	cursorAPI.ID = cursor.ID

	// insertion point for basic fields copy operations
	cursorAPI.Name = cursor.Name
	cursorAPI.StackName = cursor.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCursorAPIToCursor update basic, pointers and slice of pointers fields of cursor
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cursorAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCursorAPIToCursor(cursorAPI: CursorAPI, cursor: Cursor, frontRepo: FrontRepo) {

	cursor.CreatedAt = cursorAPI.CreatedAt
	cursor.DeletedAt = cursorAPI.DeletedAt
	cursor.ID = cursorAPI.ID

	// insertion point for basic fields copy operations
	cursor.Name = cursorAPI.Name
	cursor.StackName = cursorAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
