// insertion point for imports
import { GongBasicFieldDB } from './gongbasicfield-db'
import { GongTimeFieldDB } from './gongtimefield-db'
import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongStructDB {

	static GONGSTRUCT_NAME = "GongStruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasOnAfterUpdateSignature: boolean = false

	// insertion point for other declarations
	GongBasicFields?: Array<GongBasicFieldDB>
	GongTimeFields?: Array<GongTimeFieldDB>
	PointerToGongStructFields?: Array<PointerToGongStructFieldDB>
	SliceOfPointerToGongStructFields?: Array<SliceOfPointerToGongStructFieldDB>
}
