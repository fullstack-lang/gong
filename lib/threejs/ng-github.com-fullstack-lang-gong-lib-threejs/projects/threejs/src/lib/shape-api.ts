// insertion point for imports
import { Vector2API } from './vector2-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ShapeAPI {

	static GONGSTRUCT_NAME = "Shape"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	ShapePointersEncoding: ShapePointersEncoding = new ShapePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ShapePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Points: number[] = []
}
