// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TorusGeometryAPI {

	static GONGSTRUCT_NAME = "TorusGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Radius: number = 0
	Tube: number = 0
	RadialSegments: number = 0
	TubularSegments: number = 0
	Arc: number = 0

	// insertion point for other decls

	TorusGeometryPointersEncoding: TorusGeometryPointersEncoding = new TorusGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TorusGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
