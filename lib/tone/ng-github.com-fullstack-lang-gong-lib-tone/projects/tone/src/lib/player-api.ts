// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PlayerAPI {

	static GONGSTRUCT_NAME = "Player"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Status: string = ""

	// insertion point for other decls

	PlayerPointersEncoding: PlayerPointersEncoding = new PlayerPointersEncoding
}

export class PlayerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
