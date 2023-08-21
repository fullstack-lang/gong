// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellBooleanDB {

	static GONGSTRUCT_NAME = "CellBoolean"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for other declarations
}
