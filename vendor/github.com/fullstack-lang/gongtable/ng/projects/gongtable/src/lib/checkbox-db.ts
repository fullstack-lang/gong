// insertion point for imports
import { FormDivDB } from './formdiv-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CheckBoxDB {

	static GONGSTRUCT_NAME = "CheckBox"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for other declarations
	FormDiv_CheckBoxsDBID: NullInt64 = new NullInt64
	FormDiv_CheckBoxsDBID_Index: NullInt64  = new NullInt64 // store the index of the checkbox instance in FormDiv.CheckBoxs
	FormDiv_CheckBoxs_reverse?: FormDivDB 

}
