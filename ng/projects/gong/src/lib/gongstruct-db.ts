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
	IsIgnoredForFront: boolean = false

	// insertion point for pointers and slices of pointers declarations
	GongBasicFields: Array<GongBasicFieldDB> = []
	GongTimeFields: Array<GongTimeFieldDB> = []
	PointerToGongStructFields: Array<PointerToGongStructFieldDB> = []
	SliceOfPointerToGongStructFields: Array<SliceOfPointerToGongStructFieldDB> = []

	GongStructPointersEncoding: GongStructPointersEncoding = new GongStructPointersEncoding
}

export class GongStructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongBasicFields: number[] = []
	GongTimeFields: number[] = []
	PointerToGongStructFields: number[] = []
	SliceOfPointerToGongStructFields: number[] = []
}
