// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FavIconAPI {

	static GONGSTRUCT_NAME = "FavIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for other decls

	FavIconPointersEncoding: FavIconPointersEncoding = new FavIconPointersEncoding
}

export class FavIconPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
