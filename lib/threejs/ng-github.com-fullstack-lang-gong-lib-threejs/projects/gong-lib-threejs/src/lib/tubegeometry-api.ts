// insertion point for imports
import { CurveAPI } from './curve-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TubeGeometryAPI {

	static GONGSTRUCT_NAME = "TubeGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	TubularSegments: number = 0
	Radius: number = 0
	RadialSegments: number = 0
	Closed: boolean = false

	// insertion point for other decls

	TubeGeometryPointersEncoding: TubeGeometryPointersEncoding = new TubeGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TubeGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	PathID: NullInt64 = new NullInt64 // if pointer is null, Path.ID = 0

}
