// insertion point for imports
import { OptionAPI } from './option-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldSelectAPI {

	static GONGSTRUCT_NAME = "FormFieldSelect"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CanBeEmpty: boolean = false

	// insertion point for other decls

	FormFieldSelectPointersEncoding: FormFieldSelectPointersEncoding = new FormFieldSelectPointersEncoding
}

export class FormFieldSelectPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ValueID: NullInt64 = new NullInt64 // if pointer is null, Value.ID = 0

	Options: number[] = []
}
