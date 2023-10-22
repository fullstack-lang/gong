// insertion point for imports
import { GongStructDB } from './gongstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliceOfPointerToGongStructFieldDB {

	static GONGSTRUCT_NAME = "SliceOfPointerToGongStructField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for pointers and slices of pointers declarations
	GongStruct?: GongStructDB


	SliceOfPointerToGongStructFieldPointersEncoding: SliceOfPointerToGongStructFieldPointersEncoding = new SliceOfPointerToGongStructFieldPointersEncoding
}

export class SliceOfPointerToGongStructFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongStructID: NullInt64 = new NullInt64 // if pointer is null, GongStruct.ID = 0

}
