// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDateTimeAPI {

	static GONGSTRUCT_NAME = "FormFieldDateTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for other decls

	FormFieldDateTimePointersEncoding: FormFieldDateTimePointersEncoding = new FormFieldDateTimePointersEncoding
}

export class FormFieldDateTimePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
