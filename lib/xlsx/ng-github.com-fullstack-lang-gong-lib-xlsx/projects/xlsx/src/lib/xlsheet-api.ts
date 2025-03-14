// insertion point for imports
import { XLRowAPI } from './xlrow-api'
import { XLCellAPI } from './xlcell-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLSheetAPI {

	static GONGSTRUCT_NAME = "XLSheet"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	MaxRow: number = 0
	MaxCol: number = 0
	NbRows: number = 0

	// insertion point for other decls

	XLSheetPointersEncoding: XLSheetPointersEncoding = new XLSheetPointersEncoding
}

export class XLSheetPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Rows: number[] = []
	SheetCells: number[] = []
}
