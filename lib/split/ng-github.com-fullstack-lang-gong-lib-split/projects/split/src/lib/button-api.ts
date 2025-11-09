// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonAPI {

	static GONGSTRUCT_NAME = "Button"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	ButtonPointersEncoding: ButtonPointersEncoding = new ButtonPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
