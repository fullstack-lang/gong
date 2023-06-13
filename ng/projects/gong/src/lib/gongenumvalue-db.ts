// insertion point for imports
import { GongEnumDB } from './gongenum-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValueDB {

	static GONGSTRUCT_NAME = "GongEnumValue"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for other declarations
	GongEnum_GongEnumValuesDBID: NullInt64 = new NullInt64
	GongEnum_GongEnumValuesDBID_Index: NullInt64  = new NullInt64 // store the index of the gongenumvalue instance in GongEnum.GongEnumValues
	GongEnum_GongEnumValues_reverse?: GongEnumDB 

}
