// insertion point for imports
import { ClassshapeDB } from './classshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FieldDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Fieldname: string = ""
	FieldTypeAsString: string = ""
	Structname: string = ""
	Fieldtypename: string = ""

	// insertion point for other declarations
	Classshape_FieldsDBID: NullInt64 = new NullInt64
	Classshape_FieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the field instance in Classshape.Fields
	Classshape_Fields_reverse?: ClassshapeDB 

}
