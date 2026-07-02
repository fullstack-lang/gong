// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BoxGeometryAPI {

	static GONGSTRUCT_NAME = "BoxGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	Depth: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0
	DepthSegments: number = 0

	// insertion point for other decls

	BoxGeometryPointersEncoding: BoxGeometryPointersEncoding = new BoxGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class BoxGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
