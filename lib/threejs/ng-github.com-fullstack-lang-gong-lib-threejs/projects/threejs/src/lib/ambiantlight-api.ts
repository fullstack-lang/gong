// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AmbiantLightAPI {

	static GONGSTRUCT_NAME = "AmbiantLight"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Intensity: number = 0

	// insertion point for other decls

	AmbiantLightPointersEncoding: AmbiantLightPointersEncoding = new AmbiantLightPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class AmbiantLightPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
