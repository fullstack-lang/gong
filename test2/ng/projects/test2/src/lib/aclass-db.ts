// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AclassDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Date: Date = new (Date)
	Booleanfield: string = ""
	Floatfield: number = 0
	Intfield: number = 0
	Anotherbooleanfield: string = ""
	Duration1: number = 0

	// insertion point for other declarations
	Duration1_string?: string
}
