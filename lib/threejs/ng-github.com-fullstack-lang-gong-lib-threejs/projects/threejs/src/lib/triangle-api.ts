// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TriangleAPI {

	static GONGSTRUCT_NAME = "Triangle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	V1: number = 0
	V2: number = 0
	V3: number = 0

	// insertion point for other decls

	TrianglePointersEncoding: TrianglePointersEncoding = new TrianglePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TrianglePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
