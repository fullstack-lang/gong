// insertion point for imports
import { FormFieldSelectDB } from './formfieldselect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class OptionDB {

	static GONGSTRUCT_NAME = "Option"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations

	OptionPointersEncoding: OptionPointersEncoding = new OptionPointersEncoding
}

export class OptionPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	// reverse pointers encoding (to be removed)
	FormFieldSelect_OptionsDBID: NullInt64 = new NullInt64
	FormFieldSelect_OptionsDBID_Index: NullInt64  = new NullInt64 // store the index of the option instance in FormFieldSelect.Options
	FormFieldSelect_Options_reverse?: FormFieldSelectDB 

}
