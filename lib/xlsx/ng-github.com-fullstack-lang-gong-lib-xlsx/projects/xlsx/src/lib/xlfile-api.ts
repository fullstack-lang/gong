// insertion point for imports
import { XLSheetAPI } from './xlsheet-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLFileAPI {

	static GONGSTRUCT_NAME = "XLFile"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NbSheets: number = 0

	// insertion point for other decls

	XLFilePointersEncoding: XLFilePointersEncoding = new XLFilePointersEncoding
}

export class XLFilePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Sheets: number[] = []
}
