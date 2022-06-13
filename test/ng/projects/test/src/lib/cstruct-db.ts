// insertion point for imports
import { BstructDB } from './bstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CstructDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	CName: string = ""
	CFloatfield: number = 0

	// insertion point for other declarations
	Bstruct?: BstructDB
	BstructID: NullInt64 = new NullInt64 // if pointer is null, Bstruct.ID = 0

}
