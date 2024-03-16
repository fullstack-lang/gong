// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldTimeAPI {

	static GONGSTRUCT_NAME = "FormFieldTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date
	Step: number = 0

	// insertion point for other decls

	FormFieldTimePointersEncoding: FormFieldTimePointersEncoding = new FormFieldTimePointersEncoding
}

export class FormFieldTimePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
