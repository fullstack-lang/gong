// insertion point for imports
import { BstructAPI } from './bstruct-api'
import { GstructAPI } from './gstruct-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DstructAPI {

	static GONGSTRUCT_NAME = "Dstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	DstructPointersEncoding: DstructPointersEncoding = new DstructPointersEncoding
}

export class DstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Anarrayofb: number[] = []
	GstructID: NullInt64 = new NullInt64 // if pointer is null, Gstruct.ID = 0

	Gstructs: number[] = []
}
