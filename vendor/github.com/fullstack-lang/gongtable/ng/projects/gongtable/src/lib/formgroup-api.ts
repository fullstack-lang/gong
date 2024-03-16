// insertion point for imports
import { FormDivAPI } from './formdiv-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormGroupAPI {

	static GONGSTRUCT_NAME = "FormGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasSuppressButton: boolean = false
	HasSuppressButtonBeenPressed: boolean = false

	// insertion point for other decls

	FormGroupPointersEncoding: FormGroupPointersEncoding = new FormGroupPointersEncoding
}

export class FormGroupPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormDivs: number[] = []
}
