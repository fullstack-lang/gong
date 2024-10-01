// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GstructAPI {

	static GONGSTRUCT_NAME = "Gstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for other decls

	GstructPointersEncoding: GstructPointersEncoding = new GstructPointersEncoding
}

export class GstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
