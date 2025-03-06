// insertion point for imports
import { AsSplitAreaAPI } from './assplitarea-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplitAPI {

	static GONGSTRUCT_NAME = "AsSplit"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Direction: string = ""

	// insertion point for other decls

	AsSplitPointersEncoding: AsSplitPointersEncoding = new AsSplitPointersEncoding
}

export class AsSplitPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	AsSplitAreas: number[] = []
}
