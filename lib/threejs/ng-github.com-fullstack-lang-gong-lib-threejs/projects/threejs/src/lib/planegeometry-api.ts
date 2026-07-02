// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PlaneGeometryAPI {

	static GONGSTRUCT_NAME = "PlaneGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0

	// insertion point for other decls

	PlaneGeometryPointersEncoding: PlaneGeometryPointersEncoding = new PlaneGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class PlaneGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
