// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValueDB {

	static GONGSTRUCT_NAME = "GongEnumValue"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for pointers and slices of pointers declarations

	GongEnumValuePointersEncoding: GongEnumValuePointersEncoding = new GongEnumValuePointersEncoding
}

export class GongEnumValuePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
