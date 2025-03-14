// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLCellAPI {

	static GONGSTRUCT_NAME = "XLCell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for other decls

	XLCellPointersEncoding: XLCellPointersEncoding = new XLCellPointersEncoding
}

export class XLCellPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
