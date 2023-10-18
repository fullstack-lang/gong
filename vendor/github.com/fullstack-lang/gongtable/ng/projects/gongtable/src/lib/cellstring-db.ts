// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellStringDB {

	static GONGSTRUCT_NAME = "CellString"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for pointers and slices of pointers declarations

	CellStringPointersEncoding: CellStringPointersEncoding = new CellStringPointersEncoding
}

export class CellStringPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
