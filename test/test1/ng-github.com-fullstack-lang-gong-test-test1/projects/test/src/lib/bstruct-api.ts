// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BstructAPI {

	static GONGSTRUCT_NAME = "Bstruct"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for other decls

	BstructPointersEncoding: BstructPointersEncoding = new BstructPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class BstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
