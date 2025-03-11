// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DummyAgentAPI {

	static GONGSTRUCT_NAME = "DummyAgent"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""
	Name: string = ""

	// insertion point for other decls

	DummyAgentPointersEncoding: DummyAgentPointersEncoding = new DummyAgentPointersEncoding
}

export class DummyAgentPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
