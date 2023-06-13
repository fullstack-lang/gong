// insertion point for imports
import { GongEnumDB } from './gongenum-db'
import { GongStructDB } from './gongstruct-db'

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

	// insertion point for other declarations
	GongEnum?: GongEnumDB
	GongEnumID: NullInt64 = new NullInt64 // if pointer is null, GongEnum.ID = 0

	GongStruct_GongBasicFieldsDBID: NullInt64 = new NullInt64
	GongStruct_GongBasicFieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the gongbasicfield instance in GongStruct.GongBasicFields
	GongStruct_GongBasicFields_reverse?: GongStructDB 

}
