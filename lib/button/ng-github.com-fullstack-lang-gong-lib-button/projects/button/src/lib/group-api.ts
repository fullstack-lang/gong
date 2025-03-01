// insertion point for imports
import { ButtonAPI } from './button-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GroupAPI {

	static GONGSTRUCT_NAME = "Group"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Percentage: number = 0

	// insertion point for other decls

	GroupPointersEncoding: GroupPointersEncoding = new GroupPointersEncoding
}

export class GroupPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Buttons: number[] = []
}
