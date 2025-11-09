// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PointAPI {

	static GONGSTRUCT_NAME = "Point"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for other decls

	PointPointersEncoding: PointPointersEncoding = new PointPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class PointPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
