// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGIconAPI {

	static GONGSTRUCT_NAME = "SVGIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for other decls

	SVGIconPointersEncoding: SVGIconPointersEncoding = new SVGIconPointersEncoding
}

export class SVGIconPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
