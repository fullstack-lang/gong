// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldIntDB {

	static GONGSTRUCT_NAME = "FormFieldInt"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0
	HasMinValidator: boolean = false
	MinValue: number = 0
	HasMaxValidator: boolean = false
	MaxValue: number = 0

	// insertion point for pointers and slices of pointers declarations

	FormFieldIntPointersEncoding: FormFieldIntPointersEncoding = new FormFieldIntPointersEncoding
}

export class FormFieldIntPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
