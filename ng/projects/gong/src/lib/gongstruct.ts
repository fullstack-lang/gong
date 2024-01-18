// generated code - do not edit
// insertion point for imports
import { GongBasicField } from './gongbasicfield'
import { GongTimeField } from './gongtimefield'
import { PointerToGongStructField } from './pointertogongstructfield'
import { SliceOfPointerToGongStructField } from './sliceofpointertogongstructfield'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongStruct {

	static GONGSTRUCT_NAME = "GongStruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasOnAfterUpdateSignature: boolean = false
	IsIgnoredForFront: boolean = false

	// insertion point for pointers and slices of pointers declarations
	GongBasicFields: Array<GongBasicField> = []
	GongTimeFields: Array<GongTimeField> = []
	PointerToGongStructFields: Array<PointerToGongStructField> = []
	SliceOfPointerToGongStructFields: Array<SliceOfPointerToGongStructField> = []
}
