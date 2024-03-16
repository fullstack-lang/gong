// insertion point for imports
import { GongLinkAPI } from './gonglink-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongNoteAPI {

	static GONGSTRUCT_NAME = "GongNote"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Body: string = ""
	BodyHTML: string = ""

	// insertion point for other decls

	GongNotePointersEncoding: GongNotePointersEncoding = new GongNotePointersEncoding
}

export class GongNotePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Links: number[] = []
}
