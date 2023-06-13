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

	// insertion point for other declarations
	Links?: Array<GongLinkDB>
}
