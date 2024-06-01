// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongTimeFieldAPI {

	static GONGSTRUCT_NAME = "GongTimeField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for other decls

	GongTimeFieldPointersEncoding: GongTimeFieldPointersEncoding = new GongTimeFieldPointersEncoding
}

export class GongTimeFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
