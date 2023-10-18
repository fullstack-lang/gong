// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellIconDB {

	static GONGSTRUCT_NAME = "CellIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations

	CellIconPointersEncoding: CellIconPointersEncoding = new CellIconPointersEncoding
}

export class CellIconPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
