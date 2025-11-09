// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SvgImageAPI {

	static GONGSTRUCT_NAME = "SvgImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other decls

	SvgImagePointersEncoding: SvgImagePointersEncoding = new SvgImagePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class SvgImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
