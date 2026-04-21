// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredSvgImageAPI {

	static GONGSTRUCT_NAME = "RectAnchoredSvgImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other decls

	RectAnchoredSvgImagePointersEncoding: RectAnchoredSvgImagePointersEncoding = new RectAnchoredSvgImagePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class RectAnchoredSvgImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
