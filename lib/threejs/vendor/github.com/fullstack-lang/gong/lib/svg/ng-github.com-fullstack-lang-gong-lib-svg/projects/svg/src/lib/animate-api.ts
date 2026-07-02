// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnimateAPI {

	static GONGSTRUCT_NAME = "Animate"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AttributeName: string = ""
	Values: string = ""
	From: string = ""
	To: string = ""
	Dur: string = ""
	RepeatCount: string = ""

	// insertion point for other decls

	AnimatePointersEncoding: AnimatePointersEncoding = new AnimatePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class AnimatePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
