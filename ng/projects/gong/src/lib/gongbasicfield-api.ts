// insertion point for imports
import { GongEnumAPI } from './gongenum-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongBasicFieldAPI {

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
	IsBespokeWidth: boolean = false
	BespokeWidth: number = 0
	IsBespokeHeight: boolean = false
	BespokeHeight: number = 0

	// insertion point for other decls

	GongBasicFieldPointersEncoding: GongBasicFieldPointersEncoding = new GongBasicFieldPointersEncoding
}

export class GongBasicFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	GongEnumID: NullInt64 = new NullInt64 // if pointer is null, GongEnum.ID = 0

}
