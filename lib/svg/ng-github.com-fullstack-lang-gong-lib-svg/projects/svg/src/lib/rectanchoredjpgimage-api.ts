// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredJpgImageAPI {

	static GONGSTRUCT_NAME = "RectAnchoredJpgImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for other decls

	RectAnchoredJpgImagePointersEncoding: RectAnchoredJpgImagePointersEncoding = new RectAnchoredJpgImagePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class RectAnchoredJpgImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
