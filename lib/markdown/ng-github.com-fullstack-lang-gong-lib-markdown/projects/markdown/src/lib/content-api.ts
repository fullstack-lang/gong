// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ContentAPI {

	static GONGSTRUCT_NAME = "Content"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other decls

	ContentPointersEncoding: ContentPointersEncoding = new ContentPointersEncoding
}

export class ContentPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
