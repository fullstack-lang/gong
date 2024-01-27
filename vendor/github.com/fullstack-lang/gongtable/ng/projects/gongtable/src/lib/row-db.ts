// insertion point for imports
import { CellDB } from './cell-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RowDB {

	static GONGSTRUCT_NAME = "Row"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsChecked: boolean = false

	// insertion point for other decls

	RowPointersEncoding: RowPointersEncoding = new RowPointersEncoding
}

export class RowPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Cells: number[] = []
}
