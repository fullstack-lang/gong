// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValueAPI {

	static GONGSTRUCT_NAME = "GongEnumValue"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for other decls

	GongEnumValuePointersEncoding: GongEnumValuePointersEncoding = new GongEnumValuePointersEncoding
}

export class GongEnumValuePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
