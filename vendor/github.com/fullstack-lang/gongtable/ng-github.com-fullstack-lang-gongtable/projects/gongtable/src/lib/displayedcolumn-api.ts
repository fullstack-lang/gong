// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplayedColumnAPI {

	static GONGSTRUCT_NAME = "DisplayedColumn"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	DisplayedColumnPointersEncoding: DisplayedColumnPointersEncoding = new DisplayedColumnPointersEncoding
}

export class DisplayedColumnPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
