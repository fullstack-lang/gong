// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormAPI {

	static GONGSTRUCT_NAME = "Form"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	FormPointersEncoding: FormPointersEncoding = new FormPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class FormPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
