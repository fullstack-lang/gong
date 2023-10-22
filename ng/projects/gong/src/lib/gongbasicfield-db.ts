// insertion point for imports
import { GongEnumDB } from './gongenum-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongBasicFieldDB {

	static GONGSTRUCT_NAME = "GongBasicField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	BasicKindName: string = ""
	DeclaredType: string = ""
	CompositeStructName: string = ""
	Index: number = 0
	IsDocLink: boolean = false
	IsTextArea: boolean = false

	// insertion point for pointers and slices of pointers declarations
	GongEnum?: GongEnumDB


	GongBasicFieldPointersEncoding: GongBasicFieldPointersEncoding = new GongBasicFieldPointersEncoding
}

export class GongBasicFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongEnumID: NullInt64 = new NullInt64 // if pointer is null, GongEnum.ID = 0

}
