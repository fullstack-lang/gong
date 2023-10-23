// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BstructDB {

	static GONGSTRUCT_NAME = "Bstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for pointers and slices of pointers declarations

	BstructPointersEncoding: BstructPointersEncoding = new BstructPointersEncoding
}

export class BstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
