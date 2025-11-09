// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableAPI {

	static GONGSTRUCT_NAME = "Table"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	TablePointersEncoding: TablePointersEncoding = new TablePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TablePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
