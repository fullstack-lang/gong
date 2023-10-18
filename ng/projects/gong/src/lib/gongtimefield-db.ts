// insertion point for imports
import { GongStructDB } from './gongstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongTimeFieldDB {

	static GONGSTRUCT_NAME = "GongTimeField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for pointers and slices of pointers declarations

	GongTimeFieldPointersEncoding: GongTimeFieldPointersEncoding = new GongTimeFieldPointersEncoding
}

export class GongTimeFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	// reverse pointers encoding (to be removed)
	GongStruct_GongTimeFieldsDBID: NullInt64 = new NullInt64
	GongStruct_GongTimeFieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the gongtimefield instance in GongStruct.GongTimeFields
	GongStruct_GongTimeFields_reverse?: GongStructDB 

}
