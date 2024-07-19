// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BAPI {

	static GONGSTRUCT_NAME = "B"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	BPointersEncoding: BPointersEncoding = new BPointersEncoding
}

export class BPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
