// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredPngImageAPI {

	static GONGSTRUCT_NAME = "RectAnchoredPngImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for other decls

	RectAnchoredPngImagePointersEncoding: RectAnchoredPngImagePointersEncoding = new RectAnchoredPngImagePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class RectAnchoredPngImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
