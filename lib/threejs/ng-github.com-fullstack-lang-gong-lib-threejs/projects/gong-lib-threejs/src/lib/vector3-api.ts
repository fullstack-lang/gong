// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Vector3API {

	static GONGSTRUCT_NAME = "Vector3"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0

	// insertion point for other decls

	Vector3PointersEncoding: Vector3PointersEncoding = new Vector3PointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class Vector3PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
