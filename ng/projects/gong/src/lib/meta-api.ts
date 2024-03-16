// insertion point for imports
import { MetaReferenceAPI } from './metareference-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MetaAPI {

	static GONGSTRUCT_NAME = "Meta"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""

	// insertion point for other decls

	MetaPointersEncoding: MetaPointersEncoding = new MetaPointersEncoding
}

export class MetaPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	MetaReferences: number[] = []
}
