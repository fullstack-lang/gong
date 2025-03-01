// insertion point for imports
import { GroupAPI } from './group-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayoutAPI {

	static GONGSTRUCT_NAME = "Layout"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	LayoutPointersEncoding: LayoutPointersEncoding = new LayoutPointersEncoding
}

export class LayoutPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Groups: number[] = []
}
