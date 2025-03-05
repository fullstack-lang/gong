// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SplitAreaAPI {

	static GONGSTRUCT_NAME = "SplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	SplitAreaPointersEncoding: SplitAreaPointersEncoding = new SplitAreaPointersEncoding
}

export class SplitAreaPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
