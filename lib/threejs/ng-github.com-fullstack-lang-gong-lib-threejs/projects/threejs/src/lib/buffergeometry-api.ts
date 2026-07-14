// insertion point for imports
import { Vector3API } from './vector3-api'
import { TriangleAPI } from './triangle-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BufferGeometryAPI {

	static GONGSTRUCT_NAME = "BufferGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	BufferGeometryPointersEncoding: BufferGeometryPointersEncoding = new BufferGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class BufferGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Vertices: number[] = []
	Faces: number[] = []
}
