// insertion point for imports
import { GongStructAPI } from './gongstruct-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliceOfPointerToGongStructFieldAPI {

	static GONGSTRUCT_NAME = "SliceOfPointerToGongStructField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for other decls

	SliceOfPointerToGongStructFieldPointersEncoding: SliceOfPointerToGongStructFieldPointersEncoding = new SliceOfPointerToGongStructFieldPointersEncoding
}

export class SliceOfPointerToGongStructFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongStructID: NullInt64 = new NullInt64 // if pointer is null, GongStruct.ID = 0

}
