// insertion point for imports
import { BstructAPI } from './bstruct-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstruct2UseAPI {

	static GONGSTRUCT_NAME = "AstructBstruct2Use"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	AstructBstruct2UsePointersEncoding: AstructBstruct2UsePointersEncoding = new AstructBstruct2UsePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class AstructBstruct2UsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Bstrcut2ID: NullInt64 = new NullInt64 // if pointer is null, Bstrcut2.ID = 0

}
