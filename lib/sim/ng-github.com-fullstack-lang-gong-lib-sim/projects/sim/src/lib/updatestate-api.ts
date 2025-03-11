// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UpdateStateAPI {

	static GONGSTRUCT_NAME = "UpdateState"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0
	Period: number = 0

	// insertion point for other decls
	Duration_string?: string
	Period_string?: string

	UpdateStatePointersEncoding: UpdateStatePointersEncoding = new UpdateStatePointersEncoding
}

export class UpdateStatePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
