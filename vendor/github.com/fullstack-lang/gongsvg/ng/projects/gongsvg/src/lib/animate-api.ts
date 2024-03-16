// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnimateAPI {

	static GONGSTRUCT_NAME = "Animate"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AttributeName: string = ""
	Values: string = ""
	Dur: string = ""
	RepeatCount: string = ""

	// insertion point for other decls

	AnimatePointersEncoding: AnimatePointersEncoding = new AnimatePointersEncoding
}

export class AnimatePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
