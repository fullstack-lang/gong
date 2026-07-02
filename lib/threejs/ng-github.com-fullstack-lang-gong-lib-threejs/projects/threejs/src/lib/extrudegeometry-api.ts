// insertion point for imports
import { ShapeAPI } from './shape-api'
import { CurveAPI } from './curve-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ExtrudeGeometryAPI {

	static GONGSTRUCT_NAME = "ExtrudeGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Steps: number = 0

	// insertion point for other decls

	ExtrudeGeometryPointersEncoding: ExtrudeGeometryPointersEncoding = new ExtrudeGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ExtrudeGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeID: NullInt64 = new NullInt64 // if pointer is null, Shape.ID = 0

	ExtrudePathID: NullInt64 = new NullInt64 // if pointer is null, ExtrudePath.ID = 0

}
