// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormEditAssocButtonAPI {

	static GONGSTRUCT_NAME = "FormEditAssocButton"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for other decls

	FormEditAssocButtonPointersEncoding: FormEditAssocButtonPointersEncoding = new FormEditAssocButtonPointersEncoding
}

export class FormEditAssocButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
