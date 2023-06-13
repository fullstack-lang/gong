// insertion point for imports
import { MetaReferenceDB } from './metareference-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MetaDB {

	static GONGSTRUCT_NAME = "Meta"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""

	// insertion point for other declarations
	MetaReferences?: Array<MetaReferenceDB>
}
