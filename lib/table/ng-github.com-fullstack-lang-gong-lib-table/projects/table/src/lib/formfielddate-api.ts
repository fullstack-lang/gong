// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDateAPI {

	static GONGSTRUCT_NAME = "FormFieldDate"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for other decls

	FormFieldDatePointersEncoding: FormFieldDatePointersEncoding = new FormFieldDatePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class FormFieldDatePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
