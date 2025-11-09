// insertion point for imports
import { AsSplitAreaAPI } from './assplitarea-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ViewAPI {

	static GONGSTRUCT_NAME = "View"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ShowViewName: boolean = false
	IsSelectedView: boolean = false

	// insertion point for other decls

	ViewPointersEncoding: ViewPointersEncoding = new ViewPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ViewPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	RootAsSplitAreas: number[] = []
}
