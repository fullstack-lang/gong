// insertion point for imports
import { GongStructShapeDB } from './gongstructshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FieldDB {

	static GONGSTRUCT_NAME = "Field"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	FieldTypeAsString: string = ""
	Structname: string = ""
	Fieldtypename: string = ""

	// insertion point for other declarations
	GongStructShape_FieldsDBID: NullInt64 = new NullInt64
	GongStructShape_FieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the field instance in GongStructShape.Fields
	GongStructShape_Fields_reverse?: GongStructShapeDB 

}
