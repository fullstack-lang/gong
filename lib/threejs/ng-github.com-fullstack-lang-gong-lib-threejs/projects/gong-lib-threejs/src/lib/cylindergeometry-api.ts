// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CylinderGeometryAPI {

	static GONGSTRUCT_NAME = "CylinderGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RadiusTop: number = 0
	RadiusBottom: number = 0
	Height: number = 0
	RadialSegments: number = 0
	HeightSegments: number = 0
	OpenEnded: boolean = false
	ThetaStart: number = 0
	ThetaLength: number = 0

	// insertion point for other decls

	CylinderGeometryPointersEncoding: CylinderGeometryPointersEncoding = new CylinderGeometryPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class CylinderGeometryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
