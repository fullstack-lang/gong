// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CheckboxAPI {

	static GONGSTRUCT_NAME = "Checkbox"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ValueBool: boolean = false
	LabelForTrue: string = ""
	LabelForFalse: string = ""

	// insertion point for other decls

	CheckboxPointersEncoding: CheckboxPointersEncoding = new CheckboxPointersEncoding
}

export class CheckboxPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
