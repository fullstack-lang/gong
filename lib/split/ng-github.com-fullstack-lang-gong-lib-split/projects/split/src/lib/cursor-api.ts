// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CursorAPI {

	static GONGSTRUCT_NAME = "Cursor"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	Style: string = ""

	// insertion point for other decls

	CursorPointersEncoding: CursorPointersEncoding = new CursorPointersEncoding
}

export class CursorPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
