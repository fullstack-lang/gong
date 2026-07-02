// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ToneAPI {

	static GONGSTRUCT_NAME = "Tone"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	TonePointersEncoding: TonePointersEncoding = new TonePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TonePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
