// insertion point for imports
import { RectAPI } from './rect-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ControlPointAPI {

	static GONGSTRUCT_NAME = "ControlPoint"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X_Relative: number = 0
	Y_Relative: number = 0

	// insertion point for other decls

	ControlPointPointersEncoding: ControlPointPointersEncoding = new ControlPointPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ControlPointPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ClosestRectID: NullInt64 = new NullInt64 // if pointer is null, ClosestRect.ID = 0

}
