// insertion point for imports
import { Vector3API } from './vector3-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CurveAPI {

	static GONGSTRUCT_NAME = "Curve"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	CurvePointersEncoding: CurvePointersEncoding = new CurvePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class CurvePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Points: number[] = []
}
