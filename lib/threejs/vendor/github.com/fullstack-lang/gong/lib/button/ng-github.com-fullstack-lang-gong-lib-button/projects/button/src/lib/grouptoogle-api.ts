// insertion point for imports
import { ButtonToggleAPI } from './buttontoggle-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GroupToogleAPI {

	static GONGSTRUCT_NAME = "GroupToogle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Percentage: number = 0
	IsSingleSelector: boolean = false

	// insertion point for other decls

	GroupTooglePointersEncoding: GroupTooglePointersEncoding = new GroupTooglePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class GroupTooglePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ButtonToggles: number[] = []
}
