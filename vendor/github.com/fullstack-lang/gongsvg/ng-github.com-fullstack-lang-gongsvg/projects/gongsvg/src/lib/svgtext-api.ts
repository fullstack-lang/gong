// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SvgTextAPI {

	static GONGSTRUCT_NAME = "SvgText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""

	// insertion point for other decls

	SvgTextPointersEncoding: SvgTextPointersEncoding = new SvgTextPointersEncoding
}

export class SvgTextPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
