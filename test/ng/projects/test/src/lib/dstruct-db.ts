// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DstructDB {
	CreatedAt?: string
	DeletedAt?: string

	// With the multiple stacks, each sack has its own DB, therefore, 
	// the unique identifier of an instance is a combination of the ID and the stack path
	ID: number = 0
	GONG__StackPath: string = ""

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
}
