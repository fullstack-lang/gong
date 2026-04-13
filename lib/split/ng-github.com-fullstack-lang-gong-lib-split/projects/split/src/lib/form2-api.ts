// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Form2API {

	static GONGSTRUCT_NAME = "Form2"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	Form2PointersEncoding: Form2PointersEncoding = new Form2PointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class Form2PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
