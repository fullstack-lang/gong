// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TreeAPI {

	static GONGSTRUCT_NAME = "Tree"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	TreeName: string = ""

	// insertion point for other decls

	TreePointersEncoding: TreePointersEncoding = new TreePointersEncoding
}

export class TreePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
