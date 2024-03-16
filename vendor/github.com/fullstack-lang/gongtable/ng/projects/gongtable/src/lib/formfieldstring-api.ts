// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldStringAPI {

	static GONGSTRUCT_NAME = "FormFieldString"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""
	IsTextArea: boolean = false

	// insertion point for other decls

	FormFieldStringPointersEncoding: FormFieldStringPointersEncoding = new FormFieldStringPointersEncoding
}

export class FormFieldStringPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
