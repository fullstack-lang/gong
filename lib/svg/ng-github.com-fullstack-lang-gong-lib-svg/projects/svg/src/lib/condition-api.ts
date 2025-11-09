// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ConditionAPI {

	static GONGSTRUCT_NAME = "Condition"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	ConditionPointersEncoding: ConditionPointersEncoding = new ConditionPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ConditionPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
