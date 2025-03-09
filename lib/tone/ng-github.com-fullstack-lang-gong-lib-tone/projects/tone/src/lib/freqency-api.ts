// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FreqencyAPI {

	static GONGSTRUCT_NAME = "Freqency"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	FreqencyPointersEncoding: FreqencyPointersEncoding = new FreqencyPointersEncoding
}

export class FreqencyPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
