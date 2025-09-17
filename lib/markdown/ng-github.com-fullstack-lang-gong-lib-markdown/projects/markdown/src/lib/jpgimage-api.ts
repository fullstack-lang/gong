// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class JpgImageAPI {

	static GONGSTRUCT_NAME = "JpgImage"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for other decls

	JpgImagePointersEncoding: JpgImagePointersEncoding = new JpgImagePointersEncoding
}

export class JpgImagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
