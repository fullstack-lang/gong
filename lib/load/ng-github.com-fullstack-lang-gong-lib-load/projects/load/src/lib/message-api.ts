// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MessageAPI {

	static GONGSTRUCT_NAME = "Message"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	MessagePointersEncoding: MessagePointersEncoding = new MessagePointersEncoding
}

export class MessagePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
