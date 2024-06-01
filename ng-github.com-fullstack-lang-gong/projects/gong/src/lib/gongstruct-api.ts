// insertion point for imports
import { GongBasicFieldAPI } from './gongbasicfield-api'
import { GongTimeFieldAPI } from './gongtimefield-api'
import { PointerToGongStructFieldAPI } from './pointertogongstructfield-api'
import { SliceOfPointerToGongStructFieldAPI } from './sliceofpointertogongstructfield-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongStructAPI {

	static GONGSTRUCT_NAME = "GongStruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasOnAfterUpdateSignature: boolean = false
	IsIgnoredForFront: boolean = false

	// insertion point for other decls

	GongStructPointersEncoding: GongStructPointersEncoding = new GongStructPointersEncoding
}

export class GongStructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongBasicFields: number[] = []
	GongTimeFields: number[] = []
	PointerToGongStructFields: number[] = []
	SliceOfPointerToGongStructFields: number[] = []
}
