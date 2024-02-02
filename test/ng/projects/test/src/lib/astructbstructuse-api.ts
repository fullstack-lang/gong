// insertion point for imports
import { BstructAPI } from './bstruct-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstructUseAPI {

	static GONGSTRUCT_NAME = "AstructBstructUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	AstructBstructUsePointersEncoding: AstructBstructUsePointersEncoding = new AstructBstructUsePointersEncoding
}

export class AstructBstructUsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Bstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Bstruct2.ID = 0

}
