// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Vector2API {

	static GONGSTRUCT_NAME = "Vector2"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for other decls

	Vector2PointersEncoding: Vector2PointersEncoding = new Vector2PointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class Vector2PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
