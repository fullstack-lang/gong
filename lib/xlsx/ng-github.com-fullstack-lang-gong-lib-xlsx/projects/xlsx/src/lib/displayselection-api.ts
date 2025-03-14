// insertion point for imports
import { XLFileAPI } from './xlfile-api'
import { XLSheetAPI } from './xlsheet-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplaySelectionAPI {

	static GONGSTRUCT_NAME = "DisplaySelection"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	DisplaySelectionPointersEncoding: DisplaySelectionPointersEncoding = new DisplaySelectionPointersEncoding
}

export class DisplaySelectionPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	XLFileID: NullInt64 = new NullInt64 // if pointer is null, XLFile.ID = 0

	XLSheetID: NullInt64 = new NullInt64 // if pointer is null, XLSheet.ID = 0

}
