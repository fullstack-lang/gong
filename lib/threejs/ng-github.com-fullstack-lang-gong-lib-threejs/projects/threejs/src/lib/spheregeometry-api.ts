// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SphereGeometryAPI {

	static GONGSTRUCT_NAME = "SphereGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Radius: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0
	PhiStart: number = 0
	PhiLength: number = 0
	ThetaStart: number = 0
	ThetaLength: number = 0

	// insertion point for other decls

	SphereGeometryPointersEncoding: SphereGeometryPointersEncoding = new SphereGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class SphereGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
