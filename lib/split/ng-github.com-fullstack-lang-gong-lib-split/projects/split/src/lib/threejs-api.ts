// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ThreejsAPI {

	static GONGSTRUCT_NAME = "Threejs"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	ThreejsPointersEncoding: ThreejsPointersEncoding = new ThreejsPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ThreejsPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
