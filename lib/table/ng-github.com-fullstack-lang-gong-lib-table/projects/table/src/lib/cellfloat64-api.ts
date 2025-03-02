// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellFloat64API {

	static GONGSTRUCT_NAME = "CellFloat64"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0

	// insertion point for other decls

	CellFloat64PointersEncoding: CellFloat64PointersEncoding = new CellFloat64PointersEncoding
}

export class CellFloat64PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
