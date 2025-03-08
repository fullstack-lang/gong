// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DocAPI {

	static GONGSTRUCT_NAME = "Doc"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	DocPointersEncoding: DocPointersEncoding = new DocPointersEncoding
}

export class DocPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
