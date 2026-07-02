// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellIntAPI {

	static GONGSTRUCT_NAME = "CellInt"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for other decls

	CellIntPointersEncoding: CellIntPointersEncoding = new CellIntPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class CellIntPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
