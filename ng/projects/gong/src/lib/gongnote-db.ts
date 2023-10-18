// insertion point for imports
import { GongLinkDB } from './gonglink-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongNoteDB {

	static GONGSTRUCT_NAME = "GongNote"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Body: string = ""
	BodyHTML: string = ""

	// insertion point for pointers and slices of pointers declarations
	Links: Array<GongLinkDB> = []

	GongNotePointersEncoding: GongNotePointersEncoding = new GongNotePointersEncoding
}

export class GongNotePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Links: number[] = []
}
