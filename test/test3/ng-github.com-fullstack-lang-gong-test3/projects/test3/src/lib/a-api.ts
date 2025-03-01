// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AAPI {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	APointersEncoding: APointersEncoding = new APointersEncoding
}

export class APointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
