// insertion point for imports
import { FormDivDB } from './formdiv-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormGroupDB {

	static GONGSTRUCT_NAME = "FormGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasSuppressButton: boolean = false
	HasSuppressButtonBeenPressed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	FormDivs: Array<FormDivDB> = []

	FormGroupPointersEncoding: FormGroupPointersEncoding = new FormGroupPointersEncoding
}

export class FormGroupPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormDivs: number[] = []
}
