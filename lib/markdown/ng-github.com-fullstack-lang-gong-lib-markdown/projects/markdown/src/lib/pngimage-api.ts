// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PngImageAPI {

	static GONGSTRUCT_NAME = "PngImage"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for other decls

	PngImagePointersEncoding: PngImagePointersEncoding = new PngImagePointersEncoding
}

export class PngImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
