// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CheckBoxAPI {

	static GONGSTRUCT_NAME = "CheckBox"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for other decls

	CheckBoxPointersEncoding: CheckBoxPointersEncoding = new CheckBoxPointersEncoding
}

export class CheckBoxPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
