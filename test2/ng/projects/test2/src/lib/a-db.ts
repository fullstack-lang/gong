// insertion point for imports
import { BDB } from './b-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ADB {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Bs: Array<BDB> = []

	APointersEncoding: APointersEncoding = new APointersEncoding
}

export class APointersEncoding {
	Bs: number[] = []
}