// generated code - do not edit
// insertion point for imports
import { GongStruct } from './gongstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliceOfPointerToGongStructField {

	static GONGSTRUCT_NAME = "SliceOfPointerToGongStructField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for pointers and slices of pointers declarations
	GongStruct?: GongStruct

}
