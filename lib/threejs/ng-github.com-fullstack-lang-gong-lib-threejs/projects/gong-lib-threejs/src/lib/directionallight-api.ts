// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DirectionalLightAPI {

	static GONGSTRUCT_NAME = "DirectionalLight"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0
	Intensity: number = 0
	IsWithCastShadow: boolean = false

	// insertion point for other decls

	DirectionalLightPointersEncoding: DirectionalLightPointersEncoding = new DirectionalLightPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class DirectionalLightPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
