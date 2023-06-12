// insertion point for imports
import { GongEnumShapeDB } from './gongenumshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValueEntryDB {

	static GONGSTRUCT_NAME = "GongEnumValueEntry"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""

	// insertion point for other declarations
	GongEnumShape_GongEnumValueEntrysDBID: NullInt64 = new NullInt64
	GongEnumShape_GongEnumValueEntrysDBID_Index: NullInt64  = new NullInt64 // store the index of the gongenumvalueentry instance in GongEnumShape.GongEnumValueEntrys
	GongEnumShape_GongEnumValueEntrys_reverse?: GongEnumShapeDB 

}
