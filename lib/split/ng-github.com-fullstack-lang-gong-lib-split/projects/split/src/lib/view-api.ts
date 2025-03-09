// insertion point for imports
import { AsSplitAreaAPI } from './assplitarea-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ViewAPI {

	static GONGSTRUCT_NAME = "View"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	ViewPointersEncoding: ViewPointersEncoding = new ViewPointersEncoding
}

export class ViewPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	RootAsSplitAreas: number[] = []
}
