// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class EventAPI {

	static GONGSTRUCT_NAME = "Event"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0

	// insertion point for other decls
	Duration_string?: string

	EventPointersEncoding: EventPointersEncoding = new EventPointersEncoding
}

export class EventPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
