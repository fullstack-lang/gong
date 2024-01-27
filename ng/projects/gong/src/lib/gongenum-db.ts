// insertion point for imports
import { GongEnumValueDB } from './gongenumvalue-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumDB {

	static GONGSTRUCT_NAME = "GongEnum"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: number = 0

	// insertion point for other decls
	Type_string?: string

	GongEnumPointersEncoding: GongEnumPointersEncoding = new GongEnumPointersEncoding
}

export class GongEnumPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongEnumValues: number[] = []
}
