// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormSortAssocButtonDB {

	static GONGSTRUCT_NAME = "FormSortAssocButton"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for pointers and slices of pointers declarations

	FormSortAssocButtonPointersEncoding: FormSortAssocButtonPointersEncoding = new FormSortAssocButtonPointersEncoding
}

export class FormSortAssocButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
