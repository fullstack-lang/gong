// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SvgAPI {

	static GONGSTRUCT_NAME = "Svg"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	SvgPointersEncoding: SvgPointersEncoding = new SvgPointersEncoding
}

export class SvgPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
