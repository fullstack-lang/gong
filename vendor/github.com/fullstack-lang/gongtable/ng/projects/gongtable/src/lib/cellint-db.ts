// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellIntDB {

	static GONGSTRUCT_NAME = "CellInt"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for pointers and slices of pointers declarations

	CellIntPointersEncoding: CellIntPointersEncoding = new CellIntPointersEncoding
}

export class CellIntPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
