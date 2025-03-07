// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableAPI {

	static GONGSTRUCT_NAME = "Table"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	TableName: string = ""

	// insertion point for other decls

	TablePointersEncoding: TablePointersEncoding = new TablePointersEncoding
}

export class TablePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
