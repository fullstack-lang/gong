// insertion point for imports
import { BAPI } from './b-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AAPI {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NumberField: number = 0

	// insertion point for other decls

	APointersEncoding: APointersEncoding = new APointersEncoding
}

export class APointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	BID: NullInt64 = new NullInt64 // if pointer is null, B.ID = 0

	Bs: number[] = []
}
