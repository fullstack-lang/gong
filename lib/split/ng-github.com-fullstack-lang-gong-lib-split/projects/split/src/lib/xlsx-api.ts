// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XlsxAPI {

	static GONGSTRUCT_NAME = "Xlsx"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	XlsxPointersEncoding: XlsxPointersEncoding = new XlsxPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class XlsxPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
