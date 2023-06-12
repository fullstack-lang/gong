// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PositionDB {

	static GONGSTRUCT_NAME = "Position"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	X: number = 0
	Y: number = 0
	Name: string = ""

	// insertion point for other declarations
}
