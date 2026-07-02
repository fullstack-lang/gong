// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellIconAPI {

	static GONGSTRUCT_NAME = "CellIcon"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""
	NeedsConfirmation: boolean = false
	ConfirmationMessage: string = ""

	// insertion point for other decls

	CellIconPointersEncoding: CellIconPointersEncoding = new CellIconPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class CellIconPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
