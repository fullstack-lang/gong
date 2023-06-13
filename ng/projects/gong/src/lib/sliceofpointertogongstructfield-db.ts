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

	// insertion point for other declarations
	GongStruct?: GongStructDB
	GongStructID: NullInt64 = new NullInt64 // if pointer is null, GongStruct.ID = 0

	GongStruct_SliceOfPointerToGongStructFieldsDBID: NullInt64 = new NullInt64
	GongStruct_SliceOfPointerToGongStructFieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the sliceofpointertogongstructfield instance in GongStruct.SliceOfPointerToGongStructFields
	GongStruct_SliceOfPointerToGongStructFields_reverse?: GongStructDB 

}
