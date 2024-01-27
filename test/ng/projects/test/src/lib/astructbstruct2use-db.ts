// insertion point for imports
import { BstructDB } from './bstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstruct2UseDB {

	static GONGSTRUCT_NAME = "AstructBstruct2Use"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	AstructBstruct2UsePointersEncoding: AstructBstruct2UsePointersEncoding = new AstructBstruct2UsePointersEncoding
}

export class AstructBstruct2UsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Bstrcut2ID: NullInt64 = new NullInt64 // if pointer is null, Bstrcut2.ID = 0

}
