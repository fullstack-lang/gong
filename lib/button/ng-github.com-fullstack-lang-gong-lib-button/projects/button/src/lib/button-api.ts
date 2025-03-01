// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonAPI {

	static GONGSTRUCT_NAME = "Button"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	Icon: string = ""

	// insertion point for other decls

	ButtonPointersEncoding: ButtonPointersEncoding = new ButtonPointersEncoding
}

export class ButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
